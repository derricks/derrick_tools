/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"math/rand"
	"strconv"

	"github.com/spf13/cobra"
)

// greekCmd represents the greek command
var greekCmd = &cobra.Command{
	Use:   "greek",
	Short: "Test memory of the Greek alphabet",
	Run:   quizGreekAlphabet,
}

var greekAlphabet = []string{
	"alpha",
	"beta",
	"gamma",
	"delta",
	"epsilon",
	"zeta",
	"eta",
	"theta",
	"iota",
	"kappa",
	"lambda",
	"mu",
	"nu",
	"xi",
	"omicron",
	"pi",
	"rho",
	"sigma",
	"tau",
	"upsilon",
	"phi",
	"chi",
	"psi",
	"omega",
}

type quizGreekFunc func([]string) promptAndResponse

func quizGreekAlphabet(cmd *cobra.Command, args []string) {
	funcs := []quizGreekFunc{
		quizPositionFromGreekLetter,
		quizGreekLetterFromPosition,
		quizGreekLetterBefore,
		quizGreekLetterAfter,
	}

	function := funcs[rand.Intn(len(funcs))]
	promptAndCheckResponse(function(greekAlphabet))
}

func quizPositionFromGreekLetter(alphabet []string) promptAndResponse {
	index := rand.Intn(len(alphabet))
	return promptAndResponse{fmt.Sprintf("What position is letter %s?", alphabet[index]), strconv.Itoa(index + 1)}
}

func quizGreekLetterFromPosition(alphabet []string) promptAndResponse {
	index := rand.Intn(len(alphabet))
	return promptAndResponse{fmt.Sprintf("What letter is at position %d?", index+1), alphabet[index]}
}

func quizGreekLetterBefore(alphabet []string) promptAndResponse {
	index := 0
	for index == 0 {
		index = rand.Intn(len(alphabet))
	}
	return promptAndResponse{fmt.Sprintf("What letter comes before %s?", alphabet[index]), alphabet[index-1]}
}

func quizGreekLetterAfter(alphabet []string) promptAndResponse {
	index := rand.Intn(len(alphabet) - 1) // -1 to ensure we don't get the last item
	return promptAndResponse{fmt.Sprintf("What letter comes after %s?", alphabet[index]), alphabet[index+1]}
}

func init() {
	memoryquizCmd.AddCommand(greekCmd)
}
