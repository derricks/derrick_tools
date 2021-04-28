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
	nobleGas     bool
}

var elements = []elementInfo{
	elementInfo{1, "Hydrogen", false},
	elementInfo{2, "Helium", true},
	elementInfo{3, "Lithium", false},
	elementInfo{4, "Beryllium", false},
	elementInfo{5, "Boron", false},
	elementInfo{6, "Carbon", false},
	elementInfo{7, "Nitrogen", false},
	elementInfo{8, "Oxygen", false},
	elementInfo{9, "Fluorine", false},
	elementInfo{10, "Neon", true},
	elementInfo{11, "Sodium", false},
	elementInfo{12, "Magnesium", false},
	elementInfo{13, "Aluminum", false},
	elementInfo{14, "Silcon", false},
	elementInfo{15, "Phosphorous", false},
	elementInfo{16, "Sulfur", false},
	elementInfo{17, "Chlorine", false},
	elementInfo{18, "Argon", true},
	elementInfo{19, "Potassium", false},
	elementInfo{20, "Calcium", false},
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