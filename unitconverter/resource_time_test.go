package unitconverter

import (
	"testing"
)

func TestTimeTypes(t *testing.T) {
	tmpV := 1.0
	value = &tmpV
	types := GetTimeTypes()(false)
	testFindObjectByName(types, t)
}

func TestTime(t *testing.T) {
	// Test types
	types := GetTimeTypes()

	// Cases to test
	cases := []TestCases{
		{
			wanted:   "Second",
			original: "H",
			input:    1,
			expected: 3600,
			sublist:  false,
		},
	}

	testConvertObject(cases, types, t)
}
