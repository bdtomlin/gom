package forms

import (
	"math"
	"strconv"
	"strings"
)

func FloatToInt(f float64) int {
	return int(math.Round(f))
}

func FmtInt(i int) string {
	return NumFmt(strconv.Itoa(i))
}

func FmtFloat(f float64) string {
	return NumFmt(strconv.FormatFloat(f, 'f', 'g', 64))
}

func ParseInt(numStr string) (string, int, error) {
	s, f, err := ParseFloat(numStr)
	if err != nil {
		return numStr, int(f), err
	}

	i := FloatToInt(f)
	s = NumFmt(strconv.Itoa(i))

	return s, i, nil
}

func ParseFloat(numStr string) (string, float64, error) {
	norm := strings.ReplaceAll(numStr, ",", "")

	f, err := strconv.ParseFloat(norm, 64)
	if err != nil {
		return numStr, f, err
	}

	return NumFmt(norm), f, nil
}

// AddCommas adds comma separators to the integer part of a numeric string.
// It handles both integers and decimals, and supports negative numbers.
// Examples:
// - "1234567" -> "1,234,567"
// - "1234567.89" -> "1,234,567.89"
// - "-1234567.89" -> "-1,234,567.89"
// - "0.123" -> "0.123" (adds leading 0 if no integer part)
func NumFmt(numStr string) string {
	if numStr == "" {
		return ""
	}

	// Handle sign
	sign := ""
	if numStr[0] == '-' {
		sign = "-"
		numStr = numStr[1:]
	}

	// Handle leading decimal point (e.g., ".123" -> "0.123")
	if strings.HasPrefix(numStr, ".") {
		numStr = "0" + numStr
	}

	// Split into integer and fractional parts
	dotPos := strings.Index(numStr, ".")
	var integerPart, fractionalPart string
	if dotPos == -1 {
		integerPart = numStr
		fractionalPart = ""
	} else {
		integerPart = numStr[:dotPos]
		fractionalPart = numStr[dotPos:]
	}

	// Add commas to integer part
	integerWithCommas := addCommasToInteger(integerPart)

	return sign + integerWithCommas + fractionalPart
}

// addCommasToInteger adds commas every three digits from the right to a string representing an integer.
func addCommasToInteger(part string) string {
	if len(part) <= 3 {
		return part
	}

	var parts []string
	i := len(part)
	for i > 3 {
		parts = append([]string{part[i-3 : i]}, parts...)
		parts = append([]string{","}, parts...)
		i -= 3
	}
	parts = append([]string{part[:i]}, parts...)
	return strings.Join(parts, "")
}
