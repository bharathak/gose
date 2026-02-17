package main

import (
	"fmt"
	"github.com/bharathak/gose"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Starting Swiss Ephemeris Example...")

	// Get library version
	sweVer := make([]byte, 256)
	gose.Version(sweVer)
	// Find null terminator
	n := 0
	for i, b := range sweVer {
		if b == 0 {
			n = i
			break
		}
	}
	fmt.Printf("Library used: Swiss Ephemeris v%s\n", string(sweVer[:n]))

	// Set ephemeris path
	// By default, we look for 'ephe' in the current directory or parent
	ephePath := "./ephe"
	if _, err := os.Stat(ephePath); os.IsNotExist(err) {
		// Fallback to source location if running in this environment
		ephePath = "/users/ephe"
	}

	absPath, _ := filepath.Abs(ephePath)
	fmt.Printf("Setting ephemeris path to: %s\n", absPath)
	gose.SetEphePath([]byte(absPath))

	// Test Julday
	tjdUt := gose.Julday(2024, 5, 20, 12.0, gose.SeGregCal)
	fmt.Printf("Julian Day for 2024-05-20 12:00 UTC: %.6f\n", tjdUt)

	// Test CalcUt (Sun)
	xx := make([]float64, 6)
	serr := make([]byte, 256)
	iflgret := gose.CalcUt(tjdUt, gose.SeSun, gose.SeflgSpeed, xx, serr)

	sn := 0
	for i, b := range serr {
		if b == 0 {
			sn = i
			break
		}
	}

	if iflgret < 0 {
		fmt.Printf("CalcUt error: %s\n", string(serr[:sn]))
	} else {
		fmt.Printf("Sun Longitude: %.6f, Latitude: %.6f, Distance: %.6f AU\n", xx[0], xx[1], xx[2])
	}

	// Test New Function: SolcrossUt
	// Find next time Sun reaches 0 degrees Longitude (Aries Ingress)
	serrCross := make([]byte, 256)
	jdCross := gose.SolcrossUt(0, tjdUt, 0, serrCross)

	snc := 0
	for i, b := range serrCross {
		if b == 0 {
			snc = i
			break
		}
	}

	if jdCross < 0 && snc > 0 {
		fmt.Printf("SolcrossUt error: %s\n", string(serrCross[:snc]))
	} else {
		fmt.Printf("Next Sun 0° Longitude (UT): %.6f\n", jdCross)

		// Convert back to date
		jyear := make([]int, 1)
		jmon := make([]int, 1)
		jday := make([]int, 1)
		jut := make([]float64, 1)
		gose.Revjul(jdCross, gose.SeGregCal, jyear, jmon, jday, jut)
		fmt.Printf("Date of Ingress: %d-%02d-%02d %.2f hours\n", jyear[0], jmon[0], jday[0], jut[0])
	}

	gose.Close()
	fmt.Println("Done.")
}
