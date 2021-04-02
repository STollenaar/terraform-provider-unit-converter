package unitconverter

import (
	"testing"
)

func TestByteTypes(t *testing.T) {
	tmpV := 1.0
	types := GetByteTypes()(tmpV, false)
	testFindObjectByName(types, t)
}

func TestByte(t *testing.T) {
	// Test types
	types := GetByteTypes()

	// Cases to test
	// TODO add more test cases

	cases := []TestCasesTypes{
		{
			wanted:   "Bit",
			original: "Byte",
			input:    1,
			expected: 8,
			sublist:  false,
		},
	}

	testConvertObject(cases, types, t)
}
