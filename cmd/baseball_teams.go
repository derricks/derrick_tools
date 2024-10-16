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

var baseballCmd = &cobra.Command{
	Use:   "baseball-teams",
	Short: "Quiz baseball teams",
	Run:   quizBaseballTeams,
}

type baseballTeam struct {
	index int    `crossquery:"all"`
	name  string `crossquery:"all"`
	// some areas, such as New York, have more than one team
	area    string   `crossquery:"guess"`
	mascots []string `crossquery:"given"`
}

var baseballTeams = []baseballTeam{
	{1, "Diamondbacks", "Arizona", []string{"D. Baxter the Bobcat", "D-backs Luchador"}},
	{2, "Braves", "Atlanta", []string{"Blooper"}},
	{3, "Orioles", "Baltimore", []string{"The Oriole Bird"}},
	{4, "Red Sox", "Boston", []string{"Wally the Green Monster", "Tessie"}},
	{5, "Cubs", "Chicago", []string{"Clark"}},
	{6, "White Sox", "Chicago", []string{"Southpaw"}},
	{7, "Reds", "Cincinnati", []string{"Mr. Red", "Gapper", "Mr. Redlegs", "Rosie Red"}},
	{8, "Guardians", "Cleveland", []string{"Slider"}},
	{9, "Rockies", "Colorado", []string{}},
	{10, "Tigers", "Detroit", []string{}},
	{11, "Astros", "Houston", []string{}},
	{12, "Royals", "Kansas City", []string{}},
	{13, "Angels", "Los Angeles", []string{}},
	{14, "Dodgers", "Los Angeles", []string{}},
	{15, "Marlins", "Miami", []string{}},
	{16, "Brewers", "Milwaukee", []string{}},
	{17, "Twins", "Minnesota", []string{}},
	{18, "Mets", "New York", []string{}},
	{19, "Yankees", "New York", []string{}},
	{20, "Athletics", "Oakland", []string{}},
	{21, "Phillies", "Philadelphia", []string{}},
	{22, "Pirates", "Pittsburgh", []string{}},
	{23, "Padres", "San Diego", []string{}},
	{24, "Giants", "San Francisco", []string{}},
	{25, "Mariners", "Seattle", []string{}},
	{26, "Cardinals", "St. Louis", []string{}},
	{27, "Rays", "Tampa Bay", []string{}},
	{28, "Rangers", "Texas", []string{}},
	{29, "Blue Jays", "Toronto", []string{}},
	{30, "Nationals", "Washington", []string{}},
}

type baseballQuestion func([]baseballTeam) promptAndResponse

func quizBaseballTeams(cmd *cobra.Command, args []string) {

	var promptFuncs = []baseballQuestion{
		crossQueryBaseballTeamInfo,
	}

	function := randomItemFromSlice(promptFuncs)
	promptAndCheckResponse(function(baseballTeams))
}

func crossQueryBaseballTeamInfo(teams []baseballTeam) promptAndResponse {
	foundTeam := randomItemFromSlice(teams)
	return constructCrossQuery("baseball team", foundTeam)
}

func init() {
	memoryquizCmd.AddCommand(baseballCmd)
}
