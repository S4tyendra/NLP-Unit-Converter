package converter

import "strings"

type Unit struct {
	Name         string
	Symbol       string
	Aliases      []string
	ToBaseFunc   func(float64) float64
	FromBaseFunc func(float64) float64
}

type UnitSystem struct {
	Name     string
	BaseUnit string
	Units    map[string]Unit
}

func NewVolumeSystem() UnitSystem {
	flOzToMl := 29.5735295625

	return UnitSystem{
		Name:     "Volume",
		BaseUnit: "Milliliters",
		Units: map[string]Unit{
			"Milliliters": {
				Name:         "Milliliters",
				Symbol:       "mL",
				Aliases:      []string{"ml", "milliliter", "milliliters", "millilitre", "millilitres", "milli"},
				ToBaseFunc:   func(val float64) float64 { return val },
				FromBaseFunc: func(val float64) float64 { return val },
			},
			"Liters": {
				Name:         "Liters",
				Symbol:       "L",
				Aliases:      []string{"l", "liter", "liters", "litre", "litres"},
				ToBaseFunc:   func(val float64) float64 { return val * 1000.0 },
				FromBaseFunc: func(val float64) float64 { return val / 1000.0 },
			},
			"Cubic meters": {
				Name:         "Cubic meters",
				Symbol:       "m³",
				Aliases:      []string{"m3", "m^3", "cubicmeter", "cubicmeters"},
				ToBaseFunc:   func(val float64) float64 { return val * 1000000.0 },
				FromBaseFunc: func(val float64) float64 { return val / 1000000.0 },
			},
			"Cubic centimeters": {
				Name:         "Cubic centimeters",
				Symbol:       "cm³",
				Aliases:      []string{"cm3", "cm^3", "cubiccentimeter", "cubiccentimeters", "cc"},
				ToBaseFunc:   func(val float64) float64 { return val },
				FromBaseFunc: func(val float64) float64 { return val },
			},
			"Fluid Ounce": {
				Name:         "Fluid Ounce",
				Symbol:       "fl oz",
				Aliases:      []string{"floz", "fluidounce", "fluidounces", "oz"},
				ToBaseFunc:   func(val float64) float64 { return val * flOzToMl },
				FromBaseFunc: func(val float64) float64 { return val / flOzToMl },
			},
			"Cup": {
				Name:         "Cup",
				Symbol:       "c",
				Aliases:      []string{"cup", "cups"},
				ToBaseFunc:   func(val float64) float64 { return val * flOzToMl * 8 },
				FromBaseFunc: func(val float64) float64 { return val / (flOzToMl * 8) },
			},
			"Pint": {
				Name:         "Pint",
				Symbol:       "pt",
				Aliases:      []string{"pint", "pints"},
				ToBaseFunc:   func(val float64) float64 { return val * flOzToMl * 16 },
				FromBaseFunc: func(val float64) float64 { return val / (flOzToMl * 16) },
			},
			"Quart": {
				Name:         "Quart",
				Symbol:       "qt",
				Aliases:      []string{"quart", "quarts"},
				ToBaseFunc:   func(val float64) float64 { return val * flOzToMl * 32 },
				FromBaseFunc: func(val float64) float64 { return val / (flOzToMl * 32) },
			},
			"Gallon": {
				Name:         "Gallon",
				Symbol:       "gal",
				Aliases:      []string{"gallon", "gallons"},
				ToBaseFunc:   func(val float64) float64 { return val * flOzToMl * 128 },
				FromBaseFunc: func(val float64) float64 { return val / (flOzToMl * 128) },
			},
			"Cubic feet": {
				Name:         "Cubic feet",
				Symbol:       "ft³",
				Aliases:      []string{"ft3", "ft^3", "cubicfoot", "cubicfeet"},
				ToBaseFunc:   func(val float64) float64 { return val * 28316.8 },
				FromBaseFunc: func(val float64) float64 { return val / 28316.8 },
			},
			"Barrels": {
				Name:         "Barrels",
				Symbol:       "bbl",
				Aliases:      []string{"barrel", "barrels"},
				ToBaseFunc:   func(val float64) float64 { return val * 158987.3 },
				FromBaseFunc: func(val float64) float64 { return val / 158987.3 },
			},
		},
	}
}

