// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/gpl-3.0.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gose

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo darwin LDFLAGS: -L${SRCDIR}/../swisseph -Wl,-rpath,${SRCDIR}/../swisseph -L/opt/homebrew/lib -L/usr/local/lib -lswe -lm
#cgo linux LDFLAGS: -L${SRCDIR}/../swisseph -Wl,-rpath,${SRCDIR}/../swisseph -L/usr/local/lib -L/usr/lib -lswe -lm
#cgo !darwin,!linux,!windows LDFLAGS: -L${SRCDIR}/../swisseph -lswe -lm
#cgo windows LDFLAGS: -L${SRCDIR}/../swisseph -lswedll64
#include "swephexp.h"
#include "sweph.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"bytes"
	"unsafe"
)

// emptyCChar is a static NUL byte used when a nil or empty string is passed to C,
// ensuring C functions never dereference a NULL char* pointer.
var emptyCChar C.char

// goseScratch holds pre-allocated buffers and scalar variables protected by goseMutex.
// Using package-level scratch memory under goseMutex eliminates CGo pointer-escape heap
// allocations on hot paths while guaranteeing that Swiss Ephemeris C functions always
// receive valid, sufficiently sized buffers.
var goseScratch struct {
	serr   [256]byte
	star   [512]byte
	samod  [256]byte
	sdet   [4096]byte
	spname [256]byte
	csbuf  [64]byte
	cstr   [1024]byte

	out1 [50]float64
	out2 [50]float64
	out3 [50]float64
	out4 [50]float64

	in1 [6]float64
	in2 [6]float64
	in3 [6]float64

	d0, d1, d2         C.double
	i0, i1, i2, i3, i4 C.int32
	ci0, ci1, ci2      C.int
}

// sliceToDoublePtr returns a pointer to the first element of slice.
// If slice is empty or nil, it returns nil.
func sliceToDoublePtr(s []float64) *C.double {
	if len(s) == 0 {
		return nil
	}
	return (*C.double)(unsafe.Pointer(unsafe.SliceData(s)))
}

// doubleOutputPtr returns a *C.double backed directly by s if len(s) >= minLen,
// or backed by scratch[:minLen] if len(s) < minLen.
// Caller must hold goseMutex when scratch belongs to goseScratch.
func doubleOutputPtr(s []float64, minLen int, scratch []float64) *C.double {
	if len(s) >= minLen {
		return (*C.double)(unsafe.Pointer(unsafe.SliceData(s)))
	}
	clear(scratch[:minLen])
	return (*C.double)(unsafe.Pointer(unsafe.SliceData(scratch)))
}

// doubleInOutPtr returns a *C.double backed directly by s if len(s) >= minLen,
// or backed by scratch[:minLen] (initialized with the elements of s) if len(s) < minLen.
// Caller must hold goseMutex when scratch belongs to goseScratch.
func doubleInOutPtr(s []float64, minLen int, scratch []float64) *C.double {
	if len(s) >= minLen {
		return (*C.double)(unsafe.Pointer(unsafe.SliceData(s)))
	}
	clear(scratch[:minLen])
	copy(scratch[:minLen], s)
	return (*C.double)(unsafe.Pointer(unsafe.SliceData(scratch)))
}

// copyDoubleOutput copies scratch[:minLen] back into dst when 0 < len(dst) < minLen.
func copyDoubleOutput(dst []float64, minLen int, scratch []float64) {
	if len(dst) > 0 && len(dst) < minLen {
		copy(dst, scratch[:minLen])
	}
}

// doubleInputPtr ensures an input slice has at least minLen elements and never returns NULL.
// Caller must hold goseMutex when scratch belongs to goseScratch.
func doubleInputPtr(s []float64, minLen int, scratch []float64) *C.double {
	if len(s) >= minLen {
		return (*C.double)(unsafe.Pointer(unsafe.SliceData(s)))
	}
	clear(scratch[:minLen])
	copy(scratch[:minLen], s)
	return (*C.double)(unsafe.Pointer(unsafe.SliceData(scratch)))
}

// errorBufPtr returns a 256-byte *C.char buffer for Swiss Ephemeris error messages.
// Caller must hold goseMutex.
func errorBufPtr(serr []byte) *C.char {
	if len(serr) >= 256 {
		serr[0] = 0
		return (*C.char)(unsafe.Pointer(unsafe.SliceData(serr)))
	}
	goseScratch.serr[0] = 0
	return (*C.char)(unsafe.Pointer(&goseScratch.serr[0]))
}

// copyErrorBuf copies goseScratch.serr into serr when 0 < len(serr) < 256.
// Caller must hold goseMutex.
func copyErrorBuf(serr []byte) {
	if len(serr) > 0 && len(serr) < 256 {
		copy(serr, goseScratch.serr[:])
	}
}

// starInOutPtr copies star into goseScratch.star (512 bytes = 2*SeMaxStname),
// guarantees NUL-termination, and returns a pointer to goseScratch.star[0].
// Caller must hold goseMutex.
func starInOutPtr(star []byte) (*C.char, bool) {
	clear(goseScratch.star[:])
	inLen := 0
	if len(star) > 0 {
		if idx := bytes.IndexByte(star, 0); idx >= 0 {
			inLen = idx
		} else {
			inLen = len(star)
		}
		n := inLen
		if n >= len(goseScratch.star) {
			n = len(goseScratch.star) - 1
		}
		copy(goseScratch.star[:n], star[:n])
	}
	hasOutSpace := len(star) > inLen+1
	return (*C.char)(unsafe.Pointer(&goseScratch.star[0])), hasOutSpace
}

