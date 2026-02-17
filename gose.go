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
	"runtime"
	"sync"
	"unsafe"
)

var goseMutex sync.Mutex

// HeliacalUt function as declared in swephexp.h:676
func HeliacalUt(tjdstartUt float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, typeEvent int, iflag int, dret []float64, serr []byte) int32 {
	ctjdstartUt, ctjdstartUtAllocMap := (C.double)(tjdstartUt), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	cdatm, cdatmAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&datm)))
	cdobs, cdobsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dobs)))
	cobjectName, cobjectNameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&objectName)))
	ctypeEvent, ctypeEventAllocMap := (C.int32)(typeEvent), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cdret, cdretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_heliacal_ut(ctjdstartUt, cgeopos, cdatm, cdobs, cobjectName, ctypeEvent, ciflag, cdret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdretAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctypeEventAllocMap)
	runtime.KeepAlive(cobjectNameAllocMap)
	runtime.KeepAlive(cdobsAllocMap)
	runtime.KeepAlive(cdatmAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ctjdstartUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// HeliacalPhenoUt function as declared in swephexp.h:677
func HeliacalPhenoUt(tjdUt float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, typeEvent int, helflag int, darr []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	cdatm, cdatmAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&datm)))
	cdobs, cdobsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dobs)))
	cobjectName, cobjectNameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&objectName)))
	ctypeEvent, ctypeEventAllocMap := (C.int32)(typeEvent), cgoAllocsUnknown
	chelflag, chelflagAllocMap := (C.int32)(helflag), cgoAllocsUnknown
	cdarr, cdarrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&darr)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_heliacal_pheno_ut(ctjdUt, cgeopos, cdatm, cdobs, cobjectName, ctypeEvent, chelflag, cdarr, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdarrAllocMap)
	runtime.KeepAlive(chelflagAllocMap)
	runtime.KeepAlive(ctypeEventAllocMap)
	runtime.KeepAlive(cobjectNameAllocMap)
	runtime.KeepAlive(cdobsAllocMap)
	runtime.KeepAlive(cdatmAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// VisLimitMag function as declared in swephexp.h:678
func VisLimitMag(tjdut float64, geopos []float64, datm []float64, dobs []float64, objectName []byte, helflag int, dret []float64, serr []byte) int32 {
	ctjdut, ctjdutAllocMap := (C.double)(tjdut), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	cdatm, cdatmAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&datm)))
	cdobs, cdobsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dobs)))
	cobjectName, cobjectNameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&objectName)))
	chelflag, chelflagAllocMap := (C.int32)(helflag), cgoAllocsUnknown
	cdret, cdretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_vis_limit_mag(ctjdut, cgeopos, cdatm, cdobs, cobjectName, chelflag, cdret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdretAllocMap)
	runtime.KeepAlive(chelflagAllocMap)
	runtime.KeepAlive(cobjectNameAllocMap)
	runtime.KeepAlive(cdobsAllocMap)
	runtime.KeepAlive(cdatmAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ctjdutAllocMap)
	__v := (int32)(__ret)
	return __v
}

// HeliacalAngle function as declared in swephexp.h:681
func HeliacalAngle(tjdut float64, dgeo []float64, datm []float64, dobs []float64, helflag int, mag float64, aziObj float64, aziSun float64, aziMoon float64, altMoon float64, dret []float64, serr []byte) int32 {
	ctjdut, ctjdutAllocMap := (C.double)(tjdut), cgoAllocsUnknown
	cdgeo, cdgeoAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dgeo)))
	cdatm, cdatmAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&datm)))
	cdobs, cdobsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dobs)))
	chelflag, chelflagAllocMap := (C.int32)(helflag), cgoAllocsUnknown
	cmag, cmagAllocMap := (C.double)(mag), cgoAllocsUnknown
	caziObj, caziObjAllocMap := (C.double)(aziObj), cgoAllocsUnknown
	caziSun, caziSunAllocMap := (C.double)(aziSun), cgoAllocsUnknown
	caziMoon, caziMoonAllocMap := (C.double)(aziMoon), cgoAllocsUnknown
	caltMoon, caltMoonAllocMap := (C.double)(altMoon), cgoAllocsUnknown
	cdret, cdretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_heliacal_angle(ctjdut, cdgeo, cdatm, cdobs, chelflag, cmag, caziObj, caziSun, caziMoon, caltMoon, cdret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdretAllocMap)
	runtime.KeepAlive(caltMoonAllocMap)
	runtime.KeepAlive(caziMoonAllocMap)
	runtime.KeepAlive(caziSunAllocMap)
	runtime.KeepAlive(caziObjAllocMap)
	runtime.KeepAlive(cmagAllocMap)
	runtime.KeepAlive(chelflagAllocMap)
	runtime.KeepAlive(cdobsAllocMap)
	runtime.KeepAlive(cdatmAllocMap)
	runtime.KeepAlive(cdgeoAllocMap)
	runtime.KeepAlive(ctjdutAllocMap)
	__v := (int32)(__ret)
	return __v
}

// TopoArcusVisionis function as declared in swephexp.h:682
func TopoArcusVisionis(tjdut float64, dgeo []float64, datm []float64, dobs []float64, helflag int, mag float64, aziObj float64, altObj float64, aziSun float64, aziMoon float64, altMoon float64, dret []float64, serr []byte) int32 {
	ctjdut, ctjdutAllocMap := (C.double)(tjdut), cgoAllocsUnknown
	cdgeo, cdgeoAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dgeo)))
	cdatm, cdatmAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&datm)))
	cdobs, cdobsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dobs)))
	chelflag, chelflagAllocMap := (C.int32)(helflag), cgoAllocsUnknown
	cmag, cmagAllocMap := (C.double)(mag), cgoAllocsUnknown
	caziObj, caziObjAllocMap := (C.double)(aziObj), cgoAllocsUnknown
	caltObj, caltObjAllocMap := (C.double)(altObj), cgoAllocsUnknown
	caziSun, caziSunAllocMap := (C.double)(aziSun), cgoAllocsUnknown
	caziMoon, caziMoonAllocMap := (C.double)(aziMoon), cgoAllocsUnknown
	caltMoon, caltMoonAllocMap := (C.double)(altMoon), cgoAllocsUnknown
	cdret, cdretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_topo_arcus_visionis(ctjdut, cdgeo, cdatm, cdobs, chelflag, cmag, caziObj, caltObj, caziSun, caziMoon, caltMoon, cdret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdretAllocMap)
	runtime.KeepAlive(caltMoonAllocMap)
	runtime.KeepAlive(caziMoonAllocMap)
	runtime.KeepAlive(caziSunAllocMap)
	runtime.KeepAlive(caltObjAllocMap)
	runtime.KeepAlive(caziObjAllocMap)
	runtime.KeepAlive(cmagAllocMap)
	runtime.KeepAlive(chelflagAllocMap)
	runtime.KeepAlive(cdobsAllocMap)
	runtime.KeepAlive(cdatmAllocMap)
	runtime.KeepAlive(cdgeoAllocMap)
	runtime.KeepAlive(ctjdutAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SetAstroModels function as declared in swephexp.h:686
func SetAstroModels(samod []byte, iflag int32) {
	csamod, csamodAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&samod)))
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_astro_models(csamod, ciflag)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(csamodAllocMap)
}

// GetAstroModels function as declared in swephexp.h:687
func GetAstroModels(samod []byte, sdet []byte, iflag int32) {
	csamod, csamodAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&samod)))
	csdet, csdetAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&sdet)))
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_get_astro_models(csamod, csdet, ciflag)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(csdetAllocMap)
	runtime.KeepAlive(csamodAllocMap)
}

