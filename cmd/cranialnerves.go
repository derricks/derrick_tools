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
	"github.com/spf13/cobra"
)

// cranialnervesCmd represents the cranialnerves command
var cranialnervesCmd = &cobra.Command{
	Use:   "cranialnerves",
	Short: "Quiz about the cranial nerves",
	Run:   quizCranialNerves,
}

var cranialNerves = []string{
	"olfactory",
	"optic",
	"oculomotor",
	"trochlear",
	"trigeminal",
	"abducens",
	"facial",
	"auditory",
	"glossopharyngeal",
	"vagus",
	"spinal accessory",
	"hypoglossal",
}

type quizNerveFunc func([]string) promptAndResponse

func quizCranialNerves(cmd *cobra.Command, args []string) {
	nerveQuizzes := []quizNerveFunc{
		quizCranialNerveByIndex,
		quizCranialNerveIndexFromName,
	}

	quizFunction := randomItemFromSlice(nerveQuizzes)
	promptAndCheckResponse(quizFunction(cranialNerves))
}

func quizCranialNerveByIndex(nerves []string) promptAndResponse {
	return quizStringAtIndexInList("cranial nerve", nerves)
}

func quizCranialNerveIndexFromName(nerves []string) promptAndResponse {
	return quizIndexOfStringInList(nerves)
}

func init() {
	memoryquizCmd.AddCommand(cranialnervesCmd)
}
