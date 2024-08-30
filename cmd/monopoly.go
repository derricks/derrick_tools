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

// countriesCmd represents the countries command
var monopolyCmd = &cobra.Command{
	Use:   "monopoly",
	Short: "Memory quizzes about spaces in Monopoly",
	Run:   quizMonopoly,
}

type monopolySquare struct {
	position      int
	name          string
	color         string
	purchasePrice int
}

const (
	NO_COLOR     string = "none"
	PURPLE       string = "purple"
	LIGHT_BLUE   string = "light blue"
	LIGHT_PURPLE string = "light purple"
	ORANGE       string = "orange"
	RED          string = "red"
	YELLOW       string = "yellow"
	GREEN        string = "green"
	DARK_BLUE    string = "dark blue"
)

var monopolyBoard = []monopolySquare{
	{1, "Go", NO_COLOR, 0},
	{2, "Mediterranean Avenue", PURPLE, 60},
	{3, "Community Chest (1)", NO_COLOR, 0},
	{4, "Baltic Avenue", PURPLE, 60},
	{5, "Income Tax", NO_COLOR, 0},
	{6, "Reading Railroad", NO_COLOR, 200},
	{7, "Oriental Avenue", LIGHT_BLUE, 100},
	{8, "Chance (1)", NO_COLOR, 0},
	{9, "Vermont Avenue", LIGHT_BLUE, 100},
	{10, "Connecticut Avenue", LIGHT_BLUE, 120},
	{11, "Jail", NO_COLOR, 0},
	{12, "St. Charles Place", LIGHT_PURPLE, 140},
	{13, "Electric Company", NO_COLOR, 150},
	{14, "States Avenue", LIGHT_PURPLE, 140},
	{15, "Virginia Avenue", LIGHT_PURPLE, 160},
	{16, "Pennsylvania Railroad", NO_COLOR, 200},
	{17, "St. James Place", ORANGE, 180},
	{18, "Community Chest (2)", NO_COLOR, 0},
	{19, "Tennessee Avenue", ORANGE, 180},
	{20, "New York Avenue", ORANGE, 200},
	{21, "Free Parking", NO_COLOR, 0},
	{22, "Kentucky Avenue", RED, 220},
	{23, "Chance (2)", NO_COLOR, 0},
	{24, "Indiana Avenue", RED, 220},
	{25, "Illinois Avenue", RED, 240},
	{26, "B&O Railroad", NO_COLOR, 200},
	{27, "Atlantic Avenue", YELLOW, 260},
	{28, "Ventnor Avenue", YELLOW, 260},
	{29, "Water Works", NO_COLOR, 150},
	{30, "Marvin Gardens", YELLOW, 280},
	{31, "Go to Jail", NO_COLOR, 0},
	{32, "Pacific Avenue", GREEN, 300},
	{33, "North Carolina Avenue", GREEN, 300},
	{34, "Community Chest (3)", NO_COLOR, 0},
	{35, "Pennsylvania Avenue", GREEN, 320},
	{36, "Short Line Railroad", NO_COLOR, 200},
	{37, "Chance (3)", NO_COLOR, 0},
	{38, "Park Place", DARK_BLUE, 350},
	{39, "Luxury Tax", NO_COLOR, 0},
	{40, "Boardwalk", DARK_BLUE, 400},
}

// the usual pattern is to just pass a list to various quiz functions
// but because we only want to run certain quizzes (color and purchase price)
// on certain properties, we pass just the property to the quiz function,
type monopolyQuery func(monopolySquare) promptAndResponse

func quizMonopoly(cmd *cobra.Command, args []string) {

	quizFuncs := []monopolyQuery{quizMonopolyNameFromPosition, quizMonopolyPositionFromName}
	property := monopolyBoard[rand.Intn(len(monopolyBoard))]
	if property.color != NO_COLOR {
		quizFuncs = append(quizFuncs, quizMonopolyColorForProperty)
	}

	if property.purchasePrice != 0 {
		quizFuncs = append(quizFuncs, quizMonopolyPurchasePriceForProperty)
	}

	// special case for quizzing all the properties of a specific color
	// that can't work off a single property, so it's handled separately.
	// we do this by picking a random number between 0 and one more the length of the board.
	// if the random number = the len of the board, then use the special question
	if rand.Intn(len(monopolyBoard)+1) == len(monopolyBoard) {
		promptAndCheckResponse(quizMonopolyPropertiesByColor(monopolyBoard))
	} else {
		function := quizFuncs[rand.Intn(len(quizFuncs))]
		promptAndCheckResponse(function(property))
	}
}

func quizMonopolyPropertiesByColor(monopolyBoard []monopolySquare) promptAndResponse {
	colors := []string{PURPLE, LIGHT_BLUE, LIGHT_PURPLE, ORANGE, RED, YELLOW, GREEN, DARK_BLUE}
	randomColor := randomItemFromSlice(colors)

	properties := make([]string, 0)
	for _, property := range monopolyBoard {
		if property.color == randomColor {
			properties = append(properties, property.name)
		}
	}
	return promptAndResponse{fmt.Sprintf("Name the %s properties (in order)", randomColor), strings.Join(properties, ",")}
}

func quizMonopolyNameFromPosition(property monopolySquare) promptAndResponse {
	return promptAndResponse{fmt.Sprintf("What is the name of the square at position %d", property.position), property.name}
}

func quizMonopolyPositionFromName(property monopolySquare) promptAndResponse {
	return promptAndResponse{fmt.Sprintf("What position is %s at?", property.name), strconv.Itoa(property.position)}
}

func quizMonopolyColorForProperty(property monopolySquare) promptAndResponse {
	return promptAndResponse{fmt.Sprintf("What color is %s?", property.name), property.color}
}

func quizMonopolyPurchasePriceForProperty(property monopolySquare) promptAndResponse {
	return promptAndResponse{fmt.Sprintf("What is the purchase price for %s?", property.name), strconv.Itoa(property.purchasePrice)}
}

func init() {
	memoryquizCmd.AddCommand(monopolyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countriesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countriesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