func NewLengthSystem() UnitSystem {
	return UnitSystem{
		Name:     "Length",
		BaseUnit: "Meters",
		Units: map[string]Unit{
			"Meters": {
				Name:         "Meters",
				Symbol:       "m",
				Aliases:      []string{"m", "meter", "meters", "metre", "metres"},
				ToBaseFunc:   func(val float64) float64 { return val },
				FromBaseFunc: func(val float64) float64 { return val },
			},
			"Kilometers": {
				Name:         "Kilometers",
				Symbol:       "km",
				Aliases:      []string{"km", "kilometer", "kilometers", "kilometre", "kilometres"},
				ToBaseFunc:   func(val float64) float64 { return val * 1000.0 },
				FromBaseFunc: func(val float64) float64 { return val / 1000.0 },
			},
			"Centimeters": {
				Name:         "Centimeters",
				Symbol:       "cm",
				Aliases:      []string{"cm", "centimeter", "centimeters", "centimetre", "centimetres"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.01 },
				FromBaseFunc: func(val float64) float64 { return val / 0.01 },
			},
			"Millimeters": {
				Name:         "Millimeters",
				Symbol:       "mm",
				Aliases:      []string{"mm", "millimeter", "millimeters", "millimetre", "millimetres"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.001 },
				FromBaseFunc: func(val float64) float64 { return val / 0.001 },
			},
			"Inches": {
				Name:         "Inches",
				Symbol:       "in",
				Aliases:      []string{"in", "inch", "inches"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.0254 },
				FromBaseFunc: func(val float64) float64 { return val / 0.0254 },
			},
			"Feet": {
				Name:         "Feet",
				Symbol:       "ft",
				Aliases:      []string{"ft", "foot", "feet"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.3048 },
				FromBaseFunc: func(val float64) float64 { return val / 0.3048 },
			},
			"Yards": {
				Name:         "Yards",
				Symbol:       "yd",
				Aliases:      []string{"yd", "yard", "yards"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.9144 },
				FromBaseFunc: func(val float64) float64 { return val / 0.9144 },
			},
			"Miles": {
				Name:         "Miles",
				Symbol:       "mi",
				Aliases:      []string{"mi", "mile", "miles"},
				ToBaseFunc:   func(val float64) float64 { return val * 1609.34 },
				FromBaseFunc: func(val float64) float64 { return val / 1609.34 },
			},
		},
	}
}

func NewWeightSystem() UnitSystem {
	return UnitSystem{
		Name:     "Weight",
		BaseUnit: "Grams",
		Units: map[string]Unit{
			"Grams": {
				Name:         "Grams",
				Symbol:       "g",
				Aliases:      []string{"g", "gram", "grams"},
				ToBaseFunc:   func(val float64) float64 { return val },
				FromBaseFunc: func(val float64) float64 { return val },
			},
			"Kilograms": {
				Name:         "Kilograms",
				Symbol:       "kg",
				Aliases:      []string{"kg", "kilogram", "kilograms"},
				ToBaseFunc:   func(val float64) float64 { return val * 1000.0 },
				FromBaseFunc: func(val float64) float64 { return val / 1000.0 },
			},
			"Milligrams": {
				Name:         "Milligrams",
				Symbol:       "mg",
				Aliases:      []string{"mg", "milligram", "milligrams"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.001 },
				FromBaseFunc: func(val float64) float64 { return val / 0.001 },
			},
			"Pounds": {
				Name:         "Pounds",
				Symbol:       "lb",
				Aliases:      []string{"lb", "lbs", "pound", "pounds"},
				ToBaseFunc:   func(val float64) float64 { return val * 453.592 },
				FromBaseFunc: func(val float64) float64 { return val / 453.592 },
			},
			"Ounces": {
				Name:         "Ounces",
				Symbol:       "oz",
				Aliases:      []string{"ounce", "ounces"},
				ToBaseFunc:   func(val float64) float64 { return val * 28.3495 },
				FromBaseFunc: func(val float64) float64 { return val / 28.3495 },
			},
		},
	}
}

