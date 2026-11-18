package gose_test

import (
	"bytes"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/bharathak/gose"
)

func init() {
	// Set ephemeris path to swisseph/ephe if available
	ephePath := "../swisseph/ephe"
	if abs, err := filepath.Abs(ephePath); err == nil {
		if _, err := os.Stat(abs); err == nil {
			gose.SetEphePathStr(abs)
		}
	}
}

// 1. Library Version, Denum, and Constants Parity
func TestVersion(t *testing.T) {
	buf := make([]byte, 256)
	ptr := gose.Version(buf)
	if ptr != &buf[0] {
		t.Fatal("expected Version(buf) to return pointer to buf[0]")
	}
	verStr := gose.VersionStr()
	if verStr == "" {
		t.Fatal("expected non-empty version string")
	}
	if verStr != "2.10.03" {
		t.Fatalf("expected version 2.10.03, got %s", verStr)
	}

	// Test short buffer and nil buffer
	shortBuf := make([]byte, 4)
	shortPtr := gose.Version(shortBuf)
	if shortPtr != &shortBuf[0] || string(shortBuf) != "2.10" {
		t.Fatalf("unexpected short buffer version: %q", string(shortBuf))
	}
	if gose.Version(nil) == nil {
		t.Fatal("expected non-nil pointer from Version(nil)")
	}

	// Test GetLibraryPath and GetLibraryPathStr
	libBuf := make([]byte, 256)
	if gose.GetLibraryPath(libBuf) != &libBuf[0] {
		t.Fatal("expected GetLibraryPath(libBuf) to return &libBuf[0]")
	}
	_ = gose.GetLibraryPathStr()
	if gose.GetLibraryPath(nil) == nil {
		t.Fatal("expected non-nil pointer from GetLibraryPath(nil)")
	}

	// Test GetDenum (synced from latest swisseph)
	var xx [6]float64
	var serr [256]byte
	_ = gose.CalcUt(2451545.0, gose.SeSun, gose.SeflgSwieph, xx[:], serr[:])
	denum := gose.GetDenum()
	// With Moshier fallback or SE, denum is 404 or 431
	if denum != 404 && denum != 431 && denum != 441 && denum != 0 {
		t.Fatalf("unexpected GetDenum value: %d", denum)
	}

	// Verify newly synced swephexp.h constants
	if gose.Se3DeNumber != 441 {
		t.Fatalf("expected Se3DeNumber=441, got %d", gose.Se3DeNumber)
	}
	if gose.SeFnameDe441 != "de441.eph" || gose.SeFnameDft3 != "de441.eph" {
		t.Fatalf("unexpected DE441 constants: %s / %s", gose.SeFnameDe441, gose.SeFnameDft3)
	}
	if gose.SeEphePath != ".:/home/ephe/" {
		t.Fatalf("expected SeEphePath='.:/home/ephe/', got %q", gose.SeEphePath)
	}
}

// 2. Calendar & Julian Day Calculations
func TestJuldayAndRevjul(t *testing.T) {
	// J2000.0: 2000-01-01 12:00 UTC is exactly JD 2451545.0
	jd := gose.Julday(2000, 1, 1, 12.0, gose.SeGregCal)
	if math.Abs(jd-2451545.0) > 1e-6 {
		t.Fatalf("expected JD 2451545.0 for J2000.0, got %f", jd)
	}

	// Test Julian Calendar conversion (1582-10-04 is JD 2299160.5)
	jdJul := gose.Julday(1582, 10, 4, 12.0, gose.SeJulCal)
	if math.Abs(jdJul-2299160.0) > 1e-6 {
		t.Fatalf("expected JD 2299160.0 for Julian 1582-10-04 12:00, got %f", jdJul)
	}

	// Test Revjul round-trip
	y := make([]int, 1)
	m := make([]int, 1)
	d := make([]int, 1)
	ut := make([]float64, 1)
	gose.Revjul(jd, gose.SeGregCal, y, m, d, ut)
	if y[0] != 2000 || m[0] != 1 || d[0] != 1 || math.Abs(ut[0]-12.0) > 1e-6 {
		t.Fatalf("Revjul failed: got %d-%d-%d %f", y[0], m[0], d[0], ut[0])
	}

	// Test DayOfWeek (Swiss Ephemeris: Monday = 0, ..., Saturday = 5, Sunday = 6)
	dow := gose.DayOfWeek(jd)
	if dow != 5 {
		t.Fatalf("expected day of week 5 (Saturday), got %d", dow)
	}

	// Test DateConversion
	tjdRet := make([]float64, 1)
	retDate := gose.DateConversion(2000, 1, 1, 12.0, 'g', tjdRet)
	if retDate != 0 || math.Abs(tjdRet[0]-2451545.0) > 1e-6 {
		t.Fatalf("DateConversion failed: ret=%d, jd=%f", retDate, tjdRet[0])
	}

	// Test time.Time integration
	tm := time.Date(2024, 5, 20, 12, 0, 0, 0, time.UTC)
	jdTm := gose.JuldayTime(tm, gose.SeGregCal)
	expectedJD := 2460451.0
	if math.Abs(jdTm-expectedJD) > 1e-5 {
		t.Fatalf("expected JD %f, got %f", expectedJD, jdTm)
	}

	revTm := gose.RevjulTime(jdTm, gose.SeGregCal)
	if revTm.Year() != 2024 || revTm.Month() != 5 || revTm.Day() != 20 || revTm.Hour() != 12 {
		t.Fatalf("RevjulTime round-trip failed: got %v", revTm)
	}

	// Test RevjulTime exact second precision (no float truncation to 59.999999s)
	for _, tc := range []time.Time{
		time.Date(2024, 5, 20, 15, 30, 45, 0, time.UTC),
		time.Date(2024, 5, 20, 15, 1, 0, 0, time.UTC),
		time.Date(1990, 11, 22, 6, 45, 30, 0, time.UTC),
	} {
		jdVal := gose.JuldayTime(tc, gose.SeGregCal)
		got := gose.RevjulTime(jdVal, gose.SeGregCal)
		if !got.Round(time.Millisecond).Equal(tc) {
			t.Fatalf("RevjulTime precision mismatch: want %v, got %v", tc, got)
		}
	}
}

