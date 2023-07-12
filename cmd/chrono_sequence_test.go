package cmd

import (
	"testing"
	"time"
)

type durationTranslateTest struct {
	input       string
	expected    string
	expectedErr bool
}

func TestTranslateDuration(t *testing.T) {
	tests := []durationTranslateTest{
		durationTranslateTest{"2d", "48h", false},
		durationTranslateTest{"3h", "3h", false},
		durationTranslateTest{"3xd", "", true},
	}

	for _, test := range tests {
		actual, err := translateDuration(test.input)

		if err == nil && test.expectedErr {
			t.Logf("Expected an error but didn't get one for input: %s", test.input)
			t.Fail()
		}

		if err != nil && !test.expectedErr {
			t.Logf("Did not expect an error for input %s but got: %v", test.input, err)
			t.Fail()
		}

		if test.expected != actual {
			t.Logf("Expected %s from input %s but got %s", test.expected, test.input, actual)
			t.Fail()
		}
	}
}

type translateFormatTest struct {
	input    string
	expected string
}

func TestTranslateFormat(t *testing.T) {
	tests := []translateFormatTest{
		translateFormatTest{"DateOnly", time.DateOnly},
		translateFormatTest{"SomeNewTime", "SomeNewTime"},
	}

	for _, test := range tests {
		actual := translateFormat(test.input)
		if test.expected != actual {
			t.Logf("Expected %s for input %s but got %s", test.expected, test.input, actual)
			t.Fail()
		}
	}

}
