/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// elementsCmd represents the elements command
var elementsCmd = &cobra.Command{
	Use:   "elements",
	Short: "Test recall of periodic table of elements information",
	Run:   quizElements,
}

type elementInfo struct {
	atomicNumber int
	name         string
}

var elements = []elementInfo{
	elementInfo{1, "Hydrogen"},
	elementInfo{2, "Helium"},
	elementInfo{3, "Lithium"},
	elementInfo{4, "Beryllium"},
	elementInfo{5, "Boron"},
	elementInfo{6, "Carbon"},
	elementInfo{7, "Nitrogen"},
	elementInfo{8, "Oxygen"},
	elementInfo{9, "Fluorine"},
	elementInfo{10, "Neon"},
	elementInfo{11, "Sodium"},
	elementInfo{12, "Magnesium"},
	elementInfo{13, "Aluminum"},
	elementInfo{14, "Silicon"},
	elementInfo{15, "Phosphorous"},
	elementInfo{16, "Sulfur"},
	elementInfo{17, "Chlorine"},
	elementInfo{18, "Argon"},
	elementInfo{19, "Potassium"},
	elementInfo{20, "Calcium"},
	elementInfo{21, "Scandium"},
	elementInfo{22, "Titanium"},
	elementInfo{23, "Vanadium"},
	elementInfo{24, "Chromium"},
	elementInfo{25, "Manganese"},
	elementInfo{26, "Iron"},
	elementInfo{27, "Cobalt"},
	elementInfo{28, "Nickel"},
	elementInfo{29, "Copper"},
	elementInfo{30, "Zinc"},
	elementInfo{31, "Gallium"},
	elementInfo{32, "Germanium"},
	elementInfo{33, "Arsenic"},
	elementInfo{34, "Selenium"},
	elementInfo{35, "Bromine"},
	elementInfo{36, "Krypton"},
	elementInfo{37, "Rubidium"},
	elementInfo{38, "Strontium"},
	elementInfo{39, "Yttrium"},
	elementInfo{40, "Zirconium"},
	elementInfo{41, "Niobium"},
	elementInfo{42, "Molybdenum"},
	elementInfo{43, "Technetium"},
	elementInfo{44, "Ruthenium"},
	elementInfo{45, "Rhodium"},
	elementInfo{46, "Palladium"},
	elementInfo{47, "Silver"},
	elementInfo{48, "Cadmium"},
	elementInfo{49, "Indium"},
	elementInfo{50, "Tin"},
	elementInfo{51, "Antimony"},
	elementInfo{52, "Tellurium"},
	elementInfo{53, "Iodine"},
	elementInfo{54, "Xenon"},
	elementInfo{55, "Cesium"},
	elementInfo{56, "Barium"},
	elementInfo{57, "Lanthanum"},
	elementInfo{58, "Cerium"},
	elementInfo{59, "Praseodymium"},
	elementInfo{60, "Neodymium"},
	elementInfo{61, "Promethium"},
	elementInfo{62, "Samarium"},
	elementInfo{63, "Europium"},
	elementInfo{64, "Gadolinium"},
	elementInfo{65, "Terbium"},
	elementInfo{66, "Dysprosium"},
	elementInfo{67, "Holium"},
	elementInfo{68, "Erbium"},
	elementInfo{69, "Thulium"},
	elementInfo{70, "Ytterbium"},
	elementInfo{71, "Lutetium"},
	elementInfo{72, "Hafnium"},
	elementInfo{73, "Tantalum"},
	elementInfo{74, "Tungsten"},
	elementInfo{75, "Rhenium"},
	elementInfo{76, "Osmium"},
	elementInfo{77, "Iridium"},
	elementInfo{78, "Platinum"},
	elementInfo{79, "Gold"},
	elementInfo{80, "Mercury"},
	elementInfo{81, "Thalium"},
	elementInfo{82, "Lead"},
	elementInfo{83, "Bismuth"},
	elementInfo{84, "Polonium"},
	elementInfo{85, "Astatine"},
	elementInfo{86, "Radon"},
	elementInfo{87, "Francium"},
	elementInfo{88, "Radium"},
	elementInfo{89, "Actinium"},
	elementInfo{90, "Thorium"},
	elementInfo{91, "Protactinium"},
	elementInfo{92, "Uranium"},
	elementInfo{93, "Neptunium"},
	elementInfo{94, "Plutonium"},
	elementInfo{95, "Americium"},
	elementInfo{96, "Curium"},
	elementInfo{97, "Berkelium"},
	elementInfo{98, "Californium"},
	elementInfo{99, "Einsteinium"},
	elementInfo{100, "Fermium"},
	elementInfo{101, "Mendelevium"},
	elementInfo{102, "Nobelium"},
	elementInfo{103, "Lawrencium"},
	elementInfo{104, "Rutherfordium"},
	elementInfo{105, "Dubnium"},
	elementInfo{106, "Seaborgium"},
	elementInfo{107, "Bohrium"},
	elementInfo{108, "Hassium"},
	elementInfo{109, "Meitnerium"},
	elementInfo{110, "Darmstadtium"},
	elementInfo{111, "Roentgenium"},
	elementInfo{112, "Copernicium"},
	elementInfo{113, "Nihonium"},
	elementInfo{114, "Flerovium"},
	elementInfo{115, "Moscovium"},
	elementInfo{116, "Livermorium"},
	elementInfo{117, "Tennessine"},
	elementInfo{118, "Oganesson"},
}

func quizElements(cmd *cobra.Command, args []string) {
	funcs := []func([]elementInfo) promptAndResponse{
		quizElementNameFromNumber,
		quizElementNumberFromName,
	}

	quizFunc := funcs[rand.Intn(len(funcs))]
	promptAndCheckResponse(quizFunc(elements))
}

func quizElementNameFromNumber(elements []elementInfo) promptAndResponse {
	element := elements[rand.Intn(len(elements))]
	return promptAndResponse{fmt.Sprintf("What is the name of element %d?", element.atomicNumber), element.name}
}

func quizElementNumberFromName(elements []elementInfo) promptAndResponse {
	element := elements[rand.Intn(len(elements))]
	return promptAndResponse{fmt.Sprintf("What is the atomic number of %s?", element.name), strconv.Itoa(element.atomicNumber)}
}

func init() {
	memoryquizCmd.AddCommand(elementsCmd)
}
