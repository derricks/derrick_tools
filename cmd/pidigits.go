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

// pidigitsCmd represents the pidigits command
var pidigitsCmd = &cobra.Command{
	Use:   "pidigits",
	Short: "Quiz recall of chunks of pi",
	Run:   quizPiDigts,
}

var piChunks = []string{
	"3141",
	"5926",
	"5358",
	"9793",
	"2384",
	"6264",
	"3383",
	"2795",
	"0288",
	"4197",
	"1693",
	"9937",
	"5105",
	"8209",
	"7494",
	"4592",
	"3078",
	"1640",
	"6286",
	"2089",
	"9862",
	"8034",
	"8253",
	"4211",
	"7067",
	"9821",
	"4808",
	"6513",
	"2823",
	"0664",
	"7093",
	"8446",
	"0955",
	"0582",
	"2317",
	"2535",
	"9408",
	"1284",
	"8111",
	"7450",
	"2841",
	"0270",
	"1938",
	"5211",
	"0555",
	"9644",
	"6229",
	"4895",
	"4930",
	"3819",
	"6442",
	"8810",
}

type quizPi func([]string) promptAndResponse

func quizPiDigts(cmd *cobra.Command, args []string) {
	quizzes := []quizPi{
		quizIndexOfPiChunk,
		quizPiChunkByIndex,
		quizPiDigitByIndex,
	}

	quiz := quizzes[rand.Intn(len(quizzes))]
	promptAndCheckResponse(quiz(piChunks))
}

func quizIndexOfPiChunk(chunks []string) promptAndResponse {
	return quizIndexOfStringInList(chunks)
}

func quizPiChunkByIndex(chunks []string) promptAndResponse {
	return quizStringAtIndexInList("pi chunk", chunks)
}

func quizPiDigitByIndex(chunks []string) promptAndResponse {
	chunkIndex := rand.Intn(len(piChunks))
	digits := strings.Split(chunks[chunkIndex], "")
	indexInChunk := rand.Intn(len(digits))
	return promptAndResponse{fmt.Sprintf("What pi digit is at position %d?", (chunkIndex*4)+indexInChunk+1), digits[indexInChunk]}
}

func init() {
	memoryquizCmd.AddCommand(pidigitsCmd)
}
