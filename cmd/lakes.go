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
	{1, "Caspian Sea", true, []string{"Russia", "Kazakhstan", "Turkmenistan", "Iran", "Azerbaijan"}},
	{2, "Superior", false, []string{"Canada", "United States"}},
	{3, "Victoria", false, []string{"Kenya", "Tanzania", "Uganda"}},
	{4, "Huron", false, []string{"Canada", "United States"}},
	{5, "Michigan", false, []string{"United States"}},
	{6, "Tanganyika", false, []string{"Burundi", "Zambia", "Democratic Republic of the Congo", "Tanzania"}},
	{7, "Baikal", false, []string{"Russia"}},
	{8, "Great Bear Lake", false, []string{"Canada"}},
	{9, "Malawi", false, []string{"Mozambique", "Tanzania"}},
	{10, "Great Slave Lake", false, []string{"Canada"}},
	{11, "Erie", false, []string{"Canada", "United States"}},
	{12, "Winnipeg", false, []string{"Canada"}},
	{13, "Ontario", false, []string{"Canada", "United States"}},
	{14, "Ladoga", false, []string{"Russia"}},
	{15, "Balkhash", true, []string{"Kazakhstan"}},
	{16, "Bangweulu", false, []string{"Zambia"}},
	{17, "Vostok", false, []string{"Antarctica"}},
	{18, "Onega", false, []string{"Russia"}},
	{19, "Titicaca", false, []string{"Bolivia", "Peru"}},
	{20, "Nicaragua", false, []string{"Nicaragua"}},
	{21, "Athabasca", false, []string{"Canada"}},
	{22, "Turkana", true, []string{"Kenya", "Ethiopia"}},
	{23, "Reindeer Lake", false, []string{"Canada"}},
	{24, "Issyk-Kul", true, []string{"Kyrgyzstan"}},
	{25, "Urmia", true, []string{"Iran"}},
	{26, "Vanern", false, []string{"Sweden"}},
	{27, "Winnipegosis", false, []string{"Canada"}},
	{28, "Albert", false, []string{"Uganda", "Democratic Republic of the Congo"}},
	{29, "Mweru", false, []string{"Zambia"}},
	{30, "Nettilling", false, []string{"Canada"}},
	{31, "Nipigon", false, []string{"Canada"}},
	{32, "Manitoba", false, []string{"Canada"}},
	{33, "Taymyr", false, []string{"Russia"}},
	{34, "Qinghai Lake", true, []string{"China"}},
	{35, "Saimaa", false, []string{"Finland"}},
	{36, "Lake of the Woods", false, []string{"Canada", "United States"}},
	{37, "Khanka", false, []string{"China"}},
	{38, "Sarygamyish", false, []string{"Uzbekistan", "Turkmenistan"}},
	{39, "Dubawnt", false, []string{"Canada"}},
	{40, "Van", true, []string{"Turkey"}},
	{41, "Peipus", false, []string{"Russia", "Estonia"}},
	{42, "Uvs", true, []string{"Mongolia"}},
	{43, "Poyang", false, []string{"China"}},
	{44, "Tana", false, []string{"Ethiopia"}},
	{45, "Amadjuak", false, []string{"Canada"}},
	{46, "Melville", true, []string{"Canada"}},
}

type lakeQuiz func([]lakeInfo) promptAndResponse

func quizLakes(cmd *cobra.Command, args []string) {
	quizzes := []lakeQuiz{
		quizLakeBySizeRank,
		quizSizeByLake,
		quizLakeSalinity,
		quizLakeInCountry,
	}

	quiz := randomItemFromSlice(quizzes)
	promptAndCheckResponse(quiz(lakes))
}

func quizLakeBySizeRank(lakes []lakeInfo) promptAndResponse {
	lake := randomItemFromSlice(lakes)
	return promptAndResponse{fmt.Sprintf("What is the name of the lake at position %d", lake.sizeOrder), lake.name}
}

func quizSizeByLake(lakes []lakeInfo) promptAndResponse {
	lake := randomItemFromSlice(lakes)
	return promptAndResponse{fmt.Sprintf("What is the size rank of lake %s?", lake.name), strconv.Itoa(lake.sizeOrder)}
}

func quizLakeSalinity(lakes []lakeInfo) promptAndResponse {
	lake := randomItemFromSlice(lakes)
	return promptAndResponse{fmt.Sprintf("%s is a saline lake, true or false?", lake.name), strconv.FormatBool(lake.isSaline)}
}

func quizLakeInCountry(lakes []lakeInfo) promptAndResponse {
	lake1 := randomItemFromSlice(lakes)
	country := lake1.countries[rand.Intn(len(lake1.countries))]
	lake2 := randomItemFromSlice(lakes)
	return promptAndResponse{fmt.Sprintf("Lake %s touches %s, true or false?", lake2.name, country), strconv.FormatBool(isStringInSlice(country, lake2.countries))}
}

func init() {
	memoryquizCmd.AddCommand(lakesCmd)
}
