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
	region     string
	currency   string
}

var countries = []countryInfo{
	countryInfo{1, "Russia", "Moscow", "", "ruble"},
	countryInfo{2, "Canada", "Ottawa", "", "dollar"},
	countryInfo{3, "United States", "Washington, D.C.", "", "US dollar"},
	countryInfo{4, "China", "Beijing", "", "renminbi"},
	countryInfo{5, "Brazil", "Brasilia", "", "real"},
	countryInfo{6, "Australia", "Canberra", "", "Australian dollar"},
	countryInfo{7, "India", "New Delhi", "", "rupee"},
	countryInfo{8, "Argentina", "Buenos Aires", "", "peso"},
	countryInfo{9, "Kazakhstan", "Nur-Sultan", "", "tenge"},
	countryInfo{10, "Algeria", "Algiers", "Africa", "dinar"},
	countryInfo{11, "Democratic Republic of Congo", "Kinshasa", "Africa", "franc"},
	countryInfo{12, "Saudi Arabia", "Riyadh", "", "riyal"},
	countryInfo{13, "Mexico", "Mexico City", "", "peso"},
	countryInfo{14, "Indonesia", "Jakarta", "", "rupiah"},
	countryInfo{15, "Sudan", "Khartoum", "Africa", "pound"},
	countryInfo{16, "Libya", "Tripoli", "Africa", "dinar"},
	countryInfo{17, "Iran", "Tehran", "", "rial"},
	countryInfo{18, "Mongolia", "Ulaanbataar", "", "tugrik"},
	countryInfo{19, "Peru", "Lima", "", "sol"},
	countryInfo{20, "Chad", "N'Djamena", "Africa", "franc"},
	countryInfo{21, "Niger", "Niamey", "Africa", "franc"},
	countryInfo{22, "Angola", "Luanda", "Africa", "kwanza"},
	countryInfo{23, "Mali", "Bamako", "Africa", "franc"},
	countryInfo{24, "South Africa", "Pretoria", "Africa", "rand"},
	countryInfo{25, "Colombia", "Bogota", "", "peso"},
	countryInfo{26, "Ethiopia", "Addis Ababa", "Africa", "birr"},
	countryInfo{27, "Bolivia", "Sucre", "", "boliviano"},
	countryInfo{28, "Mauritania", "Nouakchott", "Africa", "ouguiya"},
	countryInfo{29, "Egypt", "Cairo", "Africa", "pound"},
	countryInfo{30, "Tanzania", "Dodoma", "Africa", "shilling"},
	countryInfo{31, "Nigeria", "Abuja", "Africa", "naira"},
	countryInfo{32, "Venezuela", "Caracas", "", "bolivar"},
	countryInfo{33, "Pakistan", "Islamabad", "", "rupee"},
	countryInfo{34, "Namibia", "Windhoek", "Africa", "dollar"},
	countryInfo{35, "Mozambique", "Maputo", "Africa", "metical"},
	countryInfo{36, "Turkey", "Ankara", "", "lira"},
	countryInfo{37, "Chile", "Santiago", "", "peso"},
	countryInfo{38, "Zambia", "Lusaka", "Africa", "kwacha"},
	countryInfo{39, "Myanmar", "Naypyidaw", "", "kyat"},
	countryInfo{40, "Afghanistan", "Kabul", "", "afghani"},
	countryInfo{41, "South Sudan", "Juba", "Africa", "pound"},
	countryInfo{42, "France", "Paris", "", "euro"},
	countryInfo{43, "Somalia", "Mogadishu", "Africa", "shilling"},
	countryInfo{44, "Central African Republic", "Bangui", "Africa", "franc"},
	countryInfo{45, "Ukraine", "Kyiv", "", "hryvnia"},
	countryInfo{46, "Madagascar", "Antananarivo", "Africa", "ariary"},
	countryInfo{47, "Botswana", "Gaborone", "Africa", "pula"},
	countryInfo{48, "Kenya", "Nairobi", "Africa", "shilling"},
	countryInfo{49, "Yemen", "Sana'a", "", "rial"},
	countryInfo{50, "Thailand", "Bangkok", "", "baht"},
	countryInfo{51, "Spain", "Madrid", "", "euro"},
	countryInfo{52, "Turkmenistan", "Ashgabat", "", "manat"},
	countryInfo{53, "Cameroon", "Yaounde", "Africa", "franc"},
	countryInfo{54, "Papua New Guinea", "Port Moresby", "", "kina"},
	countryInfo{55, "Sweden", "Stockholm", "", "krona"},
	countryInfo{56, "Uzbekistan", "Tashkent", "", "som"},
	countryInfo{57, "Morocco", "Rabat", "Africa", "dirham"},
	countryInfo{58, "Iraq", "Baghdad", "", "dinar"},
	countryInfo{59, "Paraguay", "Asuncion", "", "guarani"},
	countryInfo{60, "Zimbabwe", "Harare", "Africa", "US dollar"},
	countryInfo{61, "Norway", "Oslo", "", "krone"},
	countryInfo{62, "Japan", "Tokyo", "", "yen"},
	countryInfo{63, "Germany", "Berlin", "", "euro"},
	countryInfo{64, "Republic of the Congo", "Brazzaville", "Africa", "franc"},
	countryInfo{65, "Finland", "Helsinki", "", "euro"},
	countryInfo{66, "Vietnam", "Hanoi", "", "dong"},
	countryInfo{67, "Malaysia", "Kuala Lumpur", "", "ringgit"},
	countryInfo{68, "Ivory Coast", "Yamoussoukro", "Africa", "franc"},
	countryInfo{69, "Poland", "Warsaw", "", "zloty"},
	countryInfo{70, "Oman", "Muscat", "", "rial"},
	countryInfo{71, "Italy", "Rome", "", "euro"},
	countryInfo{72, "Phillipines", "Manila", "", "peso"},
	countryInfo{73, "Ecuador", "Quito", "", "US dollar"},
	countryInfo{74, "Burkina Faso", "Ouagadougou", "Africa", "franc"},
	countryInfo{75, "New Zealand", "Wellington", "", "dollar"},
	countryInfo{76, "Gabon", "Libreville", "Africa", "franc"},
	countryInfo{77, "Guinea", "Conakry", "Africa", "franc"},
	countryInfo{78, "United Kingdom", "London", "", "pound"},
	countryInfo{79, "Uganda", "Kampala", "Africa", "shilling"},
	countryInfo{80, "Ghana", "Accra", "Africa", "cedi"},
	countryInfo{81, "Romania", "Bucharest", "", "leu"},
	countryInfo{82, "Laos", "Vientiane", "", "kip"},
	countryInfo{83, "Guyana", "Georgetown", "", "dollar"},
	countryInfo{84, "Belarus", "Minsk", "", "ruble"},
	countryInfo{85, "Kyrgyzstan", "Bishkek", "", "som"},
	countryInfo{86, "Senegal", "Dakar", "Africa", "franc"},
	countryInfo{87, "Syria", "Damascus", "", "pound"},
	countryInfo{88, "Cambodia", "Phnom Penh", "", "riel"},
	countryInfo{89, "Uruguay", "Montevideo", "", "peso"},
	countryInfo{90, "Suriname", "Paramaribo", "", "dollar"},
	countryInfo{91, "Tunisia", "Tunis", "Africa", "dinar"},
	countryInfo{92, "Bangladesh", "Dhaka", "", "taka"},
	countryInfo{93, "Nepal", "Kathmandu", "", "rupee"},
	countryInfo{94, "Tajikistan", "Dusharbe", "", "somoni"},
	countryInfo{95, "Greece", "Athens", "", "euro"},
	countryInfo{96, "Nicaragua", "Managua", "", "cordoba"},
	countryInfo{97, "North Korea", "Pyongyang", "", "won"},
	countryInfo{98, "Malawi", "Lilongwe", "Africa", "kwacha"},
	countryInfo{99, "Eritrea", "Asmara", "Africa", "nakfa"},
	countryInfo{100, "Benin", "Porto-Novo", "Africa", "franc"},
	countryInfo{101, "Honduras", "Tegucigalpa", "", "lempira"},
	countryInfo{102, "Liberia", "Monrovia", "Africa", "dollar"},
	countryInfo{103, "Bulgaria", "Sofia", "", "lev"},
	countryInfo{104, "Cuba", "Havana", "", "peso"},
	countryInfo{105, "Guatemala", "Guatemala City", "", "quetzal"},
	countryInfo{106, "Iceland", "Reykjavik", "", "krona"},
	countryInfo{107, "South Korea", "Seoul", "", "won"},
	countryInfo{108, "Hungary", "Budapest", "", "forint"},
	countryInfo{109, "Portugal", "Lisbon", "", "euro"},
	countryInfo{110, "Jordan", "Amman", "", "dinar"},
	countryInfo{111, "Serbia", "Belgrade", "", "dinar"},
	countryInfo{112, "Azerbaijan", "Baku", "", "manat"},
	countryInfo{113, "Austria", "Vienna", "", "euro"},
	countryInfo{114, "United Arab Emirates", "Abu Dhabi", "", "dirham"},
	countryInfo{115, "Czech Republic", "Prague", "", "koruna"},
	countryInfo{116, "Panama", "Panama City", "", "US dollar"},
	countryInfo{117, "Sierra Leone", "Freetown", "Africa", "leone"},
	countryInfo{118, "Ireland", "Dublin", "", "euro"},
	countryInfo{119, "Georgia", "Tbilisi", "", "lari"},
	countryInfo{120, "Sri Lanka", "Sri Jayawardenepura Kotte", "", "rupee"},
	countryInfo{121, "Lithuania", "Vilnius", "", "euro"},
	countryInfo{122, "Latvia", "Riga", "", "euro"},
	countryInfo{123, "Togo", "Lome", "Africa", "franc"},
	countryInfo{124, "Croatia", "Zagreb", "", "kona"},
	countryInfo{125, "Bosnia and Herzegovina", "Sarajevo", "", "mark"},
	countryInfo{126, "Costa Rica", "San Jose", "", "colon"},
	countryInfo{127, "Slovakia", "Bratislava", "", "euro"},
	countryInfo{128, "Dominican Republic", "Santo Domingo", "", "peso"},
	countryInfo{129, "Estonia", "Tallinn", "", "kroon"},
	countryInfo{130, "Denmark", "Copenhagen", "", "krone"},
	countryInfo{131, "Netherlands", "Amsterdam", "", "euro"},
	countryInfo{132, "Switzerland", "Bern", "", "Swiss franc"},
	countryInfo{133, "Bhutan", "Thimphu", "", "ngultrum"},
	countryInfo{134, "Guinea-Bissau", "Bissau", "Africa", "franc"},
	countryInfo{135, "Moldova", "Kishinev", "", "leu"},
	countryInfo{136, "Belgium", "Brussels", "", "euro"},
	countryInfo{137, "Lesotho", "Maseru", "Africa", "loti"},
	countryInfo{138, "Armenia", "Yerevan", "", "dram"},
	countryInfo{139, "Solomon Islands", "Honiara", "", "dollar"},
	countryInfo{140, "Albania", "Tirana", "", "lek"},
	countryInfo{141, "Equatorial Guinea", "Malabo", "Africa", "franc"},
	countryInfo{142, "Burundi", "Gitega", "Africa", "franc"},
	countryInfo{143, "Haiti", "Port-au-Prince", "", "gourde"},
	countryInfo{144, "Rwanda", "Kigali", "Africa", "franc"},
	countryInfo{145, "North Macedonia", "Skopje", "", "denar"},
	countryInfo{146, "Djibouti", "Djibouti", "Africa", "franc"},
	countryInfo{147, "Belize", "Belmopan", "", "dollar"},
	countryInfo{148, "El Salvador", "San Salvador", "", "US dollar"},
	countryInfo{149, "Israel", "Jerusalem", "", "shekel"},
	countryInfo{150, "Slovenia", "Ljubljana", "", "euro"},
	countryInfo{151, "Fiji", "Suva", "", "dollar"},
	countryInfo{152, "Kuwait", "Kuwait City", "", "dinar"},
	countryInfo{153, "Eswatini", "Mbabane", "Africa", "rand"},
	countryInfo{154, "East Timor", "Dili", "", "timor-leste"},
	countryInfo{155, "The Bahamas", "Nassau", "", "dollar"},
	countryInfo{156, "Montenegro", "Podgorica", "", "euro"},
	countryInfo{157, "Vanuatu", "Port Vila", "", "vatu"},
	countryInfo{158, "Qatar", "Doha", "", "rial"},
	countryInfo{159, "The Gambia", "Banjul", "Africa", "dalasi"},
	countryInfo{160, "Jamaica", "Kingston", "", "dollar"},
	countryInfo{161, "Lebanon", "Beirut", "", "pound"},
	countryInfo{162, "Cyprus", "Nicosia", "", "euro"},
	countryInfo{163, "Brunei", "Bandar Seri Begawan", "", "dollar"},
	countryInfo{164, "Trinidad and Tobago", "Port of Spain", "", "dollar"},
	countryInfo{165, "Cape Verde", "Praia", "Africa", "escudo"},
	countryInfo{166, "Samoa", "Apia", "", "tala"},
	countryInfo{167, "Luxembourg", "Luxembourg City", "", "euro"},
	countryInfo{168, "Mauritius", "Port Louis", "Africa", "rupee"},
	countryInfo{169, "Comoros", "Moroni", "Africa", "franc"},
	countryInfo{170, "Sao Tome and Principe", "Sao Tome", "Africa", "dobra"},
	countryInfo{171, "Kiribati", "South Tarawa", "", "Australian dollar"},
	countryInfo{172, "Bahrain", "Manama", "", "dinar"},
	countryInfo{173, "Dominica", "Roseau", "", "Eastern Caribbean dollar"},
	countryInfo{174, "Tonga", "Nuku'alofa", "", "pa'anga"},
	countryInfo{175, "Singapore", "Singapore", "", "dollar"},
	countryInfo{176, "Federated States of Micronesia", "Palikir", "", "US dollar"},
	countryInfo{177, "Saint Lucia", "Castries", "", "Eastern Caribbean dollar"},
	countryInfo{178, "Andorra", "Andorra la Vella", "", "euro"},
	countryInfo{179, "Palau", "Ngerulmud", "", "US dollar"},
	countryInfo{180, "Seychelles", "Victoria", "Africa", "rupee"},
	countryInfo{181, "Antigua and Barbuda", "St. John's", "", "Eastern Caribbean dollar"},
	countryInfo{182, "Barbados", "Bridgetown", "", "dollar"},
	countryInfo{183, "Saint Vincent and the Grenadines", "Kingstown", "", "Eastern Caribbean dollar"},
	countryInfo{184, "Grenada", "St. George's", "", "Eastern Caribbean dollar"},
	countryInfo{185, "Malta", "Valletta", "", "euro"},
	countryInfo{186, "Maldives", "Male", "", "rufiyaa"},
	countryInfo{187, "Saint Kitts and Nevis", "Basseterre", "", "Eastern Caribbean dollar"},
	countryInfo{188, "Marshall Islands", "Majuro", "", "US dollar"},
	countryInfo{189, "Liechtenstein", "Vaduz", "", "Swiss franc"},
	countryInfo{190, "San Marino", "San Marino", "", "euro"},
	countryInfo{191, "Tuvalu", "Funafuti", "", "US dollar"},
	countryInfo{192, "Nauru", "Yaren", "", "Australian dollar"},
	countryInfo{193, "Monaco", "Monaco", "", "euro"},
	countryInfo{194, "Vatican City", "Vatican City", "", "euro"},
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
		quizCurrencyFromCountry,
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

func quizCurrencyFromCountry(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return promptAndResponse{fmt.Sprintf("What is the currency of %s?", country.name), country.currency}
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
