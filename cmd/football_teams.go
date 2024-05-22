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
	"math/rand"

	"github.com/spf13/cobra"
)

var footballCmd = &cobra.Command{
	Use:   "football-teams",
	Short: "Quiz US football teams",
	Run:   quizFootballTeams,
}

type footballTeam struct {
	index int    `crossquery:"all"`
	name  string `crossquery:"all"`
	// some areas, such as New York, have more than one team
	area   string `crossquery:"guess"`
	league string `crossquery:"guess"`
}

const (
	AFC_EAST  = "AFC East"
	AFC_NORTH = "AFC North"
	AFC_SOUTH = "AFC South"
	AFC_WEST  = "AFC West"
	NFC_EAST  = "NFC East"
	NFC_NORTH = "NFC North"
	NFC_SOUTH = "NFC South"
	NFC_WEST  = "NFC West"
)

var footballTeams = []footballTeam{
	{1, "Bills", "Buffalo", AFC_EAST},
	{2, "Dolphins", "Miami", AFC_EAST},
	{3, "Patriots", "New England", AFC_EAST},
	{4, "Jets", "New York", AFC_EAST},
	{5, "Ravens", "Baltimore", AFC_NORTH},
	{6, "Bengals", "Cincinnati", AFC_NORTH},
	{7, "Browns", "Cleveland", AFC_NORTH},
	{8, "Steelers", "Pittsburgh", AFC_NORTH},
	{9, "Texans", "Houston", AFC_SOUTH},
	{10, "Colts", "Indianapolis", AFC_SOUTH},
	{11, "Jaguars", "Jacksonville", AFC_SOUTH},
	{12, "Titans", "Tennessee", AFC_SOUTH},
	{13, "Broncos", "Denver", AFC_WEST},
	{14, "Chiefs", "Kansas City", AFC_WEST},
	{15, "Raiders", "Las Vegas", AFC_WEST},
	{16, "Chargers", "Los Angeles", AFC_WEST},
	{17, "Cowboys", "Dallas", NFC_EAST},
	{18, "Giants", "New York", NFC_EAST},
	{19, "Eagles", "Philadelphia", NFC_EAST},
	{20, "Commanders", "Washington", NFC_EAST},
	{21, "Bears", "Chicago", NFC_NORTH},
	{22, "Lions", "Detroit", NFC_NORTH},
	{23, "Packers", "Green Bay", NFC_NORTH},
	{24, "Vikings", "Minnesota", NFC_NORTH},
	{25, "Falcons", "Atlanta", NFC_SOUTH},
	{26, "Panthers", "Carolina", NFC_SOUTH},
	{27, "Saints", "New Orleans", NFC_SOUTH},
	{28, "Buccaneers", "Tampa Bay", NFC_SOUTH},
	{29, "Cardinals", "Arizona", NFC_WEST},
	{30, "Rams", "Los Angeles", NFC_WEST},
	{31, "49ers", "San Francisco", NFC_WEST},
	{32, "Seahawks", "Seattle", NFC_WEST},
}

type footballQuestion func([]footballTeam) promptAndResponse

func quizFootballTeams(cmd *cobra.Command, args []string) {

	var promptFuncs = []footballQuestion{
		crossQueryFootballTeamInfo,
	}

	function := promptFuncs[rand.Intn(len(promptFuncs))]
	promptAndCheckResponse(function(footballTeams))
}

func crossQueryFootballTeamInfo(teams []footballTeam) promptAndResponse {
	foundTeam := teams[rand.Intn(len(teams))]
	return constructCrossQuery("football team", foundTeam)
}

func init() {
	memoryquizCmd.AddCommand(footballCmd)
}
