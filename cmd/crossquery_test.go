/*
Copyright © 2022 Derrick Schneider derrick.schneider@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"math/rand"
	"testing"
	"time"
)

// this can only produce one response string
type crossQuery1 struct {
	value1 string `crossquery:"given"`
	value2 string `crossquery:"guess"`
}

func TestCrossQueryOneGivenOneGuess(test *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := crossQuery1{"abc", "def"}
	result := constructCrossQuery("test1", s)
	if result.prompt != "What is the value2 of the test1 with value1 of abc?" {
		test.Errorf("Incorrect prompt. Was \"%s\"", result.prompt)
	}

	if result.response != "def" {
		test.Errorf("Incorrect response. Was \"%s\"", result.response)
	}
}

type crossQuery2 struct {
	v1 string `crossquery:""`
	v2 int    `crossquery:"all"`
}

func TestCrossQueryAll(test *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := crossQuery2{"xyz", 34}
	result := constructCrossQuery("test2", s)

	// can be one of two
	if result.prompt != "What is the v2 of the test2 with v1 of xyz?" &&
		result.prompt != "What is the v1 of the test2 with v2 of 34?" {
		test.Errorf("Invalid prompt. Was %s\n", result.prompt)
	}

	if result.response != "xyz" && result.response != "34" {
		test.Errorf("Invalid response. Was \"%s\"\n", result.response)
	}
}

type crossQuery3 struct {
	givens  []string `crossquery:"given"`
	guesses []string `crossquery:"guess"`
}

func TestCrossQuerySlice(test *testing.T) {
	s := crossQuery3{[]string{"givenslice"}, []string{"guessslice"}}
	result := constructCrossQuery("test3", s)
	if result.prompt != "What is the guesses of the test3 with givens of givenslice?" {
		test.Errorf("Invalid prompt. Was %s", result.prompt)
	}

	if result.response != "guessslice" {
		test.Errorf("Invalid response. Was %s", result.response)
	}
}
