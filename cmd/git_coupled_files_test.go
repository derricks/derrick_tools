package cmd

import (
	"testing"
)

func TestTrackerCouplings(test *testing.T) {
	tracker := newTracker("abcd")
	tracker.coupleFile("xyz")
	tracker.coupleFile("xyz")
	tracker.coupleFile("def")

	expectedCounts := map[string]int{
		"xyz": 2,
		"def": 1,
	}

	couplings := tracker.getCoupledFiles()
	if len(couplings) != len(expectedCounts) {
		test.Errorf("Expected %d items but got %d items", len(expectedCounts), len(couplings))
	}

  for _, coupling := range couplings {
    expectedCount := expectedCounts[coupling.name]
    if coupling.coupledCount != expectedCount {
      test.Errorf("Expected %d couplings for %s but was %d", expectedCount, coupling.name, coupling.coupledCount)
    }
  }
}