// 3. Core Planetary Calculations
func TestCalcUt(t *testing.T) {
	jd := 2451545.0 // J2000.0
	xx := make([]float64, 6)
	serr := make([]byte, 256)

	// Calculate Sun position
	ret := gose.CalcUt(jd, gose.SeSun, gose.SeflgSpeed, xx, serr)
	if ret < 0 {
		t.Fatalf("CalcUt Sun failed: %s", string(serr))
	}
	// Sun at J2000 is ~280.46 degrees
	if xx[0] < 280.0 || xx[0] > 281.0 {
		t.Fatalf("Sun longitude out of expected range: %f", xx[0])
	}
	// Distance should be approx 0.983 AU (perihelion is in early January)
	if xx[2] < 0.98 || xx[2] > 0.99 {
		t.Fatalf("Sun distance out of expected range: %f", xx[2])
	}
	// Sun daily motion ~0.9856 deg/day
	if xx[3] < 0.95 || xx[3] > 1.05 {
		t.Fatalf("Sun speed out of expected range: %f", xx[3])
	}

	// Test CalcUtSimple ergonomic wrapper
	pos, iflgret, err := gose.CalcUtSimple(jd, gose.SeSun, gose.SeflgSpeed)
	if err != nil || iflgret < 0 {
		t.Fatalf("CalcUtSimple Sun failed: %v", err)
	}
	if math.Abs(pos.Longitude-xx[0]) > 1e-9 {
		t.Fatalf("mismatch between CalcUt and CalcUtSimple: %f vs %f", xx[0], pos.Longitude)
	}

	// Calculate Moon position
	ret = gose.CalcUt(jd, gose.SeMoon, gose.SeflgSpeed, xx, serr)
	if ret < 0 {
		t.Fatalf("CalcUt Moon failed: %s", string(serr))
	}
	if xx[0] < 0 || xx[0] >= 360 {
		t.Fatalf("Moon longitude invalid: %f", xx[0])
	}
	// Moon speed ~12 to 15 deg/day
	if xx[3] < 11.0 || xx[3] > 16.0 {
		t.Fatalf("Moon speed out of expected range: %f", xx[3])
	}

	// Test Planet Names
	planets := []struct {
		id   int
		name string
	}{
		{gose.SeSun, "Sun"},
		{gose.SeMoon, "Moon"},
		{gose.SeMercury, "Mercury"},
		{gose.SeVenus, "Venus"},
		{gose.SeMars, "Mars"},
		{gose.SeJupiter, "Jupiter"},
		{gose.SeSaturn, "Saturn"},
		{gose.SeUranus, "Uranus"},
		{gose.SeNeptune, "Neptune"},
		{gose.SePluto, "Pluto"},
	}
	for _, p := range planets {
		name := gose.GetPlanetNameStr(p.id)
		if name != p.name {
			t.Fatalf("expected '%s', got '%s'", p.name, name)
		}
	}
}

// 4. House Systems & Angles
func TestHouses(t *testing.T) {
	jd := 2451545.0
	geolat := 51.5074 // London
	geolon := -0.1278

	cusps := make([]float64, 13)
	ascmc := make([]float64, 10)

	// Test Placidus ('P'), Koch ('K'), Whole Sign ('W'), Equal ('E')
	systems := []byte{'P', 'K', 'W', 'E'}
	for _, sys := range systems {
		ret := gose.Houses(jd, geolat, geolon, int(sys), cusps, ascmc)
		if ret < 0 {
			t.Fatalf("Houses calculation failed for system %c", sys)
		}

		// Check that all 12 cusps are in [0, 360)
		for i := 1; i <= 12; i++ {
			if cusps[i] < 0 || cusps[i] >= 360 {
				t.Fatalf("system %c cusp %d invalid: %f", sys, i, cusps[i])
			}
		}
		// Ascendant and MC
		if ascmc[0] < 0 || ascmc[0] >= 360 {
			t.Fatalf("system %c Ascendant invalid: %f", sys, ascmc[0])
		}
		if ascmc[1] < 0 || ascmc[1] >= 360 {
			t.Fatalf("system %c MC invalid: %f", sys, ascmc[1])
		}
	}

	// Test HousesSimple
	res, iflgret, err := gose.HousesSimple(jd, 0, geolat, geolon, int('P'))
	if err != nil || iflgret < 0 {
		t.Fatalf("HousesSimple failed: %v", err)
	}
	if math.Abs(res.Ascendant-ascmc[0]) > 1e-9 {
		t.Fatalf("HousesSimple Ascendant mismatch: %f vs %f", res.Ascendant, ascmc[0])
	}

	// Test HousePos
	xpin := []float64{100.0, 0.0}
	serr := make([]byte, 256)
	hpos := gose.HousePos(ascmc[2], geolat, 23.44, int('P'), xpin, serr)
	if hpos < 1.0 || hpos > 13.0 {
		t.Fatalf("HousePos out of expected range [1, 13]: %f", hpos)
	}

	// Test HouseName
	hname := gose.HouseNameStr(int32('P'))
	if hname == "" {
		t.Fatal("expected non-empty Placidus house name")
	}
}

