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

	"github.com/spf13/cobra"
)

// shakespeareCmd represents the shakespeare command
var lakesCmd = &cobra.Command{
	Use:   "lakes",
	Short: "Test recall of the names, sizes, and salinity of the world's largest lakes",
	Run:   quizLakes,
}

type lakeInfo struct {
	sizeOrder int
	name      string
	isSaline  bool
}

var lakes = []lakeInfo{
	lakeInfo{1, "Caspian Sea", true},
	lakeInfo{2, "Superior", false},
	lakeInfo{3, "Victoria", false},
	lakeInfo{4, "Huron", false},
	lakeInfo{5, "Michigan", false},
	lakeInfo{6, "Tanganyika", false},
	lakeInfo{7, "Baikal", false},
	lakeInfo{8, "Great Bear Lake", false},
	lakeInfo{9, "Malawi", false},
	lakeInfo{10, "Great Slave Lake", false},
	lakeInfo{11, "Erie", false},
	lakeInfo{12, "Winnipeg", false},
	lakeInfo{13, "Ontario", false},
	lakeInfo{14, "Ladoga", false},
	lakeInfo{15, "Balkhash", true},
	lakeInfo{16, "Bangweulu", false},
	lakeInfo{17, "Vostok", false},
	lakeInfo{18, "Onega", false},
	lakeInfo{19, "Titicaca", false},
	lakeInfo{20, "Nicaragua", false},
	lakeInfo{21, "Athabasca", false},
	lakeInfo{22, "Turkana", true},
	lakeInfo{23, "Reindeer Lake", false},
	lakeInfo{24, "Issyk-Kul", true},
	lakeInfo{25, "Urmia", true},
	lakeInfo{26, "Vanern", false},
	lakeInfo{27, "Winnpegosis", false},
	lakeInfo{28, "Albert", false},
	lakeInfo{29, "Mweru", false},
	lakeInfo{30, "Nettilling", false},
	lakeInfo{31, "Nipigon", false},
	lakeInfo{32, "Manitoba", false},
	lakeInfo{33, "Taymyr", false},
	lakeInfo{34, "Qinghai Lake", true},
	lakeInfo{35, "Saimaa", false},
	lakeInfo{36, "Lake of the Woods", false},
	lakeInfo{37, "Khanka", false},
	lakeInfo{38, "Sarygamyish", false},
	lakeInfo{39, "Dubawnt", false},
	lakeInfo{40, "Van", true},
	lakeInfo{41, "Peipus", false},
	lakeInfo{42, "Uvs", true},
	lakeInfo{43, "Poyang", false},
	lakeInfo{44, "Tana", false},
	lakeInfo{45, "Amadjuak", false},
	lakeInfo{46, "Melville", true},
}

type lakeQuiz func([]lakeInfo) promptAndResponse

func quizLakes(cmd *cobra.Command, args []string) {
	quizzes := []lakeQuiz{
		quizLakeBySizeRank,
		quizSizeByLake,
		quizLakeSalinity,
	}

	quiz := quizzes[rand.Intn(len(quizzes))]
	promptAndCheckResponse(quiz(lakes))
}

func randomLake(lakes []lakeInfo) lakeInfo {
	return lakes[rand.Intn(len(lakes))]
}

func quizLakeBySizeRank(lakes []lakeInfo) promptAndResponse {
	lake := randomLake(lakes)
	return promptAndResponse{fmt.Sprintf("What is the name of the lake at position %d", lake.sizeOrder), lake.name}
}

func quizSizeByLake(lakes []lakeInfo) promptAndResponse {
	lake := randomLake(lakes)
	return promptAndResponse{fmt.Sprintf("What is the size rank of lake %s?", lake.name), strconv.Itoa(lake.sizeOrder)}
}

func quizLakeSalinity(lakes []lakeInfo) promptAndResponse {
	lake := randomLake(lakes)
	return promptAndResponse{fmt.Sprintf("%s is a saline lake, true or false?", lake.name), strconv.FormatBool(lake.isSaline)}
}

func init() {
	memoryquizCmd.AddCommand(lakesCmd)
}
