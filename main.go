package main

import (
	"flag"
	"fmt"
	"nlpconverter/converter"
	"os"
	"strings"
)

var testCases = []string{
	// Volume
	"1L & 23 ml",
	"1Liter + 100.87 milli in cm^3",
	".5 gal in L",
	"1.5e3 ml in L",
	"1/2 gallon + 1/4 pint in cups",
	"two pints and a half cup in floz",
	"500ml - .25L",
	"1 leter in ml",
	"2 gallens in L",
	"one gallon and 2.5 litres in ml",

	// Length
	"1 km in miles",
	"a foot and 5 inches in cm",
	"100 meters + 0.1km in ft",

	// Weight
	"1kg in lbs",
	"two pounds and 8 ounces in grams",
	"100g + .5kg",

	// Temperature
	"100 C in F",
	"212 f in C",
	"0c in k",

	// Area
	"100 sqft in m2",
	"2 acres in ha",

	// Speed
	"60 mph in kph",
	"100 km/h in knots",

	// Compound
	"10 km / 2 hr in m/s",

	// Time
	"1 day in hours",

	// Previously failing
	"two pints + a half cup in floz",
	"one gallon + 2.5 litres in ml",
	"a foot + 5 inches in cm",
	"two pounds + 8 ounces in grams",
}

func printHelp() {
	fmt.Println("Usage: nlp-unit-converter [expression]")
	fmt.Println("\nEvaluates a natural language expression of units and converts them.")
	fmt.Println("\nFlags:")
	fmt.Println("  -h, --help\t\tPrints this help message.")
	fmt.Println("\nExamples:")
	fmt.Println("| Expression                           | Result                                  |")
	fmt.Println("|------------------------------------|-----------------------------------------|")
	unitMap := converter.MustRegisterSystems()
	conv := converter.NewConverter(unitMap)
	for _, tc := range testCases {
		result, err := conv.Process(tc)
		if err != nil {
			fmt.Printf("| %-34s | Error: %-29s |\n", tc, err.Error())
		} else {
			resultStr := fmt.Sprintf("%g %s (%s)", result.Value, result.UnitSymbol, result.UnitName)
			fmt.Printf("| %-34s | %-39s |\n", tc, resultStr)
		}
	}
	fmt.Println("|------------------------------------|-----------------------------------------|")
}

func main() {
	help := flag.Bool("h", false, "Prints the help message.")
	flag.BoolVar(help, "help", false, "Prints the help message.")
	flag.Parse()

	if *help {
		printHelp()
		os.Exit(0)
	}

	unitMap := converter.MustRegisterSystems()
	conv := converter.NewConverter(unitMap)

	if len(os.Args) == 1 {
		fmt.Println("No expression provided. Use -h or --help for usage information.")
		os.Exit(1)
	}

	expression := strings.Join(os.Args[1:], " ")
	result, err := conv.Process(expression)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%g %s (%s)\n", result.Value, result.UnitSymbol, result.UnitName)
		os.Exit(0)
	}
}