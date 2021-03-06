
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const miTokm float64 = 1.60934
const miToMeters float64 = 1609.34
const miToFeet float64 = 5280
const miTomi float64 = 1

const kmtomi float64 = 0.621371
const meterstomi float64 = 0.000621371
const feettomi float64 = 0.000189394

func main() {
	myFrom := os.Args[1]
	myTo := os.Args[2]

	switch {
	case strings.HasSuffix(myFrom, "mi"):
		switch myTo {
		case "miles":
			distance, _ := strconv.ParseFloat(myFrom[:len(myFrom)-2], 64)
			fmt.Println(distance * miTomi, "miles")
		case "km":
			distance, _ := strconv.ParseFloat(myFrom[:len(myFrom)-2], 64)
			fmt.Println(distance * miTokm, "kilometers")
			fmt.Println(distance / kmtomi, "kilometers")
		case "m":
			distance, _ := strconv.ParseFloat(myFrom[:len(myFrom)-2], 64)
			fmt.Println(distance * miToMeters, "meters")
			fmt.Println(distance / meterstomi, "meters")
		case "ft":
			distance, _ := strconv.ParseFloat(myFrom[:len(myFrom)-2], 64)
			fmt.Println(distance * miToFeet, "feet")
			fmt.Println(distance / feettomi, "feet")
		}
	}
}

/*

at terminal:
go install
08_03_challenge 50mi km

returns this:
80.467

*/

/*
Create a program which parses a query to do distance conversions.

For example, from a terminal:
	$ distance_converter 50mi km
Should produce:
	80.47km

It should support miles (mi), kilometers (km), feet (ft) and meters (m)

*/