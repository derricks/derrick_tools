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

// shakespeareCmd represents the shakespeare command
var musesCmd = &cobra.Command{
	Use:   "muses",
	Short: "Test recall of the names and areas of the nine muses",
	Run:   quizMuses,
}

type muse struct {
	name  string
	areas []string
}

func (m muse) areasFormatted() string {
	return strings.Join(m.areas, " and ")
}

var muses = []muse{
	{"Calliope", []string{"epic poetry"}},
	{"Clio", []string{"history"}},
	{"Euterpe", []string{"flutes", "music"}},
	{"Thalia", []string{"pastoral poetry"}},
	{"Melpomene", []string{"tragedy"}},
	{"Terpsichore", []string{"dance"}},
	{"Erato", []string{"love poetry"}},
	{"Polyhymnia", []string{"sacred poetry"}},
	{"Urania", []string{"astronomy"}},
}

type museQuiz func([]muse) promptAndResponse

func quizMuses(cmd *cobra.Command, args []string) {
	quizzes := []museQuiz{
		quizMuseByArea,
		quizAreaByMuse,
		quizAllMuses,
	}

	quiz := randomItemFromSlice(quizzes)
	promptAndCheckResponse(quiz(muses))
}

func quizMuseByArea(muses []muse) promptAndResponse {
	muse := randomMuse(muses)
	return promptAndResponse{fmt.Sprintf("Who is the muse of %s?", muse.areasFormatted()), muse.name}
}

func quizAreaByMuse(muses []muse) promptAndResponse {
	muse := randomMuse(muses)
	return promptAndResponse{fmt.Sprintf("What is %s the muse of?", muse.name), muse.areasFormatted()}
}

func quizAllMuses(muses []muse) promptAndResponse {
	museNames := []string{}
	for i := 0; i < len(muses); i++ {
		museNames = append(museNames, muses[i].name)
	}
	return promptAndResponse{"Name all the muses in memory walk order", strings.Join(museNames, ",")}
}

func randomMuse(muses []muse) muse {
	return muses[rand.Intn(len(muses))]
}

func init() {
	memoryquizCmd.AddCommand(musesCmd)
}