func NewTemperatureSystem() UnitSystem {
	return UnitSystem{
		Name:     "Temperature",
		BaseUnit: "Celsius",
		Units: map[string]Unit{
			"Celsius": {
				Name:         "Celsius",
				Symbol:       "°C",
				Aliases:      []string{"c", "celsius"},
				ToBaseFunc:   func(val float64) float64 { return val },
				FromBaseFunc: func(val float64) float64 { return val },
			},
			"Fahrenheit": {
				Name:         "Fahrenheit",
				Symbol:       "°F",
				Aliases:      []string{"f", "fahrenheit"},
				ToBaseFunc:   func(val float64) float64 { return (val - 32) * 5 / 9 },
				FromBaseFunc: func(val float64) float64 { return (val * 9 / 5) + 32 },
			},
			"Kelvin": {
				Name:         "Kelvin",
				Symbol:       "K",
				Aliases:      []string{"k", "kelvin"},
				ToBaseFunc:   func(val float64) float64 { return val - 273.15 },
				FromBaseFunc: func(val float64) float64 { return val + 273.15 },
			},
		},
	}
}

func NewAreaSystem() UnitSystem {
	return UnitSystem{
		Name:     "Area",
		BaseUnit: "Square Meters",
		Units: map[string]Unit{
			"Square Meters": {
				Name:         "Square Meters",
				Symbol:       "m²",
				Aliases:      []string{"m2", "sqm", "squaremeter", "squaremeters"},
				ToBaseFunc:   func(val float64) float64 { return val },
				FromBaseFunc: func(val float64) float64 { return val },
			},
			"Square Kilometers": {
				Name:         "Square Kilometers",
				Symbol:       "km²",
				Aliases:      []string{"km2", "sqkm", "squarekilometer", "squarekilometers"},
				ToBaseFunc:   func(val float64) float64 { return val * 1000000.0 },
				FromBaseFunc: func(val float64) float64 { return val / 1000000.0 },
			},
			"Hectares": {
				Name:         "Hectares",
				Symbol:       "ha",
				Aliases:      []string{"ha", "hectare", "hectares"},
				ToBaseFunc:   func(val float64) float64 { return val * 10000.0 },
				FromBaseFunc: func(val float64) float64 { return val / 10000.0 },
			},
			"Square Miles": {
				Name:         "Square Miles",
				Symbol:       "mi²",
				Aliases:      []string{"mi2", "sqmi", "squaremile", "squaremiles"},
				ToBaseFunc:   func(val float64) float64 { return val * 2589988.11 },
				FromBaseFunc: func(val float64) float64 { return val / 2589988.11 },
			},
			"Acres": {
				Name:         "Acres",
				Symbol:       "ac",
				Aliases:      []string{"ac", "acre", "acres"},
				ToBaseFunc:   func(val float64) float64 { return val * 4046.86 },
				FromBaseFunc: func(val float64) float64 { return val / 4046.86 },
			},
			"Square Yards": {
				Name:         "Square Yards",
				Symbol:       "yd²",
				Aliases:      []string{"yd2", "sqyd", "squareyard", "squareyards"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.836127 },
				FromBaseFunc: func(val float64) float64 { return val / 0.836127 },
			},
			"Square Feet": {
				Name:         "Square Feet",
				Symbol:       "ft²",
				Aliases:      []string{"ft2", "sqft", "squarefoot", "squarefeet"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.092903 },
				FromBaseFunc: func(val float64) float64 { return val / 0.092903 },
			},
			"Square Inches": {
				Name:         "Square Inches",
				Symbol:       "in²",
				Aliases:      []string{"in2", "sqin", "squareinch", "squareinches"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.00064516 },
				FromBaseFunc: func(val float64) float64 { return val / 0.00064516 },
			},
		},
	}
}

