package unitconverter

import (
	"testing"
)

func TestTimeTypes(t *testing.T) {
	tmpV := 1.0
	types := GetTimeTypes()(tmpV, false)
	testFindObjectByName(types, t)
}

func TestTime(t *testing.T) {
	// Test types
	types := GetTimeTypes()

	// Cases to test
	// TODO add more test cases

	cases := []TestCasesTypes{
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
