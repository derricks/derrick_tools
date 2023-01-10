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
var whosonfirstCmd = &cobra.Command{
	Use:   "whosonfirst",
	Short: "Test recall of the players in Who's On First",
	Run:   quizWhosonfirst,
}

type whosonfirst struct {
	name     string
	position string
}

var whosOnFirstPlayers = []whosonfirst{
	whosonfirst{"Today", "Catcher"},
	whosonfirst{"Tomorrow", "Pitcher"},
	whosonfirst{"Who", "First Base"},
	whosonfirst{"What", "Second Base"},
	whosonfirst{"Because", "Center Field"},
	whosonfirst{"Shortstop", "I Don't Care"},
	whosonfirst{"I Don't Know", "Third Base"},
	whosonfirst{"Why", "Left Field"},
}

type whosonfirstQuiz func([]whosonfirst) promptAndResponse

func quizWhosonfirst(cmd *cobra.Command, args []string) {
	quizzes := []whosonfirstQuiz{
		quizPlayerByPosition,
		quizPositionByPlayer,
	}

	quiz := quizzes[rand.Intn(len(quizzes))]
	promptAndCheckResponse(quiz(whosOnFirstPlayers))
}

func quizPlayerByPosition(players []whosonfirst) promptAndResponse {
	player := randomPlayer(players)
	return promptAndResponse{fmt.Sprintf("Name the player at %s", player.position), player.name}
}

func quizPositionByPlayer(players []whosonfirst) promptAndResponse {
	player := randomPlayer(players)
	return promptAndResponse{fmt.Sprintf("What position does %s play?", player.name), player.position}
}

func randomPlayer(players []whosonfirst) whosonfirst {
	return players[rand.Intn(len(players))]
}

func init() {
	memoryquizCmd.AddCommand(whosonfirstCmd)
}
