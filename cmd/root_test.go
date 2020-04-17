package cmd

import (
	"bufio"
	"strings"
	"testing"
)

func TestReaderToChannel(test *testing.T) {
	testData := "line1\nline2 and some  \nline3 is the best"
	output := make(chan string)
	results := make([]string, 0, 3)
	go pushReaderToChannel(bufio.NewReader(strings.NewReader(testData)), output)

	for line := range output {
		results = append(results, line)
	}

	if len(results) != 3 {
		test.Errorf("Expected 3 results but got %d", len(results))
	}

	expectedStrings := []string{"line1", "line2 and some", "line3 is the best"}
	for index, curString := range expectedStrings {
		if results[index] != curString {
			test.Errorf("Expected item %d to be %s but was %s", index, curString, results[index])
		}
	}

}

func TestStringInSlice(test *testing.T) {
	testData := []string{"a", "b", "c"}
	for _, curTest := range testData {
		if !isStringInSlice(curTest, testData) {
			test.Errorf("Expected %s to be in %v but it is not", curTest, testData)
		}
	}

	if isStringInSlice("d", testData) {
		test.Errorf("Did not expect to find d in %v, but did", testData)
	}
}
