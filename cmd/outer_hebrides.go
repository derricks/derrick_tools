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

var outerHebridesCmd = &cobra.Command{
	Use:   "outer-hebrides",
	Short: "Quiz Outer Hebrides islands",
	Run:   quizOuterHebrides,
}

// use a generic struct since this will apply to
// Inner Hebrides, Orkney, and other groups of islands
type island struct {
	index int    `crossquery:"all"`
	name  string `crossquery:"all"`
}

var outerHebrides = []island{
	{1, "Lewis and Harris"},
	{2, "Great Bernera"},
	{3, "Mealasta Island"},
	{4, "Scarp"},
	{5, "Soay Mor"},
	{6, "Taransay"},
	{7, "Ensay"},
	{8, "Shillay"},
	{9, "Pabbay"},
	{10, "Boreray"},
	{11, "Vallay"},
	{12, "North Uist"},
	{13, "St. Kilda"},
	{14, "Kirkibost Island"},
	{15, "Monach Islands"},
	{16, "Baleshare"},
	{17, "Benbecula"},
	{18, "South Uist"},
	{19, "Fuday"},
	{20, "Barra"},
	{21, "Vatersay"},
	{22, "Mingulay"},
	{23, "Berneray"},
	{24, "Rosinish"},
	{25, "Sandray"},
	{26, "Erisay"},
	{27, "Wiay"},
	{28, "Grimsay"},
	{29, "Ronay"},
	{30, "Hermetray"},
	{31, "Berneray"}, // yes, again
	{32, "Killegray"},
	{33, "Scalpay"},
	{34, "Seaforth Island"},
	{35, "Shiant Islands"},
}

func quizOuterHebrides(cmd *cobra.Command, args []string) {

	chosenIsland := randomItemFromSlice(outerHebrides)
	promptAndCheckResponse(constructCrossQuery("Outer Hebrides", chosenIsland))
}

func init() {
	memoryquizCmd.AddCommand(outerHebridesCmd)
}
