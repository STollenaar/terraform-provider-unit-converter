package unitconverter

import "testing"

type TestCases struct {
	wanted, original string
	expected, input  float64
	sublist          bool
}

func testFindObjectByName(types []Object, t *testing.T) {
	for _, c := range types {
		wantName := c.Name
		wantNameShort := c.NameShort

		gotName := FindObjectByName(wantName, types).Name
		gotNameShort := FindObjectByName(wantNameShort, types).NameShort

		if gotName != wantName {
			t.Fatalf("Error matching gotName and wantName: %#v vs %#v", gotName, wantName)
		}
		if wantNameShort != gotNameShort {
			t.Fatalf("Error matching gotNameShort and wantNameShort: %#v vs %#v", gotNameShort, wantNameShort)

		}
	}

}

func testConvertObject(cases []TestCases, types func(bool) []Object, t *testing.T) {

	// Doing the tests
	for _, c := range cases {
		value = &c.input
		original := FindObjectByName(c.original, types(c.sublist))
		wanted := FindObjectByName(c.wanted, types(c.sublist))
		out := ConvertFuncMath(c.input, original.Unit, wanted.Unit)

		if out != c.expected {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}

}
