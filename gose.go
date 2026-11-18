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
#include "swephexp.h"
#include "sweph.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"bytes"
	"errors"
	"math"
	"sync"
	"time"
	"unsafe"
)

var goseMutex sync.Mutex

// HeliacalUt function as declared in swephexp.h:679
func HeliacalUt(tjdstartUt float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, typeEvent int, iflag int, dret []float64, serr []byte) int32 {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	cdatm := doubleInOutPtr(datm, 4, goseScratch.in2[:])
	cdobs := doubleInOutPtr(dobs, 6, goseScratch.in3[:])
	cobjectName, hasOutSpace := starInOutPtr(objectName)
	cdret := doubleOutputPtr(dret, 50, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_heliacal_ut(C.double(tjdstartUt), cgeopos, cdatm, cdobs, cobjectName, C.int32(typeEvent), C.int32(iflag), cdret, pserr)
	copyDoubleOutput(datm, 4, goseScratch.in2[:])
	copyDoubleOutput(dobs, 6, goseScratch.in3[:])
	copyStarOut(objectName, hasOutSpace)
	copyDoubleOutput(dret, 50, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// HeliacalPhenoUt function as declared in swephexp.h:680
func HeliacalPhenoUt(tjdUt float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, typeEvent int, helflag int, darr []float64, serr []byte) int32 {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	cdatm := doubleInOutPtr(datm, 4, goseScratch.in2[:])
	cdobs := doubleInOutPtr(dobs, 6, goseScratch.in3[:])
	cobjectName, hasOutSpace := starInOutPtr(objectName)
	cdarr := doubleOutputPtr(darr, 50, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_heliacal_pheno_ut(C.double(tjdUt), cgeopos, cdatm, cdobs, cobjectName, C.int32(typeEvent), C.int32(helflag), cdarr, pserr)
	copyDoubleOutput(datm, 4, goseScratch.in2[:])
	copyDoubleOutput(dobs, 6, goseScratch.in3[:])
	copyStarOut(objectName, hasOutSpace)
	copyDoubleOutput(darr, 50, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// VisLimitMag function as declared in swephexp.h:681
func VisLimitMag(tjdut float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, helflag int, dret []float64, serr []byte) int32 {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	cdatm := doubleInOutPtr(datm, 4, goseScratch.in2[:])
	cdobs := doubleInOutPtr(dobs, 6, goseScratch.in3[:])
	cobjectName, hasOutSpace := starInOutPtr(objectName)
	cdret := doubleOutputPtr(dret, 50, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_vis_limit_mag(C.double(tjdut), cgeopos, cdatm, cdobs, cobjectName, C.int32(helflag), cdret, pserr)
	copyDoubleOutput(datm, 4, goseScratch.in2[:])
	copyDoubleOutput(dobs, 6, goseScratch.in3[:])
	copyStarOut(objectName, hasOutSpace)
	copyDoubleOutput(dret, 50, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// HeliacalAngle function as declared in swephexp.h:684
func HeliacalAngle(tjdut float64, dgeo []float64, datm []float64, dobs []float64, helflag int, mag float64, aziObj float64, aziSun float64, aziMoon float64, altMoon float64, dret []float64, serr []byte) int32 {
	goseMutex.Lock()
	cdgeo := doubleInputPtr(dgeo, 3, goseScratch.in1[:])
	cdatm := doubleInOutPtr(datm, 4, goseScratch.in2[:])
	cdobs := doubleInOutPtr(dobs, 6, goseScratch.in3[:])
	cdret := doubleOutputPtr(dret, 50, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_heliacal_angle(C.double(tjdut), cdgeo, cdatm, cdobs, C.int32(helflag), C.double(mag), C.double(aziObj), C.double(aziSun), C.double(aziMoon), C.double(altMoon), cdret, pserr)
	copyDoubleOutput(datm, 4, goseScratch.in2[:])
	copyDoubleOutput(dobs, 6, goseScratch.in3[:])
	copyDoubleOutput(dret, 50, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// TopoArcusVisionis function as declared in swephexp.h:685
func TopoArcusVisionis(tjdut float64, dgeo []float64, datm []float64, dobs []float64, helflag int, mag float64, aziObj float64, altObj float64, aziSun float64, aziMoon float64, altMoon float64, dret []float64, serr []byte) int32 {
	goseMutex.Lock()
	cdgeo := doubleInputPtr(dgeo, 3, goseScratch.in1[:])
	cdatm := doubleInOutPtr(datm, 4, goseScratch.in2[:])
	cdobs := doubleInOutPtr(dobs, 6, goseScratch.in3[:])
	cdret := doubleOutputPtr(dret, 50, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_topo_arcus_visionis(C.double(tjdut), cdgeo, cdatm, cdobs, C.int32(helflag), C.double(mag), C.double(aziObj), C.double(altObj), C.double(aziSun), C.double(aziMoon), C.double(altMoon), cdret, pserr)
	copyDoubleOutput(datm, 4, goseScratch.in2[:])
	copyDoubleOutput(dobs, 6, goseScratch.in3[:])
	copyDoubleOutput(dret, 50, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// SetAstroModels function as declared in swephexp.h:689
func SetAstroModels(samod []byte, iflag int32) {
	goseMutex.Lock()
	csamod := cStringInputPtr(samod)
	C.swe_set_astro_models(csamod, C.int32(iflag))
	goseMutex.Unlock()
}

// GetAstroModels function as declared in swephexp.h:690
func GetAstroModels(samod []byte, sdet []byte, iflag int32) {
	goseMutex.Lock()
	clear(goseScratch.samod[:])
	goseScratch.sdet[0] = 0
	if len(samod) > 0 {
		n := bytes.IndexByte(samod, 0)
		if n < 0 {
			n = len(samod)
		}
		if n >= len(goseScratch.samod) {
			n = len(goseScratch.samod) - 1
		}
		copy(goseScratch.samod[:n], samod[:n])
	}
	csamod := (*C.char)(unsafe.Pointer(&goseScratch.samod[0]))
	csdet := (*C.char)(unsafe.Pointer(&goseScratch.sdet[0]))
	C.swe_get_astro_models(csamod, csdet, C.int32(iflag))
	if len(samod) > 0 {
		outLen := bytes.IndexByte(goseScratch.samod[:], 0)
		if outLen < 0 {
			outLen = len(goseScratch.samod) - 1
		}
		if len(samod) > outLen {
			copy(samod, goseScratch.samod[:outLen+1])
		} else {
			copy(samod, goseScratch.samod[:len(samod)])
		}
	}
	if len(sdet) > 0 {
		outLen := bytes.IndexByte(goseScratch.sdet[:], 0)
		if outLen < 0 {
			outLen = len(goseScratch.sdet) - 1
		}
		if len(sdet) > outLen {
			copy(sdet, goseScratch.sdet[:outLen+1])
		} else {
			copy(sdet, goseScratch.sdet[:len(sdet)])
		}
	}
	goseMutex.Unlock()
}

// Version function as declared in swephexp.h:696
func Version(arg0 []byte) *byte {
	if len(arg0) >= 256 {
		ret := C.swe_version((*C.char)(unsafe.Pointer(unsafe.SliceData(arg0))))
		return cCharToBytePtr(ret)
	}
	goseMutex.Lock()
	goseScratch.spname[0] = 0
	ret := C.swe_version((*C.char)(unsafe.Pointer(&goseScratch.spname[0])))
	if len(arg0) > 0 {
		copy(arg0, goseScratch.spname[:])
		goseMutex.Unlock()
		return &arg0[0]
	}
	goseMutex.Unlock()
	return cCharToBytePtr(ret)
}

// GetDenum function as declared in swephexp.h:697
func GetDenum() int32 {
	goseMutex.Lock()
	ret := C.swe_get_denum()
	goseMutex.Unlock()
	return int32(ret)
}

// GetLibraryPath function as declared in swephexp.h:698
func GetLibraryPath(arg0 []byte) *byte {
	goseMutex.Lock()
	if len(arg0) >= 256 {
		ret := C.swe_get_library_path((*C.char)(unsafe.Pointer(unsafe.SliceData(arg0))))
		goseMutex.Unlock()
		return cCharToBytePtr(ret)
	}
	goseScratch.spname[0] = 0
	ret := C.swe_get_library_path((*C.char)(unsafe.Pointer(&goseScratch.spname[0])))
	if len(arg0) > 0 {
		copy(arg0, goseScratch.spname[:])
		goseMutex.Unlock()
		return &arg0[0]
	}
	goseMutex.Unlock()
	return cCharToBytePtr(ret)
}

// Calc function as declared in swephexp.h:701
func Calc(tjd float64, ipl int, iflag int, xx []float64, serr []byte) int32 {
	goseMutex.Lock()
	pxx := doubleOutputPtr(xx, 6, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_calc(C.double(tjd), C.int(ipl), C.int32(iflag), pxx, pserr)
	copyDoubleOutput(xx, 6, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// CalcUt function as declared in swephexp.h:706
func CalcUt(tjdUt float64, ipl int, iflag int, xx []float64, serr []byte) int32 {
	goseMutex.Lock()
	pxx := doubleOutputPtr(xx, 6, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_calc_ut(C.double(tjdUt), C.int32(ipl), C.int32(iflag), pxx, pserr)
	copyDoubleOutput(xx, 6, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// CalcPctr function as declared in swephexp.h:709
func CalcPctr(tjd float64, ipl int, iplctr int, iflag int, xxret []float64, serr []byte) int32 {
	goseMutex.Lock()
	pxx := doubleOutputPtr(xxret, 6, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_calc_pctr(C.double(tjd), C.int32(ipl), C.int32(iplctr), C.int32(iflag), pxx, pserr)
	copyDoubleOutput(xxret, 6, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Solcross function as declared in swephexp.h:711
func Solcross(x2cross float64, jdEt float64, flag int32, serr []byte) float64 {
	goseMutex.Lock()
	pserr := errorBufPtr(serr)
	ret := C.swe_solcross(C.double(x2cross), C.double(jdEt), C.int32(flag), pserr)
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return float64(ret)
}

// SolcrossUt function as declared in swephexp.h:712
func SolcrossUt(x2cross float64, jdUt float64, flag int32, serr []byte) float64 {
	goseMutex.Lock()
	pserr := errorBufPtr(serr)
	ret := C.swe_solcross_ut(C.double(x2cross), C.double(jdUt), C.int32(flag), pserr)
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return float64(ret)
}

// Mooncross function as declared in swephexp.h:713
func Mooncross(x2cross float64, jdEt float64, flag int32, serr []byte) float64 {
	goseMutex.Lock()
	pserr := errorBufPtr(serr)
	ret := C.swe_mooncross(C.double(x2cross), C.double(jdEt), C.int32(flag), pserr)
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return float64(ret)
}

// MooncrossUt function as declared in swephexp.h:714
func MooncrossUt(x2cross float64, jdUt float64, flag int32, serr []byte) float64 {
	goseMutex.Lock()
	pserr := errorBufPtr(serr)
	ret := C.swe_mooncross_ut(C.double(x2cross), C.double(jdUt), C.int32(flag), pserr)
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return float64(ret)
}

// MooncrossNode function as declared in swephexp.h:715
func MooncrossNode(jdEt float64, flag int32, xlon []float64, xlat []float64, serr []byte) float64 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	goseScratch.d1 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_mooncross_node(C.double(jdEt), C.int32(flag), &goseScratch.d0, &goseScratch.d1, pserr)
	if len(xlon) > 0 {
		xlon[0] = float64(goseScratch.d0)
	}
	if len(xlat) > 0 {
		xlat[0] = float64(goseScratch.d1)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return float64(ret)
}

// MooncrossNodeUt function as declared in swephexp.h:716
func MooncrossNodeUt(jdUt float64, flag int32, xlon []float64, xlat []float64, serr []byte) float64 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	goseScratch.d1 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_mooncross_node_ut(C.double(jdUt), C.int32(flag), &goseScratch.d0, &goseScratch.d1, pserr)
	if len(xlon) > 0 {
		xlon[0] = float64(goseScratch.d0)
	}
	if len(xlat) > 0 {
		xlat[0] = float64(goseScratch.d1)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return float64(ret)
}

// HelioCross function as declared in swephexp.h:717
func HelioCross(ipl int32, x2cross float64, jdEt float64, iflag int32, dir int32, jdCross []float64, serr []byte) int32 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_helio_cross(C.int32(ipl), C.double(x2cross), C.double(jdEt), C.int32(iflag), C.int32(dir), &goseScratch.d0, pserr)
	if len(jdCross) > 0 {
		jdCross[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// HelioCrossUt function as declared in swephexp.h:718
func HelioCrossUt(ipl int32, x2cross float64, jdUt float64, iflag int32, dir int32, jdCross []float64, serr []byte) int32 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_helio_cross_ut(C.int32(ipl), C.double(x2cross), C.double(jdUt), C.int32(iflag), C.int32(dir), &goseScratch.d0, pserr)
	if len(jdCross) > 0 {
		jdCross[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Fixstar function as declared in swephexp.h:721
func Fixstar(star []byte, tjd float64, iflag int, xx []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstar, hasOutSpace := starInOutPtr(star)
	pxx := doubleOutputPtr(xx, 6, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_fixstar(cstar, C.double(tjd), C.int32(iflag), pxx, pserr)
	copyStarOut(star, hasOutSpace)
	copyDoubleOutput(xx, 6, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// FixstarUt function as declared in swephexp.h:726
func FixstarUt(star []byte, tjdUt float64, iflag int, xx []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstar, hasOutSpace := starInOutPtr(star)
	pxx := doubleOutputPtr(xx, 6, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_fixstar_ut(cstar, C.double(tjdUt), C.int32(iflag), pxx, pserr)
	copyStarOut(star, hasOutSpace)
	copyDoubleOutput(xx, 6, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// FixstarMag function as declared in swephexp.h:729
func FixstarMag(star []byte, mag []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstar, hasOutSpace := starInOutPtr(star)
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_fixstar_mag(cstar, &goseScratch.d0, pserr)
	copyStarOut(star, hasOutSpace)
	if len(mag) > 0 {
		mag[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Fixstar2 function as declared in swephexp.h:731
func Fixstar2(star []byte, tjd float64, iflag int, xx []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstar, hasOutSpace := starInOutPtr(star)
	pxx := doubleOutputPtr(xx, 6, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_fixstar2(cstar, C.double(tjd), C.int32(iflag), pxx, pserr)
	copyStarOut(star, hasOutSpace)
	copyDoubleOutput(xx, 6, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Fixstar2Ut function as declared in swephexp.h:736
func Fixstar2Ut(star []byte, tjdUt float64, iflag int, xx []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstar, hasOutSpace := starInOutPtr(star)
	pxx := doubleOutputPtr(xx, 6, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_fixstar2_ut(cstar, C.double(tjdUt), C.int32(iflag), pxx, pserr)
	copyStarOut(star, hasOutSpace)
	copyDoubleOutput(xx, 6, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Fixstar2Mag function as declared in swephexp.h:739
func Fixstar2Mag(star []byte, mag []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstar, hasOutSpace := starInOutPtr(star)
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_fixstar2_mag(cstar, &goseScratch.d0, pserr)
	copyStarOut(star, hasOutSpace)
	if len(mag) > 0 {
		mag[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Close function as declared in swephexp.h:742
func Close() {
	goseMutex.Lock()
	C.swe_close()
	goseMutex.Unlock()
}

// SetEphePath function as declared in swephexp.h:745
func SetEphePath(path []byte) {
	goseMutex.Lock()
	cpath := cStringInputPtr(path)
	C.swe_set_ephe_path(cpath)
	goseMutex.Unlock()
}

// SetJplFile function as declared in swephexp.h:748
func SetJplFile(fname []byte) {
	goseMutex.Lock()
	cfname := cStringInputPtr(fname)
	C.swe_set_jpl_file(cfname)
	goseMutex.Unlock()
}

// GetPlanetName function as declared in swephexp.h:751
func GetPlanetName(ipl int, spname []byte) *byte {
	goseMutex.Lock()
	if len(spname) >= 256 {
		ret := C.swe_get_planet_name(C.int(ipl), (*C.char)(unsafe.Pointer(unsafe.SliceData(spname))))
		goseMutex.Unlock()
		return cCharToBytePtr(ret)
	}
	goseScratch.spname[0] = 0
	ret := C.swe_get_planet_name(C.int(ipl), (*C.char)(unsafe.Pointer(&goseScratch.spname[0])))
	if len(spname) > 0 {
		copy(spname, goseScratch.spname[:])
		goseMutex.Unlock()
		return &spname[0]
	}
	goseMutex.Unlock()
	return cCharToBytePtr(ret)
}

// SetTopo function as declared in swephexp.h:754
func SetTopo(geolon float64, geolat float64, geoalt float64) {
	goseMutex.Lock()
	C.swe_set_topo(C.double(geolon), C.double(geolat), C.double(geoalt))
	goseMutex.Unlock()
}

// SetSidMode function as declared in swephexp.h:757
func SetSidMode(sidMode int, t0 float64, ayanT0 float64) {
	goseMutex.Lock()
	C.swe_set_sid_mode(C.int32(sidMode), C.double(t0), C.double(ayanT0))
	goseMutex.Unlock()
}

// GetAyanamsaEx function as declared in swephexp.h:760
func GetAyanamsaEx(tjdEt float64, iflag int, daya []float64, serr []byte) int32 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_get_ayanamsa_ex(C.double(tjdEt), C.int32(iflag), &goseScratch.d0, pserr)
	if len(daya) > 0 {
		daya[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// GetAyanamsaExUt function as declared in swephexp.h:761
func GetAyanamsaExUt(tjdUt float64, iflag int, daya []float64, serr []byte) int32 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_get_ayanamsa_ex_ut(C.double(tjdUt), C.int32(iflag), &goseScratch.d0, pserr)
	if len(daya) > 0 {
		daya[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// GetAyanamsa function as declared in swephexp.h:762
func GetAyanamsa(tjdEt float64) float64 {
	goseMutex.Lock()
	ret := C.swe_get_ayanamsa(C.double(tjdEt))
	goseMutex.Unlock()
	return float64(ret)
}

// GetAyanamsaUt function as declared in swephexp.h:763
func GetAyanamsaUt(tjdUt float64) float64 {
	goseMutex.Lock()
	ret := C.swe_get_ayanamsa_ut(C.double(tjdUt))
	goseMutex.Unlock()
	return float64(ret)
}

// GetAyanamsaName function as declared in swephexp.h:766
func GetAyanamsaName(isidmode int32) string {
	if isidmode < 0 {
		return ""
	}
	ret := C.swe_get_ayanamsa_name(C.int32(isidmode))
	return cCharToString(ret)
}

// GetCurrentFileData function as declared in swephexp.h:767
func GetCurrentFileData(ifno int, tfstart []float64, tfend []float64, denum []int32) string {
	goseMutex.Lock()
	goseScratch.d0 = 0
	goseScratch.d1 = 0
	goseScratch.ci0 = 0
	ret := C.swe_get_current_file_data(C.int(ifno), &goseScratch.d0, &goseScratch.d1, &goseScratch.ci0)
	if len(tfstart) > 0 {
		tfstart[0] = float64(goseScratch.d0)
	}
	if len(tfend) > 0 {
		tfend[0] = float64(goseScratch.d1)
	}
	if len(denum) > 0 {
		denum[0] = int32(goseScratch.ci0)
	}
	s := cCharToString(ret)
	goseMutex.Unlock()
	return s
}

// DateConversion function as declared in swephexp.h:775
func DateConversion(y int, m int, d int, utime float64, c byte, tjd []float64) int32 {
	gregflag := int32(SeJulCal)
	if c == 'g' {
		gregflag = SeGregCal
	}
	jd := Julday(y, m, d, utime, gregflag)
	ryear, rmon, rday, _ := revjulCore(jd, int(gregflag))
	if len(tjd) > 0 {
		tjd[0] = jd
	}
	if rmon == m && rday == d && ryear == y {
		return Ok
	}
	return Err
}

// Julday function as declared in swephexp.h:781
func Julday(year int, month int, day int, hour float64, gregflag int32) float64 {
	u := float64(year)
	if month < 3 {
		u -= 1.0
	}
	u0 := u + 4712.0
	u1 := float64(month) + 1.0
	if u1 < 4.0 {
		u1 += 12.0
	}
	jd := math.Floor(u0*365.25) + math.Floor(30.6*u1+0.000001) + float64(day) + hour/24.0 - 63.5
	if gregflag == SeGregCal {
		u2 := math.Floor(math.Abs(u)/100.0) - math.Floor(math.Abs(u)/400.0)
		if u < 0.0 {
			u2 = -u2
		}
		jd = jd - u2 + 2.0
		if (u < 0.0) && (u/100.0 == math.Floor(u/100.0)) && (u/400.0 != math.Floor(u/400.0)) {
			jd -= 1.0
		}
	}
	return jd
}

func revjulCore(jd float64, gregflag int) (jyear, jmon, jday int, jut float64) {
	u0 := jd + 32082.5
	if gregflag == SeGregCal {
		u1 := u0 + math.Floor(u0/36525.0) - math.Floor(u0/146100.0) - 38.0
		if jd >= 1830691.5 {
			u1 += 1.0
		}
		u0 = u0 + math.Floor(u1/36525.0) - math.Floor(u1/146100.0) - 38.0
	}
	u2 := math.Floor(u0 + 123.0)
	u3 := math.Floor((u2 - 122.2) / 365.25)
	u4 := math.Floor((u2 - math.Floor(365.25*u3)) / 30.6001)
	jmon = int(u4 - 1.0)
	if jmon > 12 {
		jmon -= 12
	}
	jday = int(u2 - math.Floor(365.25*u3) - math.Floor(30.6001*u4))
	jyear = int(u3 + math.Floor((u4-2.0)/12.0) - 4800.0)
	jut = (jd - math.Floor(jd+0.5) + 0.5) * 24.0
	return jyear, jmon, jday, jut
}

// Revjul function as declared in swephexp.h:785
func Revjul(jd float64, gregflag int, jyear []int, jmon []int, jday []int, jut []float64) {
	y, m, d, ut := revjulCore(jd, gregflag)
	if len(jyear) > 0 {
		jyear[0] = y
	}
	if len(jmon) > 0 {
		jmon[0] = m
	}
	if len(jday) > 0 {
		jday[0] = d
	}
	if len(jut) > 0 {
		jut[0] = ut
	}
}

// UtcToJd function as declared in swephexp.h:790
func UtcToJd(iyear int, imonth int, iday int, ihour int, imin int, dsec float64, gregflag int, dret []float64, serr []byte) int32 {
	goseMutex.Lock()
	pdret := doubleOutputPtr(dret, 2, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_utc_to_jd(C.int32(iyear), C.int32(imonth), C.int32(iday), C.int32(ihour), C.int32(imin), C.double(dsec), C.int32(gregflag), pdret, pserr)
	copyDoubleOutput(dret, 2, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// JdetToUtc function as declared in swephexp.h:795
func JdetToUtc(tjdEt float64, gregflag int, iyear []int, imonth []int, iday []int, ihour []int, imin []int, dsec []float64) {
	goseMutex.Lock()
	goseScratch.i0, goseScratch.i1, goseScratch.i2, goseScratch.i3, goseScratch.i4 = 0, 0, 0, 0, 0
	goseScratch.d0 = 0
	C.swe_jdet_to_utc(C.double(tjdEt), C.int32(gregflag), &goseScratch.i0, &goseScratch.i1, &goseScratch.i2, &goseScratch.i3, &goseScratch.i4, &goseScratch.d0)
	if len(iyear) > 0 {
		iyear[0] = int(goseScratch.i0)
	}
	if len(imonth) > 0 {
		imonth[0] = int(goseScratch.i1)
	}
	if len(iday) > 0 {
		iday[0] = int(goseScratch.i2)
	}
	if len(ihour) > 0 {
		ihour[0] = int(goseScratch.i3)
	}
	if len(imin) > 0 {
		imin[0] = int(goseScratch.i4)
	}
	if len(dsec) > 0 {
		dsec[0] = float64(goseScratch.d0)
	}
	goseMutex.Unlock()
}

// Jdut1ToUtc function as declared in swephexp.h:800
func Jdut1ToUtc(tjdUt float64, gregflag int, iyear []int, imonth []int, iday []int, ihour []int, imin []int, dsec []float64) {
	goseMutex.Lock()
	goseScratch.i0, goseScratch.i1, goseScratch.i2, goseScratch.i3, goseScratch.i4 = 0, 0, 0, 0, 0
	goseScratch.d0 = 0
	C.swe_jdut1_to_utc(C.double(tjdUt), C.int32(gregflag), &goseScratch.i0, &goseScratch.i1, &goseScratch.i2, &goseScratch.i3, &goseScratch.i4, &goseScratch.d0)
	if len(iyear) > 0 {
		iyear[0] = int(goseScratch.i0)
	}
	if len(imonth) > 0 {
		imonth[0] = int(goseScratch.i1)
	}
	if len(iday) > 0 {
		iday[0] = int(goseScratch.i2)
	}
	if len(ihour) > 0 {
		ihour[0] = int(goseScratch.i3)
	}
	if len(imin) > 0 {
		imin[0] = int(goseScratch.i4)
	}
	if len(dsec) > 0 {
		dsec[0] = float64(goseScratch.d0)
	}
	goseMutex.Unlock()
}

// UtcTimeZone function as declared in swephexp.h:805
func UtcTimeZone(iyear int, imonth int, iday int, ihour int, imin int, dsec float64, dTimezone float64, iyearOut []int, imonthOut []int, idayOut []int, ihourOut []int, iminOut []int, dsecOut []float64) {
	haveLeapsec := false
	if dsec >= 60.0 {
		haveLeapsec = true
		dsec -= 1.0
	}
	dhour := float64(ihour) + float64(imin)/60.0 + dsec/3600.0
	tjd := Julday(iyear, imonth, iday, 0, SeGregCal)
	dhour -= dTimezone
	if dhour < 0.0 {
		tjd -= 1.0
		dhour += 24.0
	}
	if dhour >= 24.0 {
		tjd += 1.0
		dhour -= 24.0
	}
	yOut, mOut, dOut, _ := revjulCore(tjd+0.001, SeGregCal)
	hOut := int(dhour)
	d := (dhour - float64(hOut)) * 60.0
	minOut := int(d)
	secOut := (d - float64(minOut)) * 60.0
	if haveLeapsec {
		secOut += 1.0
	}
	if len(iyearOut) > 0 {
		iyearOut[0] = yOut
	}
	if len(imonthOut) > 0 {
		imonthOut[0] = mOut
	}
	if len(idayOut) > 0 {
		idayOut[0] = dOut
	}
	if len(ihourOut) > 0 {
		ihourOut[0] = hOut
	}
	if len(iminOut) > 0 {
		iminOut[0] = minOut
	}
	if len(dsecOut) > 0 {
		dsecOut[0] = secOut
	}
}

// Houses function as declared in swephexp.h:816
func Houses(tjdUt float64, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64) int32 {
	minCusps := houseCuspsMinLen(hsys)
	goseMutex.Lock()
	pcusps := doubleOutputPtr(cusps, minCusps, goseScratch.out1[:])
	pascmc := doubleInOutPtr(ascmc, 10, goseScratch.out2[:])
	ret := C.swe_houses(C.double(tjdUt), C.double(geolat), C.double(geolon), C.int(hsys), pcusps, pascmc)
	copyDoubleOutput(cusps, minCusps, goseScratch.out1[:])
	copyDoubleOutput(ascmc, 10, goseScratch.out2[:])
	goseMutex.Unlock()
	return int32(ret)
}

// HousesEx function as declared in swephexp.h:820
func HousesEx(tjdUt float64, iflag int, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64) int32 {
	minCusps := houseCuspsMinLen(hsys)
	goseMutex.Lock()
	pcusps := doubleOutputPtr(cusps, minCusps, goseScratch.out1[:])
	pascmc := doubleInOutPtr(ascmc, 10, goseScratch.out2[:])
	ret := C.swe_houses_ex(C.double(tjdUt), C.int32(iflag), C.double(geolat), C.double(geolon), C.int(hsys), pcusps, pascmc)
	copyDoubleOutput(cusps, minCusps, goseScratch.out1[:])
	copyDoubleOutput(ascmc, 10, goseScratch.out2[:])
	goseMutex.Unlock()
	return int32(ret)
}

// HousesEx2 function as declared in swephexp.h:824
func HousesEx2(tjdUt float64, iflag int, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64, cuspSpeed []float64, ascmcSpeed []float64, serr []byte) int32 {
	minCusps := houseCuspsMinLen(hsys)
	goseMutex.Lock()
	pcusps := doubleOutputPtr(cusps, minCusps, goseScratch.out1[:])
	pascmc := doubleInOutPtr(ascmc, 10, goseScratch.out2[:])
	var pcuspSpeed *C.double
	if len(cuspSpeed) > 0 {
		pcuspSpeed = doubleOutputPtr(cuspSpeed, minCusps, goseScratch.out3[:])
	}
	var pascmcSpeed *C.double
	if len(ascmcSpeed) > 0 {
		pascmcSpeed = doubleOutputPtr(ascmcSpeed, 10, goseScratch.out4[:])
	}
	pserr := errorBufPtr(serr)
	ret := C.swe_houses_ex2(C.double(tjdUt), C.int32(iflag), C.double(geolat), C.double(geolon), C.int(hsys), pcusps, pascmc, pcuspSpeed, pascmcSpeed, pserr)
	copyDoubleOutput(cusps, minCusps, goseScratch.out1[:])
	copyDoubleOutput(ascmc, 10, goseScratch.out2[:])
	copyDoubleOutput(cuspSpeed, minCusps, goseScratch.out3[:])
	copyDoubleOutput(ascmcSpeed, 10, goseScratch.out4[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// HousesArmc function as declared in swephexp.h:828
func HousesArmc(armc float64, geolat float64, eps float64, hsys int, cusps []float64, ascmc []float64) int32 {
	minCusps := houseCuspsMinLen(hsys)
	goseMutex.Lock()
	pcusps := doubleOutputPtr(cusps, minCusps, goseScratch.out1[:])
	pascmc := doubleInOutPtr(ascmc, 10, goseScratch.out2[:])
	ret := C.swe_houses_armc(C.double(armc), C.double(geolat), C.double(eps), C.int(hsys), pcusps, pascmc)
	copyDoubleOutput(cusps, minCusps, goseScratch.out1[:])
	copyDoubleOutput(ascmc, 10, goseScratch.out2[:])
	goseMutex.Unlock()
	return int32(ret)
}

// HousesArmcEx2 function as declared in swephexp.h:832
func HousesArmcEx2(armc float64, geolat float64, eps float64, hsys int, cusps []float64, ascmc []float64, cuspSpeed []float64, ascmcSpeed []float64, serr []byte) int32 {
	minCusps := houseCuspsMinLen(hsys)
	goseMutex.Lock()
	pcusps := doubleOutputPtr(cusps, minCusps, goseScratch.out1[:])
	pascmc := doubleInOutPtr(ascmc, 10, goseScratch.out2[:])
	var pcuspSpeed *C.double
	if len(cuspSpeed) > 0 {
		pcuspSpeed = doubleOutputPtr(cuspSpeed, minCusps, goseScratch.out3[:])
	}
	var pascmcSpeed *C.double
	if len(ascmcSpeed) > 0 {
		pascmcSpeed = doubleOutputPtr(ascmcSpeed, 10, goseScratch.out4[:])
	}
	pserr := errorBufPtr(serr)
	ret := C.swe_houses_armc_ex2(C.double(armc), C.double(geolat), C.double(eps), C.int(hsys), pcusps, pascmc, pcuspSpeed, pascmcSpeed, pserr)
	copyDoubleOutput(cusps, minCusps, goseScratch.out1[:])
	copyDoubleOutput(ascmc, 10, goseScratch.out2[:])
	copyDoubleOutput(cuspSpeed, minCusps, goseScratch.out3[:])
	copyDoubleOutput(ascmcSpeed, 10, goseScratch.out4[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// HousePos function as declared in swephexp.h:836
func HousePos(armc float64, geolat float64, eps float64, hsys int, xpin []float64, serr []byte) float64 {
	goseMutex.Lock()
	pxpin := doubleInputPtr(xpin, 2, goseScratch.in1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_house_pos(C.double(armc), C.double(geolat), C.double(eps), C.int(hsys), pxpin, pserr)
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return float64(ret)
}

// HouseName function as declared in swephexp.h:839
func HouseName(hsys int32) *byte {
	ret := C.swe_house_name(C.int(hsys))
	return cCharToBytePtr(ret)
}

// GauquelinSector function as declared in swephexp.h:847
func GauquelinSector(tUt float64, ipl int, starname []byte, iflag int, imeth int, geopos []float64, atpress float64, attemp float64, dgsect []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstarname, hasOutSpace := optionalStarInOutPtr(starname)
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_gauquelin_sector(C.double(tUt), C.int32(ipl), cstarname, C.int32(iflag), C.int32(imeth), cgeopos, C.double(atpress), C.double(attemp), &goseScratch.d0, pserr)
	copyStarOut(starname, hasOutSpace)
	if len(dgsect) > 0 {
		dgsect[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// SolEclipseWhere function as declared in swephexp.h:851
func SolEclipseWhere(tjd float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	goseMutex.Lock()
	cgeopos := doubleOutputPtr(geopos, 2, goseScratch.out1[:])
	cattr := doubleOutputPtr(attr, 20, goseScratch.out2[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_sol_eclipse_where(C.double(tjd), C.int32(ifl), cgeopos, cattr, pserr)
	copyDoubleOutput(geopos, 2, goseScratch.out1[:])
	copyDoubleOutput(attr, 20, goseScratch.out2[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunOccultWhere function as declared in swephexp.h:853
func LunOccultWhere(tjd float64, ipl int, starname []byte, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstarname, hasOutSpace := optionalStarInOutPtr(starname)
	cgeopos := doubleOutputPtr(geopos, 2, goseScratch.out1[:])
	cattr := doubleOutputPtr(attr, 20, goseScratch.out2[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_lun_occult_where(C.double(tjd), C.int32(ipl), cstarname, C.int32(ifl), cgeopos, cattr, pserr)
	copyStarOut(starname, hasOutSpace)
	copyDoubleOutput(geopos, 2, goseScratch.out1[:])
	copyDoubleOutput(attr, 20, goseScratch.out2[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// SolEclipseHow function as declared in swephexp.h:856
func SolEclipseHow(tjd float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	cattr := doubleOutputPtr(attr, 20, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_sol_eclipse_how(C.double(tjd), C.int32(ifl), cgeopos, cattr, pserr)
	copyDoubleOutput(attr, 20, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// SolEclipseWhenLoc function as declared in swephexp.h:859
func SolEclipseWhenLoc(tjdStart float64, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	ctret := doubleOutputPtr(tret, 10, goseScratch.out1[:])
	cattr := doubleOutputPtr(attr, 20, goseScratch.out2[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_sol_eclipse_when_loc(C.double(tjdStart), C.int32(ifl), cgeopos, ctret, cattr, C.int32(backward), pserr)
	copyDoubleOutput(tret, 10, goseScratch.out1[:])
	copyDoubleOutput(attr, 20, goseScratch.out2[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunOccultWhenLoc function as declared in swephexp.h:861
func LunOccultWhenLoc(tjdStart float64, ipl int, starname []byte, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	goseMutex.Lock()
	cstarname, hasOutSpace := optionalStarInOutPtr(starname)
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	ctret := doubleOutputPtr(tret, 10, goseScratch.out1[:])
	cattr := doubleOutputPtr(attr, 20, goseScratch.out2[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_lun_occult_when_loc(C.double(tjdStart), C.int32(ipl), cstarname, C.int32(ifl), cgeopos, ctret, cattr, C.int32(backward), pserr)
	copyStarOut(starname, hasOutSpace)
	copyDoubleOutput(tret, 10, goseScratch.out1[:])
	copyDoubleOutput(attr, 20, goseScratch.out2[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// SolEclipseWhenGlob function as declared in swephexp.h:865
func SolEclipseWhenGlob(tjdStart float64, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	goseMutex.Lock()
	ctret := doubleOutputPtr(tret, 10, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_sol_eclipse_when_glob(C.double(tjdStart), C.int32(ifl), C.int32(ifltype), ctret, C.int32(backward), pserr)
	copyDoubleOutput(tret, 10, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunOccultWhenGlob function as declared in swephexp.h:869
func LunOccultWhenGlob(tjdStart float64, ipl int, starname []byte, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	goseMutex.Lock()
	cstarname, hasOutSpace := optionalStarInOutPtr(starname)
	ctret := doubleOutputPtr(tret, 10, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_lun_occult_when_glob(C.double(tjdStart), C.int32(ipl), cstarname, C.int32(ifl), C.int32(ifltype), ctret, C.int32(backward), pserr)
	copyStarOut(starname, hasOutSpace)
	copyDoubleOutput(tret, 10, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunEclipseHow function as declared in swephexp.h:873
func LunEclipseHow(tjdUt float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	cattr := doubleOutputPtr(attr, 20, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_lun_eclipse_how(C.double(tjdUt), C.int32(ifl), cgeopos, cattr, pserr)
	copyDoubleOutput(attr, 20, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunEclipseWhen function as declared in swephexp.h:880
func LunEclipseWhen(tjdStart float64, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	goseMutex.Lock()
	ctret := doubleOutputPtr(tret, 10, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_lun_eclipse_when(C.double(tjdStart), C.int32(ifl), C.int32(ifltype), ctret, C.int32(backward), pserr)
	copyDoubleOutput(tret, 10, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunEclipseWhenLoc function as declared in swephexp.h:883
func LunEclipseWhenLoc(tjdStart float64, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	ctret := doubleOutputPtr(tret, 10, goseScratch.out1[:])
	cattr := doubleOutputPtr(attr, 20, goseScratch.out2[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_lun_eclipse_when_loc(C.double(tjdStart), C.int32(ifl), cgeopos, ctret, cattr, C.int32(backward), pserr)
	copyDoubleOutput(tret, 10, goseScratch.out1[:])
	copyDoubleOutput(attr, 20, goseScratch.out2[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Pheno function as declared in swephexp.h:887
func Pheno(tjd float64, ipl int, iflag int, attr []float64, serr []byte) int32 {
	goseMutex.Lock()
	cattr := doubleOutputPtr(attr, 20, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_pheno(C.double(tjd), C.int32(ipl), C.int32(iflag), cattr, pserr)
	copyDoubleOutput(attr, 20, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// PhenoUt function as declared in swephexp.h:889
func PhenoUt(tjdUt float64, ipl int, iflag int, attr []float64, serr []byte) int32 {
	goseMutex.Lock()
	cattr := doubleOutputPtr(attr, 20, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_pheno_ut(C.double(tjdUt), C.int32(ipl), C.int32(iflag), cattr, pserr)
	copyDoubleOutput(attr, 20, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Refrac function as declared in swephexp.h:891
func Refrac(inalt float64, atpress float64, attemp float64, calcFlag int32) float64 {
	ret := C.swe_refrac(C.double(inalt), C.double(atpress), C.double(attemp), C.int32(calcFlag))
	return float64(ret)
}

// RefracExtended function as declared in swephexp.h:893
func RefracExtended(inalt float64, geoalt float64, atpress float64, attemp float64, lapseRate float64, calcFlag int, dret []float64) float64 {
	if len(dret) == 0 {
		ret := C.swe_refrac_extended(C.double(inalt), C.double(geoalt), C.double(atpress), C.double(attemp), C.double(lapseRate), C.int32(calcFlag), nil)
		return float64(ret)
	}
	if len(dret) >= 4 {
		cdret := (*C.double)(unsafe.Pointer(unsafe.SliceData(dret)))
		ret := C.swe_refrac_extended(C.double(inalt), C.double(geoalt), C.double(atpress), C.double(attemp), C.double(lapseRate), C.int32(calcFlag), cdret)
		return float64(ret)
	}
	goseMutex.Lock()
	cdret := doubleOutputPtr(dret, 4, goseScratch.out1[:])
	ret := C.swe_refrac_extended(C.double(inalt), C.double(geoalt), C.double(atpress), C.double(attemp), C.double(lapseRate), C.int32(calcFlag), cdret)
	copyDoubleOutput(dret, 4, goseScratch.out1[:])
	goseMutex.Unlock()
	return float64(ret)
}

// SetLapseRate function as declared in swephexp.h:895
func SetLapseRate(lapseRate float64) {
	goseMutex.Lock()
	C.swe_set_lapse_rate(C.double(lapseRate))
	goseMutex.Unlock()
}

// Azalt function as declared in swephexp.h:897
func Azalt(tjdUt float64, calcFlag int, geopos []float64, atpress float64, attemp float64, xin []float64, xaz []float64) {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	cxin := doubleInputPtr(xin, 3, goseScratch.in2[:])
	cxaz := doubleOutputPtr(xaz, 3, goseScratch.out1[:])
	C.swe_azalt(C.double(tjdUt), C.int32(calcFlag), cgeopos, C.double(atpress), C.double(attemp), cxin, cxaz)
	copyDoubleOutput(xaz, 3, goseScratch.out1[:])
	goseMutex.Unlock()
}

// AzaltRev function as declared in swephexp.h:906
func AzaltRev(tjdUt float64, calcFlag int, geopos []float64, xin []float64, xout []float64) {
	goseMutex.Lock()
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	cxin := doubleInputPtr(xin, 3, goseScratch.in2[:])
	cxout := doubleOutputPtr(xout, 3, goseScratch.out1[:])
	C.swe_azalt_rev(C.double(tjdUt), C.int32(calcFlag), cgeopos, cxin, cxout)
	copyDoubleOutput(xout, 3, goseScratch.out1[:])
	goseMutex.Unlock()
}

// RiseTransTrueHor function as declared in swephexp.h:913
func RiseTransTrueHor(tjdUt float64, ipl int, starname []byte, epheflag int, rsmi int, geopos []float64, atpress float64, attemp float64, horhgt float64, tret []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstarname, hasOutSpace := optionalStarInOutPtr(starname)
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	ctret := doubleOutputPtr(tret, 1, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_rise_trans_true_hor(C.double(tjdUt), C.int32(ipl), cstarname, C.int32(epheflag), C.int32(rsmi), cgeopos, C.double(atpress), C.double(attemp), C.double(horhgt), ctret, pserr)
	copyStarOut(starname, hasOutSpace)
	copyDoubleOutput(tret, 1, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// RiseTrans function as declared in swephexp.h:922
func RiseTrans(tjdUt float64, ipl int, starname []byte, epheflag int, rsmi int, geopos []float64, atpress float64, attemp float64, tret []float64, serr []byte) int32 {
	goseMutex.Lock()
	cstarname, hasOutSpace := optionalStarInOutPtr(starname)
	cgeopos := doubleInputPtr(geopos, 3, goseScratch.in1[:])
	ctret := doubleOutputPtr(tret, 1, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_rise_trans(C.double(tjdUt), C.int32(ipl), cstarname, C.int32(epheflag), C.int32(rsmi), cgeopos, C.double(atpress), C.double(attemp), ctret, pserr)
	copyStarOut(starname, hasOutSpace)
	copyDoubleOutput(tret, 1, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// NodAps function as declared in swephexp.h:930
func NodAps(tjdEt float64, ipl int, iflag int, method int, xnasc []float64, xndsc []float64, xperi []float64, xaphe []float64, serr []byte) int32 {
	goseMutex.Lock()
	cxnasc := doubleOutputPtr(xnasc, 6, goseScratch.out1[:])
	cxndsc := doubleOutputPtr(xndsc, 6, goseScratch.out2[:])
	cxperi := doubleOutputPtr(xperi, 6, goseScratch.out3[:])
	cxaphe := doubleOutputPtr(xaphe, 6, goseScratch.out4[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_nod_aps(C.double(tjdEt), C.int32(ipl), C.int32(iflag), C.int32(method), cxnasc, cxndsc, cxperi, cxaphe, pserr)
	copyDoubleOutput(xnasc, 6, goseScratch.out1[:])
	copyDoubleOutput(xndsc, 6, goseScratch.out2[:])
	copyDoubleOutput(xperi, 6, goseScratch.out3[:])
	copyDoubleOutput(xaphe, 6, goseScratch.out4[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// NodApsUt function as declared in swephexp.h:936
func NodApsUt(tjdUt float64, ipl int, iflag int, method int, xnasc []float64, xndsc []float64, xperi []float64, xaphe []float64, serr []byte) int32 {
	goseMutex.Lock()
	cxnasc := doubleOutputPtr(xnasc, 6, goseScratch.out1[:])
	cxndsc := doubleOutputPtr(xndsc, 6, goseScratch.out2[:])
	cxperi := doubleOutputPtr(xperi, 6, goseScratch.out3[:])
	cxaphe := doubleOutputPtr(xaphe, 6, goseScratch.out4[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_nod_aps_ut(C.double(tjdUt), C.int32(ipl), C.int32(iflag), C.int32(method), cxnasc, cxndsc, cxperi, cxaphe, pserr)
	copyDoubleOutput(xnasc, 6, goseScratch.out1[:])
	copyDoubleOutput(xndsc, 6, goseScratch.out2[:])
	copyDoubleOutput(xperi, 6, goseScratch.out3[:])
	copyDoubleOutput(xaphe, 6, goseScratch.out4[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// GetOrbitalElements function as declared in swephexp.h:941
func GetOrbitalElements(tjdEt float64, ipl int, iflag int, dret []float64, serr []byte) int32 {
	goseMutex.Lock()
	cdret := doubleOutputPtr(dret, 50, goseScratch.out1[:])
	pserr := errorBufPtr(serr)
	ret := C.swe_get_orbital_elements(C.double(tjdEt), C.int32(ipl), C.int32(iflag), cdret, pserr)
	copyDoubleOutput(dret, 50, goseScratch.out1[:])
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// OrbitMaxMinTrueDistance function as declared in swephexp.h:944
func OrbitMaxMinTrueDistance(tjdEt float64, ipl int, iflag int, dmax []float64, dmin []float64, dtrue []float64, serr []byte) int32 {
	goseMutex.Lock()
	goseScratch.d0, goseScratch.d1, goseScratch.d2 = 0, 0, 0
	pserr := errorBufPtr(serr)
	ret := C.swe_orbit_max_min_true_distance(C.double(tjdEt), C.int32(ipl), C.int32(iflag), &goseScratch.d0, &goseScratch.d1, &goseScratch.d2, pserr)
	if len(dmax) > 0 {
		dmax[0] = float64(goseScratch.d0)
	}
	if len(dmin) > 0 {
		dmin[0] = float64(goseScratch.d1)
	}
	if len(dtrue) > 0 {
		dtrue[0] = float64(goseScratch.d2)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Deltat function as declared in swephexp.h:951
func Deltat(tjd float64) float64 {
	goseMutex.Lock()
	ret := C.swe_deltat(C.double(tjd))
	goseMutex.Unlock()
	return float64(ret)
}

// DeltatEx function as declared in swephexp.h:952
func DeltatEx(tjd float64, iflag int, serr []byte) float64 {
	goseMutex.Lock()
	pserr := errorBufPtr(serr)
	ret := C.swe_deltat_ex(C.double(tjd), C.int32(iflag), pserr)
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return float64(ret)
}

// TimeEqu function as declared in swephexp.h:955
func TimeEqu(tjd float64, te []float64, serr []byte) int32 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_time_equ(C.double(tjd), &goseScratch.d0, pserr)
	if len(te) > 0 {
		te[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// LmtToLat function as declared in swephexp.h:956
func LmtToLat(tjdLmt float64, geolon float64, tjdLat []float64, serr []byte) int32 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_lmt_to_lat(C.double(tjdLmt), C.double(geolon), &goseScratch.d0, pserr)
	if len(tjdLat) > 0 {
		tjdLat[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// LatToLmt function as declared in swephexp.h:957
func LatToLmt(tjdLat float64, geolon float64, tjdLmt []float64, serr []byte) int32 {
	goseMutex.Lock()
	goseScratch.d0 = 0
	pserr := errorBufPtr(serr)
	ret := C.swe_lat_to_lmt(C.double(tjdLat), C.double(geolon), &goseScratch.d0, pserr)
	if len(tjdLmt) > 0 {
		tjdLmt[0] = float64(goseScratch.d0)
	}
	copyErrorBuf(serr)
	goseMutex.Unlock()
	return int32(ret)
}

// Sidtime0 function as declared in swephexp.h:960
func Sidtime0(tjdUt float64, eps float64, nut float64) float64 {
	goseMutex.Lock()
	ret := C.swe_sidtime0(C.double(tjdUt), C.double(eps), C.double(nut))
	goseMutex.Unlock()
	return float64(ret)
}

// Sidtime function as declared in swephexp.h:961
func Sidtime(tjdUt float64) float64 {
	goseMutex.Lock()
	ret := C.swe_sidtime(C.double(tjdUt))
	goseMutex.Unlock()
	return float64(ret)
}

// SetInterpolateNut function as declared in swephexp.h:962
func SetInterpolateNut(doInterpolate int32) {
	goseMutex.Lock()
	C.swe_set_interpolate_nut(C.AS_BOOL(doInterpolate))
	goseMutex.Unlock()
}

// Cotrans function as declared in swephexp.h:965
func Cotrans(xpo []float64, xpn []float64, eps float64) {
	if len(xpo) >= 3 && len(xpn) >= 3 {
		C.swe_cotrans((*C.double)(unsafe.Pointer(unsafe.SliceData(xpo))), (*C.double)(unsafe.Pointer(unsafe.SliceData(xpn))), C.double(eps))
		return
	}
	goseMutex.Lock()
	cxpo := doubleInputPtr(xpo, 3, goseScratch.in1[:])
	cxpn := doubleOutputPtr(xpn, 3, goseScratch.out1[:])
	C.swe_cotrans(cxpo, cxpn, C.double(eps))
	copyDoubleOutput(xpn, 3, goseScratch.out1[:])
	goseMutex.Unlock()
}

// CotransSp function as declared in swephexp.h:966
func CotransSp(xpo []float64, xpn []float64, eps float64) {
	if len(xpo) >= 6 && len(xpn) >= 6 {
		C.swe_cotrans_sp((*C.double)(unsafe.Pointer(unsafe.SliceData(xpo))), (*C.double)(unsafe.Pointer(unsafe.SliceData(xpn))), C.double(eps))
		return
	}
	goseMutex.Lock()
	cxpo := doubleInputPtr(xpo, 6, goseScratch.in1[:])
	cxpn := doubleOutputPtr(xpn, 6, goseScratch.out1[:])
	C.swe_cotrans_sp(cxpo, cxpn, C.double(eps))
	copyDoubleOutput(xpn, 6, goseScratch.out1[:])
	goseMutex.Unlock()
}

// GetTidAcc function as declared in swephexp.h:969
func GetTidAcc() float64 {
	goseMutex.Lock()
	ret := C.swe_get_tid_acc()
	goseMutex.Unlock()
	return float64(ret)
}

// SetTidAcc function as declared in swephexp.h:970
func SetTidAcc(tAcc float64) {
	goseMutex.Lock()
	C.swe_set_tid_acc(C.double(tAcc))
	goseMutex.Unlock()
}

// SetDeltaTUserdef function as declared in swephexp.h:974
func SetDeltaTUserdef(dt float64) {
	goseMutex.Lock()
	C.swe_set_delta_t_userdef(C.double(dt))
	goseMutex.Unlock()
}

// Degnorm function as declared in swephexp.h:976
func Degnorm(x float64) float64 {
	y := math.Mod(x, 360.0)
	if math.Abs(y) < 1e-13 {
		y = 0
	}
	if y < 0.0 {
		y += 360.0
	}
	return y
}

// Radnorm function as declared in swephexp.h:977
func Radnorm(x float64) float64 {
	y := math.Mod(x, Twopi)
	if math.Abs(y) < 1e-13 {
		y = 0
	}
	if y < 0.0 {
		y += Twopi
	}
	return y
}

// RadMidp function as declared in swephexp.h:978
func RadMidp(x1 float64, x0 float64) float64 {
	return Degtorad * DegMidp(x1*Radtodeg, x0*Radtodeg)
}

// DegMidp function as declared in swephexp.h:979
func DegMidp(x1 float64, x0 float64) float64 {
	d := Difdeg2n(x1, x0)
	return Degnorm(x0 + d/2.0)
}

// SplitDeg function as declared in swephexp.h:981
func SplitDeg(ddeg float64, roundflag int, ideg []int, imin []int, isec []int, dsecfr []float64, isgn []int32) {
	if ddeg >= 0 && (roundflag&SeSplitDegNakshatra) != 0 {
		goseMutex.Lock()
		goseScratch.i0, goseScratch.i1, goseScratch.i2, goseScratch.i3 = 0, 0, 0, 0
		goseScratch.d0 = 0
		C.swe_split_deg(C.double(ddeg), C.int32(roundflag), &goseScratch.i0, &goseScratch.i1, &goseScratch.i2, &goseScratch.d0, &goseScratch.i3)
		if len(ideg) > 0 {
			ideg[0] = int(goseScratch.i0)
		}
		if len(imin) > 0 {
			imin[0] = int(goseScratch.i1)
		}
		if len(isec) > 0 {
			isec[0] = int(goseScratch.i2)
		}
		if len(dsecfr) > 0 {
			dsecfr[0] = float64(goseScratch.d0)
		}
		if len(isgn) > 0 {
			isgn[0] = int32(goseScratch.i3)
		}
		goseMutex.Unlock()
		return
	}

	var dadd float64
	sgn := int32(1)
	if ddeg < 0 {
		sgn = -1
		ddeg = -ddeg
	}
	if (roundflag & SeSplitDegRoundDeg) != 0 {
		dadd = 0.5
	} else if (roundflag & SeSplitDegRoundMin) != 0 {
		dadd = 0.5 / 60.0
	} else if (roundflag & SeSplitDegRoundSec) != 0 {
		dadd = 0.5 / 3600.0
	}
	if (roundflag & SeSplitDegKeepDeg) != 0 {
		if int32(ddeg+dadd)-int32(ddeg) > 0 {
			dadd = 0
		}
	} else if (roundflag & SeSplitDegKeepSign) != 0 {
		if math.Mod(ddeg, 30.0)+dadd >= 30.0 {
			dadd = 0
		}
	}
	ddeg += dadd
	if (roundflag & SeSplitDegZodiacal) != 0 {
		sgn = int32(ddeg / 30.0)
		if sgn == 12 {
			sgn = 0
		}
		ddeg = math.Mod(ddeg, 30.0)
	}
	degOut := int32(ddeg)
	ddeg -= float64(degOut)
	minOut := int32(ddeg * 60.0)
	ddeg -= float64(minOut) / 60.0
	secOut := int32(ddeg * 3600.0)
	var secfrOut float64
	if (roundflag & (SeSplitDegRoundDeg | SeSplitDegRoundMin | SeSplitDegRoundSec)) == 0 {
		secfrOut = ddeg*3600.0 - float64(secOut)
	}
	if (roundflag & SeSplitDegRoundDeg) != 0 {
		minOut = 0
	}
	if (roundflag & (SeSplitDegRoundDeg | SeSplitDegRoundMin)) != 0 {
		secOut = 0
	}
	if len(ideg) > 0 {
		ideg[0] = int(degOut)
	}
	if len(imin) > 0 {
		imin[0] = int(minOut)
	}
	if len(isec) > 0 {
		isec[0] = int(secOut)
	}
	if len(dsecfr) > 0 {
		dsecfr[0] = secfrOut
	}
	if len(isgn) > 0 {
		isgn[0] = sgn
	}
}

// Csnorm function as declared in swephexp.h:990
func Csnorm(p int32) int32 {
	r := int64(p) % Deg360
	if r < 0 {
		r += Deg360
	}
	return int32(r)
}

// Difcsn function as declared in swephexp.h:993
func Difcsn(p1 int, p2 int32) int32 {
	return Csnorm(int32(p1) - p2)
}

// Difdegn function as declared in swephexp.h:995
func Difdegn(p1 float64, p2 float64) float64 {
	return Degnorm(p1 - p2)
}

// Difcs2n function as declared in swephexp.h:998
func Difcs2n(p1 int, p2 int32) int32 {
	dif := Csnorm(int32(p1) - p2)
	if dif >= Deg180 {
		return dif - Deg360
	}
	return dif
}

// Difdeg2n function as declared in swephexp.h:1000
func Difdeg2n(p1 float64, p2 float64) float64 {
	dif := Degnorm(p1 - p2)
	if dif >= 180.0 {
		return dif - 360.0
	}
	return dif
}

// Difrad2n function as declared in swephexp.h:1001
func Difrad2n(p1 float64, p2 float64) float64 {
	dif := Radnorm(p1 - p2)
	if dif >= Twopi/2.0 {
		return dif - Twopi
	}
	return dif
}

// Csroundsec function as declared in swephexp.h:1004
func Csroundsec(x int32) int32 {
	t := (x + 50) / 100 * 100
	if t > x && t%Deg30 == 0 {
		t = x / 100 * 100
	}
	return t
}

// D2l function as declared in swephexp.h:1007
func D2l(x float64) int32 {
	if x >= 0 {
		return int32(x + 0.5)
	}
	return -int32(0.5 - x)
}

// DayOfWeek function as declared in swephexp.h:1010
func DayOfWeek(jd float64) int32 {
	return ((int32(math.Floor(jd-2433282.0-1.5)) % 7) + 7) % 7
}

// Cs2timestr function as declared in swephexp.h:1012
func Cs2timestr(t int, sep int, suppressZero int, a []byte) *byte {
	if len(a) >= 32 {
		ret := C.swe_cs2timestr(C.CSEC(t), C.int(sep), C.AS_BOOL(suppressZero), (*C.char)(unsafe.Pointer(unsafe.SliceData(a))))
		return cCharToBytePtr(ret)
	}
	goseMutex.Lock()
	clear(goseScratch.csbuf[:])
	ret := C.swe_cs2timestr(C.CSEC(t), C.int(sep), C.AS_BOOL(suppressZero), (*C.char)(unsafe.Pointer(&goseScratch.csbuf[0])))
	if len(a) > 0 {
		copy(a, goseScratch.csbuf[:])
		goseMutex.Unlock()
		return &a[0]
	}
	goseMutex.Unlock()
	return cCharToBytePtr(ret)
}

// Cs2lonlatstr function as declared in swephexp.h:1014
func Cs2lonlatstr(t int, pchar byte, mchar byte, s []byte) *byte {
	if len(s) >= 32 {
		ret := C.swe_cs2lonlatstr(C.CSEC(t), C.char(pchar), C.char(mchar), (*C.char)(unsafe.Pointer(unsafe.SliceData(s))))
		return cCharToBytePtr(ret)
	}
	goseMutex.Lock()
	clear(goseScratch.csbuf[:])
	ret := C.swe_cs2lonlatstr(C.CSEC(t), C.char(pchar), C.char(mchar), (*C.char)(unsafe.Pointer(&goseScratch.csbuf[0])))
	if len(s) > 0 {
		copy(s, goseScratch.csbuf[:])
		goseMutex.Unlock()
		return &s[0]
	}
	goseMutex.Unlock()
	return cCharToBytePtr(ret)
}

// Cs2degstr function as declared in swephexp.h:1016
func Cs2degstr(t int, a []byte) *byte {
	if len(a) >= 32 {
		ret := C.swe_cs2degstr(C.CSEC(t), (*C.char)(unsafe.Pointer(unsafe.SliceData(a))))
		return cCharToBytePtr(ret)
	}
	goseMutex.Lock()
	clear(goseScratch.csbuf[:])
	ret := C.swe_cs2degstr(C.CSEC(t), (*C.char)(unsafe.Pointer(&goseScratch.csbuf[0])))
	if len(a) > 0 {
		copy(a, goseScratch.csbuf[:])
		goseMutex.Unlock()
		return &a[0]
	}
	goseMutex.Unlock()
	return cCharToBytePtr(ret)
}

// --- Idiomatic High-Level Go API Helpers ---

// VersionStr returns the Swiss Ephemeris library version as a Go string.
func VersionStr() string {
	goseMutex.Lock()
	goseScratch.spname[0] = 0
	ret := C.swe_version((*C.char)(unsafe.Pointer(&goseScratch.spname[0])))
	s := cCharToString(ret)
	goseMutex.Unlock()
	return s
}

// GetLibraryPathStr returns the Swiss Ephemeris library path as a Go string.
func GetLibraryPathStr() string {
	goseMutex.Lock()
	goseScratch.spname[0] = 0
	ret := C.swe_get_library_path((*C.char)(unsafe.Pointer(&goseScratch.spname[0])))
	s := cCharToString(ret)
	goseMutex.Unlock()
	return s
}

// SetEphePathStr sets the ephemeris path using a Go string.
func SetEphePathStr(path string) {
	goseMutex.Lock()
	cpath := goStringInputPtr(path)
	C.swe_set_ephe_path(cpath)
	goseMutex.Unlock()
}

// SetJplFileStr sets the JPL ephemeris file name using a Go string.
func SetJplFileStr(fname string) {
	goseMutex.Lock()
	cfname := goStringInputPtr(fname)
	C.swe_set_jpl_file(cfname)
	goseMutex.Unlock()
}

// GetPlanetNameStr returns the planet name directly as a Go string.
func GetPlanetNameStr(ipl int) string {
	goseMutex.Lock()
	goseScratch.spname[0] = 0
	ret := C.swe_get_planet_name(C.int(ipl), (*C.char)(unsafe.Pointer(&goseScratch.spname[0])))
	s := cCharToString(ret)
	goseMutex.Unlock()
	return s
}

// HouseNameStr returns the name of a house system as a Go string.
func HouseNameStr(hsys int32) string {
	ret := C.swe_house_name(C.int(hsys))
	return cCharToString(ret)
}

// JuldayTime converts a standard Go time.Time to Julian Day (UT).
func JuldayTime(t time.Time, gregflag int32) float64 {
	ut := t.UTC()
	hour := float64(ut.Hour()) + float64(ut.Minute())/60.0 + float64(ut.Second())/3600.0 + float64(ut.Nanosecond())/3.6e12
	return Julday(ut.Year(), int(ut.Month()), ut.Day(), hour, gregflag)
}

// RevjulTime converts a Julian Day (UT) to a standard Go time.Time in UTC.
func RevjulTime(jd float64, gregflag int) time.Time {
	y, m, d, ut := revjulCore(jd, gregflag)
	totalNsec := int64(math.Round(ut * 3.6e12))
	hour := int(totalNsec / 3_600_000_000_000)
	rem := totalNsec % 3_600_000_000_000
	min := int(rem / 60_000_000_000)
	rem %= 60_000_000_000
	sec := int(rem / 1_000_000_000)
	nsec := int(rem % 1_000_000_000)
	return time.Date(y, time.Month(m), d, hour, min, sec, nsec, time.UTC)
}

// CalcUtSimple provides an ergonomic wrapper around CalcUt, returning a structured Coordinates object.
func CalcUtSimple(tjdUt float64, ipl int, iflag int) (pos Coordinates, iflgret int32, err error) {
	goseMutex.Lock()
	goseScratch.serr[0] = 0
	pxx := (*C.double)(unsafe.Pointer(&goseScratch.out1[0]))
	pserr := (*C.char)(unsafe.Pointer(&goseScratch.serr[0]))
	ret := C.swe_calc_ut(C.double(tjdUt), C.int32(ipl), C.int32(iflag), pxx, pserr)
	iflgret = int32(ret)
	if iflgret < 0 {
		errMsg := cBytesToString(goseScratch.serr[:])
		goseMutex.Unlock()
		return Coordinates{}, iflgret, errors.New(errMsg)
	}
	pos = Coordinates{
		Longitude: goseScratch.out1[0],
		Latitude:  goseScratch.out1[1],
		Distance:  goseScratch.out1[2],
		SpeedLong: goseScratch.out1[3],
		SpeedLat:  goseScratch.out1[4],
		SpeedDist: goseScratch.out1[5],
	}
	goseMutex.Unlock()
	return pos, iflgret, nil
}

// HousesSimple provides an ergonomic wrapper around HousesEx, returning a structured HousesResult.
func HousesSimple(tjdUt float64, iflag int, geolat, geolon float64, hsys int) (res HousesResult, iflgret int32, err error) {
	goseMutex.Lock()
	pcusps := (*C.double)(unsafe.Pointer(&goseScratch.out1[0]))
	pascmc := (*C.double)(unsafe.Pointer(&goseScratch.out2[0]))
	ret := C.swe_houses_ex(C.double(tjdUt), C.int32(iflag), C.double(geolat), C.double(geolon), C.int(hsys), pcusps, pascmc)
	iflgret = int32(ret)
	if iflgret < 0 {
		goseMutex.Unlock()
		return HousesResult{}, iflgret, errors.New("houses calculation failed")
	}
	copy(res.Cusps[:], goseScratch.out1[:13])
	res.Ascendant = goseScratch.out2[0]
	res.MC = goseScratch.out2[1]
	res.ARMC = goseScratch.out2[2]
	res.Vertex = goseScratch.out2[3]
	res.Equasc = goseScratch.out2[4]
	res.Coasc1 = goseScratch.out2[5]
	res.Coasc2 = goseScratch.out2[6]
	res.Polasc = goseScratch.out2[7]
	goseMutex.Unlock()
	return res, iflgret, nil
}
