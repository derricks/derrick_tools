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
	"github.com/spf13/cobra"
)

var nbaCmd = &cobra.Command{
	Use:   "nba-teams",
	Short: "Quiz US nba teams",
	Run:   quizNbaTeams,
}

type nbaTeam struct {
	index int    `crossquery:"all"`
	name  string `crossquery:"all"`
	// some areas, such as New York, have more than one team
	area     string `crossquery:"guess"`
	division string `crossquery:"guess"`
}

const (
	NBA_ATLANTIC  = "Atlantic"
	NBA_SOUTHEAST = "Southeast"
	NBA_PACIFIC   = "Pacific"
	NBA_CENTRAL   = "Central"
	NBA_NORTHWEST = "Northwest"
	NBA_SOUTHWEST = "Southwest"
)

var nbaTeams = []nbaTeam{
	{1, "Celtics", "Boston", NBA_ATLANTIC},
	{2, "Nets", "Brooklyn", NBA_ATLANTIC},
	{3, "Knicks", "New York", NBA_ATLANTIC},
	{4, "76ers", "Philadelphia", NBA_ATLANTIC},
	{5, "Raptors", "Toronto", NBA_ATLANTIC},
	{6, "Hawks", "Atlanta", NBA_SOUTHEAST},
	{7, "Hornets", "Charlotte", NBA_SOUTHEAST},
	{8, "Heat", "Miami", NBA_SOUTHEAST},
	{9, "Magic", "Orlando", NBA_SOUTHEAST},
	{10, "Wizards", "Washington", NBA_SOUTHEAST},
	{11, "Warriors", "Golden State", NBA_PACIFIC},
	{12, "Clippers", "LA", NBA_PACIFIC},
	{13, "Lakers", "Los Angeles", NBA_PACIFIC},
	{14, "Suns", "Phoenix", NBA_PACIFIC},
	{15, "Kings", "Sacramento", NBA_PACIFIC},
	{16, "Bulls", "Chicago", NBA_CENTRAL},
	{17, "Cavaliers", "Cleveland", NBA_CENTRAL},
	{18, "Pistons", "Detroit", NBA_CENTRAL},
	{19, "Pacers", "Indiana", NBA_CENTRAL},
	{20, "Bucks", "Milwaukee", NBA_CENTRAL},
	{21, "Nuggets", "Denver", NBA_NORTHWEST},
	{22, "Timberwolves", "Minnesota", NBA_NORTHWEST},
	{23, "Thunder", "Oklahoma City", NBA_NORTHWEST},
	{24, "Trailblazers", "Portland", NBA_NORTHWEST},
	{25, "Jazz", "Utah", NBA_NORTHWEST},
	{26, "Mavericks", "Dallas", NBA_SOUTHWEST},
	{27, "Rockets", "Houston", NBA_SOUTHWEST},
	{28, "Grizzlies", "Memphis", NBA_SOUTHWEST},
	{29, "Pelicans", "New Orleans", NBA_SOUTHWEST},
	{30, "Spurs", "San Antonio", NBA_SOUTHWEST},
}

func quizNbaTeams(cmd *cobra.Command, args []string) {
	promptAndCheckResponse(crossQueryNbaTeamInfo(nbaTeams))
}

func crossQueryNbaTeamInfo(teams []nbaTeam) promptAndResponse {
	foundTeam := randomItemFromSlice(teams)
	return constructCrossQuery("NBA team", foundTeam)
}

func init() {
	memoryquizCmd.AddCommand(nbaCmd)
}
