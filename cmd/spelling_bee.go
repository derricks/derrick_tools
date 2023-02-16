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

var spellingBeeCmd = &cobra.Command{
	Use:   "spellingbee",
	Short: "Quiz Spelling Bee word sets",
	Run:   quizSpellingBee,
}

var spellingBeeSets = [][]string{
	{"FAIR", "FRIAR", "AFFAIR", "RIFFRAFF", "RAFFIA"},
	{"LATHE", "ATHLETE", "LETHAL", "HEALTH", "TELEHEALTH"},
	{"ACORN", "CORONA", "RANCOR", "RACCOON"},
	{"HEAD", "AHEAD", "HEADED", "DEADHEAD", "DEADHEADED"},
	{"TRIAL", "RATTAIL", "LARIAT", "TRAIL", "ATRIAL"},
	{"RATIO", "TRAITOR", "ORATORIO", "TRATTORIA"},
	{"FOAL", "ALOOF", "OFFAL", "LOAF", "LOOFA", "FALLOFF"},
	{"ELBOW", "BELLOW", "BELOW", "WOBBLE", "BOWEL"},
	{"DINE", "DENIED", "INDEED", "DINED", "INDIE"},
	{"LINGO", "OGLING", "GOOGLING", "GOGGLING", "LOGGING", "OILING", "LOGIN"},
	{"DICE", "DICED", "DEICE", "DEICED", "DECIDE", "DECIDED", "DIECIDE", "ICED"},
	{"INGOT", "TOGGING", "TOTING", "TOOTING", "IGNITION", "INTONING", "TONING", "NOTING"},
	{"ACHED", "HEADACHE", "CACHED"},
	{"BINGO", "GIBBON", "BONING", "BOING", "BOOING"},
	{"NOTICE", "CONCEIT", "INNOCENT", "CONTINENT", "TECTONIC", "CONNECTION", "CONTENTION", "INCONTINENT"},
	{"GENIE", "EGGING", "GENII", "ENGINE"},
	{"PIPET", "PIPETTE", "PETIT", "TIPPET", "PETITE"},
	{"CART", "TRACT", "CATARACT", "ATTRACT", "CARAT"},
	{"TACIT", "TACTIC", "CACTI", "ATTIC"},
	{"AWING", "GNAWING", "AWNING", "WAGING", "WAGGING"},
	{"CLEAN", "ENLACE", "LANCE", "CANCEL"},
	{"OWING", "WOOING", "GOWNING", "OWNING", "WOWING"},
	{"GALE", "EAGLE", "LEGAL", "ALLEGE"},
	{"MEAN", "NAME", "ENEMA", "MANE", "AMEN"},
	{"TEAM", "TEAMMATE", "MATTE", "MEAT", "META", "MATE"},
	{"TRAP", "APPARAT", "RAPT", "APART", "PART", "TARP", "RATTRAP"},
	{"PEND", "PENDED", "DEPEND", "DEPENDED", "DEEPEN", "DEEPENED", "PENNED"},
	{"PACE", "APACE", "PEACE", "CAPE"},
	{"PEAL", "APPLE", "PALE", "LAPEL", "LEAP", "APPEAL", "PAELLA", "APPELLEE", "PLEA"},
	{"HOLE", "HELLO", "HELLHOLE"},
	{"PLEAD", "PLEADED", "DAPPLED", "PADDLE", "PEDAL", "LEAPED", "PALED", "PALLED", "LAPPED", "DAPPLE", "PEALED", "APPEALED"},
	{"PECAN", "PENANCE", "PANACEA", "CANAPE"},
	{"ICON", "COIN", "IONIC", "CONIC", "ICONIC"},
	{"TEEN", "TENT", "TENET", "ENTENTE"},
	{"ATOM", "MOAT", "TOMATO"},
	{"NICE", "NIECE"},
	{"ACID", "ACIDIC", "CICADA"},
	{"INTEL", "LENTIL", "LINTEL", "LINNET", "INLET", "ENTITLE", "LENIENT"},
	{"LAIC", "LAICAL", "LILAC", "ILIAC"},
	{"DANCE", "CANED", "CANNED", "DECADENCE", "DANCED", "CADENCE"},
	{"LAMA", "LLAMA", "MAMMAL", "MALL"},
	{"ABLE", "BALE", "LABEL", "BABBLE", "BABEL"},
	{"BILE", "BIBLE", "BELIE", "LIBEL", "LIBELEE"},
	{"PAINT", "PATINA", "PINATA", "INAPT"},
	{"LOOT", "LOTTO", "TOOL", "TOLL"},
	{"PALL", "APPALL", "PAPAL", "PALAPA"},
	{"VIAL", "AVAIL", "VILLA"},
	{"ANVIL", "VILLAIN", "VANILLA", "VANILLIN"},
	{"LUNA", "ULNA", "ANNUL", "ANNUAL", "LUNULA"},
	{"GAVE", "AGAVE", "GAVAGE"},
	{"OUGHT", "OUTTHOUGHT", "TOUGH", "THOUGHT"},
	{"RANT", "RATTAN", "TARTAN", "TANTRA", "TANTARA"},
	{"GAIN", "AGAIN", "AGING", "GAINING", "NAGGING", "ANGINA", "GAGGING", "GANGING"},
	{"DOGE", "GEODE", "DODGE", "DOGGED", "DODGED"},
}

func quizSpellingBee(cmd *cobra.Command, args []string) {
	wordSet := spellingBeeSets[rand.Intn(len(spellingBeeSets))]
	word := wordSet[rand.Intn(len(wordSet))]
	inputSet := responseFromPrompt(promptAndResponse{fmt.Sprintf("What are other Spelling Bee words for %s (separate by commas)?", word), ""})

	enteredWords := strings.Split(inputSet, ",")

	error := false
	// verify that entered words doesn't have entries not in the list
	for _, enteredWord := range enteredWords {
		if !isStringInSlice(enteredWord, wordSet) {
			error = true
			fmt.Printf("%s is not in the list of words for %s\n", enteredWord, word)
		}
	}

	// and now verify that the list isn't missing any words
	for _, validWord := range wordSet {
		if !isStringInSlice(validWord, enteredWords) && validWord != word {
			error = true
			fmt.Printf("You missed %s\n", validWord)
		}
	}

	if !error {
		fmt.Println("You got them all!")
	}
}

func init() {
	memoryquizCmd.AddCommand(spellingBeeCmd)
}
