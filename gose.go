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
	"errors"
	"sync"
	"time"
	"unsafe"
)

var goseMutex sync.Mutex

// HeliacalUt function as declared in swephexp.h:676
func HeliacalUt(tjdstartUt float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, typeEvent int, iflag int, dret []float64, serr []byte) int32 {
	cgeopos := safeDoubleInput(geopos, 3)
	cdatm := safeDoubleInput(datm, 4)
	cdobs := safeDoubleInput(dobs, 6)
	cobjectName := safeCString(objectName)
	cdret, cleanupRet := safeDoubleOutput(dret, 50)
	if cleanupRet != nil {
		defer cleanupRet()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_heliacal_ut(C.double(tjdstartUt), cgeopos, cdatm, cdobs, cobjectName, C.int32(typeEvent), C.int32(iflag), cdret, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// HeliacalPhenoUt function as declared in swephexp.h:677
func HeliacalPhenoUt(tjdUt float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, typeEvent int, helflag int, darr []float64, serr []byte) int32 {
	cgeopos := safeDoubleInput(geopos, 3)
	cdatm := safeDoubleInput(datm, 4)
	cdobs := safeDoubleInput(dobs, 6)
	cobjectName := safeCString(objectName)
	cdarr, cleanupArr := safeDoubleOutput(darr, 50)
	if cleanupArr != nil {
		defer cleanupArr()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_heliacal_pheno_ut(C.double(tjdUt), cgeopos, cdatm, cdobs, cobjectName, C.int32(typeEvent), C.int32(helflag), cdarr, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// VisLimitMag function as declared in swephexp.h:678
func VisLimitMag(tjdut float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, helflag int, dret []float64, serr []byte) int32 {
	cgeopos := safeDoubleInput(geopos, 3)
	cdatm := safeDoubleInput(datm, 4)
	cdobs := safeDoubleInput(dobs, 6)
	cobjectName := safeCString(objectName)
	cdret, cleanupRet := safeDoubleOutput(dret, 50)
	if cleanupRet != nil {
		defer cleanupRet()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_vis_limit_mag(C.double(tjdut), cgeopos, cdatm, cdobs, cobjectName, C.int32(helflag), cdret, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// HeliacalAngle function as declared in swephexp.h:681
func HeliacalAngle(tjdut float64, dgeo []float64, datm []float64, dobs []float64, helflag int, mag float64, aziObj float64, aziSun float64, aziMoon float64, altMoon float64, dret []float64, serr []byte) int32 {
	cdgeo := safeDoubleInput(dgeo, 3)
	cdatm := safeDoubleInput(datm, 4)
	cdobs := safeDoubleInput(dobs, 6)
	cdret, cleanupRet := safeDoubleOutput(dret, 50)
	if cleanupRet != nil {
		defer cleanupRet()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_heliacal_angle(C.double(tjdut), cdgeo, cdatm, cdobs, C.int32(helflag), C.double(mag), C.double(aziObj), C.double(aziSun), C.double(aziMoon), C.double(altMoon), cdret, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// TopoArcusVisionis function as declared in swephexp.h:682
func TopoArcusVisionis(tjdut float64, dgeo []float64, datm []float64, dobs []float64, helflag int, mag float64, aziObj float64, altObj float64, aziSun float64, aziMoon float64, altMoon float64, dret []float64, serr []byte) int32 {
	cdgeo := safeDoubleInput(dgeo, 3)
	cdatm := safeDoubleInput(datm, 4)
	cdobs := safeDoubleInput(dobs, 6)
	cdret, cleanupRet := safeDoubleOutput(dret, 50)
	if cleanupRet != nil {
		defer cleanupRet()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_topo_arcus_visionis(C.double(tjdut), cdgeo, cdatm, cdobs, C.int32(helflag), C.double(mag), C.double(aziObj), C.double(altObj), C.double(aziSun), C.double(aziMoon), C.double(altMoon), cdret, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// SetAstroModels function as declared in swephexp.h:686
func SetAstroModels(samod []byte, iflag int32) {
	csamod := safeCString(samod)
	goseMutex.Lock()
	C.swe_set_astro_models(csamod, C.int32(iflag))
	goseMutex.Unlock()
}

// GetAstroModels function as declared in swephexp.h:687
func GetAstroModels(samod []byte, sdet []byte, iflag int32) {
	var bufMod [256]byte
	var bufDet [256]byte
	var csamod, csdet *C.char
	if len(samod) >= 256 {
		csamod = (*C.char)(unsafe.Pointer(unsafe.SliceData(samod)))
	} else {
		csamod = (*C.char)(unsafe.Pointer(&bufMod[0]))
	}
	if len(sdet) >= 256 {
		csdet = (*C.char)(unsafe.Pointer(unsafe.SliceData(sdet)))
	} else {
		csdet = (*C.char)(unsafe.Pointer(&bufDet[0]))
	}
	goseMutex.Lock()
	C.swe_get_astro_models(csamod, csdet, C.int32(iflag))
	goseMutex.Unlock()
	if len(samod) < 256 && len(samod) > 0 {
		copy(samod, bufMod[:])
	}
	if len(sdet) < 256 && len(sdet) > 0 {
		copy(sdet, bufDet[:])
	}
}

// Version function as declared in swephexp.h:693
func Version(arg0 []byte) *byte {
	var carg0 *C.char
	var buf [256]byte
	if len(arg0) >= 256 {
		carg0 = (*C.char)(unsafe.Pointer(unsafe.SliceData(arg0)))
	} else if len(arg0) > 0 {
		carg0 = (*C.char)(unsafe.Pointer(&buf[0]))
	}
	ret := C.swe_version(carg0)
	if len(arg0) > 0 && len(arg0) < 256 {
		copy(arg0, buf[:])
	}
	return cCharToBytePtr(ret)
}

// GetLibraryPath function as declared in swephexp.h:694
func GetLibraryPath(arg0 []byte) *byte {
	var carg0 *C.char
	var buf [256]byte
	if len(arg0) >= 256 {
		carg0 = (*C.char)(unsafe.Pointer(unsafe.SliceData(arg0)))
	} else if len(arg0) > 0 {
		carg0 = (*C.char)(unsafe.Pointer(&buf[0]))
	}
	goseMutex.Lock()
	ret := C.swe_get_library_path(carg0)
	goseMutex.Unlock()
	if len(arg0) > 0 && len(arg0) < 256 {
		copy(arg0, buf[:])
	}
	return cCharToBytePtr(ret)
}

// Calc function as declared in swephexp.h:697
func Calc(tjd float64, ipl int, iflag int, xx []float64, serr []byte) int32 {
	pxx, cleanupXX := safeDoubleOutput(xx, 6)
	if cleanupXX != nil {
		defer cleanupXX()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_calc(C.double(tjd), C.int(ipl), C.int32(iflag), pxx, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// CalcUt function as declared in swephexp.h:702
func CalcUt(tjdUt float64, ipl int, iflag int, xx []float64, serr []byte) int32 {
	pxx, cleanupXX := safeDoubleOutput(xx, 6)
	if cleanupXX != nil {
		defer cleanupXX()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_calc_ut(C.double(tjdUt), C.int32(ipl), C.int32(iflag), pxx, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// CalcPctr function as declared in swephexp.h:705
func CalcPctr(tjd float64, ipl int, iplctr int, iflag int, xxret []float64, serr []byte) int32 {
	pxx, cleanupXX := safeDoubleOutput(xxret, 6)
	if cleanupXX != nil {
		defer cleanupXX()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_calc_pctr(C.double(tjd), C.int32(ipl), C.int32(iplctr), C.int32(iflag), pxx, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// Solcross function as declared in swephexp.h:707
func Solcross(x2cross float64, jdEt float64, flag int32, serr []byte) float64 {
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_solcross(C.double(x2cross), C.double(jdEt), C.int32(flag), pserr)
	goseMutex.Unlock()
	return float64(ret)
}

// SolcrossUt function as declared in swephexp.h:708
func SolcrossUt(x2cross float64, jdUt float64, flag int32, serr []byte) float64 {
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_solcross_ut(C.double(x2cross), C.double(jdUt), C.int32(flag), pserr)
	goseMutex.Unlock()
	return float64(ret)
}

// Mooncross function as declared in swephexp.h:709
func Mooncross(x2cross float64, jdEt float64, flag int32, serr []byte) float64 {
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_mooncross(C.double(x2cross), C.double(jdEt), C.int32(flag), pserr)
	goseMutex.Unlock()
	return float64(ret)
}

// MooncrossUt function as declared in swephexp.h:710
func MooncrossUt(x2cross float64, jdUt float64, flag int32, serr []byte) float64 {
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_mooncross_ut(C.double(x2cross), C.double(jdUt), C.int32(flag), pserr)
	goseMutex.Unlock()
	return float64(ret)
}

// MooncrossNode function as declared in swephexp.h:711
func MooncrossNode(jdEt float64, flag int32, xlon []float64, xlat []float64, serr []byte) float64 {
	var cxlon, cxlat C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_mooncross_node(C.double(jdEt), C.int32(flag), &cxlon, &cxlat, pserr)
	goseMutex.Unlock()
	if len(xlon) > 0 { xlon[0] = float64(cxlon) }
	if len(xlat) > 0 { xlat[0] = float64(cxlat) }
	return float64(ret)
}

// MooncrossNodeUt function as declared in swephexp.h:712
func MooncrossNodeUt(jdUt float64, flag int32, xlon []float64, xlat []float64, serr []byte) float64 {
	var cxlon, cxlat C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_mooncross_node_ut(C.double(jdUt), C.int32(flag), &cxlon, &cxlat, pserr)
	goseMutex.Unlock()
	if len(xlon) > 0 { xlon[0] = float64(cxlon) }
	if len(xlat) > 0 { xlat[0] = float64(cxlat) }
	return float64(ret)
}

// HelioCross function as declared in swephexp.h:713
func HelioCross(ipl int32, x2cross float64, jdEt float64, iflag int32, dir int32, jdCross []float64, serr []byte) int32 {
	var cjdcross C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_helio_cross(C.int32(ipl), C.double(x2cross), C.double(jdEt), C.int32(iflag), C.int32(dir), &cjdcross, pserr)
	goseMutex.Unlock()
	if len(jdCross) > 0 { jdCross[0] = float64(cjdcross) }
	return int32(ret)
}

// HelioCrossUt function as declared in swephexp.h:714
func HelioCrossUt(ipl int32, x2cross float64, jdUt float64, iflag int32, dir int32, jdCross []float64, serr []byte) int32 {
	var cjdcross C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_helio_cross_ut(C.int32(ipl), C.double(x2cross), C.double(jdUt), C.int32(iflag), C.int32(dir), &cjdcross, pserr)
	goseMutex.Unlock()
	if len(jdCross) > 0 { jdCross[0] = float64(cjdcross) }
	return int32(ret)
}

// Fixstar function as declared in swephexp.h:717
func Fixstar(star []byte, tjd float64, iflag int, xx []float64, serr []byte) int32 {
	cstar := safeCString(star)
	pxx, cleanupXX := safeDoubleOutput(xx, 6)
	if cleanupXX != nil {
		defer cleanupXX()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_fixstar(cstar, C.double(tjd), C.int32(iflag), pxx, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// FixstarUt function as declared in swephexp.h:722
func FixstarUt(star []byte, tjdUt float64, iflag int, xx []float64, serr []byte) int32 {
	cstar := safeCString(star)
	pxx, cleanupXX := safeDoubleOutput(xx, 6)
	if cleanupXX != nil {
		defer cleanupXX()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_fixstar_ut(cstar, C.double(tjdUt), C.int32(iflag), pxx, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// FixstarMag function as declared in swephexp.h:725
func FixstarMag(star []byte, mag []float64, serr []byte) int32 {
	cstar := safeCString(star)
	var cmag C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_fixstar_mag(cstar, &cmag, pserr)
	goseMutex.Unlock()
	if len(mag) > 0 { mag[0] = float64(cmag) }
	return int32(ret)
}

// Fixstar2 function as declared in swephexp.h:727
func Fixstar2(star []byte, tjd float64, iflag int, xx []float64, serr []byte) int32 {
	cstar := safeCString(star)
	pxx, cleanupXX := safeDoubleOutput(xx, 6)
	if cleanupXX != nil {
		defer cleanupXX()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_fixstar2(cstar, C.double(tjd), C.int32(iflag), pxx, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// Fixstar2Ut function as declared in swephexp.h:732
func Fixstar2Ut(star []byte, tjdUt float64, iflag int, xx []float64, serr []byte) int32 {
	cstar := safeCString(star)
	pxx, cleanupXX := safeDoubleOutput(xx, 6)
	if cleanupXX != nil {
		defer cleanupXX()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_fixstar2_ut(cstar, C.double(tjdUt), C.int32(iflag), pxx, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// Fixstar2Mag function as declared in swephexp.h:735
func Fixstar2Mag(star []byte, mag []float64, serr []byte) int32 {
	cstar := safeCString(star)
	var cmag C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_fixstar2_mag(cstar, &cmag, pserr)
	goseMutex.Unlock()
	if len(mag) > 0 { mag[0] = float64(cmag) }
	return int32(ret)
}

// Close function as declared in swephexp.h:738
func Close() {
	goseMutex.Lock()
	C.swe_close()
	goseMutex.Unlock()
}

// SetEphePath function as declared in swephexp.h:741
func SetEphePath(path []byte) {
	cpath := safeCString(path)
	goseMutex.Lock()
	C.swe_set_ephe_path(cpath)
	goseMutex.Unlock()
}

// SetJplFile function as declared in swephexp.h:744
func SetJplFile(fname []byte) {
	cfname := safeCString(fname)
	goseMutex.Lock()
	C.swe_set_jpl_file(cfname)
	goseMutex.Unlock()
}

// GetPlanetName function as declared in swephexp.h:747
func GetPlanetName(ipl int, spname []byte) *byte {
	var buf [256]byte
	var cspname *C.char
	if len(spname) >= 256 {
		cspname = (*C.char)(unsafe.Pointer(unsafe.SliceData(spname)))
	} else if len(spname) > 0 {
		cspname = (*C.char)(unsafe.Pointer(&buf[0]))
	}
	goseMutex.Lock()
	ret := C.swe_get_planet_name(C.int(ipl), cspname)
	goseMutex.Unlock()
	if len(spname) > 0 && len(spname) < 256 {
		copy(spname, buf[:])
	}
	return cCharToBytePtr(ret)
}

// SetTopo function as declared in swephexp.h:750
func SetTopo(geolon float64, geolat float64, geoalt float64) {
	goseMutex.Lock()
	C.swe_set_topo(C.double(geolon), C.double(geolat), C.double(geoalt))
	goseMutex.Unlock()
}

// SetSidMode function as declared in swephexp.h:753
func SetSidMode(sidMode int, t0 float64, ayanT0 float64) {
	goseMutex.Lock()
	C.swe_set_sid_mode(C.int32(sidMode), C.double(t0), C.double(ayanT0))
	goseMutex.Unlock()
}

// GetAyanamsaEx function as declared in swephexp.h:756
func GetAyanamsaEx(tjdEt float64, iflag int, daya []float64, serr []byte) int32 {
	var cdaya C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_get_ayanamsa_ex(C.double(tjdEt), C.int32(iflag), &cdaya, pserr)
	goseMutex.Unlock()
	if len(daya) > 0 { daya[0] = float64(cdaya) }
	return int32(ret)
}

// GetAyanamsaExUt function as declared in swephexp.h:757
func GetAyanamsaExUt(tjdUt float64, iflag int, daya []float64, serr []byte) int32 {
	var cdaya C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_get_ayanamsa_ex_ut(C.double(tjdUt), C.int32(iflag), &cdaya, pserr)
	goseMutex.Unlock()
	if len(daya) > 0 { daya[0] = float64(cdaya) }
	return int32(ret)
}

// GetAyanamsa function as declared in swephexp.h:758
func GetAyanamsa(tjdEt float64) float64 {
	goseMutex.Lock()
	ret := C.swe_get_ayanamsa(C.double(tjdEt))
	goseMutex.Unlock()
	return float64(ret)
}

// GetAyanamsaUt function as declared in swephexp.h:759
func GetAyanamsaUt(tjdUt float64) float64 {
	goseMutex.Lock()
	ret := C.swe_get_ayanamsa_ut(C.double(tjdUt))
	goseMutex.Unlock()
	return float64(ret)
}

// GetAyanamsaName function as declared in swephexp.h:762
func GetAyanamsaName(isidmode int32) string {
	ret := C.swe_get_ayanamsa_name(C.int32(isidmode))
	return cCharToString(ret)
}

// GetCurrentFileData function as declared in swephexp.h:763
func GetCurrentFileData(ifno int, tfstart []float64, tfend []float64, denum []int32) string {
	var ctfstart, ctfend C.double
	var cdenum C.int
	goseMutex.Lock()
	ret := C.swe_get_current_file_data(C.int(ifno), &ctfstart, &ctfend, &cdenum)
	goseMutex.Unlock()
	if len(tfstart) > 0 { tfstart[0] = float64(ctfstart) }
	if len(tfend) > 0 { tfend[0] = float64(ctfend) }
	if len(denum) > 0 { denum[0] = int32(cdenum) }
	return cCharToString(ret)
}

// DateConversion function as declared in swephexp.h:771
func DateConversion(y int, m int, d int, utime float64, c byte, tjd []float64) int32 {
	var ctjd C.double
	ret := C.swe_date_conversion(C.int(y), C.int(m), C.int(d), C.double(utime), C.char(c), &ctjd)
	if len(tjd) > 0 { tjd[0] = float64(ctjd) }
	return int32(ret)
}

// Julday function as declared in swephexp.h:777
func Julday(year int, month int, day int, hour float64, gregflag int32) float64 {
	ret := C.swe_julday(C.int(year), C.int(month), C.int(day), C.double(hour), C.int(gregflag))
	return float64(ret)
}

// Revjul function as declared in swephexp.h:781
func Revjul(jd float64, gregflag int, jyear []int, jmon []int, jday []int, jut []float64) {
	var cy, cm, cd C.int
	var cjut C.double
	C.swe_revjul(C.double(jd), C.int(gregflag), &cy, &cm, &cd, &cjut)
	if len(jyear) > 0 { jyear[0] = int(cy) }
	if len(jmon) > 0 { jmon[0] = int(cm) }
	if len(jday) > 0 { jday[0] = int(cd) }
	if len(jut) > 0 { jut[0] = float64(cjut) }
}

// UtcToJd function as declared in swephexp.h:786
func UtcToJd(iyear int, imonth int, iday int, ihour int, imin int, dsec float64, gregflag int, dret []float64, serr []byte) int32 {
	pdret, cleanupRet := safeDoubleOutput(dret, 2)
	if cleanupRet != nil {
		defer cleanupRet()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_utc_to_jd(C.int32(iyear), C.int32(imonth), C.int32(iday), C.int32(ihour), C.int32(imin), C.double(dsec), C.int32(gregflag), pdret, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// JdetToUtc function as declared in swephexp.h:791
func JdetToUtc(tjdEt float64, gregflag int, iyear []int, imonth []int, iday []int, ihour []int, imin []int, dsec []float64) {
	var cy, cm, cd, ch, cmin C.int32
	var cdsec C.double
	goseMutex.Lock()
	C.swe_jdet_to_utc(C.double(tjdEt), C.int32(gregflag), &cy, &cm, &cd, &ch, &cmin, &cdsec)
	goseMutex.Unlock()
	if len(iyear) > 0 { iyear[0] = int(cy) }
	if len(imonth) > 0 { imonth[0] = int(cm) }
	if len(iday) > 0 { iday[0] = int(cd) }
	if len(ihour) > 0 { ihour[0] = int(ch) }
	if len(imin) > 0 { imin[0] = int(cmin) }
	if len(dsec) > 0 { dsec[0] = float64(cdsec) }
}

// Jdut1ToUtc function as declared in swephexp.h:796
func Jdut1ToUtc(tjdUt float64, gregflag int, iyear []int, imonth []int, iday []int, ihour []int, imin []int, dsec []float64) {
	var cy, cm, cd, ch, cmin C.int32
	var cdsec C.double
	goseMutex.Lock()
	C.swe_jdut1_to_utc(C.double(tjdUt), C.int32(gregflag), &cy, &cm, &cd, &ch, &cmin, &cdsec)
	goseMutex.Unlock()
	if len(iyear) > 0 { iyear[0] = int(cy) }
	if len(imonth) > 0 { imonth[0] = int(cm) }
	if len(iday) > 0 { iday[0] = int(cd) }
	if len(ihour) > 0 { ihour[0] = int(ch) }
	if len(imin) > 0 { imin[0] = int(cmin) }
	if len(dsec) > 0 { dsec[0] = float64(cdsec) }
}

// UtcTimeZone function as declared in swephexp.h:801
func UtcTimeZone(iyear int, imonth int, iday int, ihour int, imin int, dsec float64, dTimezone float64, iyearOut []int, imonthOut []int, idayOut []int, ihourOut []int, iminOut []int, dsecOut []float64) {
	var cy, cm, cd, ch, cmin C.int32
	var cdsec C.double
	C.swe_utc_time_zone(C.int32(iyear), C.int32(imonth), C.int32(iday), C.int32(ihour), C.int32(imin), C.double(dsec), C.double(dTimezone), &cy, &cm, &cd, &ch, &cmin, &cdsec)
	if len(iyearOut) > 0 { iyearOut[0] = int(cy) }
	if len(imonthOut) > 0 { imonthOut[0] = int(cm) }
	if len(idayOut) > 0 { idayOut[0] = int(cd) }
	if len(ihourOut) > 0 { ihourOut[0] = int(ch) }
	if len(iminOut) > 0 { iminOut[0] = int(cmin) }
	if len(dsecOut) > 0 { dsecOut[0] = float64(cdsec) }
}

// Houses function as declared in swephexp.h:812
func Houses(tjdUt float64, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64) int32 {
	pcusps, cleanupCusps := safeDoubleOutput(cusps, 13)
	if cleanupCusps != nil {
		defer cleanupCusps()
	}
	pascmc, cleanupAscmc := safeDoubleOutput(ascmc, 10)
	if cleanupAscmc != nil {
		defer cleanupAscmc()
	}
	goseMutex.Lock()
	ret := C.swe_houses(C.double(tjdUt), C.double(geolat), C.double(geolon), C.int(hsys), pcusps, pascmc)
	goseMutex.Unlock()
	return int32(ret)
}

// HousesEx function as declared in swephexp.h:816
func HousesEx(tjdUt float64, iflag int, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64) int32 {
	pcusps, cleanupCusps := safeDoubleOutput(cusps, 13)
	if cleanupCusps != nil {
		defer cleanupCusps()
	}
	pascmc, cleanupAscmc := safeDoubleOutput(ascmc, 10)
	if cleanupAscmc != nil {
		defer cleanupAscmc()
	}
	goseMutex.Lock()
	ret := C.swe_houses_ex(C.double(tjdUt), C.int32(iflag), C.double(geolat), C.double(geolon), C.int(hsys), pcusps, pascmc)
	goseMutex.Unlock()
	return int32(ret)
}

// HousesEx2 function as declared in swephexp.h:820
func HousesEx2(tjdUt float64, iflag int, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64, cuspSpeed []float64, ascmcSpeed []float64, serr []byte) int32 {
	pcusps, cleanupCusps := safeDoubleOutput(cusps, 13)
	if cleanupCusps != nil {
		defer cleanupCusps()
	}
	pascmc, cleanupAscmc := safeDoubleOutput(ascmc, 10)
	if cleanupAscmc != nil {
		defer cleanupAscmc()
	}
	pcuspSpeed, cleanupCuspSpeed := safeDoubleOutput(cuspSpeed, 13)
	if cleanupCuspSpeed != nil {
		defer cleanupCuspSpeed()
	}
	pascmcSpeed, cleanupAscmcSpeed := safeDoubleOutput(ascmcSpeed, 10)
	if cleanupAscmcSpeed != nil {
		defer cleanupAscmcSpeed()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_houses_ex2(C.double(tjdUt), C.int32(iflag), C.double(geolat), C.double(geolon), C.int(hsys), pcusps, pascmc, pcuspSpeed, pascmcSpeed, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// HousesArmc function as declared in swephexp.h:824
func HousesArmc(armc float64, geolat float64, eps float64, hsys int, cusps []float64, ascmc []float64) int32 {
	pcusps, cleanupCusps := safeDoubleOutput(cusps, 13)
	if cleanupCusps != nil {
		defer cleanupCusps()
	}
	pascmc, cleanupAscmc := safeDoubleOutput(ascmc, 10)
	if cleanupAscmc != nil {
		defer cleanupAscmc()
	}
	goseMutex.Lock()
	ret := C.swe_houses_armc(C.double(armc), C.double(geolat), C.double(eps), C.int(hsys), pcusps, pascmc)
	goseMutex.Unlock()
	return int32(ret)
}

// HousesArmcEx2 function as declared in swephexp.h:828
func HousesArmcEx2(armc float64, geolat float64, eps float64, hsys int, cusps []float64, ascmc []float64, cuspSpeed []float64, ascmcSpeed []float64, serr []byte) int32 {
	pcusps, cleanupCusps := safeDoubleOutput(cusps, 13)
	if cleanupCusps != nil {
		defer cleanupCusps()
	}
	pascmc, cleanupAscmc := safeDoubleOutput(ascmc, 10)
	if cleanupAscmc != nil {
		defer cleanupAscmc()
	}
	pcuspSpeed, cleanupCuspSpeed := safeDoubleOutput(cuspSpeed, 13)
	if cleanupCuspSpeed != nil {
		defer cleanupCuspSpeed()
	}
	pascmcSpeed, cleanupAscmcSpeed := safeDoubleOutput(ascmcSpeed, 10)
	if cleanupAscmcSpeed != nil {
		defer cleanupAscmcSpeed()
	}
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_houses_armc_ex2(C.double(armc), C.double(geolat), C.double(eps), C.int(hsys), pcusps, pascmc, pcuspSpeed, pascmcSpeed, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// HousePos function as declared in swephexp.h:832
func HousePos(armc float64, geolat float64, eps float64, hsys int, xpin []float64, serr []byte) float64 {
	pxpin := safeDoubleInput(xpin, 2)
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_house_pos(C.double(armc), C.double(geolat), C.double(eps), C.int(hsys), pxpin, pserr)
	goseMutex.Unlock()
	return float64(ret)
}

// HouseName function as declared in swephexp.h:835
func HouseName(hsys int32) *byte {
	ret := C.swe_house_name(C.int(hsys))
	return cCharToBytePtr(ret)
}

// GauquelinSector function as declared in swephexp.h:843
func GauquelinSector(tUt float64, ipl int, starname []byte, iflag int, imeth int, geopos []float64, atpress float64, attemp float64, dgsect []float64, serr []byte) int32 {
	cstarname := safeCString(starname)
	cgeopos := safeDoubleInput(geopos, 3)
	var cdgsect C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil {
		defer cleanupErr()
	}
	goseMutex.Lock()
	ret := C.swe_gauquelin_sector(C.double(tUt), C.int32(ipl), cstarname, C.int32(iflag), C.int32(imeth), cgeopos, C.double(atpress), C.double(attemp), &cdgsect, pserr)
	goseMutex.Unlock()
	if len(dgsect) > 0 { dgsect[0] = float64(cdgsect) }
	return int32(ret)
}

// SolEclipseWhere function as declared in swephexp.h:847
func SolEclipseWhere(tjd float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	cgeopos, cleanupGeo := safeDoubleOutput(geopos, 2)
	if cleanupGeo != nil { defer cleanupGeo() }
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_sol_eclipse_where(C.double(tjd), C.int32(ifl), cgeopos, cattr, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunOccultWhere function as declared in swephexp.h:849
func LunOccultWhere(tjd float64, ipl int, starname []byte, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	cstarname := safeCString(starname)
	cgeopos, cleanupGeo := safeDoubleOutput(geopos, 2)
	if cleanupGeo != nil { defer cleanupGeo() }
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_lun_occult_where(C.double(tjd), C.int32(ipl), cstarname, C.int32(ifl), cgeopos, cattr, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// SolEclipseHow function as declared in swephexp.h:852
func SolEclipseHow(tjd float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	cgeopos := safeDoubleInput(geopos, 3)
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_sol_eclipse_how(C.double(tjd), C.int32(ifl), cgeopos, cattr, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// SolEclipseWhenLoc function as declared in swephexp.h:855
func SolEclipseWhenLoc(tjdStart float64, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	cgeopos := safeDoubleInput(geopos, 3)
	ctret, cleanupTret := safeDoubleOutput(tret, 10)
	if cleanupTret != nil { defer cleanupTret() }
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_sol_eclipse_when_loc(C.double(tjdStart), C.int32(ifl), cgeopos, ctret, cattr, C.int32(backward), pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunOccultWhenLoc function as declared in swephexp.h:857
func LunOccultWhenLoc(tjdStart float64, ipl int, starname []byte, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	cstarname := safeCString(starname)
	cgeopos := safeDoubleInput(geopos, 3)
	ctret, cleanupTret := safeDoubleOutput(tret, 10)
	if cleanupTret != nil { defer cleanupTret() }
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_lun_occult_when_loc(C.double(tjdStart), C.int32(ipl), cstarname, C.int32(ifl), cgeopos, ctret, cattr, C.int32(backward), pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// SolEclipseWhenGlob function as declared in swephexp.h:861
func SolEclipseWhenGlob(tjdStart float64, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	ctret, cleanupTret := safeDoubleOutput(tret, 10)
	if cleanupTret != nil { defer cleanupTret() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_sol_eclipse_when_glob(C.double(tjdStart), C.int32(ifl), C.int32(ifltype), ctret, C.int32(backward), pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunOccultWhenGlob function as declared in swephexp.h:865
func LunOccultWhenGlob(tjdStart float64, ipl int, starname []byte, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	cstarname := safeCString(starname)
	ctret, cleanupTret := safeDoubleOutput(tret, 10)
	if cleanupTret != nil { defer cleanupTret() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_lun_occult_when_glob(C.double(tjdStart), C.int32(ipl), cstarname, C.int32(ifl), C.int32(ifltype), ctret, C.int32(backward), pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunEclipseHow function as declared in swephexp.h:869
func LunEclipseHow(tjdUt float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	cgeopos := safeDoubleInput(geopos, 3)
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_lun_eclipse_how(C.double(tjdUt), C.int32(ifl), cgeopos, cattr, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunEclipseWhen function as declared in swephexp.h:876
func LunEclipseWhen(tjdStart float64, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	ctret, cleanupTret := safeDoubleOutput(tret, 10)
	if cleanupTret != nil { defer cleanupTret() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_lun_eclipse_when(C.double(tjdStart), C.int32(ifl), C.int32(ifltype), ctret, C.int32(backward), pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// LunEclipseWhenLoc function as declared in swephexp.h:879
func LunEclipseWhenLoc(tjdStart float64, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	cgeopos := safeDoubleInput(geopos, 3)
	ctret, cleanupTret := safeDoubleOutput(tret, 10)
	if cleanupTret != nil { defer cleanupTret() }
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_lun_eclipse_when_loc(C.double(tjdStart), C.int32(ifl), cgeopos, ctret, cattr, C.int32(backward), pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// Pheno function as declared in swephexp.h:883
func Pheno(tjd float64, ipl int, iflag int, attr []float64, serr []byte) int32 {
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_pheno(C.double(tjd), C.int32(ipl), C.int32(iflag), cattr, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// PhenoUt function as declared in swephexp.h:885
func PhenoUt(tjdUt float64, ipl int, iflag int, attr []float64, serr []byte) int32 {
	cattr, cleanupAttr := safeDoubleOutput(attr, 20)
	if cleanupAttr != nil { defer cleanupAttr() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_pheno_ut(C.double(tjdUt), C.int32(ipl), C.int32(iflag), cattr, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// Refrac function as declared in swephexp.h:887
func Refrac(inalt float64, atpress float64, attemp float64, calcFlag int32) float64 {
	ret := C.swe_refrac(C.double(inalt), C.double(atpress), C.double(attemp), C.int32(calcFlag))
	return float64(ret)
}

// RefracExtended function as declared in swephexp.h:889
func RefracExtended(inalt float64, geoalt float64, atpress float64, attemp float64, lapseRate float64, calcFlag int, dret []float64) float64 {
	cdret, cleanupRet := safeDoubleOutput(dret, 4)
	if cleanupRet != nil { defer cleanupRet() }
	ret := C.swe_refrac_extended(C.double(inalt), C.double(geoalt), C.double(atpress), C.double(attemp), C.double(lapseRate), C.int32(calcFlag), cdret)
	return float64(ret)
}

// SetLapseRate function as declared in swephexp.h:891
func SetLapseRate(lapseRate float64) {
	goseMutex.Lock()
	C.swe_set_lapse_rate(C.double(lapseRate))
	goseMutex.Unlock()
}

// Azalt function as declared in swephexp.h:893
func Azalt(tjdUt float64, calcFlag int, geopos []float64, atpress float64, attemp float64, xin []float64, xaz []float64) {
	cgeopos := safeDoubleInput(geopos, 3)
	cxin := safeDoubleInput(xin, 3)
	cxaz, cleanupAz := safeDoubleOutput(xaz, 3)
	if cleanupAz != nil { defer cleanupAz() }
	goseMutex.Lock()
	C.swe_azalt(C.double(tjdUt), C.int32(calcFlag), cgeopos, C.double(atpress), C.double(attemp), cxin, cxaz)
	goseMutex.Unlock()
}

// AzaltRev function as declared in swephexp.h:902
func AzaltRev(tjdUt float64, calcFlag int, geopos []float64, xin []float64, xout []float64) {
	cgeopos := safeDoubleInput(geopos, 3)
	cxin := safeDoubleInput(xin, 3)
	cxout, cleanupOut := safeDoubleOutput(xout, 3)
	if cleanupOut != nil { defer cleanupOut() }
	goseMutex.Lock()
	C.swe_azalt_rev(C.double(tjdUt), C.int32(calcFlag), cgeopos, cxin, cxout)
	goseMutex.Unlock()
}

// RiseTransTrueHor function as declared in swephexp.h:909
func RiseTransTrueHor(tjdUt float64, ipl int, starname []byte, epheflag int, rsmi int, geopos []float64, atpress float64, attemp float64, horhgt float64, tret []float64, serr []byte) int32 {
	cstarname := safeCString(starname)
	cgeopos := safeDoubleInput(geopos, 3)
	ctret, cleanupTret := safeDoubleOutput(tret, 10)
	if cleanupTret != nil { defer cleanupTret() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_rise_trans_true_hor(C.double(tjdUt), C.int32(ipl), cstarname, C.int32(epheflag), C.int32(rsmi), cgeopos, C.double(atpress), C.double(attemp), C.double(horhgt), ctret, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// RiseTrans function as declared in swephexp.h:918
func RiseTrans(tjdUt float64, ipl int, starname []byte, epheflag int, rsmi int, geopos []float64, atpress float64, attemp float64, tret []float64, serr []byte) int32 {
	cstarname := safeCString(starname)
	cgeopos := safeDoubleInput(geopos, 3)
	ctret, cleanupTret := safeDoubleOutput(tret, 10)
	if cleanupTret != nil { defer cleanupTret() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_rise_trans(C.double(tjdUt), C.int32(ipl), cstarname, C.int32(epheflag), C.int32(rsmi), cgeopos, C.double(atpress), C.double(attemp), ctret, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// NodAps function as declared in swephexp.h:926
func NodAps(tjdEt float64, ipl int, iflag int, method int, xnasc []float64, xndsc []float64, xperi []float64, xaphe []float64, serr []byte) int32 {
	cxnasc, cleanupNasc := safeDoubleOutput(xnasc, 6)
	if cleanupNasc != nil { defer cleanupNasc() }
	cxndsc, cleanupNdsc := safeDoubleOutput(xndsc, 6)
	if cleanupNdsc != nil { defer cleanupNdsc() }
	cxperi, cleanupPeri := safeDoubleOutput(xperi, 6)
	if cleanupPeri != nil { defer cleanupPeri() }
	cxaphe, cleanupAphe := safeDoubleOutput(xaphe, 6)
	if cleanupAphe != nil { defer cleanupAphe() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_nod_aps(C.double(tjdEt), C.int32(ipl), C.int32(iflag), C.int32(method), cxnasc, cxndsc, cxperi, cxaphe, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// NodApsUt function as declared in swephexp.h:932
func NodApsUt(tjdUt float64, ipl int, iflag int, method int, xnasc []float64, xndsc []float64, xperi []float64, xaphe []float64, serr []byte) int32 {
	cxnasc, cleanupNasc := safeDoubleOutput(xnasc, 6)
	if cleanupNasc != nil { defer cleanupNasc() }
	cxndsc, cleanupNdsc := safeDoubleOutput(xndsc, 6)
	if cleanupNdsc != nil { defer cleanupNdsc() }
	cxperi, cleanupPeri := safeDoubleOutput(xperi, 6)
	if cleanupPeri != nil { defer cleanupPeri() }
	cxaphe, cleanupAphe := safeDoubleOutput(xaphe, 6)
	if cleanupAphe != nil { defer cleanupAphe() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_nod_aps_ut(C.double(tjdUt), C.int32(ipl), C.int32(iflag), C.int32(method), cxnasc, cxndsc, cxperi, cxaphe, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// GetOrbitalElements function as declared in swephexp.h:937
func GetOrbitalElements(tjdEt float64, ipl int, iflag int, dret []float64, serr []byte) int32 {
	cdret, cleanupRet := safeDoubleOutput(dret, 50)
	if cleanupRet != nil { defer cleanupRet() }
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_get_orbital_elements(C.double(tjdEt), C.int32(ipl), C.int32(iflag), cdret, pserr)
	goseMutex.Unlock()
	return int32(ret)
}

// OrbitMaxMinTrueDistance function as declared in swephexp.h:940
func OrbitMaxMinTrueDistance(tjdEt float64, ipl int, iflag int, dmax []float64, dmin []float64, dtrue []float64, serr []byte) int32 {
	var cdmax, cdmin, cdtrue C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_orbit_max_min_true_distance(C.double(tjdEt), C.int32(ipl), C.int32(iflag), &cdmax, &cdmin, &cdtrue, pserr)
	goseMutex.Unlock()
	if len(dmax) > 0 { dmax[0] = float64(cdmax) }
	if len(dmin) > 0 { dmin[0] = float64(cdmin) }
	if len(dtrue) > 0 { dtrue[0] = float64(cdtrue) }
	return int32(ret)
}

// Deltat function as declared in swephexp.h:947
func Deltat(tjd float64) float64 {
	goseMutex.Lock()
	ret := C.swe_deltat(C.double(tjd))
	goseMutex.Unlock()
	return float64(ret)
}

// DeltatEx function as declared in swephexp.h:948
func DeltatEx(tjd float64, iflag int, serr []byte) float64 {
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_deltat_ex(C.double(tjd), C.int32(iflag), pserr)
	goseMutex.Unlock()
	return float64(ret)
}

// TimeEqu function as declared in swephexp.h:951
func TimeEqu(tjd float64, te []float64, serr []byte) int32 {
	var cte C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_time_equ(C.double(tjd), &cte, pserr)
	goseMutex.Unlock()
	if len(te) > 0 { te[0] = float64(cte) }
	return int32(ret)
}

// LmtToLat function as declared in swephexp.h:952
func LmtToLat(tjdLmt float64, geolon float64, tjdLat []float64, serr []byte) int32 {
	var ctjdLat C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_lmt_to_lat(C.double(tjdLmt), C.double(geolon), &ctjdLat, pserr)
	goseMutex.Unlock()
	if len(tjdLat) > 0 { tjdLat[0] = float64(ctjdLat) }
	return int32(ret)
}

// LatToLmt function as declared in swephexp.h:953
func LatToLmt(tjdLat float64, geolon float64, tjdLmt []float64, serr []byte) int32 {
	var ctjdLmt C.double
	pserr, cleanupErr := safeErrorBuf(serr)
	if cleanupErr != nil { defer cleanupErr() }
	goseMutex.Lock()
	ret := C.swe_lat_to_lmt(C.double(tjdLat), C.double(geolon), &ctjdLmt, pserr)
	goseMutex.Unlock()
	if len(tjdLmt) > 0 { tjdLmt[0] = float64(ctjdLmt) }
	return int32(ret)
}

// Sidtime0 function as declared in swephexp.h:956
func Sidtime0(tjdUt float64, eps float64, nut float64) float64 {
	ret := C.swe_sidtime0(C.double(tjdUt), C.double(eps), C.double(nut))
	return float64(ret)
}

// Sidtime function as declared in swephexp.h:957
func Sidtime(tjdUt float64) float64 {
	goseMutex.Lock()
	ret := C.swe_sidtime(C.double(tjdUt))
	goseMutex.Unlock()
	return float64(ret)
}

// SetInterpolateNut function as declared in swephexp.h:958
func SetInterpolateNut(doInterpolate int32) {
	goseMutex.Lock()
	C.swe_set_interpolate_nut(C.AS_BOOL(doInterpolate))
	goseMutex.Unlock()
}

// Cotrans function as declared in swephexp.h:961
func Cotrans(xpo []float64, xpn []float64, eps float64) {
	cxpo := safeDoubleInput(xpo, 3)
	cxpn, cleanup := safeDoubleOutput(xpn, 3)
	if cleanup != nil { defer cleanup() }
	C.swe_cotrans(cxpo, cxpn, C.double(eps))
}

// CotransSp function as declared in swephexp.h:962
func CotransSp(xpo []float64, xpn []float64, eps float64) {
	cxpo := safeDoubleInput(xpo, 6)
	cxpn, cleanup := safeDoubleOutput(xpn, 6)
	if cleanup != nil { defer cleanup() }
	C.swe_cotrans_sp(cxpo, cxpn, C.double(eps))
}

// GetTidAcc function as declared in swephexp.h:965
func GetTidAcc() float64 {
	goseMutex.Lock()
	ret := C.swe_get_tid_acc()
	goseMutex.Unlock()
	return float64(ret)
}

// SetTidAcc function as declared in swephexp.h:966
func SetTidAcc(tAcc float64) {
	goseMutex.Lock()
	C.swe_set_tid_acc(C.double(tAcc))
	goseMutex.Unlock()
}

// SetDeltaTUserdef function as declared in swephexp.h:970
func SetDeltaTUserdef(dt float64) {
	goseMutex.Lock()
	C.swe_set_delta_t_userdef(C.double(dt))
	goseMutex.Unlock()
}

// Degnorm function as declared in swephexp.h:972
func Degnorm(x float64) float64 {
	return float64(C.swe_degnorm(C.double(x)))
}

// Radnorm function as declared in swephexp.h:973
func Radnorm(x float64) float64 {
	return float64(C.swe_radnorm(C.double(x)))
}

// RadMidp function as declared in swephexp.h:974
func RadMidp(x1 float64, x0 float64) float64 {
	return float64(C.swe_rad_midp(C.double(x1), C.double(x0)))
}

// DegMidp function as declared in swephexp.h:975
func DegMidp(x1 float64, x0 float64) float64 {
	return float64(C.swe_deg_midp(C.double(x1), C.double(x0)))
}

// SplitDeg function as declared in swephexp.h:977
func SplitDeg(ddeg float64, roundflag int, ideg []int, imin []int, isec []int, dsecfr []float64, isgn []int32) {
	var cdeg, cmin, csec, csgn C.int32
	var csecfr C.double
	C.swe_split_deg(C.double(ddeg), C.int32(roundflag), &cdeg, &cmin, &csec, &csecfr, &csgn)
	if len(ideg) > 0 { ideg[0] = int(cdeg) }
	if len(imin) > 0 { imin[0] = int(cmin) }
	if len(isec) > 0 { isec[0] = int(csec) }
	if len(dsecfr) > 0 { dsecfr[0] = float64(csecfr) }
	if len(isgn) > 0 { isgn[0] = int32(csgn) }
}

// Csnorm function as declared in swephexp.h:986
func Csnorm(p int32) int32 {
	return int32(C.swe_csnorm(C.centisec(p)))
}

// Difcsn function as declared in swephexp.h:989
func Difcsn(p1 int, p2 int32) int32 {
	return int32(C.swe_difcsn(C.centisec(p1), C.centisec(p2)))
}

// Difdegn function as declared in swephexp.h:991
func Difdegn(p1 float64, p2 float64) float64 {
	return float64(C.swe_difdegn(C.double(p1), C.double(p2)))
}

// Difcs2n function as declared in swephexp.h:994
func Difcs2n(p1 int, p2 int32) int32 {
	return int32(C.swe_difcs2n(C.centisec(p1), C.centisec(p2)))
}

// Difdeg2n function as declared in swephexp.h:996
func Difdeg2n(p1 float64, p2 float64) float64 {
	return float64(C.swe_difdeg2n(C.double(p1), C.double(p2)))
}

// Difrad2n function as declared in swephexp.h:997
func Difrad2n(p1 float64, p2 float64) float64 {
	return float64(C.swe_difrad2n(C.double(p1), C.double(p2)))
}

// Csroundsec function as declared in swephexp.h:1000
func Csroundsec(x int32) int32 {
	return int32(C.swe_csroundsec(C.centisec(x)))
}

// D2l function as declared in swephexp.h:1003
func D2l(x float64) int32 {
	return int32(C.swe_d2l(C.double(x)))
}

// DayOfWeek function as declared in swephexp.h:1006
func DayOfWeek(jd float64) int32 {
	return int32(C.swe_day_of_week(C.double(jd)))
}

// Cs2timestr function as declared in swephexp.h:1008
func Cs2timestr(t int, sep int, suppressZero int, a []byte) *byte {
	var buf [32]byte
	var ca *C.char
	if len(a) >= 32 {
		ca = (*C.char)(unsafe.Pointer(unsafe.SliceData(a)))
	} else {
		ca = (*C.char)(unsafe.Pointer(&buf[0]))
	}
	ret := C.swe_cs2timestr(C.CSEC(t), C.int(sep), C.AS_BOOL(suppressZero), ca)
	if len(a) > 0 && len(a) < 32 {
		copy(a, buf[:])
	}
	return cCharToBytePtr(ret)
}

// Cs2lonlatstr function as declared in swephexp.h:1010
func Cs2lonlatstr(t int, pchar byte, mchar byte, s []byte) *byte {
	var buf [32]byte
	var cs *C.char
	if len(s) >= 32 {
		cs = (*C.char)(unsafe.Pointer(unsafe.SliceData(s)))
	} else {
		cs = (*C.char)(unsafe.Pointer(&buf[0]))
	}
	ret := C.swe_cs2lonlatstr(C.CSEC(t), C.char(pchar), C.char(mchar), cs)
	if len(s) > 0 && len(s) < 32 {
		copy(s, buf[:])
	}
	return cCharToBytePtr(ret)
}

// Cs2degstr function as declared in swephexp.h:1012
func Cs2degstr(t int, a []byte) *byte {
	var buf [32]byte
	var ca *C.char
	if len(a) >= 32 {
		ca = (*C.char)(unsafe.Pointer(unsafe.SliceData(a)))
	} else {
		ca = (*C.char)(unsafe.Pointer(&buf[0]))
	}
	ret := C.swe_cs2degstr(C.CSEC(t), ca)
	if len(a) > 0 && len(a) < 32 {
		copy(a, buf[:])
	}
	return cCharToBytePtr(ret)
}

// --- Idiomatic High-Level Go API Helpers ---

// VersionStr returns the Swiss Ephemeris library version as a Go string.
func VersionStr() string {
	var buf [256]byte
	ret := C.swe_version((*C.char)(unsafe.Pointer(&buf[0])))
	return cCharToString(ret)
}

// SetEphePathStr sets the ephemeris path using a Go string.
func SetEphePathStr(path string) {
	cpath := stringToCString(path)
	goseMutex.Lock()
	C.swe_set_ephe_path(cpath)
	goseMutex.Unlock()
}

// SetJplFileStr sets the JPL ephemeris file name using a Go string.
func SetJplFileStr(fname string) {
	cfname := stringToCString(fname)
	goseMutex.Lock()
	C.swe_set_jpl_file(cfname)
	goseMutex.Unlock()
}

// GetPlanetNameStr returns the planet name directly as a Go string.
func GetPlanetNameStr(ipl int) string {
	var buf [256]byte
	goseMutex.Lock()
	ret := C.swe_get_planet_name(C.int(ipl), (*C.char)(unsafe.Pointer(&buf[0])))
	goseMutex.Unlock()
	return cCharToString(ret)
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
	var cy, cm, cd C.int
	var cjut C.double
	C.swe_revjul(C.double(jd), C.int(gregflag), &cy, &cm, &cd, &cjut)
	y := int(cy)
	m := int(cm)
	d := int(cd)
	ut := float64(cjut)

	hour := int(ut)
	remainder := (ut - float64(hour)) * 60.0
	min := int(remainder)
	remainder = (remainder - float64(min)) * 60.0
	sec := int(remainder)
	nsec := int((remainder - float64(sec)) * 1e9)

	return time.Date(y, time.Month(m), d, hour, min, sec, nsec, time.UTC)
}

// CalcUtSimple provides an ergonomic wrapper around CalcUt, returning a structured Coordinates object.
func CalcUtSimple(tjdUt float64, ipl int, iflag int) (pos Coordinates, iflgret int32, err error) {
	var xx [6]float64
	var serr [256]byte
	iflgret = CalcUt(tjdUt, ipl, iflag, xx[:], serr[:])
	if iflgret < 0 {
		return Coordinates{}, iflgret, errors.New(cBytesToString(serr[:]))
	}
	pos = Coordinates{
		Longitude: xx[0],
		Latitude:  xx[1],
		Distance:  xx[2],
		SpeedLong: xx[3],
		SpeedLat:  xx[4],
		SpeedDist: xx[5],
	}
	return pos, iflgret, nil
}

// HousesSimple provides an ergonomic wrapper around Houses, returning a structured HousesResult.
func HousesSimple(tjdUt float64, iflag int, geolat, geolon float64, hsys int) (res HousesResult, iflgret int32, err error) {
	var cusps [13]float64
	var ascmc [10]float64
	iflgret = HousesEx(tjdUt, iflag, geolat, geolon, hsys, cusps[:], ascmc[:])
	if iflgret < 0 {
		return HousesResult{}, iflgret, errors.New("houses calculation failed")
	}
	res = HousesResult{
		Cusps:     cusps,
		Ascendant: ascmc[0],
		MC:        ascmc[1],
		ARMC:      ascmc[2],
		Vertex:    ascmc[3],
		Equasc:    ascmc[4],
		Coasc1:    ascmc[5],
		Coasc2:    ascmc[6],
		Polasc:    ascmc[7],
	}
	return res, iflgret, nil
}