// 5. Sidereal Calculations & Ayanamsa
func TestSidereal(t *testing.T) {
	// Test Lahiri, Raman, Krishnamurti, Fagan-Bradley
	sidModes := []struct {
		mode    int32
		name    string
		minAyan float64
		maxAyan float64
	}{
		{gose.SeSidmLahiri, "Lahiri", 23.5, 24.2},
		{gose.SeSidmRaman, "Raman", 22.0, 23.0},
		{gose.SeSidmKrishnamurti, "Krishnamurti", 23.5, 24.5},
		{gose.SeSidmFaganBradley, "Fagan/Bradley", 24.2, 25.2},
	}

	tjd := 2451545.0
	for _, sm := range sidModes {
		gose.SetSidMode(int(sm.mode), 0, 0)
		ayan := gose.GetAyanamsa(tjd)
		if ayan < sm.minAyan || ayan > sm.maxAyan {
			t.Fatalf("%s ayanamsa at J2000 out of range: %f (expected [%f, %f])", sm.name, ayan, sm.minAyan, sm.maxAyan)
		}

		name := gose.GetAyanamsaName(sm.mode)
		if name == "" {
			t.Fatalf("expected non-empty ayanamsa name for mode %d", sm.mode)
		}
	}
}

// 6. Solcross and Mooncross Ingresses
func TestSolcrossAndMooncross(t *testing.T) {
	serr := make([]byte, 256)
	tjd := 2460451.0 // 2024-05-20

	// Find next Aries ingress (Sun at 0 degrees longitude)
	jdCross := gose.SolcrossUt(0, tjd, 0, serr)
	if jdCross < tjd {
		t.Fatalf("expected next crossing after %f, got %f", tjd, jdCross)
	}

	// Verify Sun is at 0 degrees at jdCross
	xx := make([]float64, 6)
	serrCalc := make([]byte, 256)
	gose.CalcUt(jdCross, gose.SeSun, 0, xx, serrCalc)
	diff := math.Abs(xx[0])
	if diff > 180 {
		diff = 360 - diff
	}
	if diff > 0.01 {
		t.Fatalf("expected Sun longitude ~0 at crossing, got %f", xx[0])
	}

	// Test Moon crossing 0 degrees
	jdMoonCross := gose.MooncrossUt(0, tjd, 0, serr)
	if jdMoonCross < tjd || jdMoonCross > tjd+30.0 {
		t.Fatalf("MooncrossUt returned unexpected date: %f", jdMoonCross)
	}

	// Test Moon node crossing
	xlon := make([]float64, 1)
	xlat := make([]float64, 1)
	jdNode := gose.MooncrossNodeUt(tjd, 0, xlon, xlat, serr)
	if jdNode < tjd || jdNode > tjd+30.0 {
		t.Fatalf("MooncrossNodeUt returned unexpected date: %f", jdNode)
	}
	if math.Abs(xlat[0]) > 0.05 {
		t.Fatalf("expected Moon latitude ~0 at node crossing, got %f", xlat[0])
	}
}

// 7. Solar and Lunar Eclipses
func TestSolarAndLunarEclipses(t *testing.T) {
	tjd := 2451545.0 // 2000-01-01
	serr := make([]byte, 256)
	tret := make([]float64, 10)

	// Global search for next total solar eclipse
	ret := gose.SolEclipseWhenGlob(tjd, 0, gose.SeEclTotal, tret, 0, serr)
	if ret < 0 {
		t.Fatalf("SolEclipseWhenGlob failed: %s", string(serr))
	}
	// The first total solar eclipse after 2000-01-01 was on 2001-06-21 (~JD 2452082)
	if tret[0] < 2452000.0 || tret[0] > 2452200.0 {
		t.Fatalf("unexpected total solar eclipse date: %f", tret[0])
	}

	// Global search for next lunar eclipse
	ret = gose.LunEclipseWhen(tjd, 0, gose.SeEclAlltypesLunar, tret, 0, serr)
	if ret < 0 {
		t.Fatalf("LunEclipseWhen failed: %s", string(serr))
	}
	// The total lunar eclipse of 2000-01-21 (~JD 2451564)
	if tret[0] < 2451550.0 || tret[0] > 2451580.0 {
		t.Fatalf("unexpected lunar eclipse date: %f", tret[0])
	}

	// Lunar eclipse details (how)
	attr := make([]float64, 20)
	geopos := []float64{-0.1278, 51.5074, 0}
	ret = gose.LunEclipseHow(tret[0], 0, geopos, attr, serr)
	if ret < 0 {
		t.Fatalf("LunEclipseHow failed: %s", string(serr))
	}
	// Umbral magnitude > 1 for total eclipse
	if attr[0] < 0.5 {
		t.Fatalf("expected significant lunar eclipse magnitude, got %f", attr[0])
	}
}

// 8. Rise, Set, Transits, and Horizontal Coordinates
func TestRiseTransAndAzalt(t *testing.T) {
	tjd := 2451545.0 // 2000-01-01
	geopos := []float64{-0.1278, 51.5074, 0} // London
	tret := make([]float64, 10)
	serr := make([]byte, 256)

	// Calculate sunrise in London
	ret := gose.RiseTrans(tjd, gose.SeSun, nil, 0, gose.SeCalcRise, geopos, 1013.25, 10.0, tret, serr)
	if ret < 0 {
		t.Fatalf("RiseTrans sunrise failed: %s", string(serr))
	}
	// Sunrise should occur on that day
	if tret[0] < tjd || tret[0] > tjd+1.5 {
		t.Fatalf("unexpected sunrise time: %f", tret[0])
	}

	// Calculate sunset in London
	ret = gose.RiseTrans(tjd, gose.SeSun, nil, 0, gose.SeCalcSet, geopos, 1013.25, 10.0, tret, serr)
	if ret < 0 {
		t.Fatalf("RiseTrans sunset failed: %s", string(serr))
	}

	// Horizontal coordinates conversion (Azalt / AzaltRev round-trip)
	xin := []float64{120.0, 15.0, 1.0} // ecliptic coordinates
	xaz := make([]float64, 3)
	gose.Azalt(tjd, gose.SeEcl2hor, geopos, 1013.25, 10.0, xin, xaz)
	if xaz[0] < 0 || xaz[0] >= 360 {
		t.Fatalf("invalid azimuth: %f", xaz[0])
	}

	xout := make([]float64, 3)
	gose.AzaltRev(tjd, gose.SeHor2ecl, geopos, xaz, xout)
	if math.Abs(xout[0]-xin[0]) > 0.001 || math.Abs(xout[1]-xin[1]) > 0.001 {
		t.Fatalf("Azalt/AzaltRev roundtrip mismatch: in=(%f, %f) out=(%f, %f)", xin[0], xin[1], xout[0], xout[1])
	}

	// Atmospheric refraction
	refr := gose.Refrac(0.0, 1013.25, 10.0, gose.SeTrueToApp)
	// Refraction at horizon is approx 34 arcminutes (~0.57 degrees, 0.48 to 0.60 depending on formula)
	if refr < 0.4 || refr > 0.7 {
		t.Fatalf("horizon refraction out of expected range: %f", refr)
	}
}

