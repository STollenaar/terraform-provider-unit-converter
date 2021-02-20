package unitconverter

import (
	"testing"
)

func TestByteTypes(t *testing.T) {
	tmpV := 1.0
	value = &tmpV
	types := GetByteTypes()(false)
	testFindObjectByName(types, t)
}

func TestByte(t *testing.T) {
	// Test types
	types := GetByteTypes()

	// Cases to test
	cases := []TestCases{
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
