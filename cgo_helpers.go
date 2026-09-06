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
#cgo CFLAGS: -I.
#cgo !windows LDFLAGS: -lswe
#cgo windows LDFLAGS: -lswedll64
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

// sliceToDoublePtr returns a pointer to the first element of slice.
// If slice is empty or nil, it returns nil.
func sliceToDoublePtr(s []float64) *C.double {
	if len(s) == 0 {
		return nil
	}
	return (*C.double)(unsafe.Pointer(unsafe.SliceData(s)))
}

// safeDoubleOutput ensures that the slice s has at least minLen elements before passing to C.
// If s has at least minLen elements, its underlying pointer is returned directly (zero allocations).
// If s has fewer than minLen elements, a temporary buffer is allocated and a copyback function is returned.
func safeDoubleOutput(s []float64, minLen int) (*C.double, func()) {
	if len(s) >= minLen {
		return (*C.double)(unsafe.Pointer(unsafe.SliceData(s))), nil
	}
	buf := make([]float64, minLen)
	return (*C.double)(unsafe.Pointer(unsafe.SliceData(buf))), func() {
		if len(s) > 0 {
			copy(s, buf)
		}
	}
}

// safeDoubleInput ensures that an input slice has at least minLen elements to prevent out-of-bounds reads.
func safeDoubleInput(s []float64, minLen int) *C.double {
	if len(s) >= minLen {
		return (*C.double)(unsafe.Pointer(unsafe.SliceData(s)))
	}
	if len(s) == 0 {
		return nil
	}
	buf := make([]float64, minLen)
	copy(buf, s)
	return (*C.double)(unsafe.Pointer(unsafe.SliceData(buf)))
}

// safeErrorBuf ensures that Swiss Ephemeris receives a buffer of at least 256 bytes for serr.
// If serr has at least 256 bytes, its pointer is used directly with zero allocations.
// If serr is nil or shorter than 256 bytes, a safe buffer is used to prevent memory corruption.
func safeErrorBuf(serr []byte) (*C.char, func()) {
	if len(serr) >= 256 {
		return (*C.char)(unsafe.Pointer(unsafe.SliceData(serr))), nil
	}
	buf := make([]byte, 256)
	return (*C.char)(unsafe.Pointer(unsafe.SliceData(buf))), func() {
		if len(serr) > 0 {
			copy(serr, buf)
		}
	}
}

// safeCString converts a byte slice into a null-terminated C string pointer.
// If the byte slice already contains a null byte, it returns a direct pointer with zero allocations.
// Otherwise, it creates a null-terminated copy to prevent buffer overreads.
func safeCString(b []byte) *C.char {
	if len(b) == 0 {
		return nil
	}
	if idx := bytes.IndexByte(b, 0); idx >= 0 {
		return (*C.char)(unsafe.Pointer(unsafe.SliceData(b)))
	}
	buf := make([]byte, len(b)+1)
	copy(buf, b)
	buf[len(b)] = 0
	return (*C.char)(unsafe.Pointer(unsafe.SliceData(buf)))
}

// stringToCString converts a Go string into a null-terminated C string pointer.
func stringToCString(s string) *C.char {
	if len(s) == 0 {
		return nil
	}
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = 0
	return (*C.char)(unsafe.Pointer(unsafe.SliceData(buf)))
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