func NewSpeedSystem() UnitSystem {
	return UnitSystem{
		Name:     "Speed",
		BaseUnit: "Meters per Second",
		Units: map[string]Unit{
			"Meters per Second": {
				Name:         "Meters per Second",
				Symbol:       "m/s",
				Aliases:      []string{"mps", "meterspersecond"},
				ToBaseFunc:   func(val float64) float64 { return val },
				FromBaseFunc: func(val float64) float64 { return val },
			},
			"Kilometers per Hour": {
				Name:         "Kilometers per Hour",
				Symbol:       "km/h",
				Aliases:      []string{"kph", "kmh", "kilometersperhour"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.277778 },
				FromBaseFunc: func(val float64) float64 { return val / 0.277778 },
			},
			"Miles per Hour": {
				Name:         "Miles per Hour",
				Symbol:       "mph",
				Aliases:      []string{"mph", "milesperhour"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.44704 },
				FromBaseFunc: func(val float64) float64 { return val / 0.44704 },
			},
			"Knots": {
				Name:         "Knots",
				Symbol:       "kt",
				Aliases:      []string{"kt", "knots"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.514444 },
				FromBaseFunc: func(val float64) float64 { return val / 0.514444 },
			},
			"Feet per Second": {
				Name:         "Feet per Second",
				Symbol:       "ft/s",
				Aliases:      []string{"fps", "feetpersecond"},
				ToBaseFunc:   func(val float64) float64 { return val * 0.3048 },
				FromBaseFunc: func(val float64) float64 { return val / 0.3048 },
			},
		},
	}
}

func NewTimeSystem() UnitSystem {
	return UnitSystem{
		Name:     "Time",
		BaseUnit: "Seconds",
		Units: map[string]Unit{
			"Seconds": {
				Name:         "Seconds",
				Symbol:       "s",
				Aliases:      []string{"s", "sec", "second", "seconds"},
				ToBaseFunc:   func(val float64) float64 { return val },
				FromBaseFunc: func(val float64) float64 { return val },
			},
			"Minutes": {
				Name:         "Minutes",
				Symbol:       "min",
				Aliases:      []string{"min", "minute", "minutes"},
				ToBaseFunc:   func(val float64) float64 { return val * 60 },
				FromBaseFunc: func(val float64) float64 { return val / 60 },
			},
			"Hours": {
				Name:         "Hours",
				Symbol:       "hr",
				Aliases:      []string{"h", "hr", "hour", "hours"},
				ToBaseFunc:   func(val float64) float64 { return val * 3600 },
				FromBaseFunc: func(val float64) float64 { return val / 3600 },
			},
			"Days": {
				Name:         "Days",
				Symbol:       "d",
				Aliases:      []string{"d", "day", "days"},
				ToBaseFunc:   func(val float64) float64 { return val * 86400 },
				FromBaseFunc: func(val float64) float64 { return val / 86400 },
			},
			"Years": {
				Name:         "Years",
				Symbol:       "yr",
				Aliases:      []string{"y", "yr", "year", "years"},
				ToBaseFunc:   func(val float64) float64 { return val * 31536000 },
				FromBaseFunc: func(val float64) float64 { return val / 31536000 },
			},
		},
	}
}

func MustRegisterSystems() map[string]Unit {
	systems := []UnitSystem{
		NewVolumeSystem(),
		NewLengthSystem(),
		NewWeightSystem(),
		NewTemperatureSystem(),
		NewAreaSystem(),
		NewSpeedSystem(),
		NewTimeSystem(),
	}

	unitMap := make(map[string]Unit)
	for _, system := range systems {
		for _, unit := range system.Units {
			for _, alias := range unit.Aliases {
				unitMap[strings.ToLower(alias)] = unit
			}
			unitMap[strings.ToLower(unit.Name)] = unit
			unitMap[strings.ToLower(unit.Symbol)] = unit
		}
	}
	return unitMap
}
