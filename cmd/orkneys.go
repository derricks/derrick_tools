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

var orkneysCmd = &cobra.Command{
	Use:   "orkneyss",
	Short: "Quiz Orkney islands",
	Run:   quizOrkneys,
}

var orkneys = []island{
	{1, "Papa Westray"},
	{2, "Westray"},
	{3, "Rousay"},
	{4, "Wyre"},
	{5, "Egilsay"},
	{6, "Gairsay"},
	{7, "Shapinsay"},
	{8, "Auskerry"},
	{9, "Stronsay"},
	{10, "Eday"},
	{11, "Sanday"},
	{12, "North Ronaldsay"},
	{13, "Mainland"},
	{14, "Copinsay"},
	{15, "Burray"},
	{16, "Hunda"},
	{17, "South Ronaldsay"},
	{18, "Swona"},
	{19, "Muckle Skerry"},
	{20, "Pentland Skerry"},
	{21, "Stroma"},
	{22, "South Walls"},
	{23, "Flotta"},
	{24, "Hoy"},
	{25, "Fara"},
	{26, "Cava"},
	{27, "Graemsay"},
}

func quizOrkneys(cmd *cobra.Command, args []string) {

	chosenIsland := randomItemFromSlice(orkneys)
	promptAndCheckResponse(constructCrossQuery("Orkneys", chosenIsland))
}

func init() {
	memoryquizCmd.AddCommand(orkneysCmd)
}