// 9. Time Systems: Delta T, Equation of Time, Sidereal Time
func TestTimeAndDeltaT(t *testing.T) {
	tjd := 2451545.0 // J2000.0

	// Delta T at J2000 is ~63.8 seconds
	dt := gose.Deltat(tjd) * 86400.0
	if dt < 63.0 || dt > 65.0 {
		t.Fatalf("Delta T at J2000 out of expected range: %f seconds", dt)
	}

	// DeltatEx
	serr := make([]byte, 256)
	dtEx := gose.DeltatEx(tjd, 0, serr) * 86400.0
	if math.Abs(dtEx-dt) > 1e-4 {
		t.Fatalf("Deltat and DeltatEx mismatch: %f vs %f", dt, dtEx)
	}

	// Equation of Time (at J2000 is approx -3.3 minutes)
	te := make([]float64, 1)
	ret := gose.TimeEqu(tjd, te, serr)
	if ret < 0 {
		t.Fatalf("TimeEqu failed: %s", string(serr))
	}
	teMin := te[0] * 1440.0
	if teMin < -4.0 || teMin > -2.5 {
		t.Fatalf("Equation of time at J2000 out of expected range: %f min", teMin)
	}

	// Sidereal Time (at J2000 12:00 UT is approx 18.697 hours)
	sid := gose.Sidtime(tjd)
	if sid < 18.65 || sid > 18.75 {
		t.Fatalf("Sidereal time at J2000 out of expected range: %f hours", sid)
	}

	// LMT to LAT round-trip
	tjdLat := make([]float64, 1)
	tjdLmt := make([]float64, 1)
	geolon := 10.0
	ret = gose.LmtToLat(tjd, geolon, tjdLat, serr)
	if ret < 0 {
		t.Fatalf("LmtToLat failed: %s", string(serr))
	}
	ret = gose.LatToLmt(tjdLat[0], geolon, tjdLmt, serr)
	if ret < 0 {
		t.Fatalf("LatToLmt failed: %s", string(serr))
	}
	if math.Abs(tjdLmt[0]-tjd) > 1e-9 {
		t.Fatalf("LMT <-> LAT roundtrip mismatch: got %f, expected %f", tjdLmt[0], tjd)
	}

	// UtcToJd and JdetToUtc round-trip
	dret := make([]float64, 2)
	ret = gose.UtcToJd(2024, 5, 20, 15, 30, 45.0, gose.SeGregCal, dret, serr)
	if ret < 0 {
		t.Fatalf("UtcToJd failed: %s", string(serr))
	}
	iyear := make([]int, 1)
	imonth := make([]int, 1)
	iday := make([]int, 1)
	ihour := make([]int, 1)
	imin := make([]int, 1)
	dsec := make([]float64, 1)
	gose.JdetToUtc(dret[0], gose.SeGregCal, iyear, imonth, iday, ihour, imin, dsec)
	if iyear[0] != 2024 || imonth[0] != 5 || iday[0] != 20 || ihour[0] != 15 || imin[0] != 30 {
		t.Fatalf("JdetToUtc mismatch: %d-%d-%d %d:%d", iyear[0], imonth[0], iday[0], ihour[0], imin[0])
	}
}

// 10. Fixed Stars
func TestFixedStars(t *testing.T) {
	tjd := 2451545.0
	xx := make([]float64, 6)
	serr := make([]byte, 256)

	// Sirius
	ret := gose.FixstarUt([]byte("Sirius"), tjd, gose.SeflgSpeed, xx, serr)
	if ret < 0 {
		t.Fatalf("FixstarUt Sirius failed: %s", string(serr))
	}
	// Sirius longitude ~104° (Cancer), latitude ~ -39.6°
	if xx[0] < 103.0 || xx[0] > 105.0 || xx[1] > -38.0 || xx[1] < -41.0 {
		t.Fatalf("Sirius coordinates out of range: lon=%f, lat=%f", xx[0], xx[1])
	}

	// Sirius magnitude is -1.46
	mag := make([]float64, 1)
	ret = gose.FixstarMag([]byte("Sirius"), mag, serr)
	if ret < 0 {
		t.Fatalf("FixstarMag Sirius failed: %s", string(serr))
	}
	if math.Abs(mag[0]-(-1.46)) > 0.1 {
		t.Fatalf("Sirius magnitude out of range: %f (expected ~ -1.46)", mag[0])
	}

	// Aldebaran
	ret = gose.FixstarUt([]byte("Aldebaran"), tjd, gose.SeflgSpeed, xx, serr)
	if ret < 0 {
		t.Fatalf("FixstarUt Aldebaran failed: %s", string(serr))
	}
	// Aldebaran longitude ~69.8° (Gemini)
	if xx[0] < 68.5 || xx[0] > 71.0 {
		t.Fatalf("Aldebaran longitude out of range: %f", xx[0])
	}
}