// Version function as declared in swephexp.h:693
func Version(arg0 []byte) *byte {
	carg0, carg0AllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&arg0)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_version(carg0)
	runtime.KeepAlive(carg0AllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// GetLibraryPath function as declared in swephexp.h:694
func GetLibraryPath(arg0 []byte) *byte {
	carg0, carg0AllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&arg0)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_library_path(carg0)
	runtime.KeepAlive(carg0AllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// Calc function as declared in swephexp.h:697
func Calc(tjd float64, ipl int, iflag int, xx []float64, serr []byte) int32 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int)(ipl), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cxx, cxxAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xx)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_calc(ctjd, cipl, ciflag, cxx, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxxAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// CalcUt function as declared in swephexp.h:702
func CalcUt(tjdUt float64, ipl int, iflag int, xx []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cxx, cxxAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xx)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_calc_ut(ctjdUt, cipl, ciflag, cxx, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxxAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// CalcPctr function as declared in swephexp.h:705
func CalcPctr(tjd float64, ipl int, iplctr int, iflag int, xxret []float64, serr []byte) int32 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	ciplctr, ciplctrAllocMap := (C.int32)(iplctr), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cxxret, cxxretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xxret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_calc_pctr(ctjd, cipl, ciplctr, ciflag, cxxret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxxretAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplctrAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Fixstar function as declared in swephexp.h:717
func Fixstar(star []byte, tjd float64, iflag int, xx []float64, serr []byte) int32 {
	cstar, cstarAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&star)))
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cxx, cxxAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xx)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_fixstar(cstar, ctjd, ciflag, cxx, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxxAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	runtime.KeepAlive(cstarAllocMap)
	__v := (int32)(__ret)
	return __v
}

// FixstarUt function as declared in swephexp.h:722
func FixstarUt(star []byte, tjdUt float64, iflag int, xx []float64, serr []byte) int32 {
	cstar, cstarAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&star)))
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cxx, cxxAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xx)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_fixstar_ut(cstar, ctjdUt, ciflag, cxx, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxxAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	runtime.KeepAlive(cstarAllocMap)
	__v := (int32)(__ret)
	return __v
}

// FixstarMag function as declared in swephexp.h:725
func FixstarMag(star []byte, mag []float64, serr []byte) int32 {
	cstar, cstarAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&star)))
	cmag, cmagAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&mag)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_fixstar_mag(cstar, cmag, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cmagAllocMap)
	runtime.KeepAlive(cstarAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Fixstar2 function as declared in swephexp.h:727
func Fixstar2(star []byte, tjd float64, iflag int, xx []float64, serr []byte) int32 {
	cstar, cstarAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&star)))
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cxx, cxxAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xx)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_fixstar2(cstar, ctjd, ciflag, cxx, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxxAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	runtime.KeepAlive(cstarAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Fixstar2Ut function as declared in swephexp.h:732
func Fixstar2Ut(star []byte, tjdUt float64, iflag int, xx []float64, serr []byte) int32 {
	cstar, cstarAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&star)))
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cxx, cxxAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xx)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_fixstar2_ut(cstar, ctjdUt, ciflag, cxx, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxxAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	runtime.KeepAlive(cstarAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Fixstar2Mag function as declared in swephexp.h:735
func Fixstar2Mag(star []byte, mag []float64, serr []byte) int32 {
	cstar, cstarAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&star)))
	cmag, cmagAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&mag)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_fixstar2_mag(cstar, cmag, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cmagAllocMap)
	runtime.KeepAlive(cstarAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Close function as declared in swephexp.h:738
func Close() {
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_close()
}

// SetEphePath function as declared in swephexp.h:741
func SetEphePath(path []byte) {
	cpath, cpathAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&path)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_ephe_path(cpath)
	runtime.KeepAlive(cpathAllocMap)
}

// SetJplFile function as declared in swephexp.h:744
func SetJplFile(fname []byte) {
	cfname, cfnameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&fname)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_jpl_file(cfname)
	runtime.KeepAlive(cfnameAllocMap)
}

// GetPlanetName function as declared in swephexp.h:747
func GetPlanetName(ipl int, spname []byte) *byte {
	cipl, ciplAllocMap := (C.int)(ipl), cgoAllocsUnknown
	cspname, cspnameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&spname)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_planet_name(cipl, cspname)
	runtime.KeepAlive(cspnameAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// SetTopo function as declared in swephexp.h:750
func SetTopo(geolon float64, geolat float64, geoalt float64) {
	cgeolon, cgeolonAllocMap := (C.double)(geolon), cgoAllocsUnknown
	cgeolat, cgeolatAllocMap := (C.double)(geolat), cgoAllocsUnknown
	cgeoalt, cgeoaltAllocMap := (C.double)(geoalt), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_topo(cgeolon, cgeolat, cgeoalt)
	runtime.KeepAlive(cgeoaltAllocMap)
	runtime.KeepAlive(cgeolatAllocMap)
	runtime.KeepAlive(cgeolonAllocMap)
}

// SetSidMode function as declared in swephexp.h:753
func SetSidMode(sidMode int, t0 float64, ayanT0 float64) {
	csidMode, csidModeAllocMap := (C.int32)(sidMode), cgoAllocsUnknown
	ct0, ct0AllocMap := (C.double)(t0), cgoAllocsUnknown
	cayanT0, cayanT0AllocMap := (C.double)(ayanT0), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_sid_mode(csidMode, ct0, cayanT0)
	runtime.KeepAlive(cayanT0AllocMap)
	runtime.KeepAlive(ct0AllocMap)
	runtime.KeepAlive(csidModeAllocMap)
}

// GetAyanamsaEx function as declared in swephexp.h:756
func GetAyanamsaEx(tjdEt float64, iflag int, daya []float64, serr []byte) int32 {
	ctjdEt, ctjdEtAllocMap := (C.double)(tjdEt), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cdaya, cdayaAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&daya)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_ayanamsa_ex(ctjdEt, ciflag, cdaya, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdayaAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdEtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetAyanamsaExUt function as declared in swephexp.h:757
func GetAyanamsaExUt(tjdUt float64, iflag int, daya []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cdaya, cdayaAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&daya)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_ayanamsa_ex_ut(ctjdUt, ciflag, cdaya, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdayaAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetAyanamsa function as declared in swephexp.h:758
func GetAyanamsa(tjdEt float64) float64 {
	ctjdEt, ctjdEtAllocMap := (C.double)(tjdEt), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_ayanamsa(ctjdEt)
	runtime.KeepAlive(ctjdEtAllocMap)
	__v := (float64)(__ret)
	return __v
}

// GetAyanamsaUt function as declared in swephexp.h:759
func GetAyanamsaUt(tjdUt float64) float64 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_ayanamsa_ut(ctjdUt)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (float64)(__ret)
	return __v
}

// GetAyanamsaName function as declared in swephexp.h:762
func GetAyanamsaName(isidmode int32) string {
	cisidmode, cisidmodeAllocMap := (C.int32)(isidmode), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_ayanamsa_name(cisidmode)
	runtime.KeepAlive(cisidmodeAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// GetCurrentFileData function as declared in swephexp.h:763
func GetCurrentFileData(ifno int, tfstart []float64, tfend []float64, denum []int32) string {
	cifno, cifnoAllocMap := (C.int)(ifno), cgoAllocsUnknown
	ctfstart, ctfstartAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tfstart)))
	ctfend, ctfendAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tfend)))
	cdenum, cdenumAllocMap := copyPIntBytes((*sliceHeader)(unsafe.Pointer(&denum)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_current_file_data(cifno, ctfstart, ctfend, cdenum)
	runtime.KeepAlive(cdenumAllocMap)
	runtime.KeepAlive(ctfendAllocMap)
	runtime.KeepAlive(ctfstartAllocMap)
	runtime.KeepAlive(cifnoAllocMap)
	__v := packPCharString(__ret)
	return __v
}

// DateConversion function as declared in swephexp.h:771
func DateConversion(y int, m int, d int, utime float64, c byte, tjd []float64) int32 {
	cy, cyAllocMap := (C.int)(y), cgoAllocsUnknown
	cm, cmAllocMap := (C.int)(m), cgoAllocsUnknown
	cd, cdAllocMap := (C.int)(d), cgoAllocsUnknown
	cutime, cutimeAllocMap := (C.double)(utime), cgoAllocsUnknown
	cc, ccAllocMap := (C.char)(c), cgoAllocsUnknown
	ctjd, ctjdAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tjd)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_date_conversion(cy, cm, cd, cutime, cc, ctjd)
	runtime.KeepAlive(ctjdAllocMap)
	runtime.KeepAlive(ccAllocMap)
	runtime.KeepAlive(cutimeAllocMap)
	runtime.KeepAlive(cdAllocMap)
	runtime.KeepAlive(cmAllocMap)
	runtime.KeepAlive(cyAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Julday function as declared in swephexp.h:777
func Julday(year int, month int, day int, hour float64, gregflag int32) float64 {
	cyear, cyearAllocMap := (C.int)(year), cgoAllocsUnknown
	cmonth, cmonthAllocMap := (C.int)(month), cgoAllocsUnknown
	cday, cdayAllocMap := (C.int)(day), cgoAllocsUnknown
	chour, chourAllocMap := (C.double)(hour), cgoAllocsUnknown
	cgregflag, cgregflagAllocMap := (C.int)(gregflag), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_julday(cyear, cmonth, cday, chour, cgregflag)
	runtime.KeepAlive(cgregflagAllocMap)
	runtime.KeepAlive(chourAllocMap)
	runtime.KeepAlive(cdayAllocMap)
	runtime.KeepAlive(cmonthAllocMap)
	runtime.KeepAlive(cyearAllocMap)
	__v := (float64)(__ret)
	return __v
}

// Revjul function as declared in swephexp.h:781
func Revjul(jd float64, gregflag int, jyear []int, jmon []int, jday []int, jut []float64) {
	cjd, cjdAllocMap := (C.double)(jd), cgoAllocsUnknown
	cgregflag, cgregflagAllocMap := (C.int)(gregflag), cgoAllocsUnknown
	cjyear, cjyearAllocMap := copyPIntBytes((*sliceHeader)(unsafe.Pointer(&jyear)))
	cjmon, cjmonAllocMap := copyPIntBytes((*sliceHeader)(unsafe.Pointer(&jmon)))
	cjday, cjdayAllocMap := copyPIntBytes((*sliceHeader)(unsafe.Pointer(&jday)))
	cjut, cjutAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&jut)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_revjul(cjd, cgregflag, cjyear, cjmon, cjday, cjut)
	runtime.KeepAlive(cjutAllocMap)
	runtime.KeepAlive(cjdayAllocMap)
	runtime.KeepAlive(cjmonAllocMap)
	runtime.KeepAlive(cjyearAllocMap)
	runtime.KeepAlive(cgregflagAllocMap)
	runtime.KeepAlive(cjdAllocMap)
}

// UtcToJd function as declared in swephexp.h:786
func UtcToJd(iyear int, imonth int, iday int, ihour int, imin int, dsec float64, gregflag int, dret []float64, serr []byte) int32 {
	ciyear, ciyearAllocMap := (C.int32)(iyear), cgoAllocsUnknown
	cimonth, cimonthAllocMap := (C.int32)(imonth), cgoAllocsUnknown
	ciday, cidayAllocMap := (C.int32)(iday), cgoAllocsUnknown
	cihour, cihourAllocMap := (C.int32)(ihour), cgoAllocsUnknown
	cimin, ciminAllocMap := (C.int32)(imin), cgoAllocsUnknown
	cdsec, cdsecAllocMap := (C.double)(dsec), cgoAllocsUnknown
	cgregflag, cgregflagAllocMap := (C.int32)(gregflag), cgoAllocsUnknown
	cdret, cdretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_utc_to_jd(ciyear, cimonth, ciday, cihour, cimin, cdsec, cgregflag, cdret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdretAllocMap)
	runtime.KeepAlive(cgregflagAllocMap)
	runtime.KeepAlive(cdsecAllocMap)
	runtime.KeepAlive(ciminAllocMap)
	runtime.KeepAlive(cihourAllocMap)
	runtime.KeepAlive(cidayAllocMap)
	runtime.KeepAlive(cimonthAllocMap)
	runtime.KeepAlive(ciyearAllocMap)
	__v := (int32)(__ret)
	return __v
}

// JdetToUtc function as declared in swephexp.h:791
func JdetToUtc(tjdEt float64, gregflag int, iyear []int, imonth []int, iday []int, ihour []int, imin []int, dsec []float64) {
	ctjdEt, ctjdEtAllocMap := (C.double)(tjdEt), cgoAllocsUnknown
	cgregflag, cgregflagAllocMap := (C.int32)(gregflag), cgoAllocsUnknown
	ciyear, ciyearAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&iyear)))
	cimonth, cimonthAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&imonth)))
	ciday, cidayAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&iday)))
	cihour, cihourAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&ihour)))
	cimin, ciminAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&imin)))
	cdsec, cdsecAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dsec)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_jdet_to_utc(ctjdEt, cgregflag, ciyear, cimonth, ciday, cihour, cimin, cdsec)
	runtime.KeepAlive(cdsecAllocMap)
	runtime.KeepAlive(ciminAllocMap)
	runtime.KeepAlive(cihourAllocMap)
	runtime.KeepAlive(cidayAllocMap)
	runtime.KeepAlive(cimonthAllocMap)
	runtime.KeepAlive(ciyearAllocMap)
	runtime.KeepAlive(cgregflagAllocMap)
	runtime.KeepAlive(ctjdEtAllocMap)
}

