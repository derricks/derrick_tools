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

// shakespeareCmd represents the shakespeare command
var shakespeareCmd = &cobra.Command{
	Use:   "shakespeare",
	Short: "Test recall of the names of Shakespeare's plays",
	Long:  `The exact chronology of Shakespeare's plays is difficult to gauge. This uses the ordering found at https://en.wikipedia.org/wiki/Chronology_of_Shakespeare%27s_plays as of 2021-03-08`,
	Run:   quizShakespeare,
}

var shakespearePlays = []string{
	"The Two Gentlemen of Verona",
	"The Taming of the Shrew",
	"Henry VI, Part 2",
	"Henry VI, Part 3",
	"Henry VI, Part 1",
	"Titus Andronicus",
	"Richard III",
	"Edward III",
	"The Comedy of Errors",
	"Love's Labour's Lost",
	"Love's Labour's Won",
	"Richard II",
	"Romeo and Juliet",
	"A Midsummer Night's Dream",
	"King John",
	"The Merchant of Venice",
	"Henry IV, Part 1",
	"The Merry Wives of Windsor",
	"Henry IV, Part 2",
	"Much Ado About Nothing",
	"Henry V",
	"Julius Caesar",
	"As You Like It",
	"Hamlet",
	"Twelfth Night",
	"Troilus and Cressida",
	"Sir Thomas More",
	"Measure for Measure",
	"Othello",
	"All's Well That Ends Well",
	"King Lear",
	"Timon of Athens",
	"MacBeth",
	"Antony and Cleopatra",
	"Pericles, Prince of Tyre",
	"Coriolanus",
	"A Winter's Tale",
	"Cymbelline",
	"The Tempest",
	"Cardenio",
	"Henry VIII",
	"The Two Noble Kinsmen",
}

type shakespeareQuiz func([]string) promptAndResponse

func quizShakespeare(cmd *cobra.Command, args []string) {
	quizzes := []shakespeareQuiz{
		quizShakespearePlayFromIndex,
		quizIndexOfShakespearePlay,
	}

	quiz := randomItemFromSlice(quizzes)
	promptAndCheckResponse(quiz(shakespearePlays))
}

func quizShakespearePlayFromIndex(plays []string) promptAndResponse {
	return quizStringAtIndexInList("Shakespeare play", shakespearePlays)
}

func quizIndexOfShakespearePlay(plays []string) promptAndResponse {
	return quizIndexOfStringInList(plays)
}

func init() {
	memoryquizCmd.AddCommand(shakespeareCmd)
}
