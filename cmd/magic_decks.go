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
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
)

var magicCmd = &cobra.Command{
	Use:   "magic",
	Short: "Quiz command of Magic the Gathering information",
	Run:   quizMagic,
}

type color string

const (
	MTG_WHITE color = "white"
	MTG_BLUE  color = "blue"
	MTG_BLACK color = "black"
	MTG_RED   color = "red"
	MTG_GREEN color = "green"
)

type magicDeck struct {
	name   string
	colors []color
}

var magicDecks = []magicDeck{
	{"Azorius", []color{MTG_WHITE, MTG_BLUE}},
	{"Orzhov Syndicate", []color{MTG_WHITE, MTG_BLACK}},
	{"Boros Legion", []color{MTG_WHITE, MTG_RED}},
	{"Selesnya Conclave", []color{MTG_WHITE, MTG_GREEN}},
	{"Dimir", []color{MTG_BLUE, MTG_BLACK}},
	{"Izzet League", []color{MTG_BLUE, MTG_RED}},
	{"Simic Combine", []color{MTG_BLUE, MTG_GREEN}},
	{"Cult of Rakdos", []color{MTG_BLACK, MTG_RED}},
	{"Golgari Swarm", []color{MTG_BLACK, MTG_GREEN}},
	{"Gruul Clans", []color{MTG_RED, MTG_GREEN}},
	{"Bant", []color{MTG_WHITE, MTG_BLUE, MTG_GREEN}},
	{"Esper", []color{MTG_WHITE, MTG_BLUE, MTG_BLACK}},
	{"Naya", []color{MTG_WHITE, MTG_RED, MTG_GREEN}},
	{"Abzan", []color{MTG_WHITE, MTG_BLACK, MTG_GREEN}},
	{"Jeskai", []color{MTG_WHITE, MTG_BLUE, MTG_RED}},
	{"Sultai", []color{MTG_BLUE, MTG_BLACK, MTG_GREEN}},
	{"Jund", []color{MTG_BLACK, MTG_RED, MTG_GREEN}},
	{"Grixis", []color{MTG_BLUE, MTG_BLACK, MTG_RED}},
	{"Mardu", []color{MTG_WHITE, MTG_BLACK, MTG_RED}},
	{"Temur", []color{MTG_BLUE, MTG_RED, MTG_GREEN}},
}

type quizMagicFunc func([]magicDeck) promptAndResponse

func quizMagic(cmd *cobra.Command, args []string) {
	funcs := []quizMagicFunc{
		quizDeckFromColors,
		quizColorsFromDeck,
	}

	function := funcs[rand.Intn(len(funcs))]
	promptAndCheckResponse(function(magicDecks))
}

func quizDeckFromColors(decks []magicDeck) promptAndResponse {
	deck := randomDeck(decks)
	colorCombo := colorArrayToString(deck.colors)
	return promptAndResponse{fmt.Sprintf("What is the deck name for %s", colorCombo), deck.name}
}

func quizColorsFromDeck(decks []magicDeck) promptAndResponse {
	deck := randomDeck(decks)
	colors := colorArrayToString(deck.colors)
	return promptAndResponse{fmt.Sprintf("What are the colors (WUBRG order) for %s", deck.name), colors}
}

func colorArrayToString(deckColors []color) string {
	var colors []string
	for _, color := range deckColors {
		colors = append(colors, string(color))
	}
	return strings.Join(colors, ",")
}

func randomDeck(decks []magicDeck) magicDeck {
	return decks[rand.Intn(len(decks))]
}

func init() {
	memoryquizCmd.AddCommand(magicCmd)
}
