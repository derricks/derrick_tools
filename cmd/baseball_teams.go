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
	{9, "Rockies", "Colorado", []string{"Dinger"}},
	{10, "Tigers", "Detroit", []string{"Paws"}},
	{11, "Astros", "Houston", []string{"Orbit"}},
	{12, "Royals", "Kansas City", []string{"Sluggerrr"}},
	{13, "Angels", "Los Angeles", []string{}},
	{14, "Dodgers", "Los Angeles", []string{}},
	{15, "Marlins", "Miami", []string{"Billy"}},
	{16, "Brewers", "Milwaukee", []string{"Bernie Brewer", "Barrelman", "The Sausages"}},
	{17, "Twins", "Minnesota", []string{"T. C. Bear"}},
	{18, "Mets", "New York", []string{"Mr. Met", "Jan Met"}},
	{19, "Yankees", "New York", []string{}},
	{20, "Athletics", "Oakland", []string{"Stomper"}},
	{21, "Phillies", "Philadelphia", []string{"Phillie Phanatic", "Phoebe Phanatic", "Phred"}},
	{22, "Pirates", "Pittsburgh", []string{"The Pierogis"}},
	{23, "Padres", "San Diego", []string{"Swinging Friar"}},
	{24, "Giants", "San Francisco", []string{"Lou Seal"}},
	{25, "Mariners", "Seattle", []string{"Mariner Moose", "The Salmon Run"}},
	{26, "Cardinals", "St. Louis", []string{"Fredbird", "Rally Squirrel"}},
	{27, "Rays", "Tampa Bay", []string{"DJ Kitty", "Raymond", "Stinger"}},
	{28, "Rangers", "Texas", []string{"Rangers Captain"}},
	{29, "Blue Jays", "Toronto", []string{"Ace", "Junior"}},
	{30, "Nationals", "Washington", []string{"Screech", "The Racing Presidents"}},
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
