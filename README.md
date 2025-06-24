# NLP Unit Converter

A natural language processing unit converter written in Go that intelligently parses and converts measurements from human-readable text input across multiple unit systems.

## Features

- **🗣️ Natural Language Input**: Parse complex expressions like "two pints and a half cup in floz", "1 km in miles"
- **📏 Multiple Unit Systems**: Supports Volume, Length, and Weight with metric, imperial, and specialized units
- **🔢 Smart Number Parsing**: Handles text numbers ("one", "two", "half"), fractions ("1/2"), and scientific notation ("1.5e3")
- **⚡ Flexible Syntax**: Supports various operators like `+`, `&`, `and`, and even `-` for subtraction
- **🎯 Target Unit Specification**: Convert to specific units using the `in [unit]` syntax
- **🤖 Intelligent Error Handling**: Provides helpful suggestions for typos and unknown units
- **⚠️ Typo Correction**: Suggests similar units when you make spelling mistakes

## Supported Units

### 📊 Volume Units
#### Metric
- **Milliliters**: ml, milliliter, milliliters, millilitre, millilitres, milli
- **Liters**: L, l, liter, liters, litre, litres
- **Cubic meters**: m³, m3, m^3, cubicmeter, cubicmeters
- **Cubic centimeters**: cm³, cm3, cm^3, cubiccentimeter, cubiccentimeters, cc

#### Imperial/US
- **Fluid Ounce**: fl oz, floz, fluidounce, fluidounces, oz
- **Cup**: c, cup, cups
- **Pint**: pt, pint, pints
- **Quart**: qt, quart, quarts
- **Gallon**: gal, gallon, gallons
- **Cubic feet**: ft³, ft3, ft^3, cubicfoot, cubicfeet

#### Specialized
- **Barrels**: bbl, barrel, barrels (US oil barrel)

### 📏 Length Units
#### Metric
- **Meters**: m, meter, meters, metre, metres
- **Kilometers**: km, kilometer, kilometers, kilometre, kilometres
- **Centimeters**: cm, centimeter, centimeters, centimetre, centimetres
- **Millimeters**: mm, millimeter, millimeters, millimetre, millimetres

#### Imperial/US
- **Inches**: in, inch, inches
- **Feet**: ft, foot, feet
- **Yards**: yd, yard, yards
- **Miles**: mi, mile, miles

### ⚖️ Weight Units
#### Metric
- **Grams**: g, gram, grams
- **Kilograms**: kg, kilogram, kilograms
- **Milligrams**: mg, milligram, milligrams

#### Imperial/US
- **Pounds**: lb, lbs, pound, pounds
- **Ounces**: oz, ounce, ounces

## Usage

### Basic Usage

```bash
go run main.go
```

### Example Inputs and Outputs

#### Volume Conversions
```
Input:  "1L & 23 ml"
Result: 1023 mL (Milliliters)

Input:  "1/2 gallon + 1/4 pint in cups"
Result: 8.5 c (Cup)

Input:  "two pints and a half cup in floz"
Result: 36 fl oz (Fluid Ounce)

Input:  "1.5e3 ml in L"
Result: 1.5 L (Liters)
```

#### Length Conversions
```
Input:  "1 km in miles"
Result: 0.621373 mi (Miles)

Input:  "a foot and 5 inches in cm"
Result: 43.18 cm (Centimeters)

Input:  "100 meters + 0.1km in ft"
Result: 656.168 ft (Feet)
```

#### Weight Conversions
```
Input:  "1kg in lbs"
Result: 2.20462 lb (Pounds)

Input:  "two pounds and 8 ounces in grams"
Result: 1133.98 g (Grams)

Input:  "100g + .5kg"
Result: 0.6 kg (Kilograms)
```

#### Smart Error Handling
```
Input:  "1 leter in ml"
Error:  unknown unit: 'leter'. Did you mean 'liter'?

Input:  "2 gallens in L"
Error:  unknown unit: 'gallens'. Did you mean 'gallons'?
```

### Programmatic Usage

```go
package main

import (
    "fmt"
    "nlpconverter/converter"
)

func main() {
    // Create different unit systems
    volumeSystem := converter.NewVolumeSystem()
    lengthSystem := converter.NewLengthSystem()
    weightSystem := converter.NewWeightSystem()
    
    // Create converters for each system
    volumeConverter := converter.NewConverter(volumeSystem)
    lengthConverter := converter.NewConverter(lengthSystem)
    weightConverter := converter.NewConverter(weightSystem)
    
    // Process natural language input
    result, err := volumeConverter.Process("1.5 gallons in liters")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Volume: %g %s (%s)\n", 
        result.Value, result.UnitSymbol, result.UnitName)
    
    result, err = lengthConverter.Process("100 meters in feet")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Length: %g %s (%s)\n", 
        result.Value, result.UnitSymbol, result.UnitName)
    
    result, err = weightConverter.Process("2 pounds in kg")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Printf("Weight: %g %s (%s)\n", 
        result.Value, result.UnitSymbol, result.UnitName)
}
```

