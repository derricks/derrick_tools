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
	"math/rand"

	"github.com/spf13/cobra"
)

var whosonfirstCmd = &cobra.Command{
	Use:   "whosonfirst",
	Short: "Test recall of the players in Who's On First",
	Run:   quizWhosonfirst,
}

type whosonfirst struct {
	name     string `crossquery:"all"`
	position string `crossquery:"all"`
}

var whosOnFirstPlayers = []whosonfirst{
	{"Today", "Catcher"},
	{"Tomorrow", "Pitcher"},
	{"Who", "First Base"},
	{"What", "Second Base"},
	{"Because", "Center Field"},
	{"I Don't Care", "Shortstop"},
	{"I Don't Know", "Third Base"},
	{"Why", "Left Field"},
}

type whosonfirstQuiz func([]whosonfirst) promptAndResponse

func quizWhosonfirst(cmd *cobra.Command, args []string) {
	player := randomPlayer(whosOnFirstPlayers)
	promptAndCheckResponse(constructCrossQuery("who's on first player", player))
}

func randomPlayer(players []whosonfirst) whosonfirst {
	return players[rand.Intn(len(players))]
}

func init() {
	memoryquizCmd.AddCommand(whosonfirstCmd)
}
