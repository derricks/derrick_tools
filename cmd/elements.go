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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// elementsCmd represents the elements command
var elementsCmd = &cobra.Command{
	Use:   "elements",
	Short: "Test recall of periodic table of elements information",
	Run:   quizElements,
}

type elementInfo struct {
	atomicNumber int    `crossquery:"all" crossqueryname:"atomic number"`
	name         string `crossquery:"all"`
	symbol       string `crossquery:"all"`
}

var elements = []elementInfo{
	elementInfo{1, "Hydrogen", "H"},
	elementInfo{2, "Helium", "He"},
	elementInfo{3, "Lithium", "Li"},
	elementInfo{4, "Beryllium", "Be"},
	elementInfo{5, "Boron", "B"},
	elementInfo{6, "Carbon", "C"},
	elementInfo{7, "Nitrogen", "N"},
	elementInfo{8, "Oxygen", "O"},
	elementInfo{9, "Fluorine", "F"},
	elementInfo{10, "Neon", "Ne"},
	elementInfo{11, "Sodium", "Na"},
	elementInfo{12, "Magnesium", "Mg"},
	elementInfo{13, "Aluminum", "Al"},
	elementInfo{14, "Silicon", "Si"},
	elementInfo{15, "Phosphorous", "P"},
	elementInfo{16, "Sulfur", "S"},
	elementInfo{17, "Chlorine", "Cl"},
	elementInfo{18, "Argon", "Ar"},
	elementInfo{19, "Potassium", "K"},
	elementInfo{20, "Calcium", "Ca"},
	elementInfo{21, "Scandium", "Sc"},
	elementInfo{22, "Titanium", "Ti"},
	elementInfo{23, "Vanadium", "V"},
	elementInfo{24, "Chromium", "Cr"},
	elementInfo{25, "Manganese", "Mn"},
	elementInfo{26, "Iron", "Fe"},
	elementInfo{27, "Cobalt", "Co"},
	elementInfo{28, "Nickel", "Ni"},
	elementInfo{29, "Copper", "Cu"},
	elementInfo{30, "Zinc", "Zn"},
	elementInfo{31, "Gallium", "Ga"},
	elementInfo{32, "Germanium", "Ge"},
	elementInfo{33, "Arsenic", "As"},
	elementInfo{34, "Selenium", "Se"},
	elementInfo{35, "Bromine", "Br"},
	elementInfo{36, "Krypton", "Kr"},
	elementInfo{37, "Rubidium", "Rb"},
	elementInfo{38, "Strontium", "Sr"},
	elementInfo{39, "Yttrium", "Y"},
	elementInfo{40, "Zirconium", "Zr"},
	elementInfo{41, "Niobium", "Nb"},
	elementInfo{42, "Molybdenum", "Mo"},
	elementInfo{43, "Technetium", "Tc"},
	elementInfo{44, "Ruthenium", "Ru"},
	elementInfo{45, "Rhodium", "Rh"},
	elementInfo{46, "Palladium", "Pd"},
	elementInfo{47, "Silver", "Ag"},
	elementInfo{48, "Cadmium", "Cd"},
	elementInfo{49, "Indium", "In"},
	elementInfo{50, "Tin", "Sn"},
	elementInfo{51, "Antimony", "Sb"},
	elementInfo{52, "Tellurium", "Te"},
	elementInfo{53, "Iodine", "I"},
	elementInfo{54, "Xenon", "Xe"},
	elementInfo{55, "Cesium", "Cs"},
	elementInfo{56, "Barium", "Ba"},
	elementInfo{57, "Lanthanum", "La"},
	elementInfo{58, "Cerium", "Ce"},
	elementInfo{59, "Praseodymium", "Pr"},
	elementInfo{60, "Neodymium", "Nd"},
	elementInfo{61, "Promethium", "Pm"},
	elementInfo{62, "Samarium", "Sm"},
	elementInfo{63, "Europium", "Eu"},
	elementInfo{64, "Gadolinium", "Gd"},
	elementInfo{65, "Terbium", "Tb"},
	elementInfo{66, "Dysprosium", "Dy"},
	elementInfo{67, "Holmium", "Ho"},
	elementInfo{68, "Erbium", "Er"},
	elementInfo{69, "Thulium", "Tm"},
	elementInfo{70, "Ytterbium", "Yb"},
	elementInfo{71, "Lutetium", "Lu"},
	elementInfo{72, "Hafnium", "Hf"},
	elementInfo{73, "Tantalum", "Ta"},
	elementInfo{74, "Tungsten", "W"},
	elementInfo{75, "Rhenium", "Re"},
	elementInfo{76, "Osmium", "Os"},
	elementInfo{77, "Iridium", "Ir"},
	elementInfo{78, "Platinum", "Pt"},
	elementInfo{79, "Gold", "Au"},
	elementInfo{80, "Mercury", "Hg"},
	elementInfo{81, "Thalium", "Tl"},
	elementInfo{82, "Lead", "Pb"},
	elementInfo{83, "Bismuth", "Bi"},
	elementInfo{84, "Polonium", "Po"},
	elementInfo{85, "Astatine", "At"},
	elementInfo{86, "Radon", "Rn"},
	elementInfo{87, "Francium", "Fr"},
	elementInfo{88, "Radium", "Ra"},
	elementInfo{89, "Actinium", "Ac"},
	elementInfo{90, "Thorium", "Th"},
	elementInfo{91, "Protactinium", "Pa"},
	elementInfo{92, "Uranium", "U"},
	elementInfo{93, "Neptunium", "Np"},
	elementInfo{94, "Plutonium", "Pu"},
	elementInfo{95, "Americium", "Am"},
	elementInfo{96, "Curium", "Cm"},
	elementInfo{97, "Berkelium", "Bk"},
	elementInfo{98, "Californium", "Cf"},
	elementInfo{99, "Einsteinium", "Es"},
	elementInfo{100, "Fermium", "Fm"},
	elementInfo{101, "Mendelevium", "Md"},
	elementInfo{102, "Nobelium", "No"},
	elementInfo{103, "Lawrencium", "Lr"},
	elementInfo{104, "Rutherfordium", "Rf"},
	elementInfo{105, "Dubnium", "Db"},
	elementInfo{106, "Seaborgium", "Sg"},
	elementInfo{107, "Bohrium", "Bh"},
	elementInfo{108, "Hassium", "Hs"},
	elementInfo{109, "Meitnerium", "Mt"},
	elementInfo{110, "Darmstadtium", "Ds"},
	elementInfo{111, "Roentgenium", "Rg"},
	elementInfo{112, "Copernicium", "Cn"},
	elementInfo{113, "Nihonium", "Nh"},
	elementInfo{114, "Flerovium", "Fl"},
	elementInfo{115, "Moscovium", "Mc"},
	elementInfo{116, "Livermorium", "Lv"},
	elementInfo{117, "Tennessine", "Ts"},
	elementInfo{118, "Oganesson", "Og"},
}

func quizElements(cmd *cobra.Command, args []string) {
	element := elements[rand.Intn(len(elements))]
	var prompt promptAndResponse
	if rand.Intn(10) < 8 {
		prompt = constructCrossQuery("atomic element", element)
	} else {
		prompt = quizElementsThatStartWithLetter(cmd, args)
	}
	promptAndCheckResponse(prompt)
}

func quizElementsThatStartWithLetter(cmd *cobra.Command, args []string) promptAndResponse {
	letterAscii := rune(65 + rand.Intn(26))
	letter := string(letterAscii)
	count := 0
	for _, element := range elements {
		if strings.HasPrefix(element.name, letter) {
			count++
		}
	}
	return promptAndResponse{fmt.Sprintf("How many elements start with %s?", letter), strconv.Itoa(count)}
}

func init() {
	memoryquizCmd.AddCommand(elementsCmd)
}
