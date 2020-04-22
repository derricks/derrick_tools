package cmd

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func TestDataSetInsert(test *testing.T) {
	dataPoints := dataSet(make(map[string]*dataPoint))
	dataPoints.addDataPoint("test", 3)
	dataPoints.addDataPoint("test", 4)

	testDataPoint := dataPoints["test"]
	if testDataPoint.label != "test" {
		test.Errorf("Expected data point to have a label of %s but was %s", "test", testDataPoint.label)
	}

	if testDataPoint.quantity != 7 {
		test.Errorf("Expected data point to have a quantity of %d but was %d", 7, testDataPoint.quantity)
	}
}

func TestDataSetLongestLabel(test *testing.T) {
	dataPoints := dataSet(make(map[string]*dataPoint))

	// longest label in an empty set is an empty string
	if dataPoints.longestLabel() != "" {
		test.Errorf("Longest label in an empty should be the empty string, was %s", dataPoints.longestLabel())
	}

	dataPoints.addDataPoint("snake", 4)
	dataPoints.addDataPoint("bat", 5)
	if dataPoints.longestLabel() != "snake" {
		test.Errorf("Longest label should have been snake was %s", dataPoints.longestLabel())
	}
}

func TestDataSetLargestValue(test *testing.T) {
	dataPoints := dataSet(make(map[string]*dataPoint))

	if dataPoints.largestValue() != math.MinInt32 {
		test.Errorf("Expected largestValue to be %v was %v", math.MinInt32, dataPoints.largestValue())
	}

	dataPoints.addDataPoint("snake", 4)
	dataPoints.addDataPoint("bat", 3)
	if dataPoints.largestValue() != 4 {
		test.Errorf("Expected larges value to be %d but was %d", 4, dataPoints.largestValue())
	}
}

type scaleBarTest struct {
	dataValue   int
	maxBarValue int
	maxBarWidth int
	expected    int
}

func TestScalingBarWdith(test *testing.T) {
	tests := []scaleBarTest{
		scaleBarTest{0, 100, 20, 0},
		scaleBarTest{100, 100, 20, 20},
		scaleBarTest{4, 8, 20, 10},
	}
	for index, curTest := range tests {
		result := scaledBarWidth(curTest.dataValue, curTest.maxBarValue, curTest.maxBarWidth)
		if result != curTest.expected {
			test.Errorf("Test %d. Expected %d but got %d", index, curTest.expected, result)
		}
	}
}

func TestCollectData(test *testing.T) {
	testFile := "a 3\na 4\nb 5"
	var currentData dataSet

	preDataCalls := 0
	postDataCalls := 0

	collectData(func(data dataSet) { preDataCalls++ }, func(data dataSet) {
		postDataCalls++
	}, bufio.NewReader(strings.NewReader(testFile)))

	if preDataCalls != 3 {
		test.Errorf("Not enough calls to preData. Should have been 3, was %d", preDataCalls)
	}

	if postDataCalls != 3 {
		test.Errorf("Not enough calls to postData. Should have been 3, was %d", postDataCalls)
	}

	test.Log(currentData)

}

type centerTest struct {
	textToCenter    string
	widthToCenterIn int
	expected        string
}

func TestCenterText(test *testing.T) {
	tests := []centerTest{
		centerTest{"ab", 10, "    ab"},
		centerTest{"abc", 2, "ab"},
		centerTest{"ab", 5, " ab"},
	}

	for index, curTest := range tests {
		returnString := centerTextInSpace(curTest.textToCenter, curTest.widthToCenterIn)
		if returnString != curTest.expected {
			test.Errorf("Test case %d: Expected '%s' but got '%s'", index, curTest.expected, returnString)
		}
	}
}
