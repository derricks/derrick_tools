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

var arrondisementsCmd = &cobra.Command{
	Use:   "arrondisements",
	Short: "Quiz command of Parisian arrondisements",
	Run:   quizArrondisements,
}

var arrondisements = []string{
	"Louvre",
	"Bourse",
	"Temple",
	"Hôtel-de-Ville",
	"Panthéon",
	"Luxembourg",
	"Palais-Bourbon",
	"Élysée",
	"Opéra",
	"Entrepôt",
	"Popincourt",
	"Reuilly",
	"Gobelins",
	"Observatoire",
	"Vaugirard",
	"Passy",
	"Batignolles-Montant",
	"Butte-Montmartre",
	"Buttes-Chaumont",
	"Ménilmontant",
}

type quizArrondisementFunc func([]string) promptAndResponse

func quizArrondisements(cmd *cobra.Command, args []string) {
	funcs := []quizArrondisementFunc{
		quizArrondisementFromIndex,
		quizIndexOfArrondisement,
	}

	function := funcs[rand.Intn(len(funcs))]
	promptAndCheckResponse(function(arrondisements))
}

func quizArrondisementFromIndex(arrondisements []string) promptAndResponse {
	return quizStringAtIndexInList("arrondisement", arrondisements)
}

func quizIndexOfArrondisement(arrondisements []string) promptAndResponse {
	return quizIndexOfStringInList(arrondisements)
}

func init() {
	memoryquizCmd.AddCommand(arrondisementsCmd)
}
