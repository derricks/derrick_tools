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
}

var constellations = []constellation{
	constellation{"Andromeda", "Princess of Ethiopia"},
	constellation{"Antlia", "Air Pump"},
	constellation{"Apus", "Bird of Paradise"},
	constellation{"Aquarius", "Water Bearer"},
	constellation{"Aquila", "Eagle"},
	constellation{"Ara", "Altar"},
	constellation{"Aries", "Ram"},
	constellation{"Auriga", "Charioteer"},
	constellation{"Bootes", "Herdsman"},
	constellation{"Caelum", "Graving Tool"},
	constellation{"Camelopardalis", "Giraffe"},
	constellation{"Cancer", "Crab"},
	constellation{"Canes Venatici", "Hunting Dogs"},
	constellation{"Canis Major", "Big Dog"},
	constellation{"Canis Minor", "Little Dog"},
	constellation{"Capricornus", "Sea Goat"},
	constellation{"Carina", "Keel of Argonauts' Ship"},
	constellation{"Cassiopeia", "Queen of Ethiopia"},
	constellation{"Centaurus", "Centaur"},
	constellation{"Cepheus", "King of Ethiopia"},
	constellation{"Cetus", "Whale"},
	constellation{"Chamaeleon", "Chameleon"},
	constellation{"Circinus", "Compasses"},
	constellation{"Columbra", "Dove"},
	constellation{"Coma Berenices", "Berenice's Hair"},
	constellation{"Corona Australis", "Southern Crown"},
	constellation{"Corona Borealis", "Northern Crown"},
	constellation{"Corvus", "Crow"},
	constellation{"Crater", "Cup"},
	constellation{"Crux", "Southern Cross"},
	constellation{"Cygnus", "Swan"},
	constellation{"Delphinus", "Porpoise"},
	constellation{"Dorado", "Swordfish"},
	constellation{"Draco", "Dragon"},
	constellation{"Equuleus", "Horse"},
	constellation{"Eridanus", "River"},
	constellation{"Fornax", "Furnace"},
	constellation{"Gemini", "Twins"},
	constellation{"Grus", "Crane"},
	constellation{"Hercules", "Hercules"},
	constellation{"Horologium", "Clock"},
	constellation{"Hydra", "Sea Serpent"},
	constellation{"Hydrus", "Water Snake"},
	constellation{"Indus", "Indian"},
	constellation{"Lacerta", "Lizard"},
	constellation{"Leo", "Lion"},
	constellation{"Leo Minor", "Little Lion"},
	constellation{"Lepus", "Hare"},
	constellation{"Libra", "Balance"},
	constellation{"Lupos", "Wolf"},
	constellation{"Lynx", "Lynx"},
	constellation{"Lyra", "Lyre"},
	constellation{"Mensa", "Tabletop Mountain"},
	constellation{"Microscopium", "Microscope"},
	constellation{"Monoceros", "Unicorn"},
	constellation{"Musca", "Fly"},
	constellation{"Norma", "Carpenter's Level"},
	constellation{"Octans", "Octant"},
	constellation{"Ophiuchus", "Holder of Serpent"},
	constellation{"Orion", "Hunter"},
	constellation{"Pavo", "Peacock"},
	constellation{"Pegasus", "Pegasus"},
	constellation{"Perseus", "Perseus"},
	constellation{"Phoenix", "Phoenix"},
	constellation{"Pictor", "Easel"},
	constellation{"Pisces", "Fishes"},
	constellation{"Piscis Austrina", "Little Fishes"},
	constellation{"Puppis", "Stern of Argonauts' Ship"},
	constellation{"Pyxis", "Compass of Argonauts' Ship"},
	constellation{"Reticulum", "Net"},
	constellation{"Sagitta", "Arrow"},
	constellation{"Sagittarius", "Archer"},
	constellation{"Scorpius", "Scorpion"},
	constellation{"Sculptor", "Sculptor's Tools"},
	constellation{"Scutum", "Shield"},
	constellation{"Serpens", "Serpent"},
	constellation{"Sextans", "Sextant"},
	constellation{"Taurus", "Bull"},
	constellation{"Telescopium", "Telescope"},
	constellation{"Triangulum", "Triangle"},
	constellation{"Triangulum Australis", "Southern Triangle"},
	constellation{"Tucana", "Toucan"},
	constellation{"Ursa Major", "Big Bear"},
	constellation{"Ursa Minor", "Little Bear"},
	constellation{"Vela", "Sail of Argonauts' Ship"},
	constellation{"Virgo", "Virgin"},
	constellation{"Volans", "Flying Fish"},
	constellation{"Vulpecula", "Fox"},
}

type constellationQuiz func([]constellation) promptAndResponse

func quizConstellations(cmd *cobra.Command, args []string) {
	quizzes := []constellationQuiz{
		quizConstellationByOrder,
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

func randomConstellation(constellations []constellation) constellation {
	return constellations[rand.Intn(len(constellations))]
}

func init() {
	memoryquizCmd.AddCommand(constellationCmd)
}
