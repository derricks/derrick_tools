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
	countryInfo{61, "Norway", "Oslo"},
	countryInfo{62, "Japan", "Tokyo"},
	countryInfo{63, "Germany", "Berlin"},
	countryInfo{64, "Republic of the Congo", "Brazzaville"},
	countryInfo{65, "Finland", "Helsinki"},
	countryInfo{66, "Vietnam", "Hanoi"},
	countryInfo{67, "Malaysia", "Kuala Lumpur"},
	countryInfo{68, "Ivory Coast", "Yamoussoukro"},
	countryInfo{69, "Poland", "Warsaw"},
	countryInfo{70, "Oman", "Muscat"},
	countryInfo{71, "Italy", "Rome"},
	countryInfo{72, "Phillipines", "Manila"},
	countryInfo{73, "Ecuador", "Quito"},
	countryInfo{74, "Burkina Faso", "Ouagadougou"},
	countryInfo{75, "New Zealand", "Wellington"},
	countryInfo{76, "Gabon", "Libreville"},
	countryInfo{77, "Guinea", "Conakry"},
	countryInfo{78, "United Kingdom", "London"},
	countryInfo{79, "Uganda", "Kampala"},
	countryInfo{80, "Ghana", "Accra"},
	countryInfo{81, "Romania", "Bucharest"},
	countryInfo{82, "Laos", "Vientiane"},
	countryInfo{83, "Guyana", "Georgetown"},
	countryInfo{84, "Belarus", "Minsk"},
	countryInfo{85, "Kyrgyzstan", "Bishkek"},
	countryInfo{86, "Senegal", "Dakar"},
	countryInfo{87, "Syria", "Damascus"},
	countryInfo{88, "Cambodia", "Phnom Penh"},
	countryInfo{89, "Uruguay", "Montevideo"},
	countryInfo{90, "Suriname", "Paramaribo"},
	countryInfo{91, "Tunisia", "Tunis"},
	countryInfo{92, "Bangladesh", "Dhaka"},
	countryInfo{93, "Nepal", "Kathmandu"},
	countryInfo{94, "Tajikistan", "Dusharbe"},
	countryInfo{95, "Greece", "Athens"},
	countryInfo{96, "Nicaragua", "Managua"},
	countryInfo{97, "North Korea", "Pyongyang"},
	countryInfo{98, "Malawi", "Lilongwe"},
	countryInfo{99, "Eritrea", "Asmara"},
	countryInfo{100, "Benin", "Porto-Novo"},
	countryInfo{101, "Honduras", "Tegucigalpa"},
	countryInfo{102, "Liberia", "Monrovia"},
	countryInfo{103, "Bulgaria", "Sofia"},
	countryInfo{104, "Cuba", "Havana"},
	countryInfo{105, "Guatemala", "Guatemala City"},
	countryInfo{106, "Iceland", "Reykjavik"},
	countryInfo{107, "South Korea", "Seoul"},
	countryInfo{108, "Hungary", "Budapest"},
	countryInfo{109, "Portugal", "Lisbon"},
	countryInfo{110, "Jordan", "Amman"},
	countryInfo{111, "Serbia", "Belgrade"},
	countryInfo{112, "Azerbaijan", "Baku"},
	countryInfo{113, "Austria", "Vienna"},
	countryInfo{114, "United Arab Emirates", "Abu Dhabi"},
	countryInfo{115, "Czech Republic", "Prague"},
	countryInfo{116, "Panama", "Panama City"},
	countryInfo{117, "Sierra Leone", "Freetown"},
	countryInfo{118, "Ireland", "Dublin"},
	countryInfo{119, "Georgia", "Tbilisi"},
	countryInfo{120, "Sri Lanka", "Sri Jayawardenepura Kotte"},
	countryInfo{121, "Lithuania", "Vilnius"},
	countryInfo{122, "Latvia", "Riga"},
	countryInfo{123, "Togo", "Lome"},
	countryInfo{124, "Croatia", "Zagreb"},
	countryInfo{125, "Bosnia and Herzegovina", "Sarajevo"},
	countryInfo{126, "Costa Rica", "San Jose"},
	countryInfo{127, "Slovakia", "Bratislava"},
	countryInfo{128, "Dominican Republic", "Santo Domingo"},
	countryInfo{129, "Estonia", "Tallinn"},
	countryInfo{130, "Denmark", "Copenhagen"},
	countryInfo{131, "Netherlands", "Amsterdam"},
	countryInfo{132, "Switzerland", "Bern"},
	countryInfo{133, "Bhutan", "Thimphu"},
	countryInfo{134, "Guinea-Bissau", "Bissau"},
	countryInfo{135, "Moldova", "Kishinev"},
	countryInfo{136, "Belgium", "Brussels"},
	countryInfo{137, "Lesotho", "Maseru"},
	countryInfo{138, "Armenia", "Yerevan"},
	countryInfo{139, "Solomon Islands", "Honiara"},
	countryInfo{140, "Albania", "Tirana"},
	countryInfo{141, "Equatorial Guinea", "Malabo"},
	countryInfo{142, "Burundi", "Gitega"},
	countryInfo{143, "Haiti", "Port-au-Prince"},
	countryInfo{144, "Rwanda", "Kigali"},
	countryInfo{145, "North Macedonia", "Skopje"},
	countryInfo{146, "Djibouti", "Djibouti"},
	countryInfo{147, "Belize", "Belmopan"},
	countryInfo{148, "El Salvador", "San Salvador"},
	countryInfo{149, "Israel", "Jerusalem"},
	countryInfo{150, "Slovenia", "Ljubjiana"},
	countryInfo{151, "Fiji", "Suva"},
	countryInfo{152, "Kuwait", "Kuwait City"},
	countryInfo{153, "Eswatini", "Mbabane"},
	countryInfo{154, "East Timor", "Dili"},
	countryInfo{155, "The Bahamas", "Nassau"},
	countryInfo{156, "Montenegro", "Podgorica"},
	countryInfo{157, "Vanuatu", "Port Vila"},
	countryInfo{158, "Qatar", "Doha"},
	countryInfo{159, "The Gambia", "Banjul"},
	countryInfo{160, "Jamaica", "Kingston"},
	countryInfo{161, "Lebanon", "Beirut"},
	countryInfo{162, "Cyprus", "Nicosia"},
	countryInfo{163, "Brunei", "Bandar Seri Begawan"},
	countryInfo{164, "Trinidad and Tobago", "Port of Spain"},
	countryInfo{165, "Cape Verde", "Praia"},
	countryInfo{166, "Samoa", "Apia"},
	countryInfo{167, "Luxembourg", "Luxembourg City"},
	countryInfo{168, "Mauritius", "Port Louis"},
	countryInfo{169, "Comoros", "Moroni"},
	countryInfo{170, "Sao Tome and Principe", "Sao Tome"},
	countryInfo{171, "Kiribati", "South Tarawa"},
	countryInfo{172, "Bahrain", "Manama"},
	countryInfo{173, "Dominica", "Roseau"},
	countryInfo{174, "Tonga", "Nuku'alofa"},
	countryInfo{175, "Singapore", "Singapore"},
	countryInfo{176, "Federated States of Micronesia", "Palikir"},
	countryInfo{177, "Saint Lucia", "Castries"},
	countryInfo{178, "Andorra", "Andorra la Vella"},
	countryInfo{179, "Palau", "Ngerulmud"},
	countryInfo{180, "Seychelles", "Victoria"},
	countryInfo{181, "Antigua and Barbuda", "St. John's"},
	countryInfo{182, "Barbados", "Bridgetown"},
	countryInfo{183, "Saint Vincent and the Grenadines", "Kingstown"},
	countryInfo{184, "Grenada", "St. George's"},
	countryInfo{185, "Malta", "Valletta"},
	countryInfo{186, "Maldives", "Male"},
	countryInfo{187, "Saint Kitts and Nevis", "Basseterre"},
	countryInfo{188, "Marshall Islands", "Majuro"},
	countryInfo{189, "Liechtenstein", "Vaduz"},
	countryInfo{190, "San Marino", "San Marino"},
	countryInfo{191, "Tuvalu", "Funafuti"},
	countryInfo{192, "Nauru", "Yaren"},
	countryInfo{193, "Monaco", "Monaco"},
	countryInfo{194, "Vatican City", "Vatican City"},
}

