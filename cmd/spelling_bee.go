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
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

var spellingBeeCmd = &cobra.Command{
	Use:   "spellingbee",
	Short: "Quiz Spelling Bee word sets",
	Run:   quizSpellingBee,
}

var spellingBeeSets = [][]string{
	{"FAIR", "FRIAR", "AFFAIR", "RIFFRAFF", "RAFFIA"},
	{"LATHE", "ATHLETE", "LETHAL", "HEALTH", "TELEHEALTH"},
}

func quizSpellingBee(cmd *cobra.Command, args []string) {
	wordSet := spellingBeeSets[rand.Intn(len(spellingBeeSets))]
	word := wordSet[rand.Intn(len(wordSet))]
	inputSet := responseFromPrompt(promptAndResponse{fmt.Sprintf("What are other Spelling Bee words for %s (separate by commas)?", word), ""})

	enteredWords := strings.Split(inputSet, ",")

	// verify that entered words doesn't have entries not in the list
	for _, enteredWord := range enteredWords {
		if !isStringInSlice(enteredWord, wordSet) {
			fmt.Printf("%s is not in the list of words for %s\n", enteredWord, word)
		}
	}

	// and now verify that the list isn't missing any words
	for _, validWord := range wordSet {
		if !isStringInSlice(validWord, enteredWords) && validWord != word {
			fmt.Printf("You missed %s\n", validWord)
		}
	}
}

func init() {
	memoryquizCmd.AddCommand(spellingBeeCmd)
}