// Jdut1ToUtc function as declared in swephexp.h:796
func Jdut1ToUtc(tjdUt float64, gregflag int, iyear []int, imonth []int, iday []int, ihour []int, imin []int, dsec []float64) {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cgregflag, cgregflagAllocMap := (C.int32)(gregflag), cgoAllocsUnknown
	ciyear, ciyearAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&iyear)))
	cimonth, cimonthAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&imonth)))
	ciday, cidayAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&iday)))
	cihour, cihourAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&ihour)))
	cimin, ciminAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&imin)))
	cdsec, cdsecAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dsec)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_jdut1_to_utc(ctjdUt, cgregflag, ciyear, cimonth, ciday, cihour, cimin, cdsec)
	runtime.KeepAlive(cdsecAllocMap)
	runtime.KeepAlive(ciminAllocMap)
	runtime.KeepAlive(cihourAllocMap)
	runtime.KeepAlive(cidayAllocMap)
	runtime.KeepAlive(cimonthAllocMap)
	runtime.KeepAlive(ciyearAllocMap)
	runtime.KeepAlive(cgregflagAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
}

// UtcTimeZone function as declared in swephexp.h:801
func UtcTimeZone(iyear int, imonth int, iday int, ihour int, imin int, dsec float64, dTimezone float64, iyearOut []int, imonthOut []int, idayOut []int, ihourOut []int, iminOut []int, dsecOut []float64) {
	ciyear, ciyearAllocMap := (C.int32)(iyear), cgoAllocsUnknown
	cimonth, cimonthAllocMap := (C.int32)(imonth), cgoAllocsUnknown
	ciday, cidayAllocMap := (C.int32)(iday), cgoAllocsUnknown
	cihour, cihourAllocMap := (C.int32)(ihour), cgoAllocsUnknown
	cimin, ciminAllocMap := (C.int32)(imin), cgoAllocsUnknown
	cdsec, cdsecAllocMap := (C.double)(dsec), cgoAllocsUnknown
	cdTimezone, cdTimezoneAllocMap := (C.double)(dTimezone), cgoAllocsUnknown
	ciyearOut, ciyearOutAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&iyearOut)))
	cimonthOut, cimonthOutAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&imonthOut)))
	cidayOut, cidayOutAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&idayOut)))
	cihourOut, cihourOutAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&ihourOut)))
	ciminOut, ciminOutAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&iminOut)))
	cdsecOut, cdsecOutAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dsecOut)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_utc_time_zone(ciyear, cimonth, ciday, cihour, cimin, cdsec, cdTimezone, ciyearOut, cimonthOut, cidayOut, cihourOut, ciminOut, cdsecOut)
	runtime.KeepAlive(cdsecOutAllocMap)
	runtime.KeepAlive(ciminOutAllocMap)
	runtime.KeepAlive(cihourOutAllocMap)
	runtime.KeepAlive(cidayOutAllocMap)
	runtime.KeepAlive(cimonthOutAllocMap)
	runtime.KeepAlive(ciyearOutAllocMap)
	runtime.KeepAlive(cdTimezoneAllocMap)
	runtime.KeepAlive(cdsecAllocMap)
	runtime.KeepAlive(ciminAllocMap)
	runtime.KeepAlive(cihourAllocMap)
	runtime.KeepAlive(cidayAllocMap)
	runtime.KeepAlive(cimonthAllocMap)
	runtime.KeepAlive(ciyearAllocMap)
}