type countryQuery func([]countryInfo) promptAndResponse

func quizCountries(cmd *cobra.Command, args []string) {
	quizFuncs := []countryQuery{
		quizCountryCapital,
		quizCountryFromCapital,
		quizCountryRankInArea,
		quizCountryFromRankInArea,
		quizCapitalFromRankInArea,
		quizWhichIsBigger,
		quizWhichIsSmaller,
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
	return promptAndResponse{fmt.Sprintf("Which country's capital is %s?", country.capital), country.name}
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

func quizWhichIsBigger(countries []countryInfo) promptAndResponse {
	country1 := randomCountry(countries)
	country2 := randomCountry(countries)

	for country1.name == country2.name {
		country2 = randomCountry(countries)
	}

	response := promptAndResponse{fmt.Sprintf("Which country is bigger: %s or %s?", country1.name, country2.name), ""}
	// rank is inverse to size: 1 means the biggest country, and 50 is bigger than 70
	if country1.rankInArea < country2.rankInArea {
		response.response = country1.name
	} else {
		response.response = country2.name
	}
	return response
}

func quizWhichIsSmaller(countries []countryInfo) promptAndResponse {
	country1 := randomCountry(countries)
	country2 := randomCountry(countries)

	for country1.name == country2.name {
		country2 = randomCountry(countries)
	}

	response := promptAndResponse{fmt.Sprintf("Which country is smaller: %s or %s?", country1.name, country2.name), ""}
	// rank is inverse to size: 1 means the biggest country, and 50 is bigger than 70
	if country1.rankInArea > country2.rankInArea {
		response.response = country1.name
	} else {
		response.response = country2.name
	}
	return response
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
