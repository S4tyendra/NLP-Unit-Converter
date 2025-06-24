package converter

type Unit struct {
	Name    string
	Symbol  string
	Aliases []string
	ToBase  float64
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
				Name:    "Milliliters",
				Symbol:  "mL",
				Aliases: []string{"ml", "milliliter", "milliliters", "millilitre", "millilitres", "milli"},
				ToBase:  1.0, // Base unit
			},
			"Liters": {
				Name:    "Liters",
				Symbol:  "L",
				Aliases: []string{"l", "liter", "liters", "litre", "litres"},
				ToBase:  1000.0, // 1 Liter = 1000 mL
			},
			"Cubic meters": {
				Name:    "Cubic meters",
				Symbol:  "m³",
				Aliases: []string{"m3", "m^3", "cubicmeter", "cubicmeters"},
				ToBase:  1000000.0, // 1 m³ = 1,000,000 mL
			},
			"Cubic centimeters": {
				Name:    "Cubic centimeters",
				Symbol:  "cm³",
				Aliases: []string{"cm3", "cm^3", "cubiccentimeter", "cubiccentimeters", "cc"},
				ToBase:  1.0, // 1 cm³ = 1 mL
			},
			"Fluid Ounce": {
				Name:    "Fluid Ounce",
				Symbol:  "fl oz",
				Aliases: []string{"floz", "fluidounce", "fluidounces", "oz"},
				ToBase:  flOzToMl,
			},
			"Cup": {
				Name:    "Cup",
				Symbol:  "c",
				Aliases: []string{"cup", "cups"},
				ToBase:  flOzToMl * 8, // 8 fl oz in a cup
			},
			"Pint": {
				Name:    "Pint",
				Symbol:  "pt",
				Aliases: []string{"pint", "pints"},
				ToBase:  flOzToMl * 16, // 16 fl oz in a pint
			},
			"Quart": {
				Name:    "Quart",
				Symbol:  "qt",
				Aliases: []string{"quart", "quarts"},
				ToBase:  flOzToMl * 32, // 32 fl oz in a quart
			},
			"Gallon": {
				Name:    "Gallon",
				Symbol:  "gal",
				Aliases: []string{"gallon", "gallons"},
				ToBase:  flOzToMl * 128, // 128 fl oz in a gallon
			},
			"Cubic feet": {
				Name:    "Cubic feet",
				Symbol:  "ft³",
				Aliases: []string{"ft3", "ft^3", "cubicfoot", "cubicfeet"},
				ToBase:  28316.8,
			},
			"Barrels": {
				Name:    "Barrels",
				Symbol:  "bbl",
				Aliases: []string{"barrel", "barrels"},
				ToBase:  158987.3, // Based on US oil barrel, see https://www.unitconverters.net/volume/barrel-oil-to-liter.htm
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
				Name:    "Meters",
				Symbol:  "m",
				Aliases: []string{"m", "meter", "meters", "metre", "metres"},
				ToBase:  1.0,
			},
			"Kilometers": {
				Name:    "Kilometers",
				Symbol:  "km",
				Aliases: []string{"km", "kilometer", "kilometers", "kilometre", "kilometres"},
				ToBase:  1000.0,
			},
			"Centimeters": {
				Name:    "Centimeters",
				Symbol:  "cm",
				Aliases: []string{"cm", "centimeter", "centimeters", "centimetre", "centimetres"},
				ToBase:  0.01,
			},
			"Millimeters": {
				Name:    "Millimeters",
				Symbol:  "mm",
				Aliases: []string{"mm", "millimeter", "millimeters", "millimetre", "millimetres"},
				ToBase:  0.001,
			},
			"Inches": {
				Name:    "Inches",
				Symbol:  "in",
				Aliases: []string{"in", "inch", "inches"},
				ToBase:  0.0254,
			},
			"Feet": {
				Name:    "Feet",
				Symbol:  "ft",
				Aliases: []string{"ft", "foot", "feet"},
				ToBase:  0.3048,
			},
			"Yards": {
				Name:    "Yards",
				Symbol:  "yd",
				Aliases: []string{"yd", "yard", "yards"},
				ToBase:  0.9144,
			},
			"Miles": {
				Name:    "Miles",
				Symbol:  "mi",
				Aliases: []string{"mi", "mile", "miles"},
				ToBase:  1609.34,
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
				Name:    "Grams",
				Symbol:  "g",
				Aliases: []string{"g", "gram", "grams"},
				ToBase:  1.0,
			},
			"Kilograms": {
				Name:    "Kilograms",
				Symbol:  "kg",
				Aliases: []string{"kg", "kilogram", "kilograms"},
				ToBase:  1000.0,
			},
			"Milligrams": {
				Name:    "Milligrams",
				Symbol:  "mg",
				Aliases: []string{"mg", "milligram", "milligrams"},
				ToBase:  0.001,
			},
			"Pounds": {
				Name:    "Pounds",
				Symbol:  "lb",
				Aliases: []string{"lb", "lbs", "pound", "pounds"},
				ToBase:  453.592,
			},
			"Ounces": {
				Name:    "Ounces",
				Symbol:  "oz",
				Aliases: []string{"ounce", "ounces"}, // "oz" can be ambiguous with "fl oz", lets see this later
				ToBase:  28.3495,
			},
		},
	}
}