// 11. Planetary Phenomena
func TestPlanetaryPheno(t *testing.T) {
	tjd := 2451545.0
	attr := make([]float64, 20)
	serr := make([]byte, 256)

	// Venus phenomena
	ret := gose.PhenoUt(tjd, gose.SeVenus, 0, attr, serr)
	if ret < 0 {
		t.Fatalf("PhenoUt Venus failed: %s", string(serr))
	}
	// Phase angle in [0, 180]
	if attr[0] < 0 || attr[0] > 180 {
		t.Fatalf("invalid phase angle: %f", attr[0])
	}
	// Phase in [0, 1]
	if attr[1] < 0 || attr[1] > 1 {
		t.Fatalf("invalid phase: %f", attr[1])
	}
}

// 12. Planetary Nodes and Apsides
func TestPlanetaryNodesAndApsides(t *testing.T) {
	tjd := 2451545.0
	xnasc := make([]float64, 6)
	xndsc := make([]float64, 6)
	xperi := make([]float64, 6)
	xaphe := make([]float64, 6)
	serr := make([]byte, 256)

	// Mars nodes and apsides
	ret := gose.NodApsUt(tjd, gose.SeMars, gose.SeflgSpeed, 0, xnasc, xndsc, xperi, xaphe, serr)
	if ret < 0 {
		t.Fatalf("NodApsUt Mars failed: %s", string(serr))
	}
	// Ascending node of Mars is near 49° in ecliptic coordinates
	if xnasc[0] < 0 || xnasc[0] >= 360 {
		t.Fatalf("invalid ascending node longitude: %f", xnasc[0])
	}
	if xperi[0] < 0 || xperi[0] >= 360 {
		t.Fatalf("invalid perihelion longitude: %f", xperi[0])
	}

	// Geocentric orbit distances (Earth-Mars)
	dmax := make([]float64, 1)
	dmin := make([]float64, 1)
	dtrue := make([]float64, 1)
	ret = gose.OrbitMaxMinTrueDistance(tjd, gose.SeMars, 0, dmax, dmin, dtrue, serr)
	if ret < 0 {
		t.Fatalf("OrbitMaxMinTrueDistance geocentric failed: %s", string(serr))
	}
	if dmin[0] < 0.35 || dmin[0] > 0.40 || dmax[0] < 2.60 || dmax[0] > 2.75 {
		t.Fatalf("Mars geocentric distances out of expected range: min=%f, max=%f", dmin[0], dmax[0])
	}

	// Heliocentric orbit distances (Sun-Mars perihelion/aphelion)
	ret = gose.OrbitMaxMinTrueDistance(tjd, gose.SeMars, gose.SeflgHelctr, dmax, dmin, dtrue, serr)
	if ret < 0 {
		t.Fatalf("OrbitMaxMinTrueDistance heliocentric failed: %s", string(serr))
	}
	if dmin[0] < 1.35 || dmin[0] > 1.42 || dmax[0] < 1.63 || dmax[0] > 1.70 {
		t.Fatalf("Mars heliocentric distances out of expected range: min=%f, max=%f", dmin[0], dmax[0])
	}
}

// 13. Orbital Elements & Heliocentric Crossings
func TestOrbitalElementsAndHeliocentric(t *testing.T) {
	tjd := 2451545.0
	dret := make([]float64, 50)
	serr := make([]byte, 256)

	// Orbital elements of Jupiter
	ret := gose.GetOrbitalElements(tjd, gose.SeJupiter, 0, dret, serr)
	if ret < 0 {
		t.Fatalf("GetOrbitalElements failed: %s", string(serr))
	}
	// Semimajor axis of Jupiter is ~5.2 AU
	semimajor := dret[0]
	if semimajor < 5.1 || semimajor > 5.3 {
		t.Fatalf("Jupiter semimajor axis out of range: %f (expected ~5.2 AU)", semimajor)
	}

	// Heliocentric crossing of Mars into Aries (0°)
	jdCross := make([]float64, 1)
	ret = gose.HelioCrossUt(gose.SeMars, 0.0, tjd, 0, 1, jdCross, serr)
	if ret < 0 || jdCross[0] < tjd {
		t.Fatalf("HelioCrossUt Mars failed: ret=%d, jd=%f", ret, jdCross[0])
	}
}

// 14. Coordinate Transformations (Cotrans & CotransSp)
func TestCoordinateTransformations(t *testing.T) {
	// Ecliptic coordinates of a point: lon=90.0, lat=0.0 (summer solstice point)
	// Obliquity of ecliptic eps = 23.44°
	eps := 23.4392911
	xpo := []float64{90.0, 0.0, 1.0}
	xpn := make([]float64, 3)

	// Swiss Ephemeris convention: ecliptic to equatorial requires -eps; equatorial to ecliptic requires +eps
	gose.Cotrans(xpo, xpn, -eps)
	// In equatorial coordinates, RA should be 90° (6h), declination should equal obliquity ~23.44°
	if math.Abs(xpn[0]-90.0) > 0.01 || math.Abs(xpn[1]-eps) > 0.01 {
		t.Fatalf("Cotrans failed: expected (90, %f), got (%f, %f)", eps, xpn[0], xpn[1])
	}

	// Reverse transformation: equatorial to ecliptic with +eps
	xrev := make([]float64, 3)
	gose.Cotrans(xpn, xrev, eps)
	if math.Abs(xrev[0]-xpo[0]) > 0.01 || math.Abs(xrev[1]-xpo[1]) > 0.01 {
		t.Fatalf("Cotrans reverse failed: expected (90, 0), got (%f, %f)", xrev[0], xrev[1])
	}
}

