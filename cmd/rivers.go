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

var riversCmd = &cobra.Command{
	Use:   "rivers",
	Short: "Quiz rivers over 1000km",
	Run:   quizRivers,
}

type river struct {
	order int    `crossquery:"all"`
	name  string `crossquery:"all"`
}

var rivers = []river{
	{1, "Nile"},
	{2, "Amazon"},
	{3, "Yangtze"},
	{4, "Mississippi"},
	{5, "Yenisey"},
	{6, "Yellow"},
	{7, "Ob"},
	{8, "Rio de la Plata"},
	{9, "Congo"},
	{10, "Amur"},
	{11, "Lena"},
	{12, "Mekong"},
	{13, "Mackenzie"},
	{14, "Niger"},
	{15, "Brahmaputra"},
	{16, "Murray"},
	{17, "Tocantins"},
	{18, "Volga"},
	{19, "Shatt al-Arab"},
	{20, "Madeira"},
	{21, "Purus"},
	{22, "Yukon"},
	{23, "Indus"},
	{24, "Sao Francisco"},
}

type riverQuestion func([]river) promptAndResponse

func quizRivers(cmd *cobra.Command, args []string) {

	var promptFuncs = []riverQuestion{
		crossQueryRiverInfo,
	}

	function := randomItemFromSlice(promptFuncs)
	promptAndCheckResponse(function(rivers))
}

func crossQueryRiverInfo(rivers []river) promptAndResponse {
	foundRiver := randomItemFromSlice(rivers)
	return constructCrossQuery("river", foundRiver)
}

func init() {
	memoryquizCmd.AddCommand(riversCmd)
}
