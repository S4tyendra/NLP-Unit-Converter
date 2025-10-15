# NLP Unit Converter

A natural language processing unit converter written in Go that intelligently parses and converts measurements from human-readable text input across multiple unit systems.

## Features

- **🗣️ Natural Language Input**: Parse complex expressions like "two pints and a half cup in floz", "1 km in miles"
- **📏 Multiple Unit Systems**: Supports Volume, Length, Weight, Temperature, Area, Speed, and Time with metric, imperial, and specialized units
- **🔢 Smart Number Parsing**: Handles text numbers ("one", "two", "half"), fractions ("1/2"), and scientific notation ("1.5e3")
- **⚡ Flexible Syntax**: Supports various operators like `+`, `&`, `and`, and even `-` for subtraction
- **🎯 Target Unit Specification**: Convert to specific units using `in [unit]` or `to [unit]` syntax
- **🌐 Web API Server**: Built-in HTTP server with interactive web interface
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

### 🌡️ Temperature Units
- **Celsius**: C, c, celsius
- **Fahrenheit**: F, f, fahrenheit
- **Kelvin**: K, k, kelvin

### 📐 Area Units
#### Metric
- **Square Meters**: m², m2, sqm, squaremeter, squaremeters
- **Square Kilometers**: km², km2, sqkm, squarekilometer, squarekilometers
- **Hectares**: ha, hectare, hectares

#### Imperial/US
- **Square Miles**: mi², mi2, sqmi, squaremile, squaremiles
- **Acres**: ac, acre, acres
- **Square Yards**: yd², yd2, sqyd, squareyard, squareyards
- **Square Feet**: ft², ft2, sqft, squarefoot, squarefeet
- **Square Inches**: in², in2, sqin, squareinch, squareinches

### 🏃 Speed Units
- **Meters per Second**: m/s, mps, meterspersecond
- **Kilometers per Hour**: km/h, kph, kmh, kilometersperhour
- **Miles per Hour**: mph, milesperhour
- **Knots**: kt, knots
- **Feet per Second**: ft/s, fps, feetpersecond

### ⏰ Time Units
- **Seconds**: s, sec, second, seconds
- **Minutes**: min, minute, minutes
- **Hours**: h, hr, hour, hours
- **Days**: d, day, days
- **Years**: y, yr, year, years

## Usage

### Command Line

#### Convert a unit
```bash
./convertunit "2l to ml"
./convertunit "convert 100 f to c"
./convertunit "5 km in miles"
```

#### Start Web Server
```bash
# Start on default port 8080
./convertunit --start-server
./convertunit -ss

# Start on custom port
./convertunit --start-server 7000
./convertunit -ss 3000
```

#### Get Help
```bash
./convertunit --help
./convertunit -h
```

### Web API

The built-in web server provides both a user-friendly HTML interface and a JSON API:

#### Access the Web Interface
1. Start the server: `./convertunit -ss`
2. Open your browser: `http://localhost:8080`
3. Enter conversions in the interactive form

#### Use the JSON API
```bash
# Volume conversion
curl "http://localhost:8080/?q=2l+to+ml"
# Response: {"value":2000,"unit_symbol":"mL","unit_name":"Milliliters"}

# Temperature conversion
curl "http://localhost:8080/?q=100+f+to+c"
# Response: {"value":37.77777777777778,"unit_symbol":"°C","unit_name":"Celsius"}

# Length conversion
curl "http://localhost:8080/?q=5+km+to+miles"
# Response: {"value":3.106863683249034,"unit_symbol":"mi","unit_name":"Miles"}

# Error handling
curl "http://localhost:8080/?q=invalid+unit"
# Response: {"error":"unknown unit: 'invalid'"}
```

**API Features:**
- Single endpoint: `/?q=your+query`
- GET requests only
- Maximum query length: 100 characters
- Returns JSON with conversion result or error
- Serves HTML page when no query parameter provided

### Example Inputs and Outputs

#### Volume Conversions
```bash
./convertunit "1L & 23 ml"
# 1023 mL (Milliliters)

./convertunit "1/2 gallon + 1/4 pint to cups"
# 8.5 c (Cup)

./convertunit "two pints and a half cup in floz"
# 36 fl oz (Fluid Ounce)

./convertunit "1.5e3 ml to L"
# 1.5 L (Liters)
```

