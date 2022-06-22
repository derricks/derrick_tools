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
	chineseZodiacInfo{"rat", 1960},
	chineseZodiacInfo{"ox", 1961},
	chineseZodiacInfo{"tiger", 1962},
	chineseZodiacInfo{"rabbit", 1963},
	chineseZodiacInfo{"dragon", 1964},
	chineseZodiacInfo{"snake", 1965},
	chineseZodiacInfo{"horse", 1966},
	chineseZodiacInfo{"goat", 1967},
	chineseZodiacInfo{"monkey", 1968},
	chineseZodiacInfo{"rooster", 1969},
	chineseZodiacInfo{"dog", 1970},
	chineseZodiacInfo{"pig", 1971},
}

type chineseZodiacQuiz func([]chineseZodiacInfo) promptAndResponse

func quizChineseZodiacs(cmd *cobra.Command, args []string) {
	quizzes := []chineseZodiacQuiz{
		quizChineseZodiacAnimalByIndex,
		quizChineseZodiacByYear,
	}

	quiz := quizzes[rand.Intn(len(quizzes))]
	promptAndCheckResponse(quiz(chineseZodiac))
}

func quizChineseZodiacAnimalByIndex(zodiac []chineseZodiacInfo) promptAndResponse {
	index := rand.Intn(len(zodiac))
	return promptAndResponse{fmt.Sprintf("What Chinese zodiac animal is at position %d", index+1), zodiac[index].animal}
}

func quizChineseZodiacByYear(zodiac []chineseZodiacInfo) promptAndResponse {
	yearOffset := rand.Intn(100)
	targetYear := zodiac[0].referenceYear + yearOffset
	animal := zodiac[0]
	var i int
	for i = 0; (targetYear-animal.referenceYear)%12 != 0; i++ {
		animal = zodiac[i]
	}
	return promptAndResponse{fmt.Sprintf("What is the chinese zodiac animal for %d?", targetYear), animal.animal}
}

func init() {
	memoryquizCmd.AddCommand(chineseZodiacCmd)
}