// 15. Math and Angle Normalization
func TestMathHelpers(t *testing.T) {
	// Degnorm
	if math.Abs(gose.Degnorm(370.0)-10.0) > 1e-9 {
		t.Fatalf("Degnorm(370) failed")
	}
	if math.Abs(gose.Degnorm(-10.0)-350.0) > 1e-9 {
		t.Fatalf("Degnorm(-10) failed")
	}

	// Radnorm
	if math.Abs(gose.Radnorm(2*math.Pi+0.5)-0.5) > 1e-9 {
		t.Fatalf("Radnorm failed")
	}

	// DegMidp & RadMidp
	if math.Abs(gose.DegMidp(10.0, 30.0)-20.0) > 1e-9 {
		t.Fatalf("DegMidp failed")
	}

	// Difdegn
	if math.Abs(gose.Difdegn(10.0, 350.0)-20.0) > 1e-9 {
		t.Fatalf("Difdegn failed")
	}
	// Difdeg2n
	if math.Abs(gose.Difdeg2n(350.0, 10.0)-(-20.0)) > 1e-9 {
		t.Fatalf("Difdeg2n failed")
	}

	// SplitDeg
	ideg := make([]int, 1)
	imin := make([]int, 1)
	isec := make([]int, 1)
	dsecfr := make([]float64, 1)
	isgn := make([]int32, 1)

	gose.SplitDeg(15.505, gose.SeSplitDegRoundSec, ideg, imin, isec, dsecfr, isgn)
	if ideg[0] != 15 || imin[0] != 30 || isec[0] != 18 || isgn[0] != 1 {
		t.Fatalf("SplitDeg failed: %d deg %d min %d sec sign=%d", ideg[0], imin[0], isec[0], isgn[0])
	}

	// Centisec helpers (Deg360 = 129,600,000 centiseconds = 360 degrees)
	cs := gose.Csnorm(gose.Deg360 + 1000)
	if cs != 1000 {
		t.Fatalf("Csnorm failed: got %d", cs)
	}
}

// 16. Memory Safety with Nil and Undersized Slices
func TestMemorySafetyNilAndUndersizedBuffers(t *testing.T) {
	tjd := 2451545.0

	// Test CalcUt with nil and empty slices (must NOT panic or segfault)
	ret := gose.CalcUt(tjd, gose.SeSun, gose.SeflgSpeed, nil, nil)
	_ = ret

	ret = gose.CalcUt(tjd, gose.SeSun, gose.SeflgSpeed, make([]float64, 2), make([]byte, 10))
	_ = ret

	// Test Revjul with nil slices (must NOT panic or segfault)
	gose.Revjul(tjd, gose.SeGregCal, nil, nil, nil, nil)

	// Test SplitDeg with nil slices
	gose.SplitDeg(15.5, 0, nil, nil, nil, nil, nil)

	// Test Houses with nil / short slices
	ret = gose.Houses(tjd, 51.5, -0.12, int('P'), nil, nil)
	_ = ret

	// TestMemorySafetyNilAndUndersizedBuffers
	// Test non-null-terminated string bytes
	gose.SetEphePath([]byte("/users/ephe")) // non-null terminated slice
	var xx [6]float64
	var serr [256]byte
	gose.Fixstar([]byte("Sirius"), tjd, 0, xx[:], serr[:]) // non-null terminated star name
}

