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
	alphabeticalOrder int    `crossquery:"all" crossqueryname:"order"`
	name              string `crossquery:"all"`
	description       string `crossquery:"all"`
	// named stars in decreasing order of brightness
	stars []string
}

var constellations = []constellation{
	constellation{1, "Andromeda", "Princess of Ethiopia", []string{"Alpheratz", "Mirach", "Almach", "Sadiradra", "Nembus", "Titawin", "Keff al Salsalat", "Adhil", "Veritate"}},
	constellation{2, "Antlia", "Air Pump", []string{"Macondo"}},
	constellation{3, "Apus", "Bird of Paradise", []string{"Karaka"}},
	constellation{4, "Aquarius", "Water Bearer", []string{"Sadalsuud", "Sadalmelik", "Skat", "Albali", "Sadachbia", "Ancha", "Bunda", "Situla", "Lionrock", "Bosona", "Marohu"}},
	constellation{5, "Aquila", "Eagle", []string{"Altair", "Tarazed", "Okab", "Alshain", "Libertas", "Chechia", "Phoenicia", "Petra"}},
	constellation{6, "Ara", "Altar", []string{"Cervantes", "Inquill"}},
	constellation{7, "Aries", "Ram", []string{"Hamal", "Sheratan", "Bharani", "Botein", "Lilii Borea", "Mesarthim"}},
	constellation{8, "Auriga", "Charioteer", []string{"Capella", "Menkalinan", "Mahasim", "Hasseleh", "Almaaz", "Haedus", "Saclatini", "Lucilinburhuc", "Nervia", "Tevel"}},
	constellation{9, "Bootes", "Herdsman", []string{"Arcturus", "Izar", "Muphrid", "Seginus", "Nekkar", "Xuange", "Alkalurops", "Merga", "Nikiwiy", "Arcalis"}},
	constellation{10, "Caelum", "Graving Tool", []string{}},
	constellation{11, "Camelopardalis", "Giraffe", []string{"Tonatiuh", "Mago"}},
	constellation{12, "Cancer", "Crab", []string{"Tarf", "Asellus Australis", "Acubens", "Asellus Borealis", "Tegmine", "Nahn", "Piautos", "Copernicus", "Meleph", "Gakyid"}},
	constellation{13, "Canes Venatici", "Hunting Dogs", []string{"Cor Caroli", "Chara", "La Superba", "Tuiren"}},
	constellation{14, "Canis Major", "Big Dog", []string{"Sirius", "Adhara", "Wezen", "Mirzam", "Aludra", "Furud", "Unurgunite", "Muliphein", "Amadioha", "Atakoraka"}},
	constellation{15, "Canis Minor", "Little Dog", []string{"Procyon", "Gomeisa"}},
	constellation{16, "Capricornus", "Sea Goat", []string{"Deneb Algedi", "Dabih", "Algedi", "Nashira", "Alshat"}},
	constellation{17, "Carina", "Keel of Argonauts' Ship", []string{"Canopus", "Miaplacidus", "Avior", "Aspidiske", "Tapecue"}},
	constellation{18, "Cassiopeia", "Queen of Ethiopia", []string{"Schedar", "Caph", "Ruchbah", "Segin", "Achird", "Fulu", "Castula", "Nushagak"}},
	constellation{19, "Centaurus", "Centaur", []string{"Rigil Kentaurus", "Hadar", "Toliman", "Menkent", "Dofida", "Uklun", "Nyamien", "Proxima Centauri"}},
	constellation{20, "Cepheus", "King of Ethiopia", []string{}},
	constellation{21, "Cetus", "Whale", []string{}},
	constellation{22, "Chamaeleon", "Chameleon", []string{}},
	constellation{23, "Circinus", "Compasses", []string{}},
	constellation{24, "Columbra", "Dove", []string{}},
	constellation{25, "Coma Berenices", "Berenice's Hair", []string{}},
	constellation{26, "Corona Australis", "Southern Crown", []string{}},
	constellation{27, "Corona Borealis", "Northern Crown", []string{}},
	constellation{28, "Corvus", "Crow", []string{}},
	constellation{29, "Crater", "Cup", []string{}},
	constellation{30, "Crux", "Southern Cross", []string{}},
	constellation{31, "Cygnus", "Swan", []string{}},
	constellation{32, "Delphinus", "Porpoise", []string{}},
	constellation{33, "Dorado", "Swordfish", []string{}},
	constellation{34, "Draco", "Dragon", []string{}},
	constellation{35, "Equuleus", "Horse", []string{}},
	constellation{36, "Eridanus", "River", []string{}},
	constellation{37, "Fornax", "Furnace", []string{}},
	constellation{38, "Gemini", "Twins", []string{}},
	constellation{39, "Grus", "Crane", []string{}},
	constellation{40, "Hercules", "Hercules", []string{}},
	constellation{41, "Horologium", "Clock", []string{}},
	constellation{42, "Hydra", "Sea Serpent", []string{}},
	constellation{43, "Hydrus", "Water Snake", []string{}},
	constellation{44, "Indus", "Indian", []string{}},
	constellation{45, "Lacerta", "Lizard", []string{}},
	constellation{46, "Leo", "Lion", []string{}},
	constellation{47, "Leo Minor", "Little Lion", []string{}},
	constellation{48, "Lepus", "Hare", []string{}},
	constellation{49, "Libra", "Balance", []string{}},
	constellation{50, "Lupus", "Wolf", []string{}},
	constellation{51, "Lynx", "Lynx", []string{}},
	constellation{52, "Lyra", "Lyre", []string{}},
	constellation{53, "Mensa", "Tabletop Mountain", []string{}},
	constellation{54, "Microscopium", "Microscope", []string{}},
	constellation{55, "Monoceros", "Unicorn", []string{}},
	constellation{56, "Musca", "Fly", []string{}},
	constellation{57, "Norma", "Carpenter's Level", []string{}},
	constellation{58, "Octans", "Octant", []string{}},
	constellation{59, "Ophiuchus", "Holder of Serpent", []string{}},
	constellation{60, "Orion", "Hunter", []string{}},
	constellation{61, "Pavo", "Peacock", []string{}},
	constellation{62, "Pegasus", "Pegasus", []string{}},
	constellation{63, "Perseus", "Perseus", []string{}},
	constellation{64, "Phoenix", "Phoenix", []string{}},
	constellation{65, "Pictor", "Easel", []string{}},
	constellation{66, "Pisces", "Fishes", []string{}},
	constellation{67, "Piscis Austrina", "Little Fishes", []string{}},
	constellation{68, "Puppis", "Stern of Argonauts' Ship", []string{}},
	constellation{69, "Pyxis", "Compass of Argonauts' Ship", []string{}},
	constellation{70, "Reticulum", "Net", []string{}},
	constellation{71, "Sagitta", "Arrow", []string{}},
	constellation{72, "Sagittarius", "Archer", []string{}},
	constellation{73, "Scorpius", "Scorpion", []string{}},
	constellation{74, "Sculptor", "Sculptor's Tools", []string{}},
	constellation{75, "Scutum", "Shield", []string{}},
	constellation{76, "Serpens", "Serpent", []string{}},
	constellation{77, "Sextans", "Sextant", []string{}},
	constellation{78, "Taurus", "Bull", []string{}},
	constellation{79, "Telescopium", "Telescope", []string{}},
	constellation{80, "Triangulum", "Triangle", []string{}},
	constellation{81, "Triangulum Australis", "Southern Triangle", []string{}},
	constellation{82, "Tucana", "Toucan", []string{}},
	constellation{83, "Ursa Major", "Big Bear", []string{}},
	constellation{84, "Ursa Minor", "Little Bear", []string{}},
	constellation{85, "Vela", "Sail of Argonauts' Ship", []string{}},
	constellation{86, "Virgo", "Virgin", []string{}},
	constellation{87, "Volans", "Flying Fish", []string{}},
	constellation{88, "Vulpecula", "Fox", []string{}},
}

