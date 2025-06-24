package main

import (
	"fmt"
	"nlpconverter/converter"
)

func runTests(title string, c *converter.Converter, testCases []string) {
	fmt.Printf("🚀 NLP Unit Converter Initialized for %s\n", title)
	fmt.Println("-------------------------------------------")

	for _, tc := range testCases {
		fmt.Printf("Input:  \"%s\"\n", tc)
		result, err := c.Process(tc)
		if err != nil {
			fmt.Printf("Error:  %v\n", err)
		} else {
			fmt.Printf("Result: %g %s (%s)\n", result.Value, result.UnitSymbol, result.UnitName)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	volumeSystem := converter.NewVolumeSystem()
	volumeConverter := converter.NewConverter(volumeSystem)
	volumeTestCases := []string{
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
	}
	runTests("Volume", volumeConverter, volumeTestCases)

	lengthSystem := converter.NewLengthSystem()
	lengthConverter := converter.NewConverter(lengthSystem)
	lengthTestCases := []string{
		"1 km in miles",
		"a foot and 5 inches in cm",
		"100 meters + 0.1km in ft",
	}
	runTests("Length", lengthConverter, lengthTestCases)

	weightSystem := converter.NewWeightSystem()
	weightConverter := converter.NewConverter(weightSystem)
	weightTestCases := []string{
		"1kg in lbs",
		"two pounds and 8 ounces in grams",
		"100g + .5kg",
	}
	runTests("Weight", weightConverter, weightTestCases)
}