#### Length Conversions
```bash
./convertunit "1 km to miles"
# 0.621373 mi (Miles)

./convertunit "a foot and 5 inches in cm"
# 43.18 cm (Centimeters)

./convertunit "100 meters + 0.1km to ft"
# 656.168 ft (Feet)
```

#### Weight Conversions
```bash
./convertunit "1kg to lbs"
# 2.20462 lb (Pounds)

./convertunit "two pounds and 8 ounces to grams"
# 1133.98 g (Grams)

./convertunit "100g + .5kg"
# 0.6 kg (Kilograms)
```

#### Temperature Conversions
```bash
./convertunit "100 c to f"
# 212 °F (Fahrenheit)

./convertunit "convert 32 f to c"
# 0 °C (Celsius)

./convertunit "0c to k"
# 273.15 K (Kelvin)
```

#### Area & Speed Conversions
```bash
./convertunit "100 sqft to m2"
# 9.2903 m² (Square Meters)

./convertunit "60 mph to kph"
# 96.56064 km/h (Kilometers per Hour)
```

#### Smart Error Handling
```bash
./convertunit "1 leter in ml"
# Error: unknown unit: 'leter'. Did you mean 'liter'?

./convertunit "2 gallens in L"
# Error: unknown unit: 'gallens'. Did you mean 'gallons'?
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
1. **Simple conversion**: `"1.5gal in ml"`, `"1.5gal to ml"`, `"100 meters to feet"`
2. **With 'convert' prefix**: `"convert 100 f to c"`, `"convert 2l to ml"`
3. **Text numbers**: `"one gallon"`, `"two pounds"`, `"half a cup"`
4. **Fractions**: `"1/2 gallon"`, `"3/4 pint"`, `"1 1/4 cups"`
5. **Scientific notation**: `"1.5e3 ml"`, `"2.5e-2 km"`
6. **Decimal variations**: `".5 gal"`, `"0.25 kg"`

### Complex Expressions
1. **Addition with operators**: `"1L + 23 ml"`, `"1L & 23 ml"`, `"2 gallons and 1l"`
2. **Mixed text and numbers**: `"one gallon and 2.5 litres"`
3. **Implicit quantities**: `"Liter + 100.87 ml"` (assumes 1 Liter)
4. **Target unit specification**: `"1Liter + 100.87 milli in cm^3"`, `"2l + 500ml to cups"`
5. **Both 'in' and 'to' keywords**: `"5 km in miles"`, `"5 km to miles"` (both work)

### Advanced Features
- **Typo tolerance**: `"1 leter"` → suggests `"liter"`
- **Multiple aliases**: `"litre"`, `"liter"`, `"L"`, `"l"` all work
- **Case insensitive**: `"ML"`, `"ml"`, `"mL"` all work
- **Flexible spacing**: `"1L"`, `"1 L"`, `"1  L"` all work
- **Optional 'convert' prefix**: `"convert 32 f to c"` works same as `"32 f to c"`
- **Dual syntax support**: Both `in` and `to` keywords supported for target units

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

2. Build the project:
```bash
go build -o convertunit
```

3. Run conversions:
```bash
./convertunit "2l to ml"
```

4. Start the web server:
```bash
./convertunit -ss        # Default port 8080
./convertunit -ss 3000   # Custom port
```

5. Use as a module in your Go project:
```bash
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
- [x] ✅ Temperature unit system (Celsius, Fahrenheit, Kelvin)
- [x] ✅ Area unit system (square meters, acres, etc.)
- [x] ✅ Speed unit system (mph, km/h, m/s)
- [x] ✅ Time unit system (seconds, minutes, hours, days, years)
- [x] ✅ Support for 'to' keyword in addition to 'in'
- [x] ✅ Support for 'convert' prefix
- [x] ✅ Web API interface with interactive HTML frontend
- [x] ✅ Configurable server port
- [ ] 🔄 Support for more mathematical operators
- [ ] 🔄 Compound units (speed, density, etc.)
- [ ] 🔄 Comprehensive test suite with edge cases
- [ ] 🔄 Docker support
- [ ] 🔄 REST API documentation with OpenAPI/Swagger

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with Go's powerful regex and string processing capabilities
- Inspired by the need for intuitive unit conversion in natural language applications