type constellationQuiz func([]constellation) promptAndResponse

func quizConstellations(cmd *cobra.Command, args []string) {
	quizzes := []constellationQuiz{
		quizConstellationByOrder,
		crossQueryConstellationInfo,
		crossQueryConstellationInfo,
		quizConstellationCountByLetter,
		quizStarInConstellation,
		quizConstellationByStar,
	}

	quiz := randomItemFromSlice(quizzes)
	promptAndCheckResponse(quiz(constellations))
}

func crossQueryConstellationInfo(constellations []constellation) promptAndResponse {
	constellation := randomItemFromSlice(constellations)
	return constructCrossQuery("constellation", constellation)
}

func quizConstellationByOrder(constellations []constellation) promptAndResponse {
	constellation := randomItemFromSlice(constellations)
	return promptAndResponse{fmt.Sprintf("Which constellation is position %d?", constellation.alphabeticalOrder), constellation.name}
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
	constellation := randomItemFromSlice(withStars)
	starIndex := rand.Intn(len(constellation.stars))
	star := constellation.stars[starIndex]
	return promptAndResponse{fmt.Sprintf("What is named star number %d in %s?", starIndex+1, constellation.name), star}
}

func quizConstellationByStar(constellations []constellation) promptAndResponse {
	withStars := constellationsWithStars(constellations)
	constellation := randomItemFromSlice(withStars)
	star := randomItemFromSlice(constellation.stars)
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

func init() {
	memoryquizCmd.AddCommand(constellationCmd)
}
