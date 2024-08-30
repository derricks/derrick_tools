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

var wineBottlesCmd = &cobra.Command{
	Use:   "bottles",
	Short: "Quiz wine bottle sizes",
	Run:   quizWineBottles,
}

type wineBottle struct {
	sizeRank     int    `crossquery:"all" crossqueryname:"order"`
	name         string `crossquery:"all"`
	bordeauxName string `crossquery:"all" crossqueryname:"Bordeaux name"`
	sizeInMl     int    `crossquery:"all" crossqueryname:"size in ml"`
}

var bottles = []wineBottle{
	{1, "Split/Piccolo", "", 187},
	{2, "Half/Demi", "", 375},
	{3, "Half-Liter/Jennie", "", 500},
	{4, "Standard", "", 750},
	{5, "Liter", "", 1000},
	{6, "Magnum", "", 1500},
	{7, "Double Magnum/Jeroboam", "", 3000},
	{8, "Rehoboam", "Jeroboam", 4500},
	{9, "Methuselah", "Imperial", 6000},
	{10, "Salmanazar", "", 9000},
	{11, "Balthazar", "", 12000},
	{12, "Nebuchadnezzar", "", 15000},
	{13, "Melchior", "", 18000},
	{14, "Solomon", "", 20000},
	{15, "Sovereign", "", 26000},
	{16, "Primat/Goliath", "", 27000},
	{17, "Melchizedek/Midas", "", 30000},
}

func quizWineBottles(cmd *cobra.Command, args []string) {

	bottle := bottles[rand.Intn(len(bottles))]
	promptAndCheckResponse(constructCrossQuery("wine bottle", bottle))
}

func init() {
	memoryquizCmd.AddCommand(wineBottlesCmd)
}
