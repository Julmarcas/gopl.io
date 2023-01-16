package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch2/tempconv"
)

type Feet float64
type Meter float64

func FeetsToMeters(f Feet) Meter { return Meter(f * 0.3048) }
func MetersToFeets(m Meter) Feet { return Feet(m / 0.3048) }
func (f Feet) String() string    { return fmt.Sprintf("%g feets", f) }
func (m Meter) String() string   { return fmt.Sprintf("%g meters", m) }

func main() {

	var values []string
	if len(os.Args) > 1 {
		fmt.Printf("args len %d", len(os.Args))
		values = os.Args[1:]
	} else {
		fmt.Print("Enter numbers separated by spaces to be transformed: ")
		scanner := bufio.NewScanner((os.Stdin))
		scanner.Scan()
		text := scanner.Text()
		values = strings.Split(text, " ")
	}

	for _, arg := range values {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		feets := Feet(t)
		meters := Meter(t)
		fmt.Printf("Temp -> %s = %s, %s = %s\nFeet-Meters -> %s = %s, %s = %s",
			f, tempconv.FToC(f), c, tempconv.CToF(c), feets, FeetsToMeters(feets), meters, MetersToFeets(meters))
	}
}
