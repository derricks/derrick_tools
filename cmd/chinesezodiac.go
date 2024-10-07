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

	"github.com/spf13/cobra"
)

// shakespeareCmd represents the shakespeare command
var chineseZodiacCmd = &cobra.Command{
	Use:   "chinesezodiac",
	Short: "Test recall of the Chinese zodiac",
	Run:   quizChineseZodiacs,
}

type chineseZodiacInfo struct {
	animal        string
	referenceYear int
}

var chineseZodiac = []chineseZodiacInfo{
	{"rat", 1960},
	{"ox", 1961},
	{"tiger", 1962},
	{"rabbit", 1963},
	{"dragon", 1964},
	{"snake", 1965},
	{"horse", 1966},
	{"goat", 1967},
	{"monkey", 1968},
	{"rooster", 1969},
	{"dog", 1970},
	{"pig", 1971},
}

type chineseZodiacQuiz func([]chineseZodiacInfo) promptAndResponse

func quizChineseZodiacs(cmd *cobra.Command, args []string) {
	quizzes := []chineseZodiacQuiz{
		quizChineseZodiacAnimalByIndex,
		quizChineseZodiacByYear,
	}

	quiz := randomItemFromSlice(quizzes)
	promptAndCheckResponse(quiz(chineseZodiac))
}

func quizChineseZodiacAnimalByIndex(zodiac []chineseZodiacInfo) promptAndResponse {
	index := rand.Intn(len(zodiac))
	return promptAndResponse{fmt.Sprintf("What Chinese zodiac animal is at position %d", index+1), zodiac[index].animal}
}

func quizChineseZodiacByYear(zodiac []chineseZodiacInfo) promptAndResponse {
	yearOffset := rand.Intn(100)
	targetYear := zodiac[0].referenceYear + yearOffset
	animal := zodiac[(targetYear-zodiac[0].referenceYear)%12]
	return promptAndResponse{fmt.Sprintf("What is the chinese zodiac animal for %d?", targetYear), animal.animal}
}

func init() {
	memoryquizCmd.AddCommand(chineseZodiacCmd)
}
