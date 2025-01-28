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
	{7, "Lower Ob"},
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
	{25, "Syr Darya"},
	{26, "Salween"},
	{27, "Saint Lawrence"},
	{28, "Rio Grande"},
	{29, "Lower Tunguska"},
	{30, "Colorado"},
	{31, "Danube"},
	{32, "Irrawaddy"},
	{33, "Zambezi"},
	{34, "Vilyuy"},
	{35, "Padma"},
	{36, "Amu Darya"},
	{37, "Japura"},
	{38, "Nelson"},
	{39, "Paraguay"},
	{40, "Kolmya"},
	{41, "Pilcomayo"},
	{42, "Upper Ob"},
	{43, "Ishim"},
	{44, "Orange"},
	{45, "Ural"},
	{46, "Jurua"},
	{47, "Arkansas"},
	{48, "Songhua"},
	{49, "Olenyok"},
	{50, "Dnieper"},
	{51, "Aldan"},
	{52, "Ubangi"},
	{53, "Negro"},
	{54, "Columbia"},
	{55, "Tapajos"},
	{56, "Pearl"},
	{57, "Red"},
	{58, "Kasai"},
	{59, "Ohio"},
	{60, "Orinoco"},
	{61, "Tarim"},
	{62, "Xingu"},
	{63, "Jubba"},
	{64, "Brazos"},
	{65, "Northern Salado"},
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
