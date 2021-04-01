package unitconverter

import (
	"testing"
)

type TestCasesMathDownstep struct {
	startUnit, expectedUnit   string
	startValue, expectedValue float64
}

func TestMathdownstep(t *testing.T) {
	// Cases to test
	cases := []TestCasesMathDownstep{
		{
			startUnit:     "Gigabit",
			startValue:    1.2,
			expectedUnit:  "Megabit",
			expectedValue: 1200,
		},
	}
	for _, c := range cases {
		resultValue, resultUnit, errorUnit := doDownstep(c.startValue, c.startUnit)

		if errorUnit != nil {
			t.Fatalf("Downstep operation returned a non nil error %s", errorUnit)
		} else if resultValue != c.expectedValue {
			t.Fatalf("Error matching resultValue and expectedValue: %#v vs %#v", resultValue, c.expectedValue)
		} else if resultUnit != c.expectedUnit {
			t.Fatalf("Error matching resultUnit and expectedUnit: %s vs %s", resultUnit, c.expectedUnit)
		}
	}
}