// Houses function as declared in swephexp.h:812
func Houses(tjdUt float64, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cgeolat, cgeolatAllocMap := (C.double)(geolat), cgoAllocsUnknown
	cgeolon, cgeolonAllocMap := (C.double)(geolon), cgoAllocsUnknown
	chsys, chsysAllocMap := (C.int)(hsys), cgoAllocsUnknown
	ccusps, ccuspsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&cusps)))
	cascmc, cascmcAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&ascmc)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_houses(ctjdUt, cgeolat, cgeolon, chsys, ccusps, cascmc)
	runtime.KeepAlive(cascmcAllocMap)
	runtime.KeepAlive(ccuspsAllocMap)
	runtime.KeepAlive(chsysAllocMap)
	runtime.KeepAlive(cgeolonAllocMap)
	runtime.KeepAlive(cgeolatAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// HousesEx function as declared in swephexp.h:816
func HousesEx(tjdUt float64, iflag int, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cgeolat, cgeolatAllocMap := (C.double)(geolat), cgoAllocsUnknown
	cgeolon, cgeolonAllocMap := (C.double)(geolon), cgoAllocsUnknown
	chsys, chsysAllocMap := (C.int)(hsys), cgoAllocsUnknown
	ccusps, ccuspsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&cusps)))
	cascmc, cascmcAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&ascmc)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_houses_ex(ctjdUt, ciflag, cgeolat, cgeolon, chsys, ccusps, cascmc)
	runtime.KeepAlive(cascmcAllocMap)
	runtime.KeepAlive(ccuspsAllocMap)
	runtime.KeepAlive(chsysAllocMap)
	runtime.KeepAlive(cgeolonAllocMap)
	runtime.KeepAlive(cgeolatAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// HousesEx2 function as declared in swephexp.h:820
func HousesEx2(tjdUt float64, iflag int, geolat float64, geolon float64, hsys int, cusps []float64, ascmc []float64, cuspSpeed []float64, ascmcSpeed []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cgeolat, cgeolatAllocMap := (C.double)(geolat), cgoAllocsUnknown
	cgeolon, cgeolonAllocMap := (C.double)(geolon), cgoAllocsUnknown
	chsys, chsysAllocMap := (C.int)(hsys), cgoAllocsUnknown
	ccusps, ccuspsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&cusps)))
	cascmc, cascmcAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&ascmc)))
	ccuspSpeed, ccuspSpeedAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&cuspSpeed)))
	cascmcSpeed, cascmcSpeedAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&ascmcSpeed)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_houses_ex2(ctjdUt, ciflag, cgeolat, cgeolon, chsys, ccusps, cascmc, ccuspSpeed, cascmcSpeed, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cascmcSpeedAllocMap)
	runtime.KeepAlive(ccuspSpeedAllocMap)
	runtime.KeepAlive(cascmcAllocMap)
	runtime.KeepAlive(ccuspsAllocMap)
	runtime.KeepAlive(chsysAllocMap)
	runtime.KeepAlive(cgeolonAllocMap)
	runtime.KeepAlive(cgeolatAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// HousesArmc function as declared in swephexp.h:824
func HousesArmc(armc float64, geolat float64, eps float64, hsys int, cusps []float64, ascmc []float64) int32 {
	carmc, carmcAllocMap := (C.double)(armc), cgoAllocsUnknown
	cgeolat, cgeolatAllocMap := (C.double)(geolat), cgoAllocsUnknown
	ceps, cepsAllocMap := (C.double)(eps), cgoAllocsUnknown
	chsys, chsysAllocMap := (C.int)(hsys), cgoAllocsUnknown
	ccusps, ccuspsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&cusps)))
	cascmc, cascmcAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&ascmc)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_houses_armc(carmc, cgeolat, ceps, chsys, ccusps, cascmc)
	runtime.KeepAlive(cascmcAllocMap)
	runtime.KeepAlive(ccuspsAllocMap)
	runtime.KeepAlive(chsysAllocMap)
	runtime.KeepAlive(cepsAllocMap)
	runtime.KeepAlive(cgeolatAllocMap)
	runtime.KeepAlive(carmcAllocMap)
	__v := (int32)(__ret)
	return __v
}

// HousesArmcEx2 function as declared in swephexp.h:828
func HousesArmcEx2(armc float64, geolat float64, eps float64, hsys int, cusps []float64, ascmc []float64, cuspSpeed []float64, ascmcSpeed []float64, serr []byte) int32 {
	carmc, carmcAllocMap := (C.double)(armc), cgoAllocsUnknown
	cgeolat, cgeolatAllocMap := (C.double)(geolat), cgoAllocsUnknown
	ceps, cepsAllocMap := (C.double)(eps), cgoAllocsUnknown
	chsys, chsysAllocMap := (C.int)(hsys), cgoAllocsUnknown
	ccusps, ccuspsAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&cusps)))
	cascmc, cascmcAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&ascmc)))
	ccuspSpeed, ccuspSpeedAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&cuspSpeed)))
	cascmcSpeed, cascmcSpeedAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&ascmcSpeed)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_houses_armc_ex2(carmc, cgeolat, ceps, chsys, ccusps, cascmc, ccuspSpeed, cascmcSpeed, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cascmcSpeedAllocMap)
	runtime.KeepAlive(ccuspSpeedAllocMap)
	runtime.KeepAlive(cascmcAllocMap)
	runtime.KeepAlive(ccuspsAllocMap)
	runtime.KeepAlive(chsysAllocMap)
	runtime.KeepAlive(cepsAllocMap)
	runtime.KeepAlive(cgeolatAllocMap)
	runtime.KeepAlive(carmcAllocMap)
	__v := (int32)(__ret)
	return __v
}

// HousePos function as declared in swephexp.h:832
func HousePos(armc float64, geolat float64, eps float64, hsys int, xpin []float64, serr []byte) float64 {
	carmc, carmcAllocMap := (C.double)(armc), cgoAllocsUnknown
	cgeolat, cgeolatAllocMap := (C.double)(geolat), cgoAllocsUnknown
	ceps, cepsAllocMap := (C.double)(eps), cgoAllocsUnknown
	chsys, chsysAllocMap := (C.int)(hsys), cgoAllocsUnknown
	cxpin, cxpinAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xpin)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_house_pos(carmc, cgeolat, ceps, chsys, cxpin, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxpinAllocMap)
	runtime.KeepAlive(chsysAllocMap)
	runtime.KeepAlive(cepsAllocMap)
	runtime.KeepAlive(cgeolatAllocMap)
	runtime.KeepAlive(carmcAllocMap)
	__v := (float64)(__ret)
	return __v
}

// HouseName function as declared in swephexp.h:835
func HouseName(hsys int32) *byte {
	chsys, chsysAllocMap := (C.int)(hsys), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_house_name(chsys)
	runtime.KeepAlive(chsysAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// GauquelinSector function as declared in swephexp.h:843
func GauquelinSector(tUt float64, ipl int, starname []byte, iflag int, imeth int, geopos []float64, atpress float64, attemp float64, dgsect []float64, serr []byte) int32 {
	ctUt, ctUtAllocMap := (C.double)(tUt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	cstarname, cstarnameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&starname)))
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cimeth, cimethAllocMap := (C.int32)(imeth), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	catpress, catpressAllocMap := (C.double)(atpress), cgoAllocsUnknown
	cattemp, cattempAllocMap := (C.double)(attemp), cgoAllocsUnknown
	cdgsect, cdgsectAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dgsect)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_gauquelin_sector(ctUt, cipl, cstarname, ciflag, cimeth, cgeopos, catpress, cattemp, cdgsect, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdgsectAllocMap)
	runtime.KeepAlive(cattempAllocMap)
	runtime.KeepAlive(catpressAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(cimethAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(cstarnameAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SolEclipseWhere function as declared in swephexp.h:847
func SolEclipseWhere(tjd float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_sol_eclipse_where(ctjd, cifl, cgeopos, cattr, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// LunOccultWhere function as declared in swephexp.h:849
func LunOccultWhere(tjd float64, ipl int, starname []byte, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	cstarname, cstarnameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&starname)))
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_lun_occult_where(ctjd, cipl, cstarname, cifl, cgeopos, cattr, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(cstarnameAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SolEclipseHow function as declared in swephexp.h:852
func SolEclipseHow(tjd float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_sol_eclipse_how(ctjd, cifl, cgeopos, cattr, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SolEclipseWhenLoc function as declared in swephexp.h:855
func SolEclipseWhenLoc(tjdStart float64, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	ctjdStart, ctjdStartAllocMap := (C.double)(tjdStart), cgoAllocsUnknown
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	ctret, ctretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tret)))
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cbackward, cbackwardAllocMap := (C.int32)(backward), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_sol_eclipse_when_loc(ctjdStart, cifl, cgeopos, ctret, cattr, cbackward, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cbackwardAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(ctretAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(ctjdStartAllocMap)
	__v := (int32)(__ret)
	return __v
}

// LunOccultWhenLoc function as declared in swephexp.h:857
func LunOccultWhenLoc(tjdStart float64, ipl int, starname []byte, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	ctjdStart, ctjdStartAllocMap := (C.double)(tjdStart), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	cstarname, cstarnameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&starname)))
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	ctret, ctretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tret)))
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cbackward, cbackwardAllocMap := (C.int32)(backward), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_lun_occult_when_loc(ctjdStart, cipl, cstarname, cifl, cgeopos, ctret, cattr, cbackward, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cbackwardAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(ctretAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(cstarnameAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdStartAllocMap)
	__v := (int32)(__ret)
	return __v
}

// SolEclipseWhenGlob function as declared in swephexp.h:861
func SolEclipseWhenGlob(tjdStart float64, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	ctjdStart, ctjdStartAllocMap := (C.double)(tjdStart), cgoAllocsUnknown
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cifltype, cifltypeAllocMap := (C.int32)(ifltype), cgoAllocsUnknown
	ctret, ctretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tret)))
	cbackward, cbackwardAllocMap := (C.int32)(backward), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_sol_eclipse_when_glob(ctjdStart, cifl, cifltype, ctret, cbackward, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cbackwardAllocMap)
	runtime.KeepAlive(ctretAllocMap)
	runtime.KeepAlive(cifltypeAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(ctjdStartAllocMap)
	__v := (int32)(__ret)
	return __v
}

