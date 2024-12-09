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

var romanNamesCmd = &cobra.Command{
	Use:   "roman-names",
	Short: "Quiz Roman names for British places",
	Run:   quizRomanNames,
}

type romanName struct {
	modernName string `crossquery:"all" crossqueryname:"modern name"`
	romanName  string `crossquery:"all" crossqueryname:"Roman name"`
}

var romanNames = []romanName{
	{"Alcester", "Alauna"},
	{"Aldborough", "Isurium Brigantium"},
	{"Bath", "Aquae Sulis"},
	{"Brough", "Petuaria"},
	{"Caernarfon", "Segontium"},
	{"Caistor", "Caistor"},
	{"Cambridge", "Duroliponte"},
	{"Canterbury", "Durovernum Cantiacorum"},
	{"Cardiff", "Tamium"},
	{"Carlisle", "Luguvalium"},
	{"Carmarthen", "Moridunum"},
	{"Catterick", "Cataractonium"},
	{"Chepstow", "Venta Silurum"},
	{"Chester", "Deva Victrix"},
	{"Chichester", "Noviomagus Reginorum"},
	{"Cirencester", "Corinium Dobunnorum"},
	{"Colchester", "Camulodunum"},
	{"Conwy", "Canofium Deceangorum"},
	{"Doncaster", "Danum"},
	{"Dorchester", "Durnovaria"},
	{"Durham", "Vinovia"},
	{"Exeter", "Isca Dumnoniorum"},
	{"Gloucester", "Glevum"},
	{"Hereford", "Magnae Dobunnorum"},
	{"Ilchester", "Lindinis"},
	{"Leicester", "Ratae Corieltauvorum"},
	{"Lichfield", "Letocetum"},
}

func quizRomanNames(cmd *cobra.Command, args []string) {
	chosenPlace := randomItemFromSlice(romanNames)
	promptAndCheckResponse(constructCrossQuery("place", chosenPlace))
}

func init() {
	memoryquizCmd.AddCommand(romanNamesCmd)
}
