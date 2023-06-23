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
	countries []string
}

var lakes = []lakeInfo{
	lakeInfo{1, "Caspian Sea", true, []string{"Russia", "Kazakhstan", "Turkmenistan", "Iran", "Azerbaijan"}},
	lakeInfo{2, "Superior", false, []string{"Canada", "United States"}},
	lakeInfo{3, "Victoria", false, []string{"Kenya", "Tanzania", "Uganda"}},
	lakeInfo{4, "Huron", false, []string{"Canada", "United States"}},
	lakeInfo{5, "Michigan", false, []string{"United States"}},
	lakeInfo{6, "Tanganyika", false, []string{"Burundi", "Zambia", "Democratic Republic of the Congo", "Tanzania"}},
	lakeInfo{7, "Baikal", false, []string{"Russia"}},
	lakeInfo{8, "Great Bear Lake", false, []string{"Canada"}},
	lakeInfo{9, "Malawi", false, []string{"Mozambique", "Tanzania"}},
	lakeInfo{10, "Great Slave Lake", false, []string{"Canada"}},
	lakeInfo{11, "Erie", false, []string{"Canada", "United States"}},
	lakeInfo{12, "Winnipeg", false, []string{"Canada"}},
	lakeInfo{13, "Ontario", false, []string{"Canada", "United States"}},
	lakeInfo{14, "Ladoga", false, []string{"Russia"}},
	lakeInfo{15, "Balkhash", true, []string{"Kazakhstan"}},
	lakeInfo{16, "Bangweulu", false, []string{"Zambia"}},
	lakeInfo{17, "Vostok", false, []string{"Antarctica"}},
	lakeInfo{18, "Onega", false, []string{"Russia"}},
	lakeInfo{19, "Titicaca", false, []string{"Bolivia", "Peru"}},
	lakeInfo{20, "Nicaragua", false, []string{"Nicaragua"}},
	lakeInfo{21, "Athabasca", false, []string{"Canada"}},
	lakeInfo{22, "Turkana", true, []string{"Kenya", "Ethiopia"}},
	lakeInfo{23, "Reindeer Lake", false, []string{"Canada"}},
	lakeInfo{24, "Issyk-Kul", true, []string{"Kyrgyzstan"}},
	lakeInfo{25, "Urmia", true, []string{"Iran"}},
	lakeInfo{26, "Vanern", false, []string{"Sweden"}},
	lakeInfo{27, "Winnipegosis", false, []string{"Canada"}},
	lakeInfo{28, "Albert", false, []string{"Uganda", "Democratic Republic of the Congo"}},
	lakeInfo{29, "Mweru", false, []string{"Zambia"}},
	lakeInfo{30, "Nettilling", false, []string{"Canada"}},
	lakeInfo{31, "Nipigon", false, []string{"Canada"}},
	lakeInfo{32, "Manitoba", false, []string{"Canada"}},
	lakeInfo{33, "Taymyr", false, []string{"Russia"}},
	lakeInfo{34, "Qinghai Lake", true, []string{"China"}},
	lakeInfo{35, "Saimaa", false, []string{"Finland"}},
	lakeInfo{36, "Lake of the Woods", false, []string{"Canada", "United States"}},
	lakeInfo{37, "Khanka", false, []string{"China"}},
	lakeInfo{38, "Sarygamyish", false, []string{"Uzbekistan", "Turkmenistan"}},
	lakeInfo{39, "Dubawnt", false, []string{"Canada"}},
	lakeInfo{40, "Van", true, []string{"Turkey"}},
	lakeInfo{41, "Peipus", false, []string{"Russia", "Estonia"}},
	lakeInfo{42, "Uvs", true, []string{"Mongolia"}},
	lakeInfo{43, "Poyang", false, []string{"China"}},
	lakeInfo{44, "Tana", false, []string{"Ethiopia"}},
	lakeInfo{45, "Amadjuak", false, []string{"Canada"}},
	lakeInfo{46, "Melville", true, []string{"Canada"}},
}

type lakeQuiz func([]lakeInfo) promptAndResponse

func quizLakes(cmd *cobra.Command, args []string) {
	quizzes := []lakeQuiz{
		quizLakeBySizeRank,
		quizSizeByLake,
		quizLakeSalinity,
		quizLakeInCountry,
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

func quizLakeInCountry(lakes []lakeInfo) promptAndResponse {
	lake1 := randomLake(lakes)
	country := lake1.countries[rand.Intn(len(lake1.countries))]
	lake2 := randomLake(lakes)
	return promptAndResponse{fmt.Sprintf("Lake %s touches %s, true or false?", lake2.name, country), strconv.FormatBool(isStringInSlice(country, lake2.countries))}
}

func init() {
	memoryquizCmd.AddCommand(lakesCmd)
}
