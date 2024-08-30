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
	"strings"

	"github.com/spf13/cobra"
)

// countriesCmd represents the countries command
var grandCrusCmd = &cobra.Command{
	Use:   "grand-crus",
	Short: "Memory quizzesa about Burgundy Grand Crus",
	Run:   quizGrandCrus,
}

const (
	GEVREY  string = "Gevrey-Chambertin"
	MOREY   string = "Morey-Saint-Denis"
	MUSIGNY string = "Chambolle-Musigny"
	FLAGEY  string = "Flagey-Echézaux"
	VOSNE   string = "Vosne-Romanée"
	PERNAND string = "Pernand-Vergelesses"
	LADOIX  string = "Ladoix-Serrigny"
	ALOXE   string = "Aloxe-Corton"
	PULIGNY string = "Puligny-Montrachet"
)

type grandCru struct {
	order   int    `crossquery:"all"`
	name    string `crossquery:"guess"`
	village string `crossquery:"guess"`
}

var grandCrus = []grandCru{
	{1, "Chablis Grand Cru", "Chablis"},
	{2, "Chambertin", GEVREY},
	{3, "Chambertin-Clos de Bèze", GEVREY},
	{4, "Chapelle-Chambertin", GEVREY},
	{5, "Charmes-Chambertin", GEVREY},
	{6, "Griotte-Chambertin", GEVREY},
	{7, "Latricières-Chambertin", GEVREY},
	{8, "Mazis-Chambertin", GEVREY},
	{9, "Mazoyères-Chambertin", GEVREY},
	{10, "Ruchottes-Chambertin", GEVREY},
	{11, "Bonnes-Mares", MOREY},
	{12, "Clos de la Roche", MOREY},
	{13, "Clos des Lambrays", MOREY},
	{14, "Clos de Tart", MOREY},
	{15, "Clos Saint-Denis", MOREY},
	{16, "Bonnes-Mares", MUSIGNY},
	{17, "Musigny", MUSIGNY},
	{18, "Clos de Vougeot", "Vougeot"},
	{19, "Échezaux", FLAGEY},
	{20, "Grands-Échezaux", FLAGEY},
	{21, "La Grand Rue", VOSNE},
	{22, "La Romanée", VOSNE},
	{23, "La Tache", VOSNE},
	{24, "Richebourg", VOSNE},
	{25, "Romanée-Conti", VOSNE},
	{26, "Romanée-Saint-Vivant", VOSNE},
	{27, "Corton", PERNAND},
	{28, "Charlemagne", PERNAND},
	{29, "Corton", LADOIX},
	{30, "Corton-Charlemagne", LADOIX},
	{31, "Corton", ALOXE},
	{32, "Charlemagne", ALOXE},
	{33, "Bâtard-Montrachet", PULIGNY},
	{34, "Bienvenues-Bâtard-Montrachet", PULIGNY},
	{35, "Chevalier-Bâtard-Montrachet", PULIGNY},
	{36, "Montrachet", PULIGNY},
	{37, "Criots-Bâtard-Montrachet", "Chassagne-Montrachet"},
}

func quizGrandCrus(cmd *cobra.Command, args []string) {
	randNumber := rand.Intn(10)
	if randNumber == 0 {
		promptAndCheckResponse(quizVineyardsForVillage(grandCrus))
	} else {
		grandCru := grandCrus[rand.Intn(len(grandCrus))]
		promptAndCheckResponse(constructCrossQuery("Grand Cru", grandCru))
	}
}

func quizVineyardsForVillage(crus []grandCru) promptAndResponse {
	randomCru := randomItemFromSlice(crus)
	village := randomCru.village
	vineyardNames := []string{}
	for _, cru := range crus {
		if cru.village == village {
			vineyardNames = append(vineyardNames, cru.name)
		}
	}
	return promptAndResponse{fmt.Sprintf("Name all the vineyards near %s", village), strings.Join(vineyardNames, ",")}
}

func init() {
	memoryquizCmd.AddCommand(grandCrusCmd)
}
