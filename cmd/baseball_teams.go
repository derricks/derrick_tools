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

var baseballCmd = &cobra.Command{
	Use:   "baseball-teams",
	Short: "Quiz baseball teams",
	Run:   quizBaseballTeams,
}

type baseballTeam struct {
	index int    `crossquery:"all"`
	name  string `crossquery:"all"`
	// some areas, such as New York, have more than one team
	area string `crossquery:"guess"`
}

var baseballTeams = []baseballTeam{
	{1, "Diamondbacks", "Arizona"},
	{2, "Braves", "Atlanta"},
	{3, "Orioles", "Baltimore"},
	{4, "Red Sox", "Boston"},
	{5, "Cubs", "Chicago"},
	{6, "White Sox", "Chicago"},
	{7, "Reds", "Cincinnati"},
	{8, "Guardians", "Cleveland"},
	{9, "Rockies", "Colorado"},
	{10, "Tigers", "Detroit"},
	{11, "Astros", "Houston"},
	{12, "Royals", "Kansas City"},
	{13, "Angels", "Los Angeles"},
	{14, "Dodgers", "Los Angeles"},
	{15, "Marlins", "Miami"},
	{16, "Brewers", "Milwaukee"},
	{17, "Twins", "Minnesota"},
	{18, "Mets", "New York"},
	{19, "Yankees", "New York"},
	{20, "Athletics", "Oakland"},
	{21, "Phillies", "Philadelphia"},
	{22, "Pirates", "Pittsburgh"},
	{23, "Padres", "San Diego"},
	{24, "Giants", "San Francisco"},
	{25, "Mariners", "Seattle"},
	{26, "Cardinals", "St. Louis"},
	{27, "Rays", "Tampa Bay"},
	{28, "Rangers", "Texas"},
	{29, "Blue Jays", "Toronto"},
	{30, "Nationals", "Washington"},
}

type baseballQuestion func([]baseballTeam) promptAndResponse

func quizBaseballTeams(cmd *cobra.Command, args []string) {

	var promptFuncs = []baseballQuestion{
		crossQueryBaseballTeamInfo,
	}

	function := promptFuncs[rand.Intn(len(promptFuncs))]
	promptAndCheckResponse(function(baseballTeams))
}

func crossQueryBaseballTeamInfo(teams []baseballTeam) promptAndResponse {
	foundTeam := teams[rand.Intn(len(teams))]
	return constructCrossQuery("baseball team", foundTeam)
}

func init() {
	memoryquizCmd.AddCommand(baseballCmd)
}
