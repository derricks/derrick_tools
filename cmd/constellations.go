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

// shakespeareCmd represents the shakespeare command
var constellationCmd = &cobra.Command{
	Use:   "constellations",
	Short: "Test recall of officially recognized constellations in alphabetical order",
	Run:   quizConstellations,
}

type constellation struct {
	name        string
	description string
	// named stars in decreasing order of brightness
	stars []string
}

var constellations = []constellation{
	constellation{"Andromeda", "Princess of Ethiopia", []string{"Alpheratz", "Mirach", "Almach", "Sadiradra", "Nembus", "Titawin", "Keff al Salsalat", "Adhil", "Veritate"}},
	constellation{"Antlia", "Air Pump", []string{"Macondo"}},
	constellation{"Apus", "Bird of Paradise", []string{}},
	constellation{"Aquarius", "Water Bearer", []string{}},
	constellation{"Aquila", "Eagle", []string{}},
	constellation{"Ara", "Altar", []string{}},
	constellation{"Aries", "Ram", []string{}},
	constellation{"Auriga", "Charioteer", []string{}},
	constellation{"Bootes", "Herdsman", []string{}},
	constellation{"Caelum", "Graving Tool", []string{}},
	constellation{"Camelopardalis", "Giraffe", []string{}},
	constellation{"Cancer", "Crab", []string{}},
	constellation{"Canes Venatici", "Hunting Dogs", []string{}},
	constellation{"Canis Major", "Big Dog", []string{}},
	constellation{"Canis Minor", "Little Dog", []string{}},
	constellation{"Capricornus", "Sea Goat", []string{}},
	constellation{"Carina", "Keel of Argonauts' Ship", []string{}},
	constellation{"Cassiopeia", "Queen of Ethiopia", []string{}},
	constellation{"Centaurus", "Centaur", []string{}},
	constellation{"Cepheus", "King of Ethiopia", []string{}},
	constellation{"Cetus", "Whale", []string{}},
	constellation{"Chamaeleon", "Chameleon", []string{}},
	constellation{"Circinus", "Compasses", []string{}},
	constellation{"Columbra", "Dove", []string{}},
	constellation{"Coma Berenices", "Berenice's Hair", []string{}},
	constellation{"Corona Australis", "Southern Crown", []string{}},
	constellation{"Corona Borealis", "Northern Crown", []string{}},
	constellation{"Corvus", "Crow", []string{}},
	constellation{"Crater", "Cup", []string{}},
	constellation{"Crux", "Southern Cross", []string{}},
	constellation{"Cygnus", "Swan", []string{}},
	constellation{"Delphinus", "Porpoise", []string{}},
	constellation{"Dorado", "Swordfish", []string{}},
	constellation{"Draco", "Dragon", []string{}},
	constellation{"Equuleus", "Horse", []string{}},
	constellation{"Eridanus", "River", []string{}},
	constellation{"Fornax", "Furnace", []string{}},
	constellation{"Gemini", "Twins", []string{}},
	constellation{"Grus", "Crane", []string{}},
	constellation{"Hercules", "Hercules", []string{}},
	constellation{"Horologium", "Clock", []string{}},
	constellation{"Hydra", "Sea Serpent", []string{}},
	constellation{"Hydrus", "Water Snake", []string{}},
	constellation{"Indus", "Indian", []string{}},
	constellation{"Lacerta", "Lizard", []string{}},
	constellation{"Leo", "Lion", []string{}},
	constellation{"Leo Minor", "Little Lion", []string{}},
	constellation{"Lepus", "Hare", []string{}},
	constellation{"Libra", "Balance", []string{}},
	constellation{"Lupus", "Wolf", []string{}},
	constellation{"Lynx", "Lynx", []string{}},
	constellation{"Lyra", "Lyre", []string{}},
	constellation{"Mensa", "Tabletop Mountain", []string{}},
	constellation{"Microscopium", "Microscope", []string{}},
	constellation{"Monoceros", "Unicorn", []string{}},
	constellation{"Musca", "Fly", []string{}},
	constellation{"Norma", "Carpenter's Level", []string{}},
	constellation{"Octans", "Octant", []string{}},
	constellation{"Ophiuchus", "Holder of Serpent", []string{}},
	constellation{"Orion", "Hunter", []string{}},
	constellation{"Pavo", "Peacock", []string{}},
	constellation{"Pegasus", "Pegasus", []string{}},
	constellation{"Perseus", "Perseus", []string{}},
	constellation{"Phoenix", "Phoenix", []string{}},
	constellation{"Pictor", "Easel", []string{}},
	constellation{"Pisces", "Fishes", []string{}},
	constellation{"Piscis Austrina", "Little Fishes", []string{}},
	constellation{"Puppis", "Stern of Argonauts' Ship", []string{}},
	constellation{"Pyxis", "Compass of Argonauts' Ship", []string{}},
	constellation{"Reticulum", "Net", []string{}},
	constellation{"Sagitta", "Arrow", []string{}},
	constellation{"Sagittarius", "Archer", []string{}},
	constellation{"Scorpius", "Scorpion", []string{}},
	constellation{"Sculptor", "Sculptor's Tools", []string{}},
	constellation{"Scutum", "Shield", []string{}},
	constellation{"Serpens", "Serpent", []string{}},
	constellation{"Sextans", "Sextant", []string{}},
	constellation{"Taurus", "Bull", []string{}},
	constellation{"Telescopium", "Telescope", []string{}},
	constellation{"Triangulum", "Triangle", []string{}},
	constellation{"Triangulum Australis", "Southern Triangle", []string{}},
	constellation{"Tucana", "Toucan", []string{}},
	constellation{"Ursa Major", "Big Bear", []string{}},
	constellation{"Ursa Minor", "Little Bear", []string{}},
	constellation{"Vela", "Sail of Argonauts' Ship", []string{}},
	constellation{"Virgo", "Virgin", []string{}},
	constellation{"Volans", "Flying Fish", []string{}},
	constellation{"Vulpecula", "Fox", []string{}},
}

