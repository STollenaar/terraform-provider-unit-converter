package unitconverter

import (
	"testing"
)

func TestLengthTypes(t *testing.T) {
	tmpV := 1.0
	types := GetLengthTypes()(tmpV, false)
	testFindObjectByName(types, t)
}

func TestLength(t *testing.T) {
	// Test types
	types := GetLengthTypes()

	// Cases to test
	cases := []TestCasesTypes{
		{
			wanted:   "Meter",
			original: "km",
			input:    1,
			expected: 1000,
			sublist:  false,
		},
	}

	testConvertObject(cases, types, t)
}
