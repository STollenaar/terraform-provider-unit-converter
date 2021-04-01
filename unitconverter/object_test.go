package unitconverter

import "testing"

type TestCasesTypes struct {
	wanted, original string
	expected, input  float64
	sublist          bool
}

func testFindObjectByName(types []Object, t *testing.T) {
	for _, c := range types {
		wantName := c.Name
		wantNameShort := c.NameShort

		gotName, err := FindObjectValueByName(wantName, types)
		gotNameShort, err := FindObjectValueByName(wantNameShort, types)

		if err != nil {
			t.Fatalf(err.Error())
		}

		if wantName != gotName.Name {
			t.Fatalf("Error matching gotName and wantName: %#v vs %#v", gotName, wantName)
		}
		if wantNameShort != gotNameShort.NameShort {
			t.Fatalf("Error matching gotNameShort and wantNameShort: %#v vs %#v", gotNameShort, wantNameShort)
		}
	}
}

func testConvertObject(cases []TestCasesTypes, types func(float64, bool) []Object, t *testing.T) {

	// Doing the tests
	for _, c := range cases {
		original, err := FindObjectValueByName(c.original, types(c.input, c.sublist))
		wanted, err := FindObjectValueByName(c.wanted, types(c.input, c.sublist))
		out := ConvertFuncMath(c.input, original.Unit, wanted.Unit)

		if out != c.expected || err != nil {
			t.Fatalf("Error matching output and expected: %#v vs %#v", out, c.expected)
		}
	}
}