// LunOccultWhenGlob function as declared in swephexp.h:865
func LunOccultWhenGlob(tjdStart float64, ipl int, starname []byte, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	ctjdStart, ctjdStartAllocMap := (C.double)(tjdStart), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	cstarname, cstarnameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&starname)))
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cifltype, cifltypeAllocMap := (C.int32)(ifltype), cgoAllocsUnknown
	ctret, ctretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tret)))
	cbackward, cbackwardAllocMap := (C.int32)(backward), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_lun_occult_when_glob(ctjdStart, cipl, cstarname, cifl, cifltype, ctret, cbackward, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cbackwardAllocMap)
	runtime.KeepAlive(ctretAllocMap)
	runtime.KeepAlive(cifltypeAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(cstarnameAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdStartAllocMap)
	__v := (int32)(__ret)
	return __v
}

// LunEclipseHow function as declared in swephexp.h:869
func LunEclipseHow(tjdUt float64, ifl int, geopos []float64, attr []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_lun_eclipse_how(ctjdUt, cifl, cgeopos, cattr, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// LunEclipseWhen function as declared in swephexp.h:876
func LunEclipseWhen(tjdStart float64, ifl int, ifltype int, tret []float64, backward int, serr []byte) int32 {
	ctjdStart, ctjdStartAllocMap := (C.double)(tjdStart), cgoAllocsUnknown
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cifltype, cifltypeAllocMap := (C.int32)(ifltype), cgoAllocsUnknown
	ctret, ctretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tret)))
	cbackward, cbackwardAllocMap := (C.int32)(backward), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_lun_eclipse_when(ctjdStart, cifl, cifltype, ctret, cbackward, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cbackwardAllocMap)
	runtime.KeepAlive(ctretAllocMap)
	runtime.KeepAlive(cifltypeAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(ctjdStartAllocMap)
	__v := (int32)(__ret)
	return __v
}

// LunEclipseWhenLoc function as declared in swephexp.h:879
func LunEclipseWhenLoc(tjdStart float64, ifl int, geopos []float64, tret []float64, attr []float64, backward int, serr []byte) int32 {
	ctjdStart, ctjdStartAllocMap := (C.double)(tjdStart), cgoAllocsUnknown
	cifl, ciflAllocMap := (C.int32)(ifl), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	ctret, ctretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tret)))
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cbackward, cbackwardAllocMap := (C.int32)(backward), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_lun_eclipse_when_loc(ctjdStart, cifl, cgeopos, ctret, cattr, cbackward, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cbackwardAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(ctretAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ciflAllocMap)
	runtime.KeepAlive(ctjdStartAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Pheno function as declared in swephexp.h:883
func Pheno(tjd float64, ipl int, iflag int, attr []float64, serr []byte) int32 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_pheno(ctjd, cipl, ciflag, cattr, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// PhenoUt function as declared in swephexp.h:885
func PhenoUt(tjdUt float64, ipl int, iflag int, attr []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cattr, cattrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&attr)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_pheno_ut(ctjdUt, cipl, ciflag, cattr, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cattrAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Refrac function as declared in swephexp.h:887
func Refrac(inalt float64, atpress float64, attemp float64, calcFlag int32) float64 {
	cinalt, cinaltAllocMap := (C.double)(inalt), cgoAllocsUnknown
	catpress, catpressAllocMap := (C.double)(atpress), cgoAllocsUnknown
	cattemp, cattempAllocMap := (C.double)(attemp), cgoAllocsUnknown
	ccalcFlag, ccalcFlagAllocMap := (C.int32)(calcFlag), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_refrac(cinalt, catpress, cattemp, ccalcFlag)
	runtime.KeepAlive(ccalcFlagAllocMap)
	runtime.KeepAlive(cattempAllocMap)
	runtime.KeepAlive(catpressAllocMap)
	runtime.KeepAlive(cinaltAllocMap)
	__v := (float64)(__ret)
	return __v
}

// RefracExtended function as declared in swephexp.h:889
func RefracExtended(inalt float64, geoalt float64, atpress float64, attemp float64, lapseRate float64, calcFlag int, dret []float64) float64 {
	cinalt, cinaltAllocMap := (C.double)(inalt), cgoAllocsUnknown
	cgeoalt, cgeoaltAllocMap := (C.double)(geoalt), cgoAllocsUnknown
	catpress, catpressAllocMap := (C.double)(atpress), cgoAllocsUnknown
	cattemp, cattempAllocMap := (C.double)(attemp), cgoAllocsUnknown
	clapseRate, clapseRateAllocMap := (C.double)(lapseRate), cgoAllocsUnknown
	ccalcFlag, ccalcFlagAllocMap := (C.int32)(calcFlag), cgoAllocsUnknown
	cdret, cdretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dret)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_refrac_extended(cinalt, cgeoalt, catpress, cattemp, clapseRate, ccalcFlag, cdret)
	runtime.KeepAlive(cdretAllocMap)
	runtime.KeepAlive(ccalcFlagAllocMap)
	runtime.KeepAlive(clapseRateAllocMap)
	runtime.KeepAlive(cattempAllocMap)
	runtime.KeepAlive(catpressAllocMap)
	runtime.KeepAlive(cgeoaltAllocMap)
	runtime.KeepAlive(cinaltAllocMap)
	__v := (float64)(__ret)
	return __v
}

// SetLapseRate function as declared in swephexp.h:891
func SetLapseRate(lapseRate float64) {
	clapseRate, clapseRateAllocMap := (C.double)(lapseRate), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_lapse_rate(clapseRate)
	runtime.KeepAlive(clapseRateAllocMap)
}

// Azalt function as declared in swephexp.h:893
func Azalt(tjdUt float64, calcFlag int, geopos []float64, atpress float64, attemp float64, xin []float64, xaz []float64) {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	ccalcFlag, ccalcFlagAllocMap := (C.int32)(calcFlag), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	catpress, catpressAllocMap := (C.double)(atpress), cgoAllocsUnknown
	cattemp, cattempAllocMap := (C.double)(attemp), cgoAllocsUnknown
	cxin, cxinAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xin)))
	cxaz, cxazAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xaz)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_azalt(ctjdUt, ccalcFlag, cgeopos, catpress, cattemp, cxin, cxaz)
	runtime.KeepAlive(cxazAllocMap)
	runtime.KeepAlive(cxinAllocMap)
	runtime.KeepAlive(cattempAllocMap)
	runtime.KeepAlive(catpressAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ccalcFlagAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
}