## Input Format

The converter supports incredibly flexible input formats:

### Basic Formats
1. **Simple conversion**: `"1.5gal in ml"`, `"100 meters in feet"`
2. **Text numbers**: `"one gallon"`, `"two pounds"`, `"half a cup"`
3. **Fractions**: `"1/2 gallon"`, `"3/4 pint"`, `"1 1/4 cups"`
4. **Scientific notation**: `"1.5e3 ml"`, `"2.5e-2 km"`
5. **Decimal variations**: `".5 gal"`, `"0.25 kg"`

### Complex Expressions
1. **Addition with operators**: `"1L + 23 ml"`, `"1L & 23 ml"`, `"2 gallons and 1l"`
2. **Mixed text and numbers**: `"one gallon and 2.5 litres"`
3. **Implicit quantities**: `"Liter + 100.87 ml"` (assumes 1 Liter)
4. **Target unit specification**: `"1Liter + 100.87 milli in cm^3"`
5. **Cross-system mixing**: Not supported (each converter handles one unit system)

### Advanced Features
- **Typo tolerance**: `"1 leter"` → suggests `"liter"`
- **Multiple aliases**: `"litre"`, `"liter"`, `"L"`, `"l"` all work
- **Case insensitive**: `"ML"`, `"ml"`, `"mL"` all work
- **Flexible spacing**: `"1L"`, `"1 L"`, `"1  L"` all work

## Architecture

The project follows a clean, modular architecture designed for extensibility:

```
nlpconverter/
├── main.go              # Demo application with comprehensive test cases
├── converter/
│   ├── converter.go     # Core conversion logic, NLP parsing, and error handling
│   └── system.go        # Unit system definitions (Volume, Length, Weight)
└── go.mod               # Go module file
```

### Key Components

- **`UnitSystem`**: Defines the base unit and all supported units with conversion factors
- **`Converter`**: Handles natural language parsing, unit conversion, and intelligent error suggestions
- **`Result`**: Contains the converted value with unit symbol and full name
- **Smart Preprocessing**: Handles text numbers, fractions, and scientific notation
- **Regex Engine**: Pre-compiled patterns for efficient parsing
- **Error Suggestions**: Levenshtein distance algorithm for typo correction

## Installation

1. Clone the repository:
```bash
git clone https://github.com/s4tyendra/NLP-Unit-Converter.git
cd NLP-Unit-Converter
```

2. Run the demo:
```bash
go run main.go
```

3. Use as a module in your Go project:
```bash
go mod init nlpconverter
go get github.com/s4tyendra/NLP-Unit-Converter
```

## Requirements

- Go 1.21 or higher

## Contributing

Contributions are welcome! Here are some ways you can contribute:

1. **Add new unit systems** (temperature, area, speed, energy, etc.)
2. **Improve NLP parsing** for more complex expressions and operators
3. **Add more unit aliases** for international and regional variations
4. **Enhance error handling** with better suggestions and context
5. **Add comprehensive tests** to ensure reliability across edge cases
6. **Implement subtraction and multiplication** operators
7. **Add support for compound units** (e.g., miles per hour, meters per second)

## Roadmap

- [x] ✅ Support for Volume, Length, and Weight unit systems
- [x] ✅ Text number parsing ("one", "two", "half")
- [x] ✅ Fraction parsing ("1/2", "3/4")
- [x] ✅ Scientific notation support
- [x] ✅ Intelligent error suggestions
- [ ] 🔄 Temperature unit system (Celsius, Fahrenheit, Kelvin)
- [ ] 🔄 Area unit system (square meters, acres, etc.)
- [ ] 🔄 Speed unit system (mph, km/h, m/s)
- [ ] 🔄 Support for subtraction and multiplication operators
- [ ] 🔄 Compound units (speed, density, etc.)
- [ ] 🔄 Web API interface
- [ ] 🔄 Command-line interface improvements
- [ ] 🔄 Comprehensive test suite with edge cases

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go's powerful regex and string processing capabilities
- Inspired by the need for intuitive unit conversion in natural language applications
