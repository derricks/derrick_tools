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
	grandCru{1, "Chablis Grand Cru", "Chablis"},
	grandCru{2, "Chambertin", GEVREY},
	grandCru{3, "Chambertin-Clos de Bèze", GEVREY},
	grandCru{4, "Chapelle-Chambertin", GEVREY},
	grandCru{5, "Charmes-Chambertin", GEVREY},
	grandCru{6, "Griotte-Chambertin", GEVREY},
	grandCru{7, "Latricières-Chambertin", GEVREY},
	grandCru{8, "Mazis-Chambertin", GEVREY},
	grandCru{9, "Mazoyères-Chambertin", GEVREY},
	grandCru{10, "Ruchottes-Chambertin", GEVREY},
	grandCru{11, "Bonnes-Mares", MOREY},
	grandCru{12, "Clos de la Roche", MOREY},
	grandCru{13, "Clos des Lambrays", MOREY},
	grandCru{14, "Clos de Tart", MOREY},
	grandCru{15, "Clos Saint-Denis", MOREY},
	grandCru{16, "Bonnes-Mares", MUSIGNY},
	grandCru{17, "Musigny", MUSIGNY},
	grandCru{18, "Clos de Vougeot", "Vougeot"},
	grandCru{19, "Échezaux", FLAGEY},
	grandCru{20, "Grands-Échezaux", FLAGEY},
	grandCru{21, "La Grand Rue", VOSNE},
	grandCru{22, "La Romanée", VOSNE},
	grandCru{23, "La Tache", VOSNE},
	grandCru{24, "Richebourg", VOSNE},
	grandCru{25, "Romanée-Conti", VOSNE},
	grandCru{26, "Romanée-Saint-Vivant", VOSNE},
	grandCru{27, "Corton", PERNAND},
	grandCru{28, "Charlemagne", PERNAND},
	grandCru{29, "Corton", LADOIX},
	grandCru{30, "Corton-Charlemagne", LADOIX},
	grandCru{31, "Corton", ALOXE},
	grandCru{32, "Charlemagne", ALOXE},
	grandCru{33, "Bâtard-Montrachet", PULIGNY},
	grandCru{34, "Bienvenues-Bâtard-Montrachet", PULIGNY},
	grandCru{35, "Chevalier-Bâtard-Montrachet", PULIGNY},
	grandCru{36, "Montrachet", PULIGNY},
	grandCru{37, "Criots-Bâtard-Montrachet", "Chassgne-Montrachet"},
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
	randomCru := crus[rand.Intn(len(crus))]
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
