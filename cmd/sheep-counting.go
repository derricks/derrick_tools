/*
Copyright © 2024 Derrick Schneider derrick.schneider@gmail.com

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

// hebrewCmd represents the hebrew command
var sheepCountingCmd = &cobra.Command{
	Use:   "sheep",
	Short: "Quiz command of English sheep counting",
	Run:   quizSheepCounting,
}

var sheepCounting = []string{
	"yain",
	"tain",
	"eddera",
	"pedera",
	"pit",
	"tayter",
	"layter",
	"overa",
	"covera",
	"dix",
	"yain-a-dix",
	"tain-a-dix",
	"eddera-dix",
	"pedera-a-dix",
	"bumfit",
	"yain-a-bumfit",
	"tain-a-bumfit",
	"eddera-bumfit",
	"pedera-a-bumfit",
	"jiggit",
}

type quizSheepCountingFunc func([]string) promptAndResponse

func quizSheepCounting(cmd *cobra.Command, args []string) {
	funcs := []quizHebrewFunc{
		quizPositionFromLetter,
		quizSheepCountingTermFromPosition,
		quizLetterBefore,
		quizLetterAfter,
	}

	function := randomItemFromSlice(funcs)
	promptAndCheckResponse(function(sheepCounting))
}

func quizSheepCountingTermFromPosition(alphabet []string) promptAndResponse {
	return quizStringAtIndexInList("sheep counting term", alphabet)
}

func init() {
	memoryquizCmd.AddCommand(sheepCountingCmd)
}
