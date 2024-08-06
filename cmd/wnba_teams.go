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

var wnbaCmd = &cobra.Command{
	Use:   "wnba-teams",
	Short: "Quiz WNBA teams",
	Run:   quizWnbaTeams,
}

type wnbaTeam struct {
	index int    `crossquery:"all"`
	name  string `crossquery:"all"`
	area  string `crossquery:"all"`
}

var wnbaTeams = []wnbaTeam{
	{1, "Dream", "Atlanta"},
	{2, "Sky", "Chicago"},
	{3, "Sun", "Connecticut"},
	{4, "Wings", "Dallas"},
	{5, "Fever", "Indiana"},
	{6, "Aces", "Las Vegas"},
	{7, "Sparks", "Los Angeles"},
	{8, "Lynx", "Minnesota"},
	{9, "Liberty", "New York"},
	{10, "Mercury", "Phoenix"},
	{11, "Storm", "Seattle"},
	{12, "Mystics", "Washington"},
}

type wnbaQuestion func([]wnbaTeam) promptAndResponse

func quizWnbaTeams(cmd *cobra.Command, args []string) {

	var promptFuncs = []wnbaQuestion{
		crossQueryWnbaTeamInfo,
	}

	function := promptFuncs[rand.Intn(len(promptFuncs))]
	promptAndCheckResponse(function(wnbaTeams))
}

func crossQueryWnbaTeamInfo(teams []wnbaTeam) promptAndResponse {
	foundTeam := teams[rand.Intn(len(teams))]
	return constructCrossQuery("WNBA team", foundTeam)
}

func init() {
	memoryquizCmd.AddCommand(wnbaCmd)
}