// AzaltRev function as declared in swephexp.h:902
func AzaltRev(tjdUt float64, calcFlag int, geopos []float64, xin []float64, xout []float64) {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	ccalcFlag, ccalcFlagAllocMap := (C.int32)(calcFlag), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	cxin, cxinAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xin)))
	cxout, cxoutAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xout)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_azalt_rev(ctjdUt, ccalcFlag, cgeopos, cxin, cxout)
	runtime.KeepAlive(cxoutAllocMap)
	runtime.KeepAlive(cxinAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(ccalcFlagAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
}

// RiseTransTrueHor function as declared in swephexp.h:909
func RiseTransTrueHor(tjdUt float64, ipl int, starname []byte, epheflag int, rsmi int, geopos []float64, atpress float64, attemp float64, horhgt float64, tret []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	cstarname, cstarnameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&starname)))
	cepheflag, cepheflagAllocMap := (C.int32)(epheflag), cgoAllocsUnknown
	crsmi, crsmiAllocMap := (C.int32)(rsmi), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	catpress, catpressAllocMap := (C.double)(atpress), cgoAllocsUnknown
	cattemp, cattempAllocMap := (C.double)(attemp), cgoAllocsUnknown
	chorhgt, chorhgtAllocMap := (C.double)(horhgt), cgoAllocsUnknown
	ctret, ctretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_rise_trans_true_hor(ctjdUt, cipl, cstarname, cepheflag, crsmi, cgeopos, catpress, cattemp, chorhgt, ctret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(ctretAllocMap)
	runtime.KeepAlive(chorhgtAllocMap)
	runtime.KeepAlive(cattempAllocMap)
	runtime.KeepAlive(catpressAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(crsmiAllocMap)
	runtime.KeepAlive(cepheflagAllocMap)
	runtime.KeepAlive(cstarnameAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// RiseTrans function as declared in swephexp.h:918
func RiseTrans(tjdUt float64, ipl int, starname []byte, epheflag int, rsmi int, geopos []float64, atpress float64, attemp float64, tret []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	cstarname, cstarnameAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&starname)))
	cepheflag, cepheflagAllocMap := (C.int32)(epheflag), cgoAllocsUnknown
	crsmi, crsmiAllocMap := (C.int32)(rsmi), cgoAllocsUnknown
	cgeopos, cgeoposAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&geopos)))
	catpress, catpressAllocMap := (C.double)(atpress), cgoAllocsUnknown
	cattemp, cattempAllocMap := (C.double)(attemp), cgoAllocsUnknown
	ctret, ctretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_rise_trans(ctjdUt, cipl, cstarname, cepheflag, crsmi, cgeopos, catpress, cattemp, ctret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(ctretAllocMap)
	runtime.KeepAlive(cattempAllocMap)
	runtime.KeepAlive(catpressAllocMap)
	runtime.KeepAlive(cgeoposAllocMap)
	runtime.KeepAlive(crsmiAllocMap)
	runtime.KeepAlive(cepheflagAllocMap)
	runtime.KeepAlive(cstarnameAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// NodAps function as declared in swephexp.h:926
func NodAps(tjdEt float64, ipl int, iflag int, method int, xnasc []float64, xndsc []float64, xperi []float64, xaphe []float64, serr []byte) int32 {
	ctjdEt, ctjdEtAllocMap := (C.double)(tjdEt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cmethod, cmethodAllocMap := (C.int32)(method), cgoAllocsUnknown
	cxnasc, cxnascAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xnasc)))
	cxndsc, cxndscAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xndsc)))
	cxperi, cxperiAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xperi)))
	cxaphe, cxapheAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xaphe)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_nod_aps(ctjdEt, cipl, ciflag, cmethod, cxnasc, cxndsc, cxperi, cxaphe, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxapheAllocMap)
	runtime.KeepAlive(cxperiAllocMap)
	runtime.KeepAlive(cxndscAllocMap)
	runtime.KeepAlive(cxnascAllocMap)
	runtime.KeepAlive(cmethodAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdEtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// NodApsUt function as declared in swephexp.h:932
func NodApsUt(tjdUt float64, ipl int, iflag int, method int, xnasc []float64, xndsc []float64, xperi []float64, xaphe []float64, serr []byte) int32 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cmethod, cmethodAllocMap := (C.int32)(method), cgoAllocsUnknown
	cxnasc, cxnascAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xnasc)))
	cxndsc, cxndscAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xndsc)))
	cxperi, cxperiAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xperi)))
	cxaphe, cxapheAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xaphe)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_nod_aps_ut(ctjdUt, cipl, ciflag, cmethod, cxnasc, cxndsc, cxperi, cxaphe, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxapheAllocMap)
	runtime.KeepAlive(cxperiAllocMap)
	runtime.KeepAlive(cxndscAllocMap)
	runtime.KeepAlive(cxnascAllocMap)
	runtime.KeepAlive(cmethodAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// GetOrbitalElements function as declared in swephexp.h:937
func GetOrbitalElements(tjdEt float64, ipl int, iflag int, dret []float64, serr []byte) int32 {
	ctjdEt, ctjdEtAllocMap := (C.double)(tjdEt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cdret, cdretAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dret)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_orbital_elements(ctjdEt, cipl, ciflag, cdret, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdretAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdEtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// OrbitMaxMinTrueDistance function as declared in swephexp.h:940
func OrbitMaxMinTrueDistance(tjdEt float64, ipl int, iflag int, dmax []float64, dmin []float64, dtrue []float64, serr []byte) int32 {
	ctjdEt, ctjdEtAllocMap := (C.double)(tjdEt), cgoAllocsUnknown
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cdmax, cdmaxAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dmax)))
	cdmin, cdminAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dmin)))
	cdtrue, cdtrueAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dtrue)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_orbit_max_min_true_distance(ctjdEt, cipl, ciflag, cdmax, cdmin, cdtrue, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cdtrueAllocMap)
	runtime.KeepAlive(cdminAllocMap)
	runtime.KeepAlive(cdmaxAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	runtime.KeepAlive(ctjdEtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Deltat function as declared in swephexp.h:947
func Deltat(tjd float64) float64 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_deltat(ctjd)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (float64)(__ret)
	return __v
}

// DeltatEx function as declared in swephexp.h:948
func DeltatEx(tjd float64, iflag int, serr []byte) float64 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_deltat_ex(ctjd, ciflag, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (float64)(__ret)
	return __v
}

// TimeEqu function as declared in swephexp.h:951
func TimeEqu(tjd float64, te []float64, serr []byte) int32 {
	ctjd, ctjdAllocMap := (C.double)(tjd), cgoAllocsUnknown
	cte, cteAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&te)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_time_equ(ctjd, cte, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cteAllocMap)
	runtime.KeepAlive(ctjdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// LmtToLat function as declared in swephexp.h:952
func LmtToLat(tjdLmt float64, geolon float64, tjdLat []float64, serr []byte) int32 {
	ctjdLmt, ctjdLmtAllocMap := (C.double)(tjdLmt), cgoAllocsUnknown
	cgeolon, cgeolonAllocMap := (C.double)(geolon), cgoAllocsUnknown
	ctjdLat, ctjdLatAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tjdLat)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_lmt_to_lat(ctjdLmt, cgeolon, ctjdLat, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(ctjdLatAllocMap)
	runtime.KeepAlive(cgeolonAllocMap)
	runtime.KeepAlive(ctjdLmtAllocMap)
	__v := (int32)(__ret)
	return __v
}

// LatToLmt function as declared in swephexp.h:953
func LatToLmt(tjdLat float64, geolon float64, tjdLmt []float64, serr []byte) int32 {
	ctjdLat, ctjdLatAllocMap := (C.double)(tjdLat), cgoAllocsUnknown
	cgeolon, cgeolonAllocMap := (C.double)(geolon), cgoAllocsUnknown
	ctjdLmt, ctjdLmtAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&tjdLmt)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_lat_to_lmt(ctjdLat, cgeolon, ctjdLmt, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(ctjdLmtAllocMap)
	runtime.KeepAlive(cgeolonAllocMap)
	runtime.KeepAlive(ctjdLatAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Sidtime0 function as declared in swephexp.h:956
func Sidtime0(tjdUt float64, eps float64, nut float64) float64 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	ceps, cepsAllocMap := (C.double)(eps), cgoAllocsUnknown
	cnut, cnutAllocMap := (C.double)(nut), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_sidtime0(ctjdUt, ceps, cnut)
	runtime.KeepAlive(cnutAllocMap)
	runtime.KeepAlive(cepsAllocMap)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (float64)(__ret)
	return __v
}

// Sidtime function as declared in swephexp.h:957
func Sidtime(tjdUt float64) float64 {
	ctjdUt, ctjdUtAllocMap := (C.double)(tjdUt), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_sidtime(ctjdUt)
	runtime.KeepAlive(ctjdUtAllocMap)
	__v := (float64)(__ret)
	return __v
}

// SetInterpolateNut function as declared in swephexp.h:958
func SetInterpolateNut(doInterpolate int32) {
	cdoInterpolate, cdoInterpolateAllocMap := (C.AS_BOOL)(doInterpolate), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_interpolate_nut(cdoInterpolate)
	runtime.KeepAlive(cdoInterpolateAllocMap)
}

// Cotrans function as declared in swephexp.h:961
func Cotrans(xpo []float64, xpn []float64, eps float64) {
	cxpo, cxpoAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xpo)))
	cxpn, cxpnAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xpn)))
	ceps, cepsAllocMap := (C.double)(eps), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_cotrans(cxpo, cxpn, ceps)
	runtime.KeepAlive(cepsAllocMap)
	runtime.KeepAlive(cxpnAllocMap)
	runtime.KeepAlive(cxpoAllocMap)
}

