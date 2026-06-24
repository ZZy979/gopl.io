// Unit-conversion program. Reads numbers and converts each number into
// temperature in °C/°F, length in ft/m, weight in lb/kg.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

const (
	MeterPerFeet = 0.3048
	KgPerLb      = 0.453592
)

func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		// temperature
		tempC := tempconv.Celsius(x)
		tempF := tempconv.Fahrenheit(x)
		fmt.Printf("%s = %s, %s = %s\n",
			tempC, tempconv.CToF(tempC), tempF, tempconv.FToC(tempF))

		// length
		lengthFeet := x / MeterPerFeet
		lengthMeter := x * MeterPerFeet
		fmt.Printf("%.6g m = %.6g ft, %.6g ft = %.6g m\n",
			x, lengthFeet, x, lengthMeter)

		// weight
		weightPound := x / KgPerLb
		weightKg := x * KgPerLb
		fmt.Printf(
			"%.6g kg = %.6g lb, %.6g lb = %.6g kg\n\n",
			x, weightPound, x, weightKg)
	}
}