// 16b. Comprehensive Regression Tests for Latest swisseph Updates & Buffer Safety
func TestHousesGauquelinSunshineAndSpeeds(t *testing.T) {
	tjd := 2451545.0
	geolat := 51.5074
	geolon := -0.1278

	// 1. Gauquelin ('G') writes 37 cusps and 37 cusp speeds (indices 0..36).
	// Verify both full 37-element buffers and undersized 13-element / nil buffers (no heap corruption).
	cusps37 := make([]float64, 37)
	ascmc10 := make([]float64, 10)
	cuspSpeed37 := make([]float64, 37)
	ascmcSpeed10 := make([]float64, 10)
	serr := make([]byte, 256)

	ret := gose.HousesEx2(tjd, 0, geolat, geolon, int('G'), cusps37, ascmc10, cuspSpeed37, ascmcSpeed10, serr)
	if ret < 0 {
		t.Fatalf("HousesEx2 Gauquelin failed: %s", string(serr))
	}
	for i := 1; i <= 36; i++ {
		if cusps37[i] < 0 || cusps37[i] >= 360 {
			t.Fatalf("Gauquelin sector %d out of range: %f", i, cusps37[i])
		}
		if cuspSpeed37[i] == 0 {
			t.Fatalf("expected non-zero Gauquelin sector %d speed", i)
		}
	}

	// Undersized 13-element slice and nil slice with 'G' and 'g' must not overflow
	shortCusps := make([]float64, 13)
	shortSpeed := make([]float64, 13)
	_ = gose.Houses(tjd, geolat, geolon, int('G'), shortCusps, nil)
	_ = gose.HousesEx(tjd, 0, geolat, geolon, int('g'), shortCusps, ascmc10)
	_ = gose.HousesEx2(tjd, 0, geolat, geolon, int('G'), shortCusps, ascmc10, shortSpeed, nil, nil)
	_ = gose.HousesArmc(ascmc10[2], geolat, 23.44, int('G'), shortCusps, ascmc10)
	_ = gose.HousesArmcEx2(ascmc10[2], geolat, 23.44, int('G'), shortCusps, ascmc10, shortSpeed, nil, nil)

	// 2. Sunshine houses ('I' and 'i') in HousesArmcEx2 use ascmc[9] as input solar declination
	cusps13 := make([]float64, 13)
	cuspSpeed13 := make([]float64, 13)
	ascmcSun := make([]float64, 10)
	ascmcSun[9] = -23.0 // valid Sun declination
	ret = gose.HousesArmcEx2(ascmc10[2], geolat, 23.44, int('I'), cusps13, ascmcSun, cuspSpeed13, ascmcSpeed10, serr)
	if ret < 0 {
		t.Fatalf("HousesArmcEx2 Sunshine ('I') failed: %s", string(serr))
	}
	// Invalid declination (|ascmc[9]| >= 90) in HousesArmc must fail and zero cusps (latest swehouse.c fix)
	ascmcSun[9] = 100.0
	ret = gose.HousesArmcEx2(ascmc10[2], geolat, 23.44, int('I'), cusps13, ascmcSun, cuspSpeed13, ascmcSpeed10, serr)
	if ret != gose.Err {
		t.Fatalf("expected Err for invalid solar declination in Sunshine houses, got %d", ret)
	}

	// 3. Test Makransky ('I', 'i') in Southern Hemisphere and APC ('Y') at Equator (latest swehouse.c fixes)
	for _, sys := range []byte{'P', 'K', 'O', 'R', 'C', 'E', 'W', 'X', 'M', 'H', 'T', 'B', 'U', 'Y', 'L', 'Q', 'F', 'S', 'I', 'i'} {
		ret = gose.HousesEx2(tjd, 0, -33.8688, 151.2093, int(sys), cusps13, ascmc10, cuspSpeed13, ascmcSpeed10, serr)
		if ret < 0 {
			t.Fatalf("HousesEx2 southern hemisphere failed for system %c: %s", sys, string(serr))
		}
		for i := 1; i <= 12; i++ {
			if cusps13[i] < 0 || cusps13[i] >= 360 {
				t.Fatalf("system %c cusp %d out of range [0, 360): %f", sys, i, cusps13[i])
			}
			// Whole Sign ('W') cusps are fixed at 0° of each sign so their cuspSpeed is 0; all other systems move continuously
			if sys != 'W' && cuspSpeed13[i] == 0 {
				t.Fatalf("system %c cuspSpeed %d is zero", sys, i)
			}
		}
		if ascmcSpeed10[0] == 0 {
			t.Fatalf("system %c Ascendant speed is zero", sys)
		}
	}

	// 4. Test SeflgRadians scaling on cuspSpeed and ascmcSpeed (latest swehouse.c fix)
	cuspsDeg := make([]float64, 13)
	speedDeg := make([]float64, 13)
	ascmcSpeedDeg := make([]float64, 10)
	_ = gose.HousesEx2(tjd, 0, geolat, geolon, int('P'), cuspsDeg, ascmc10, speedDeg, ascmcSpeedDeg, serr)

	cuspsRad := make([]float64, 13)
	speedRad := make([]float64, 13)
	ascmcSpeedRad := make([]float64, 10)
	_ = gose.HousesEx2(tjd, gose.SeflgRadians, geolat, geolon, int('P'), cuspsRad, ascmc10, speedRad, ascmcSpeedRad, serr)

	if math.Abs(speedRad[1]-speedDeg[1]*gose.Degtorad) > 1e-9 {
		t.Fatalf("SeflgRadians cuspSpeed scaling mismatch: rad=%f, deg*Degtorad=%f", speedRad[1], speedDeg[1]*gose.Degtorad)
	}
	if math.Abs(ascmcSpeedRad[0]-ascmcSpeedDeg[0]*gose.Degtorad) > 1e-9 {
		t.Fatalf("SeflgRadians ascmcSpeed scaling mismatch: rad=%f, deg*Degtorad=%f", ascmcSpeedRad[0], ascmcSpeedDeg[0]*gose.Degtorad)
	}
}