// optionalStarInOutPtr optimizes the common planet case where starname is nil or empty.
// Caller must hold goseMutex.
func optionalStarInOutPtr(star []byte) (*C.char, bool) {
	if len(star) == 0 || star[0] == 0 {
		return &emptyCChar, false
	}
	return starInOutPtr(star)
}

// copyStarOut copies the formatted star name ("trad_name,bayer_name") written by C
// back into star when star has sufficient capacity for output.
// Caller must hold goseMutex.
func copyStarOut(star []byte, hasOutSpace bool) {
	if len(star) == 0 {
		return
	}
	outLen := bytes.IndexByte(goseScratch.star[:], 0)
	if outLen < 0 {
		outLen = len(goseScratch.star) - 1
	}
	if len(star) >= outLen+1 {
		copy(star, goseScratch.star[:outLen+1])
	} else if hasOutSpace {
		copy(star, goseScratch.star[:len(star)-1])
		star[len(star)-1] = 0
	}
}

// cStringInputPtr converts a byte slice into a NUL-terminated *C.char without heap allocation
// for strings up to 1023 bytes. Never returns NULL. Caller must hold goseMutex.
func cStringInputPtr(b []byte) *C.char {
	if len(b) == 0 {
		return &emptyCChar
	}
	if bytes.IndexByte(b, 0) >= 0 {
		return (*C.char)(unsafe.Pointer(unsafe.SliceData(b)))
	}
	if len(b) < len(goseScratch.cstr) {
		copy(goseScratch.cstr[:len(b)], b)
		goseScratch.cstr[len(b)] = 0
		return (*C.char)(unsafe.Pointer(&goseScratch.cstr[0]))
	}
	buf := make([]byte, len(b)+1)
	copy(buf, b)
	buf[len(b)] = 0
	return (*C.char)(unsafe.Pointer(unsafe.SliceData(buf)))
}

// goStringInputPtr converts a Go string into a NUL-terminated *C.char without heap allocation
// for strings up to 1023 bytes. Never returns NULL. Caller must hold goseMutex.
func goStringInputPtr(s string) *C.char {
	if len(s) == 0 {
		return &emptyCChar
	}
	if len(s) < len(goseScratch.cstr) {
		copy(goseScratch.cstr[:len(s)], s)
		goseScratch.cstr[len(s)] = 0
		return (*C.char)(unsafe.Pointer(&goseScratch.cstr[0]))
	}
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = 0
	return (*C.char)(unsafe.Pointer(unsafe.SliceData(buf)))
}

// houseCuspsMinLen returns the required cusp buffer length for the given house system.
// Gauquelin sectors ('G' or 'g') write 37 elements (indices 0..36); all other systems write 13 (indices 0..12).
func houseCuspsMinLen(hsys int) int {
	if hsys == 'G' || hsys == 'g' {
		return 37
	}
	return 13
}

// cCharToString safely converts a null-terminated *C.char into a Go string.
func cCharToString(p *C.char) string {
	if p == nil {
		return ""
	}
	return C.GoString(p)
}

// cCharToBytePtr returns *byte from *C.char for backward compatibility.
func cCharToBytePtr(p *C.char) *byte {
	return (*byte)(unsafe.Pointer(p))
}

// cBytesToString extracts a Go string from a null-terminated byte slice.
func cBytesToString(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	n := bytes.IndexByte(b, 0)
	if n < 0 {
		n = len(b)
	}
	return string(b[:n])
}

// Compatibility types and methods to ensure backward compatibility
type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

type cgoAllocMap struct{}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {}
func (a *cgoAllocMap) IsEmpty() bool          { return true }
func (a *cgoAllocMap) Borrow(b *cgoAllocMap)  {}
func (a *cgoAllocMap) Free()                  {}

func copyPDoubleBytes(slice *sliceHeader) (*C.double, *cgoAllocMap) {
	if slice == nil || slice.Data == nil {
		return nil, cgoAllocsUnknown
	}
	return (*C.double)(slice.Data), cgoAllocsUnknown
}

func copyPCharBytes(slice *sliceHeader) (*C.char, *cgoAllocMap) {
	if slice == nil || slice.Data == nil {
		return nil, cgoAllocsUnknown
	}
	return (*C.char)(slice.Data), cgoAllocsUnknown
}

func copyPIntBytes(slice *sliceHeader) (*C.int, *cgoAllocMap) {
	if slice == nil || slice.Data == nil {
		return nil, cgoAllocsUnknown
	}
	return (*C.int)(slice.Data), cgoAllocsUnknown
}

func copyPInt32Bytes(slice *sliceHeader) (*C.int32, *cgoAllocMap) {
	if slice == nil || slice.Data == nil {
		return nil, cgoAllocsUnknown
	}
	return (*C.int32)(slice.Data), cgoAllocsUnknown
}

func packPCharString(p *C.char) string {
	return cCharToString(p)
}

func (x *SweData) Ref() *C.struct_swe_data {
	if x == nil {
		return nil
	}
	return x.ref84c87565
}

func (x *SweData) Free() {
	if x != nil {
		x.ref84c87565 = nil
	}
}
