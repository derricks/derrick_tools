/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"math/rand"
	"strconv"
)

// countriesCmd represents the countries command
var countriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "Memory quizzes about countries, including capitals and rank in area",
	Run:   quizCountries,
}

type countryInfo struct {
	rankInArea int
	name       string
	capital    string
}

var countries = []countryInfo{
	countryInfo{1, "Russia", "Moscow"},
	countryInfo{2, "Canada", "Ottawa"},
	countryInfo{3, "United States", "Washington, D.C."},
	countryInfo{4, "China", "Beijing"},
	countryInfo{5, "Brazil", "Brasilia"},
	countryInfo{6, "Australia", "Canberra"},
	countryInfo{7, "India", "New Delhi"},
	countryInfo{8, "Argentina", "Buenos Aires"},
	countryInfo{9, "Kazakhstan", "Nur-Sultan"},
	countryInfo{10, "Algeria", "Algiers"},
	countryInfo{11, "Democratic Republic of Congo", "Kinshasa"},
	countryInfo{12, "Saudi Arabia", "Riyadh"},
	countryInfo{13, "Mexico", "Mexico City"},
	countryInfo{14, "Indonesia", "Jakarta"},
	countryInfo{15, "Sudan", "Khartoum"},
	countryInfo{16, "Libya", "Tripoli"},
	countryInfo{17, "Iran", "Tehran"},
	countryInfo{18, "Mongolia", "Ulaanbataar"},
	countryInfo{19, "Peru", "Lima"},
	countryInfo{20, "Chad", "N'Djamena"},
	countryInfo{21, "Niger", "Niamey"},
	countryInfo{22, "Angola", "Luanda"},
	countryInfo{23, "Mali", "Bamako"},
	countryInfo{24, "South Africa", "Pretoria"},
	countryInfo{25, "Colombia", "Bogota"},
	countryInfo{26, "Ethiopia", "Addis Ababa"},
	countryInfo{27, "Bolivia", "Sucre"},
	countryInfo{28, "Mauritania", "Nouakchott"},
	countryInfo{29, "Egypt", "Cairo"},
	countryInfo{30, "Tanzania", "Dodoma"},
	countryInfo{31, "Nigeria", "Abuja"},
	countryInfo{32, "Venezuela", "Caracas"},
	countryInfo{33, "Pakistan", "Islamabad"},
	countryInfo{34, "Namibia", "Windhoek"},
	countryInfo{35, "Mozambique", "Maputo"},
	countryInfo{36, "Turkey", "Ankara"},
	countryInfo{37, "Chile", "Santiago"},
	countryInfo{38, "Zambia", "Lusaka"},
	countryInfo{39, "Myanmar", "Naypyidaw"},
	countryInfo{40, "Afghanistan", "Kabul"},
	countryInfo{41, "South Sudan", "Juba"},
	countryInfo{42, "France", "Paris"},
	countryInfo{43, "Somalia", "Mogadishu"},
	countryInfo{44, "Central African Republic", "Bangui"},
	countryInfo{45, "Ukraine", "Kyiv"},
	countryInfo{46, "Madagascar", "Antananarivo"},
	countryInfo{47, "Botswana", "Gaborone"},
	countryInfo{48, "Kenya", "Nairobi"},
	countryInfo{49, "Yemen", "Sana'a"},
	countryInfo{50, "Thailand", "Bangkok"},
	countryInfo{51, "Spain", "Madrid"},
	countryInfo{52, "Turkmenistan", "Ashgabat"},
	countryInfo{53, "Cameroon", "Yaounde"},
	countryInfo{54, "Papua New Guinea", "Port Moresby"},
	countryInfo{55, "Sweden", "Stockholm"},
	countryInfo{56, "Uzbekistan", "Tashkent"},
	countryInfo{57, "Morocco", "Rabat"},
	countryInfo{58, "Iraq", "Baghdad"},
	countryInfo{59, "Paraguay", "Asuncion"},
	countryInfo{60, "Zimbabwe", "Harare"},
}

type countryQuery func([]countryInfo) promptAndResponse

func quizCountries(cmd *cobra.Command, args []string) {
	quizFuncs := []countryQuery{
		quizCountryCapital,
		quizCountryFromCapital,
		quizCountryRankInArea,
		quizCountryFromRankInArea,
		quizCapitalFromRankInArea,
	}
	function := quizFuncs[rand.Intn(len(quizFuncs))]
	promptAndCheckResponse(function(countries))

}

func randomCountry(countries []countryInfo) countryInfo {
	return countries[rand.Intn(len(countries))]
}

func quizCountryCapital(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return promptAndResponse{fmt.Sprintf("What is the capital of %s?", country.name), country.capital}
}

func quizCountryFromCapital(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return promptAndResponse{fmt.Sprintf("Which country's capital is %?", country.capital), country.name}
}

func quizCountryRankInArea(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return promptAndResponse{fmt.Sprintf("Where does %s rank in terms of area?", country.name), strconv.Itoa(country.rankInArea)}
}

func quizCountryFromRankInArea(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return promptAndResponse{fmt.Sprintf("What country ranks number %d in area?", country.rankInArea), country.name}
}

func quizCapitalFromRankInArea(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return promptAndResponse{fmt.Sprintf("What is the capital of the country that ranks %d in area?", country.rankInArea), country.capital}
}

func init() {
	memoryquizCmd.AddCommand(countriesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// countriesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// countriesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