// CotransSp function as declared in swephexp.h:962
func CotransSp(xpo []float64, xpn []float64, eps float64) {
	cxpo, cxpoAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xpo)))
	cxpn, cxpnAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xpn)))
	ceps, cepsAllocMap := (C.double)(eps), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_cotrans_sp(cxpo, cxpn, ceps)
	runtime.KeepAlive(cepsAllocMap)
	runtime.KeepAlive(cxpnAllocMap)
	runtime.KeepAlive(cxpoAllocMap)
}

// GetTidAcc function as declared in swephexp.h:965
func GetTidAcc() float64 {
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_get_tid_acc()
	__v := (float64)(__ret)
	return __v
}

// SetTidAcc function as declared in swephexp.h:966
func SetTidAcc(tAcc float64) {
	ctAcc, ctAccAllocMap := (C.double)(tAcc), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_tid_acc(ctAcc)
	runtime.KeepAlive(ctAccAllocMap)
}

// SetDeltaTUserdef function as declared in swephexp.h:970
func SetDeltaTUserdef(dt float64) {
	cdt, cdtAllocMap := (C.double)(dt), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_set_delta_t_userdef(cdt)
	runtime.KeepAlive(cdtAllocMap)
}

// Degnorm function as declared in swephexp.h:972
func Degnorm(x float64) float64 {
	cx, cxAllocMap := (C.double)(x), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_degnorm(cx)
	runtime.KeepAlive(cxAllocMap)
	__v := (float64)(__ret)
	return __v
}

// Radnorm function as declared in swephexp.h:973
func Radnorm(x float64) float64 {
	cx, cxAllocMap := (C.double)(x), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_radnorm(cx)
	runtime.KeepAlive(cxAllocMap)
	__v := (float64)(__ret)
	return __v
}

// RadMidp function as declared in swephexp.h:974
func RadMidp(x1 float64, x0 float64) float64 {
	cx1, cx1AllocMap := (C.double)(x1), cgoAllocsUnknown
	cx0, cx0AllocMap := (C.double)(x0), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_rad_midp(cx1, cx0)
	runtime.KeepAlive(cx0AllocMap)
	runtime.KeepAlive(cx1AllocMap)
	__v := (float64)(__ret)
	return __v
}

// DegMidp function as declared in swephexp.h:975
func DegMidp(x1 float64, x0 float64) float64 {
	cx1, cx1AllocMap := (C.double)(x1), cgoAllocsUnknown
	cx0, cx0AllocMap := (C.double)(x0), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_deg_midp(cx1, cx0)
	runtime.KeepAlive(cx0AllocMap)
	runtime.KeepAlive(cx1AllocMap)
	__v := (float64)(__ret)
	return __v
}

