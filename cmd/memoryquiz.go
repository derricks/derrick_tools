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
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type promptAndResponse struct {
	prompt   string
	response string
}

// promptAndCheckResponse will use promot to pose a question to the user and wait for
// a response. If correct, it will print Correct! and return true. Otherwise it will
// print the user's answer and the right answer and return false
func promptAndCheckResponse(prompt promptAndResponse) bool {
	start := time.Now()
	userResponse := responseFromPrompt(prompt)
	fmt.Printf("You took %v to answer\n", time.Now().Sub(start))
	if userResponse != "" {
		if strings.TrimSpace(userResponse) == prompt.response {
			fmt.Println("Correct!")
			return true
		} else {
			fmt.Printf("Incorrect. The right answer was %s\n", prompt.response)
			return false
		}
	}
	return false
}

func responseFromPrompt(prompt promptAndResponse) string {
	fmt.Println(prompt.prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""

}

// quizIndexOfStringInList will ask you to identify the index of a random item within the set of items
// for instance, you might get a question such as "what position is greek letter eta?"
func quizIndexOfStringInList(items []string) promptAndResponse {
	itemIndex := rand.Intn(len(items))
	return promptAndResponse{fmt.Sprintf("What position is %s?", items[itemIndex]), strconv.Itoa(itemIndex + 1)}
}

// quizStringAtIndexInList will ask you to identify what string is at the given position in items
// for instance, you might get a question such as "which hebrew letter is at position 2"
func quizStringAtIndexInList(itemName string, items []string) promptAndResponse {
	itemIndex := rand.Intn(len(items))
	return promptAndResponse{fmt.Sprintf("What %s is at position %d?", itemName, itemIndex+1), items[itemIndex]}
}

type areaToQuizFunction struct {
	area     string
	quizFunc func(*cobra.Command, []string)
}

// memoryquizCmd represents the memoryquiz command
var memoryquizCmd = &cobra.Command{
	Use:   "memoryquiz",
	Short: "Fire up various memory quizzes",
	Run: func(cmd *cobra.Command, args []string) {
		// select an arbitrary memory quiz to run
		areaToQuizFuncs := []areaToQuizFunction{
			areaToQuizFunction{"presidents", quizPresidents},
			areaToQuizFunction{"countries", quizCountries},
			areaToQuizFunction{"digits of pi", quizPiDigts},
			areaToQuizFunction{"greek alphabet", quizGreekAlphabet},
			areaToQuizFunction{"hebrew alphabet", quizHebrewAlphabet},
			areaToQuizFunction{"cranial nerves", quizCranialNerves},
			areaToQuizFunction{"shakespeare", quizShakespeare},
			areaToQuizFunction{"powers of two", quizPowersOfTwo},
			areaToQuizFunction{"elements", quizElements},
			areaToQuizFunction{"states", quizStates},
			areaToQuizFunction{"muses", quizMuses},
			areaToQuizFunction{"chinesezodiac", quizChineseZodiacs},
			areaToQuizFunction{"bible books", quizKingJames},
			areaToQuizFunction{"canada", quizCanada},
			areaToQuizFunction{"lakes", quizLakes},
			areaToQuizFunction{"constellations", quizConstellations},
			areaToQuizFunction{"arrondisements", quizArrondisements},
			areaToQuizFunction{"spellingbee", quizSpellingBee},
			areaToQuizFunction{"monopoly", quizMonopoly},
			areaToQuizFunction{"numbers", quizLargeNumbers},
			areaToQuizFunction{"doomsday", quizDoomsday},
			areaToQuizFunction{"whosonfirst", quizWhosonfirst},
			areaToQuizFunction{"english royalty", quizEnglishRoyalty},
			areaToQuizFunction{"grand crus", quizGrandCrus},
			areaToQuizFunction{"bottles", quizWineBottles},
			areaToQuizFunction{"ca-counties", quizCaCounties},
			areaToQuizFunction{"baseball-teams", quizBaseballTeams},
			areaToQuizFunction{"football-teams", quizFootballTeams},
		}

		areaToQuiz := areaToQuizFuncs[rand.Intn(len(areaToQuizFuncs))]
		fmt.Printf("[%s]\n", areaToQuiz.area)
		areaToQuiz.quizFunc(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(memoryquizCmd)
}
