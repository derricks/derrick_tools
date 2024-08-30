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
	{1, "Andromeda", "Princess of Ethiopia", []string{"Alpheratz", "Mirach", "Almach", "Sadiradra", "Nembus", "Titawin", "Keff al Salsalat", "Adhil", "Veritate"}},
	{2, "Antlia", "Air Pump", []string{"Macondo"}},
	{3, "Apus", "Bird of Paradise", []string{"Karaka"}},
	{4, "Aquarius", "Water Bearer", []string{"Sadalsuud", "Sadalmelik", "Skat", "Albali", "Sadachbia", "Ancha", "Bunda", "Situla", "Lionrock", "Bosona", "Marohu"}},
	{5, "Aquila", "Eagle", []string{"Altair", "Tarazed", "Okab", "Alshain", "Libertas", "Chechia", "Phoenicia", "Petra"}},
	{6, "Ara", "Altar", []string{"Cervantes", "Inquill"}},
	{7, "Aries", "Ram", []string{"Hamal", "Sheratan", "Bharani", "Botein", "Lilii Borea", "Mesarthim"}},
	{8, "Auriga", "Charioteer", []string{"Capella", "Menkalinan", "Mahasim", "Hasseleh", "Almaaz", "Haedus", "Saclatini", "Lucilinburhuc", "Nervia", "Tevel"}},
	{9, "Bootes", "Herdsman", []string{"Arcturus", "Izar", "Muphrid", "Seginus", "Nekkar", "Xuange", "Alkalurops", "Merga", "Nikiwiy", "Arcalis"}},
	{10, "Caelum", "Graving Tool", []string{}},
	{11, "Camelopardalis", "Giraffe", []string{"Tonatiuh", "Mago"}},
	{12, "Cancer", "Crab", []string{"Tarf", "Asellus Australis", "Acubens", "Asellus Borealis", "Tegmine", "Nahn", "Piautos", "Copernicus", "Meleph", "Gakyid"}},
	{13, "Canes Venatici", "Hunting Dogs", []string{"Cor Caroli", "Chara", "La Superba", "Tuiren"}},
	{14, "Canis Major", "Big Dog", []string{"Sirius", "Adhara", "Wezen", "Mirzam", "Aludra", "Furud", "Unurgunite", "Muliphein", "Amadioha", "Atakoraka"}},
	{15, "Canis Minor", "Little Dog", []string{"Procyon", "Gomeisa"}},
	{16, "Capricornus", "Sea Goat", []string{"Deneb Algedi", "Dabih", "Algedi", "Nashira", "Alshat"}},
	{17, "Carina", "Keel of Argonauts' Ship", []string{"Canopus", "Miaplacidus", "Avior", "Aspidiske", "Tapecue"}},
	{18, "Cassiopeia", "Queen of Ethiopia", []string{"Schedar", "Caph", "Ruchbah", "Segin", "Achird", "Fulu", "Castula", "Nushagak"}},
	{19, "Centaurus", "Centaur", []string{"Rigil Kentaurus", "Hadar", "Toliman", "Menkent", "Dofida", "Uklun", "Nyamien", "Proxima Centauri"}},
	{20, "Cepheus", "King of Ethiopia", []string{}},
	{21, "Cetus", "Whale", []string{}},
	{22, "Chamaeleon", "Chameleon", []string{}},
	{23, "Circinus", "Compasses", []string{}},
	{24, "Columbra", "Dove", []string{}},
	{25, "Coma Berenices", "Berenice's Hair", []string{}},
	{26, "Corona Australis", "Southern Crown", []string{}},
	{27, "Corona Borealis", "Northern Crown", []string{}},
	{28, "Corvus", "Crow", []string{}},
	{29, "Crater", "Cup", []string{}},
	{30, "Crux", "Southern Cross", []string{}},
	{31, "Cygnus", "Swan", []string{}},
	{32, "Delphinus", "Porpoise", []string{}},
	{33, "Dorado", "Swordfish", []string{}},
	{34, "Draco", "Dragon", []string{}},
	{35, "Equuleus", "Horse", []string{}},
	{36, "Eridanus", "River", []string{}},
	{37, "Fornax", "Furnace", []string{}},
	{38, "Gemini", "Twins", []string{}},
	{39, "Grus", "Crane", []string{}},
	{40, "Hercules", "Hercules", []string{}},
	{41, "Horologium", "Clock", []string{}},
	{42, "Hydra", "Sea Serpent", []string{}},
	{43, "Hydrus", "Water Snake", []string{}},
	{44, "Indus", "Indian", []string{}},
	{45, "Lacerta", "Lizard", []string{}},
	{46, "Leo", "Lion", []string{}},
	{47, "Leo Minor", "Little Lion", []string{}},
	{48, "Lepus", "Hare", []string{}},
	{49, "Libra", "Balance", []string{}},
	{50, "Lupus", "Wolf", []string{}},
	{51, "Lynx", "Lynx", []string{}},
	{52, "Lyra", "Lyre", []string{}},
	{53, "Mensa", "Tabletop Mountain", []string{}},
	{54, "Microscopium", "Microscope", []string{}},
	{55, "Monoceros", "Unicorn", []string{}},
	{56, "Musca", "Fly", []string{}},
	{57, "Norma", "Carpenter's Level", []string{}},
	{58, "Octans", "Octant", []string{}},
	{59, "Ophiuchus", "Holder of Serpent", []string{}},
	{60, "Orion", "Hunter", []string{}},
	{61, "Pavo", "Peacock", []string{}},
	{62, "Pegasus", "Pegasus", []string{}},
	{63, "Perseus", "Perseus", []string{}},
	{64, "Phoenix", "Phoenix", []string{}},
	{65, "Pictor", "Easel", []string{}},
	{66, "Pisces", "Fishes", []string{}},
	{67, "Piscis Austrina", "Little Fishes", []string{}},
	{68, "Puppis", "Stern of Argonauts' Ship", []string{}},
	{69, "Pyxis", "Compass of Argonauts' Ship", []string{}},
	{70, "Reticulum", "Net", []string{}},
	{71, "Sagitta", "Arrow", []string{}},
	{72, "Sagittarius", "Archer", []string{}},
	{73, "Scorpius", "Scorpion", []string{}},
	{74, "Sculptor", "Sculptor's Tools", []string{}},
	{75, "Scutum", "Shield", []string{}},
	{76, "Serpens", "Serpent", []string{}},
	{77, "Sextans", "Sextant", []string{}},
	{78, "Taurus", "Bull", []string{}},
	{79, "Telescopium", "Telescope", []string{}},
	{80, "Triangulum", "Triangle", []string{}},
	{81, "Triangulum Australis", "Southern Triangle", []string{}},
	{82, "Tucana", "Toucan", []string{}},
	{83, "Ursa Major", "Big Bear", []string{}},
	{84, "Ursa Minor", "Little Bear", []string{}},
	{85, "Vela", "Sail of Argonauts' Ship", []string{}},
	{86, "Virgo", "Virgin", []string{}},
	{87, "Volans", "Flying Fish", []string{}},
	{88, "Vulpecula", "Fox", []string{}},
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