type constellationQuiz func([]constellation) promptAndResponse

func quizConstellations(cmd *cobra.Command, args []string) {
	quizzes := []constellationQuiz{
		quizConstellationByOrder,
		quizConstellationFromDescription,
		quizConstellationCountByLetter,
		quizStarInConstellation,
		quizConstellationByStar,
	}

	quiz := quizzes[rand.Intn(len(quizzes))]
	promptAndCheckResponse(quiz(constellations))
}

func quizConstellationByOrder(constellations []constellation) promptAndResponse {
	index := rand.Intn(len(constellations))
	constellation := constellations[index]
	return promptAndResponse{fmt.Sprintf("Which constellation is position %d?", index+1), constellation.name}
}

func quizConstellationFromDescription(constellations []constellation) promptAndResponse {
	constellation := randomConstellation(constellations)
	return promptAndResponse{fmt.Sprintf("What's the official name of the %s constellation?", constellation.description), constellation.name}
}

func quizConstellationCountByLetter(constellations []constellation) promptAndResponse {
	// pick a letter
	letterAscii := rune(65 + rand.Intn(26))
	letter := string(letterAscii)
	count := 0
	for _, constellation := range constellations {
		if strings.HasPrefix(constellation.name, letter) {
			count++
		}
	}
	return promptAndResponse{fmt.Sprintf("How many constellations start with %s?", letter), strconv.Itoa(count)}
}

func quizStarInConstellation(constellations []constellation) promptAndResponse {
	// winnow down to just constellations that have stars
	withStars := constellationsWithStars(constellations)
	constellation := randomConstellation(withStars)
	starIndex := rand.Intn(len(constellation.stars))
	return promptAndResponse{fmt.Sprintf("What is named star number %d in %s?", starIndex+1, constellation.name), constellation.stars[starIndex]}
}

func quizConstellationByStar(constellations []constellation) promptAndResponse {
	withStars := constellationsWithStars(constellations)
	constellation := randomConstellation(withStars)
	star := constellation.stars[rand.Intn(len(constellation.stars))]
	return promptAndResponse{fmt.Sprintf("Which constellation contains %s?", star), constellation.name}
}

// return a subset of constellations that have stars listed for them
// once stars are listed for all the constellations, this logic can be removed
func constellationsWithStars(constellations []constellation) []constellation {
	withStars := make([]constellation, 0, len(constellations))
	for _, constellation := range constellations {
		if len(constellation.stars) > 0 {
			withStars = append(withStars, constellation)
		}
	}
	return withStars
}

func randomConstellation(constellations []constellation) constellation {
	return constellations[rand.Intn(len(constellations))]
}

func init() {
	memoryquizCmd.AddCommand(constellationCmd)
}
