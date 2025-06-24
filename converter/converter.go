package converter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Result struct {
	Value      float64
	UnitSymbol string
	UnitName   string
}

type compiledRegexes struct {
	targetUnit *regexp.Regexp
	component  *regexp.Regexp
	fraction   *regexp.Regexp
}

type Converter struct {
	system  UnitSystem
	unitMap map[string]Unit
	regexes compiledRegexes
}

type parsedComponent struct {
	Value float64
	Unit  Unit
}

var textNumberMap = map[string]string{
	"a": "1", "an": "1", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5",
	"six": "6", "seven": "7", "eight": "8", "nine": "9", "ten": "10", "half": "0.5",
}

func NewConverter(system UnitSystem) *Converter {
	aliasMap := make(map[string]Unit)
	for _, unit := range system.Units {
		for _, alias := range unit.Aliases {
			aliasMap[strings.ToLower(alias)] = unit
		}
		aliasMap[strings.ToLower(unit.Name)] = unit
		aliasMap[strings.ToLower(unit.Symbol)] = unit
	}

	numberRegexPart := `((?:\d+(?:\.\d*)?|\.\d+)(?:[eE][+-]?\d+)?)`
	unitRegexPart := `([a-z][a-z0-9^³²]*)`

	regexes := compiledRegexes{
		targetUnit: regexp.MustCompile(`\s+in\s+([a-z0-9\s^³²]+)$`),
		component: regexp.MustCompile(fmt.Sprintf(`\s*([+-])?\s*%s?\s*%s`, numberRegexPart, unitRegexPart)),
		fraction: regexp.MustCompile(`(\d+)\s*/\s*(\d+)`),
	}

	return &Converter{
		system:  system,
		unitMap: aliasMap,
		regexes: regexes,
	}
}

func (c *Converter) preprocessInput(input string) string {
	clean := strings.ToLower(input)

	for word, numStr := range textNumberMap {
		re := regexp.MustCompile(`\b` + word + `\b`)
		clean = re.ReplaceAllString(clean, numStr)
	}

	clean = c.regexes.fraction.ReplaceAllStringFunc(clean, func(m string) string {
		parts := c.regexes.fraction.FindStringSubmatch(m)
		if len(parts) < 3 {
			return m 
		}
		num, _ := strconv.ParseFloat(parts[1], 64)
		den, _ := strconv.ParseFloat(parts[2], 64)
		if den == 0 {
			return "0"
		}
		return fmt.Sprintf("%f", num/den)
	})

	return clean
}

func (c *Converter) Process(input string) (*Result, error) {
	cleanInput := c.preprocessInput(input)

	targetMatch := c.regexes.targetUnit.FindStringSubmatch(cleanInput)
	var targetUnit *Unit
	if targetMatch != nil {
		cleanInput = strings.TrimSpace(c.regexes.targetUnit.ReplaceAllString(cleanInput, ""))
		targetUnitStr := strings.TrimSpace(targetMatch[1])
		unit, ok := c.findUnit(targetUnitStr)
		if !ok {
			return nil, c.createNotFoundError(targetUnitStr)
		}
		targetUnit = &unit
	}

	matches := c.regexes.component.FindAllStringSubmatch(cleanInput, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("no valid units found in input: '%s'", input)
	}

	var components []parsedComponent
	var lastParsedUnit Unit
	for _, match := range matches {
		signStr := match[1]
		valueStr := match[2]
		unitStr := match[3]

		if unitStr == "and" || unitStr == "&" {
			continue
		}

		unit, ok := c.findUnit(unitStr)
		if !ok {
			return nil, c.createNotFoundError(unitStr)
		}

		value := 1.0
		if valueStr != "" {
			var err error
			value, err = strconv.ParseFloat(valueStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid number: '%s'", valueStr)
			}
		}

		if signStr == "-" {
			value = -value
		}

		components = append(components, parsedComponent{Value: value, Unit: unit})
		lastParsedUnit = unit
	}

	if targetUnit == nil {
		if len(components) == 0 {
			return nil, fmt.Errorf("no processable components found in input: '%s'", input)
		}
		targetUnit = &lastParsedUnit
	}

	totalInBase := 0.0
	for _, comp := range components {
		totalInBase += comp.Value * comp.Unit.ToBase
	}

	finalValue := totalInBase / targetUnit.ToBase

	return &Result{
		Value:      finalValue,
		UnitSymbol: targetUnit.Symbol,
		UnitName:   targetUnit.Name,
	}, nil
}

func (c *Converter) findUnit(s string) (Unit, bool) {
	unit, ok := c.unitMap[strings.ToLower(s)]
	return unit, ok
}

func (c *Converter) createNotFoundError(unknownUnit string) error {
	const maxSuggestionDistance = 2
	bestMatch := ""
	minDist := maxSuggestionDistance + 1

	for alias := range c.unitMap {
		if len(alias) < 3 {
			continue
		}
		dist := levenshtein(unknownUnit, alias)
		if dist < minDist {
			minDist = dist
			bestMatch = alias
		}
	}

	if minDist <= maxSuggestionDistance {
		return fmt.Errorf("unknown unit: '%s'. Did you mean '%s'?", unknownUnit, bestMatch)
	}
	return fmt.Errorf("unknown unit: '%s'", unknownUnit)
}

// levenshtein calculates the Levenshtein distance between two strings.
func levenshtein(a, b string) int {
	f := make([]int, utf8.RuneCountInString(b)+1)
	for j := range f {
		f[j] = j
	}
	for _, ca := range a {
		j := 1
		fj1 := f[0]
		f[0]++
		for _, cb := range b {
			mn := f[j] + 1
			if cb == ca {
				mn = fj1
			}
			if f[j-1]+1 < mn {
				mn = f[j-1] + 1
			}
			fj1 = f[j]
			f[j] = mn
			j++
		}
	}
	return f[len(f)-1]
}