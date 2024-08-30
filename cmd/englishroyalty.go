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
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

var englishRoyaltyCmd = &cobra.Command{
	Use:   "english-royalty",
	Short: "Quiz English royalty",
	Run:   quizEnglishRoyalty,
}

type englishRoyal struct {
	order     int    `crossquery:"all"`
	name      string `crossquery:"guess"`
	sobriquet string
}

var royals = []englishRoyal{
	{1, "Egbert", ""},
	{2, "Aethelwulf", ""},
	{3, "Aethelbald", ""},
	{4, "Aethelbert", ""},
	{5, "Aethelred I", ""},
	{6, "Alfred", "the Great"},
	{7, "Edward", "the Elder"},
	{8, "Aethelstan", ""},
	{9, "Edmund I", ""},
	{10, "Eadred", ""},
	{11, "Eadwig", ""},
	{12, "Edgar", ""},
	{13, "Edward", "the Martyr"},
	{14, "Aethelred II", "the Unready"},
	{15, "Edmund II", "Ironside"},
	{16, "Canute", "the Dane"},
	{17, "Harold I", "Harefoot"},
	{18, "Harthacanute", ""},
	{19, "Edward", "the Confessor"},
	{20, "Harold II", ""},
	{21, "William I", "the Conqueror"},
	{22, "William II", "Rufus"},
	{23, "Henry I", ""},
	{24, "Stephen", ""},
	{25, "Henry II", ""},
	{26, "Richard I", "Lionheart"},
	{27, "John", ""},
	{28, "Henry III", ""},
	{29, "Edward I", ""},
	{30, "Edward II", ""},
	{31, "Edward III", ""},
	{32, "Richard II", ""},
	{33, "Henry IV", ""},
	{34, "Henry V", ""},
	{35, "Henry VI", ""},
	{36, "Edward IV", ""},
	{37, "Edward V", ""},
	{38, "Richard III", ""},
	{39, "Henry VII", ""},
	{40, "Henry VIII", ""},
	{41, "Edward VI", ""},
	{42, "Mary I", "Bloody Mary"},
	{43, "Elizabeth I", "Virgin Queen"},
	{44, "James I and VI", ""},
	{45, "Charles I", ""},
	{46, "Oliver Cromwell", ""},
	{47, "Richard Cromwell", ""},
	{48, "Charles II", ""},
	{49, "James II and VII", ""},
	{50, "William 3 and Mary II", ""},
	{51, "Anne", ""},
	{52, "George I", ""},
	{53, "George II", ""},
	{54, "George III", ""},
	{55, "George IV", ""},
	{56, "William IV", ""},
	{57, "Victoria", ""},
	{58, "Edward VII", ""},
	{59, "George V", ""},
	{60, "Edward VIII", ""},
	{61, "George VI", ""},
	{62, "Elizabeth II", ""},
	{63, "Charles III", ""},
}

type englishRoyalQuestion func([]englishRoyal) promptAndResponse

func quizEnglishRoyalty(cmd *cobra.Command, args []string) {

	var promptFuncs = []englishRoyalQuestion{
		crossQueryEnglishRoyal,
		crossQueryEnglishRoyal,
		crossQueryEnglishRoyal,
		crossQueryEnglishRoyal,
		crossQueryEnglishRoyal,
		crossQueryEnglishRoyal,
		quizRoyalBySobriquet,
		quizRoyalBeforeAnother,
		quizRoyalAfterAnother,
	}

	function := randomItemFromSlice(promptFuncs)
	promptAndCheckResponse(function(royals))
}

func crossQueryEnglishRoyal(royals []englishRoyal) promptAndResponse {
	royal := randomItemFromSlice(royals)
	return constructCrossQuery("English royal", royal)

}

func quizRoyalBySobriquet(royals []englishRoyal) promptAndResponse {
	// filter out royals without sobriquets
	sobriquetRoyals := make([]englishRoyal, 13)
	for _, royal := range royals {
		if royal.sobriquet != "" {
			sobriquetRoyals = append(sobriquetRoyals, royal)
		}
	}

	quizRoyal := randomItemFromSlice(sobriquetRoyals)
	return promptAndResponse{fmt.Sprintf("Which English royal had the sobriquet %s?", quizRoyal.sobriquet), quizRoyal.name}
}

func quizRoyalBeforeAnother(royals []englishRoyal) promptAndResponse {
	// exclude the first ruler, who doesn't have a predecessory	royalsWithBefore := royals[1:]
	index := rand.Intn(len(royals)-1) + 1
	return promptAndResponse{fmt.Sprintf("Who ruled England before %s?", royals[index].name), royals[index-1].name}
}

func quizRoyalAfterAnother(royals []englishRoyal) promptAndResponse {
	// exclude the last ruler, who doesn't have a successor (yet)
	index := rand.Intn(len(royals) - 1)
	return promptAndResponse{fmt.Sprintf("Who ruled England after %s?", royals[index].name), royals[index+1].name}
}

func init() {
	memoryquizCmd.AddCommand(englishRoyaltyCmd)
}