// SplitDeg function as declared in swephexp.h:977
func SplitDeg(ddeg float64, roundflag int, ideg []int, imin []int, isec []int, dsecfr []float64, isgn []int32) {
	cddeg, cddegAllocMap := (C.double)(ddeg), cgoAllocsUnknown
	croundflag, croundflagAllocMap := (C.int32)(roundflag), cgoAllocsUnknown
	cideg, cidegAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&ideg)))
	cimin, ciminAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&imin)))
	cisec, cisecAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&isec)))
	cdsecfr, cdsecfrAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&dsecfr)))
	cisgn, cisgnAllocMap := copyPInt32Bytes((*sliceHeader)(unsafe.Pointer(&isgn)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	C.swe_split_deg(cddeg, croundflag, cideg, cimin, cisec, cdsecfr, cisgn)
	runtime.KeepAlive(cisgnAllocMap)
	runtime.KeepAlive(cdsecfrAllocMap)
	runtime.KeepAlive(cisecAllocMap)
	runtime.KeepAlive(ciminAllocMap)
	runtime.KeepAlive(cidegAllocMap)
	runtime.KeepAlive(croundflagAllocMap)
	runtime.KeepAlive(cddegAllocMap)
}

// Csnorm function as declared in swephexp.h:986
func Csnorm(p int32) int32 {
	cp, cpAllocMap := (C.centisec)(p), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_csnorm(cp)
	runtime.KeepAlive(cpAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Difcsn function as declared in swephexp.h:989
func Difcsn(p1 int, p2 int32) int32 {
	cp1, cp1AllocMap := (C.centisec)(p1), cgoAllocsUnknown
	cp2, cp2AllocMap := (C.centisec)(p2), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_difcsn(cp1, cp2)
	runtime.KeepAlive(cp2AllocMap)
	runtime.KeepAlive(cp1AllocMap)
	__v := (int32)(__ret)
	return __v
}

// Difdegn function as declared in swephexp.h:991
func Difdegn(p1 float64, p2 float64) float64 {
	cp1, cp1AllocMap := (C.double)(p1), cgoAllocsUnknown
	cp2, cp2AllocMap := (C.double)(p2), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_difdegn(cp1, cp2)
	runtime.KeepAlive(cp2AllocMap)
	runtime.KeepAlive(cp1AllocMap)
	__v := (float64)(__ret)
	return __v
}

// Difcs2n function as declared in swephexp.h:994
func Difcs2n(p1 int, p2 int32) int32 {
	cp1, cp1AllocMap := (C.centisec)(p1), cgoAllocsUnknown
	cp2, cp2AllocMap := (C.centisec)(p2), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_difcs2n(cp1, cp2)
	runtime.KeepAlive(cp2AllocMap)
	runtime.KeepAlive(cp1AllocMap)
	__v := (int32)(__ret)
	return __v
}

// Difdeg2n function as declared in swephexp.h:996
func Difdeg2n(p1 float64, p2 float64) float64 {
	cp1, cp1AllocMap := (C.double)(p1), cgoAllocsUnknown
	cp2, cp2AllocMap := (C.double)(p2), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_difdeg2n(cp1, cp2)
	runtime.KeepAlive(cp2AllocMap)
	runtime.KeepAlive(cp1AllocMap)
	__v := (float64)(__ret)
	return __v
}

// Difrad2n function as declared in swephexp.h:997
func Difrad2n(p1 float64, p2 float64) float64 {
	cp1, cp1AllocMap := (C.double)(p1), cgoAllocsUnknown
	cp2, cp2AllocMap := (C.double)(p2), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_difrad2n(cp1, cp2)
	runtime.KeepAlive(cp2AllocMap)
	runtime.KeepAlive(cp1AllocMap)
	__v := (float64)(__ret)
	return __v
}

// Csroundsec function as declared in swephexp.h:1000
func Csroundsec(x int32) int32 {
	cx, cxAllocMap := (C.centisec)(x), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_csroundsec(cx)
	runtime.KeepAlive(cxAllocMap)
	__v := (int32)(__ret)
	return __v
}

// D2l function as declared in swephexp.h:1003
func D2l(x float64) int32 {
	cx, cxAllocMap := (C.double)(x), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_d2l(cx)
	runtime.KeepAlive(cxAllocMap)
	__v := (int32)(__ret)
	return __v
}

// DayOfWeek function as declared in swephexp.h:1006
func DayOfWeek(jd float64) int32 {
	cjd, cjdAllocMap := (C.double)(jd), cgoAllocsUnknown
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_day_of_week(cjd)
	runtime.KeepAlive(cjdAllocMap)
	__v := (int32)(__ret)
	return __v
}

// Cs2timestr function as declared in swephexp.h:1008
func Cs2timestr(t int, sep int, suppressZero int, a []byte) *byte {
	ct, ctAllocMap := (C.centisec)(t), cgoAllocsUnknown
	csep, csepAllocMap := (C.int)(sep), cgoAllocsUnknown
	csuppressZero, csuppressZeroAllocMap := (C.AS_BOOL)(suppressZero), cgoAllocsUnknown
	ca, caAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&a)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_cs2timestr(ct, csep, csuppressZero, ca)
	runtime.KeepAlive(caAllocMap)
	runtime.KeepAlive(csuppressZeroAllocMap)
	runtime.KeepAlive(csepAllocMap)
	runtime.KeepAlive(ctAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// Cs2lonlatstr function as declared in swephexp.h:1010
func Cs2lonlatstr(t int, pchar byte, mchar byte, s []byte) *byte {
	ct, ctAllocMap := (C.centisec)(t), cgoAllocsUnknown
	cpchar, cpcharAllocMap := (C.char)(pchar), cgoAllocsUnknown
	cmchar, cmcharAllocMap := (C.char)(mchar), cgoAllocsUnknown
	cs, csAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&s)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_cs2lonlatstr(ct, cpchar, cmchar, cs)
	runtime.KeepAlive(csAllocMap)
	runtime.KeepAlive(cmcharAllocMap)
	runtime.KeepAlive(cpcharAllocMap)
	runtime.KeepAlive(ctAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// Cs2degstr function as declared in swephexp.h:1012
func Cs2degstr(t int, a []byte) *byte {
	ct, ctAllocMap := (C.centisec)(t), cgoAllocsUnknown
	ca, caAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&a)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_cs2degstr(ct, ca)
	runtime.KeepAlive(caAllocMap)
	runtime.KeepAlive(ctAllocMap)
	__v := *(**byte)(unsafe.Pointer(&__ret))
	return __v
}

// Solcross function as declared in swephexp.h:707
func Solcross(x2cross float64, jdEt float64, flag int32, serr []byte) float64 {
	cx2cross, cx2crossAllocMap := (C.double)(x2cross), cgoAllocsUnknown
	cjdEt, cjdEtAllocMap := (C.double)(jdEt), cgoAllocsUnknown
	cflag, cflagAllocMap := (C.int32)(flag), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_solcross(cx2cross, cjdEt, cflag, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cflagAllocMap)
	runtime.KeepAlive(cjdEtAllocMap)
	runtime.KeepAlive(cx2crossAllocMap)
	return (float64)(__ret)
}

// SolcrossUt function as declared in swephexp.h:708
func SolcrossUt(x2cross float64, jdUt float64, flag int32, serr []byte) float64 {
	cx2cross, cx2crossAllocMap := (C.double)(x2cross), cgoAllocsUnknown
	cjdUt, cjdUtAllocMap := (C.double)(jdUt), cgoAllocsUnknown
	cflag, cflagAllocMap := (C.int32)(flag), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_solcross_ut(cx2cross, cjdUt, cflag, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cflagAllocMap)
	runtime.KeepAlive(cjdUtAllocMap)
	runtime.KeepAlive(cx2crossAllocMap)
	return (float64)(__ret)
}

// Mooncross function as declared in swephexp.h:709
func Mooncross(x2cross float64, jdEt float64, flag int32, serr []byte) float64 {
	cx2cross, cx2crossAllocMap := (C.double)(x2cross), cgoAllocsUnknown
	cjdEt, cjdEtAllocMap := (C.double)(jdEt), cgoAllocsUnknown
	cflag, cflagAllocMap := (C.int32)(flag), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_mooncross(cx2cross, cjdEt, cflag, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cflagAllocMap)
	runtime.KeepAlive(cjdEtAllocMap)
	runtime.KeepAlive(cx2crossAllocMap)
	return (float64)(__ret)
}

// MooncrossUt function as declared in swephexp.h:710
func MooncrossUt(x2cross float64, jdUt float64, flag int32, serr []byte) float64 {
	cx2cross, cx2crossAllocMap := (C.double)(x2cross), cgoAllocsUnknown
	cjdUt, cjdUtAllocMap := (C.double)(jdUt), cgoAllocsUnknown
	cflag, cflagAllocMap := (C.int32)(flag), cgoAllocsUnknown
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_mooncross_ut(cx2cross, cjdUt, cflag, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cflagAllocMap)
	runtime.KeepAlive(cjdUtAllocMap)
	runtime.KeepAlive(cx2crossAllocMap)
	return (float64)(__ret)
}

// MooncrossNode function as declared in swephexp.h:711
func MooncrossNode(jdEt float64, flag int32, xlon []float64, xlat []float64, serr []byte) float64 {
	cjdEt, cjdEtAllocMap := (C.double)(jdEt), cgoAllocsUnknown
	cflag, cflagAllocMap := (C.int32)(flag), cgoAllocsUnknown
	cxlon, cxlonAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xlon)))
	cxlat, cxlatAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xlat)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_mooncross_node(cjdEt, cflag, cxlon, cxlat, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxlatAllocMap)
	runtime.KeepAlive(cxlonAllocMap)
	runtime.KeepAlive(cflagAllocMap)
	runtime.KeepAlive(cjdEtAllocMap)
	return (float64)(__ret)
}

// MooncrossNodeUt function as declared in swephexp.h:712
func MooncrossNodeUt(jdUt float64, flag int32, xlon []float64, xlat []float64, serr []byte) float64 {
	cjdUt, cjdUtAllocMap := (C.double)(jdUt), cgoAllocsUnknown
	cflag, cflagAllocMap := (C.int32)(flag), cgoAllocsUnknown
	cxlon, cxlonAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xlon)))
	cxlat, cxlatAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&xlat)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_mooncross_node_ut(cjdUt, cflag, cxlon, cxlat, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cxlatAllocMap)
	runtime.KeepAlive(cxlonAllocMap)
	runtime.KeepAlive(cflagAllocMap)
	runtime.KeepAlive(cjdUtAllocMap)
	return (float64)(__ret)
}

// HelioCross function as declared in swephexp.h:713
func HelioCross(ipl int32, x2cross float64, jdEt float64, iflag int32, dir int32, jdCross []float64, serr []byte) int32 {
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	cx2cross, cx2crossAllocMap := (C.double)(x2cross), cgoAllocsUnknown
	cjdEt, cjdEtAllocMap := (C.double)(jdEt), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cdir, cdirAllocMap := (C.int32)(dir), cgoAllocsUnknown
	cjdCross, cjdCrossAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&jdCross)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_helio_cross(cipl, cx2cross, cjdEt, ciflag, cdir, cjdCross, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cjdCrossAllocMap)
	runtime.KeepAlive(cdirAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(cjdEtAllocMap)
	runtime.KeepAlive(cx2crossAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	return (int32)(__ret)
}

// HelioCrossUt function as declared in swephexp.h:714
func HelioCrossUt(ipl int32, x2cross float64, jdUt float64, iflag int32, dir int32, jdCross []float64, serr []byte) int32 {
	cipl, ciplAllocMap := (C.int32)(ipl), cgoAllocsUnknown
	cx2cross, cx2crossAllocMap := (C.double)(x2cross), cgoAllocsUnknown
	cjdUt, cjdUtAllocMap := (C.double)(jdUt), cgoAllocsUnknown
	ciflag, ciflagAllocMap := (C.int32)(iflag), cgoAllocsUnknown
	cdir, cdirAllocMap := (C.int32)(dir), cgoAllocsUnknown
	cjdCross, cjdCrossAllocMap := copyPDoubleBytes((*sliceHeader)(unsafe.Pointer(&jdCross)))
	cserr, cserrAllocMap := copyPCharBytes((*sliceHeader)(unsafe.Pointer(&serr)))
	goseMutex.Lock()
	defer goseMutex.Unlock()
	__ret := C.swe_helio_cross_ut(cipl, cx2cross, cjdUt, ciflag, cdir, cjdCross, cserr)
	runtime.KeepAlive(cserrAllocMap)
	runtime.KeepAlive(cjdCrossAllocMap)
	runtime.KeepAlive(cdirAllocMap)
	runtime.KeepAlive(ciflagAllocMap)
	runtime.KeepAlive(cjdUtAllocMap)
	runtime.KeepAlive(cx2crossAllocMap)
	runtime.KeepAlive(ciplAllocMap)
	return (int32)(__ret)
}
