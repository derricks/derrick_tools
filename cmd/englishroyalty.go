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
	englishRoyal{1, "Egbert", ""},
	englishRoyal{2, "Aethelwulf", ""},
	englishRoyal{3, "Aethelbald", ""},
	englishRoyal{4, "Aethelbert", ""},
	englishRoyal{5, "Aethelred I", ""},
	englishRoyal{6, "Alfred", "the Great"},
	englishRoyal{7, "Edward", "the Elder"},
	englishRoyal{8, "Aethelstan", ""},
	englishRoyal{9, "Edmund I", ""},
	englishRoyal{10, "Eadred", ""},
	englishRoyal{11, "Eadwig", ""},
	englishRoyal{12, "Edgar", ""},
	englishRoyal{13, "Edward", "the Martyr"},
	englishRoyal{14, "Aethelred II", "the Unready"},
	englishRoyal{15, "Edmund II", "Ironside"},
	englishRoyal{16, "Canute", "the Dane"},
	englishRoyal{17, "Harold I", "Harefoot"},
	englishRoyal{18, "Harthacanute", ""},
	englishRoyal{19, "Edward", "the Confessor"},
	englishRoyal{20, "Harold II", ""},
	englishRoyal{21, "William I", "the Conqueror"},
	englishRoyal{22, "William II", "Rufus"},
	englishRoyal{23, "Henry I", ""},
	englishRoyal{24, "Stephen", ""},
	englishRoyal{25, "Henry II", ""},
	englishRoyal{26, "Richard I", "Lionheart"},
	englishRoyal{27, "John", ""},
	englishRoyal{28, "Henry III", ""},
	englishRoyal{29, "Edward I", ""},
	englishRoyal{30, "Edward II", ""},
	englishRoyal{31, "Edward III", ""},
	englishRoyal{32, "Richard II", ""},
	englishRoyal{33, "Henry IV", ""},
	englishRoyal{34, "Henry V", ""},
	englishRoyal{35, "Henry VI", ""},
	englishRoyal{36, "Edward IV", ""},
	englishRoyal{37, "Edward V", ""},
	englishRoyal{38, "Richard III", ""},
	englishRoyal{39, "Henry VII", ""},
	englishRoyal{40, "Henry VIII", ""},
	englishRoyal{41, "Edward VI", ""},
	englishRoyal{42, "Mary I", "Bloody Mary"},
	englishRoyal{43, "Elizabeth I", "Virgin Queen"},
	englishRoyal{44, "James I and VI", ""},
	englishRoyal{45, "Charles I", ""},
	englishRoyal{46, "Oliver Cromwell", ""},
	englishRoyal{47, "Richard Cromwell", ""},
	englishRoyal{48, "Charles II", ""},
	englishRoyal{49, "James II and VII", ""},
	englishRoyal{50, "William 3 and Mary II", ""},
	englishRoyal{51, "Anne", ""},
	englishRoyal{52, "George I", ""},
	englishRoyal{53, "George II", ""},
	englishRoyal{54, "George III", ""},
	englishRoyal{55, "George IV", ""},
	englishRoyal{56, "William IV", ""},
	englishRoyal{57, "Victoria", ""},
	englishRoyal{58, "Edward VII", ""},
	englishRoyal{59, "George V", ""},
	englishRoyal{60, "Edward VIII", ""},
	englishRoyal{61, "George VI", ""},
	englishRoyal{62, "Elizabeth II", ""},
	englishRoyal{63, "Charles III", ""},
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