func TestFixstarCopybackAndAllEdgeCases(t *testing.T) {
	tjd := 2451545.0
	xx := make([]float64, 6)
	mag := make([]float64, 1)
	serr := make([]byte, 256)

	// 1. When caller passes a buffer of sufficient size (e.g., 256 or 512 bytes),
	// Fixstar* must write back the canonical "StarName,BayerName" C string.
	starBuf := make([]byte, 256)
	copy(starBuf, "Spica")
	ret := gose.Fixstar2Ut(starBuf, tjd, 0, xx, serr)
	if ret < 0 {
		t.Fatalf("Fixstar2Ut Spica failed: %s", string(serr))
	}
	n := bytes.IndexByte(starBuf, 0)
	if n < 0 {
		n = len(starBuf)
	}
	gotStar := string(starBuf[:n])
	if !strings.Contains(gotStar, "Spica,alVir") {
		t.Fatalf("expected starBuf to be updated to 'Spica,alVir', got %q", gotStar)
	}

	// 2. When caller passes a short slice ([]byte("Spica") or []byte("Spica\x00")),
	// Fixstar* must not overflow or corrupt the slice.
	shortStar := []byte("Spica")
	ret = gose.Fixstar(shortStar, tjd, 0, xx, serr)
	if ret < 0 || string(shortStar) != "Spica" {
		t.Fatalf("Fixstar with short slice failed or mutated input: ret=%d, star=%q", ret, string(shortStar))
	}
	_ = gose.FixstarMag([]byte("Spica\x00"), mag, serr)
	_ = gose.Fixstar2([]byte("Spica"), tjd, 0, xx, serr)
	_ = gose.Fixstar2Mag([]byte("Spica"), mag, serr)

	// 3. Nil and empty star slices must not segfault
	_ = gose.Fixstar(nil, tjd, 0, nil, nil)
	_ = gose.FixstarUt([]byte{}, tjd, 0, nil, nil)
	_ = gose.FixstarMag(nil, nil, nil)
	_ = gose.Fixstar2(nil, tjd, 0, nil, nil)
	_ = gose.Fixstar2Ut(nil, tjd, 0, nil, nil)
	_ = gose.Fixstar2Mag(nil, nil, nil)

	// 4. Astro models with nil, short, and "+" (detailed > 2000 bytes output)
	gose.SetAstroModels(nil, 0)
	gose.SetAstroModels([]byte("default"), 0)
	gose.GetAstroModels(nil, nil, 0)

	samod := make([]byte, 64)
	copy(samod, "+")
	sdet := make([]byte, 4096)
	gose.GetAstroModels(samod, sdet, 0)
	detLen := bytes.IndexByte(sdet, 0)
	if detLen <= 500 {
		t.Fatalf("expected detailed GetAstroModels output (>500 bytes) with '+', got %d bytes", detLen)
	}
	// Also test GetAstroModels with a small sdet buffer (e.g. 64 bytes) and '+' to ensure no overflow
	copy(samod, "+")
	smallDet := make([]byte, 64)
	gose.GetAstroModels(samod, smallDet, 0)

	// 5. Negative and out-of-range isidmode in GetAyanamsaName must not read out-of-bounds
	if name := gose.GetAyanamsaName(-1); name != "" {
		t.Fatalf("expected empty string for GetAyanamsaName(-1), got %q", name)
	}
	if name := gose.GetAyanamsaName(-256); name != "" {
		t.Fatalf("expected empty string for GetAyanamsaName(-256), got %q", name)
	}
	if name := gose.GetAyanamsaName(999); name != "" {
		t.Fatalf("expected empty string for GetAyanamsaName(999), got %q", name)
	}

	// 6. Nil input slices on all C wrappers must never segfault
	gose.SetJplFile(nil)
	gose.SetEphePath(nil)
	_ = gose.GetPlanetName(gose.SeSun, nil)
	shortName := make([]byte, 2)
	if gose.GetPlanetName(gose.SeSun, shortName) != &shortName[0] || string(shortName) != "Su" {
		t.Fatalf("unexpected GetPlanetName short buffer: %q", string(shortName))
	}
	_ = gose.Cs2timestr(360000, ':', 0, nil)
	_ = gose.Cs2lonlatstr(360000, 'E', 'W', nil)
	_ = gose.Cs2degstr(360000, nil)
	gose.Cotrans(nil, nil, 23.44)
	gose.CotransSp(nil, nil, 23.44)
	gose.Azalt(tjd, gose.SeEcl2hor, nil, 0, 0, nil, nil)
	gose.AzaltRev(tjd, gose.SeHor2ecl, nil, nil, nil)
	_ = gose.HousePos(0, 0, 23.44, int('P'), nil, nil)
	_ = gose.GauquelinSector(tjd, gose.SeSun, nil, 0, 0, nil, 0, 0, nil, nil)
	_ = gose.HeliacalUt(tjd, nil, nil, nil, nil, gose.SeHeliacalRising, 0, nil, nil)
	_ = gose.HeliacalPhenoUt(tjd, nil, nil, nil, nil, gose.SeHeliacalRising, 0, nil, nil)
	_ = gose.VisLimitMag(tjd, nil, nil, nil, nil, 0, nil, nil)
	_ = gose.HeliacalAngle(tjd, nil, nil, nil, 0, 0, 0, 0, 0, 0, nil, nil)
	_ = gose.TopoArcusVisionis(tjd, nil, nil, nil, 0, 0, 0, 0, 0, 0, 0, nil, nil)
	_ = gose.SolEclipseHow(tjd, 0, nil, nil, nil)
	_ = gose.LunEclipseHow(tjd, 0, nil, nil, nil)
	_ = gose.RiseTrans(tjd, gose.SeSun, nil, 0, gose.SeCalcRise, nil, 0, 0, nil, nil)
	_ = gose.RiseTransTrueHor(tjd, gose.SeSun, nil, 0, gose.SeCalcRise, nil, 0, 0, 0, nil, nil)

	// Restore ephemeris path
	if abs, err := filepath.Abs("../swisseph/ephe"); err == nil {
		gose.SetEphePathStr(abs)
	}
}

// 17. High-Concurrency Stress Test
func TestConcurrentAccess(t *testing.T) {
	// Verify thread safety and absence of race conditions across multiple goroutines
	var wg sync.WaitGroup
	tjdBase := 2451545.0

	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func(offset float64) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				tjd := tjdBase + offset + float64(j)
				// Call planetary calculation
				var xx [6]float64
				var serr [256]byte
				_ = gose.CalcUt(tjd, gose.SeSun, gose.SeflgSpeed, xx[:], serr[:])

				// Call house calculation
				var cusps [13]float64
				var ascmc [10]float64
				_ = gose.Houses(tjd, 51.5, -0.12, int('P'), cusps[:], ascmc[:])

				// Call Julian Day
				_ = gose.Julday(2024, 5, 20, 12.0, gose.SeGregCal)

				// Call Sidereal Time
				_ = gose.Sidtime(tjd)
			}
		}(float64(i) * 10.0)
	}

	wg.Wait()
}

// 18. Performance Benchmarks
func BenchmarkCalcUt(b *testing.B) {
	xx := make([]float64, 6)
	serr := make([]byte, 256)
	tjd := 2451545.0
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = gose.CalcUt(tjd, gose.SeSun, gose.SeflgSpeed, xx, serr)
	}
}

func BenchmarkCalcUtSimple(b *testing.B) {
	tjd := 2451545.0
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, _ = gose.CalcUtSimple(tjd, gose.SeSun, gose.SeflgSpeed)
	}
}

func BenchmarkJulday(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = gose.Julday(2024, 5, 20, 12.0, gose.SeGregCal)
	}
}

func BenchmarkHouses(b *testing.B) {
	tjd := 2451545.0
	cusps := make([]float64, 13)
	ascmc := make([]float64, 10)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = gose.Houses(tjd, 51.5, -0.12, int('P'), cusps, ascmc)
	}
}

func BenchmarkDegnorm(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = gose.Degnorm(375.25)
	}
}

func BenchmarkDeltat(b *testing.B) {
	tjd := 2451545.0
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = gose.Deltat(tjd)
	}
}
