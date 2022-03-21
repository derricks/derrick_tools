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
	rankInArea  int
	name        string
	capital     string
	region      string
	currency    string
	countryCode string
}

func (c countryInfo) flagEmoji() string {
	base := rune(0x1f1e6)
	return string(base+rune(c.countryCode[0])-'A') + string(base+rune(c.countryCode[1])-'A')
}

var countries = []countryInfo{
	countryInfo{1, "Russia", "Moscow", "", "ruble", "RU"},
	countryInfo{2, "Canada", "Ottawa", "", "dollar", "CA"},
	countryInfo{3, "United States", "Washington, D.C.", "", "US dollar", "US"},
	countryInfo{4, "China", "Beijing", "", "renminbi", "CN"},
	countryInfo{5, "Brazil", "Brasilia", "", "real", "BR"},
	countryInfo{6, "Australia", "Canberra", "", "Australian dollar", "AU"},
	countryInfo{7, "India", "New Delhi", "", "rupee", "IN"},
	countryInfo{8, "Argentina", "Buenos Aires", "", "peso", "AR"},
	countryInfo{9, "Kazakhstan", "Nur-Sultan", "", "tenge", "KZ"},
	countryInfo{10, "Algeria", "Algiers", "Africa", "dinar", "DZ"},
	countryInfo{11, "Democratic Republic of Congo", "Kinshasa", "Africa", "franc", "CD"},
	countryInfo{12, "Saudi Arabia", "Riyadh", "", "riyal", "SA"},
	countryInfo{13, "Mexico", "Mexico City", "", "peso", "MX"},
	countryInfo{14, "Indonesia", "Jakarta", "", "rupiah", "ID"},
	countryInfo{15, "Sudan", "Khartoum", "Africa", "pound", "SD"},
	countryInfo{16, "Libya", "Tripoli", "Africa", "dinar", "LY"},
	countryInfo{17, "Iran", "Tehran", "", "rial", "IR"},
	countryInfo{18, "Mongolia", "Ulaanbataar", "", "tugrik", "MN"},
	countryInfo{19, "Peru", "Lima", "", "sol", "PE"},
	countryInfo{20, "Chad", "N'Djamena", "Africa", "franc", "TD"},
	countryInfo{21, "Niger", "Niamey", "Africa", "franc", "NE"},
	countryInfo{22, "Angola", "Luanda", "Africa", "kwanza", "AO"},
	countryInfo{23, "Mali", "Bamako", "Africa", "franc", "ML"},
	countryInfo{24, "South Africa", "Pretoria", "Africa", "rand", "ZA"},
	countryInfo{25, "Colombia", "Bogota", "", "peso", "CO"},
	countryInfo{26, "Ethiopia", "Addis Ababa", "Africa", "birr", "ET"},
	countryInfo{27, "Bolivia", "Sucre", "", "boliviano", "BO"},
	countryInfo{28, "Mauritania", "Nouakchott", "Africa", "ouguiya", "MR"},
	countryInfo{29, "Egypt", "Cairo", "Africa", "pound", "EG"},
	countryInfo{30, "Tanzania", "Dodoma", "Africa", "shilling", "TZ"},
	countryInfo{31, "Nigeria", "Abuja", "Africa", "naira", "NG"},
	countryInfo{32, "Venezuela", "Caracas", "", "bolivar", "VE"},
	countryInfo{33, "Pakistan", "Islamabad", "", "rupee", "PK"},
	countryInfo{34, "Namibia", "Windhoek", "Africa", "dollar", "NA"},
	countryInfo{35, "Mozambique", "Maputo", "Africa", "metical", "MZ"},
	countryInfo{36, "Turkey", "Ankara", "", "lira", "TR"},
	countryInfo{37, "Chile", "Santiago", "", "peso", "CL"},
	countryInfo{38, "Zambia", "Lusaka", "Africa", "kwacha", "ZM"},
	countryInfo{39, "Myanmar", "Naypyidaw", "", "kyat", "MM"},
	countryInfo{40, "Afghanistan", "Kabul", "", "afghani", "AF"},
	countryInfo{41, "South Sudan", "Juba", "Africa", "pound", "SS"},
	countryInfo{42, "France", "Paris", "", "euro", "FR"},
	countryInfo{43, "Somalia", "Mogadishu", "Africa", "shilling", "SO"},
	countryInfo{44, "Central African Republic", "Bangui", "Africa", "franc", "CF"},
	countryInfo{45, "Ukraine", "Kyiv", "", "hryvnia", "UA"},
	countryInfo{46, "Madagascar", "Antananarivo", "Africa", "ariary", "MG"},
	countryInfo{47, "Botswana", "Gaborone", "Africa", "pula", "BW"},
	countryInfo{48, "Kenya", "Nairobi", "Africa", "shilling", "KE"},
	countryInfo{49, "Yemen", "Sana'a", "", "rial", "YE"},
	countryInfo{50, "Thailand", "Bangkok", "", "baht", "TH"},
	countryInfo{51, "Spain", "Madrid", "", "euro", "ES"},
	countryInfo{52, "Turkmenistan", "Ashgabat", "", "manat", "TM"},
	countryInfo{53, "Cameroon", "Yaounde", "Africa", "franc", "CM"},
	countryInfo{54, "Papua New Guinea", "Port Moresby", "", "kina", "PG"},
	countryInfo{55, "Sweden", "Stockholm", "", "krona", "SE"},
	countryInfo{56, "Uzbekistan", "Tashkent", "", "som", "UZ"},
	countryInfo{57, "Morocco", "Rabat", "Africa", "dirham", "MA"},
	countryInfo{58, "Iraq", "Baghdad", "", "dinar", "IQ"},
	countryInfo{59, "Paraguay", "Asuncion", "", "guarani", "PY"},
	countryInfo{60, "Zimbabwe", "Harare", "Africa", "US dollar", "ZW"},
	countryInfo{61, "Norway", "Oslo", "", "krone", "NO"},
	countryInfo{62, "Japan", "Tokyo", "", "yen", "JP"},
	countryInfo{63, "Germany", "Berlin", "", "euro", "DE"},
	countryInfo{64, "Republic of the Congo", "Brazzaville", "Africa", "franc", "CG"},
	countryInfo{65, "Finland", "Helsinki", "", "euro", "FI"},
	countryInfo{66, "Viet Nam", "Hanoi", "", "dong", "VN"},
	countryInfo{67, "Malaysia", "Kuala Lumpur", "", "ringgit", "MY"},
	countryInfo{68, "Ivory Coast", "Yamoussoukro", "Africa", "franc", "CI"},
	countryInfo{69, "Poland", "Warsaw", "", "zloty", "PL"},
	countryInfo{70, "Oman", "Muscat", "", "rial", "OM"},
	countryInfo{71, "Italy", "Rome", "", "euro", "IT"},
	countryInfo{72, "Phillipines", "Manila", "", "peso", "PH"},
	countryInfo{73, "Ecuador", "Quito", "", "US dollar", "EC"},
	countryInfo{74, "Burkina Faso", "Ouagadougou", "Africa", "franc", "BF"},
	countryInfo{75, "New Zealand", "Wellington", "", "dollar", "NZ"},
	countryInfo{76, "Gabon", "Libreville", "Africa", "franc", "GA"},
	countryInfo{77, "Guinea", "Conakry", "Africa", "franc", "GN"},
	countryInfo{78, "United Kingdom", "London", "", "pound", "GB"},
	countryInfo{79, "Uganda", "Kampala", "Africa", "shilling", "UG"},
	countryInfo{80, "Ghana", "Accra", "Africa", "cedi", "GH"},
	countryInfo{81, "Romania", "Bucharest", "", "leu", "RO"},
	countryInfo{82, "Laos", "Vientiane", "", "kip", "LA"},
	countryInfo{83, "Guyana", "Georgetown", "", "dollar", "GY"},
	countryInfo{84, "Belarus", "Minsk", "", "ruble", "BY"},
	countryInfo{85, "Kyrgyzstan", "Bishkek", "", "som", "KG"},
	countryInfo{86, "Senegal", "Dakar", "Africa", "franc", "SN"},
	countryInfo{87, "Syria", "Damascus", "", "pound", "SY"},
	countryInfo{88, "Cambodia", "Phnom Penh", "", "riel", "KH"},
	countryInfo{89, "Uruguay", "Montevideo", "", "peso", "UY"},
	countryInfo{90, "Suriname", "Paramaribo", "", "dollar", "SR"},
	countryInfo{91, "Tunisia", "Tunis", "Africa", "dinar", "TN"},
	countryInfo{92, "Bangladesh", "Dhaka", "", "taka", "BD"},
	countryInfo{93, "Nepal", "Kathmandu", "", "rupee", "NP"},
	countryInfo{94, "Tajikistan", "Dusharbe", "", "somoni", "TJ"},
	countryInfo{95, "Greece", "Athens", "", "euro", "GR"},
	countryInfo{96, "Nicaragua", "Managua", "", "cordoba", "NI"},
	countryInfo{97, "North Korea", "Pyongyang", "", "won", "KP"},
	countryInfo{98, "Malawi", "Lilongwe", "Africa", "kwacha", "MW"},
	countryInfo{99, "Eritrea", "Asmara", "Africa", "nakfa", "ER"},
	countryInfo{100, "Benin", "Porto-Novo", "Africa", "franc", "BJ"},
	countryInfo{101, "Honduras", "Tegucigalpa", "", "lempira", "HN"},
	countryInfo{102, "Liberia", "Monrovia", "Africa", "dollar", "LR"},
	countryInfo{103, "Bulgaria", "Sofia", "", "lev", "BG"},
	countryInfo{104, "Cuba", "Havana", "", "peso", "CU"},
	countryInfo{105, "Guatemala", "Guatemala City", "", "quetzal", "GT"},
	countryInfo{106, "Iceland", "Reykjavik", "", "krona", "IS"},
	countryInfo{107, "South Korea", "Seoul", "", "won", "KR"},
	countryInfo{108, "Hungary", "Budapest", "", "forint", "HU"},
	countryInfo{109, "Portugal", "Lisbon", "", "euro", "PT"},
	countryInfo{110, "Jordan", "Amman", "", "dinar", "JO"},
	countryInfo{111, "Serbia", "Belgrade", "", "dinar", "RS"},
	countryInfo{112, "Azerbaijan", "Baku", "", "manat", "AZ"},
	countryInfo{113, "Austria", "Vienna", "", "euro", "AT"},
	countryInfo{114, "United Arab Emirates", "Abu Dhabi", "", "dirham", "AE"},
	countryInfo{115, "Czech Republic", "Prague", "", "koruna", "CZ"},
	countryInfo{116, "Panama", "Panama City", "", "US dollar", "PA"},
	countryInfo{117, "Sierra Leone", "Freetown", "Africa", "leone", "SL"},
	countryInfo{118, "Ireland", "Dublin", "", "euro", "IE"},
	countryInfo{119, "Georgia", "Tbilisi", "", "lari", "GE"},
	countryInfo{120, "Sri Lanka", "Sri Jayawardenepura Kotte", "", "rupee", "LK"},
	countryInfo{121, "Lithuania", "Vilnius", "", "euro", "LT"},
	countryInfo{122, "Latvia", "Riga", "", "euro", "LV"},
	countryInfo{123, "Togo", "Lome", "Africa", "franc", "TG"},
	countryInfo{124, "Croatia", "Zagreb", "", "kona", "HR"},
	countryInfo{125, "Bosnia and Herzegovina", "Sarajevo", "", "mark", "BA"},
	countryInfo{126, "Costa Rica", "San Jose", "", "colon", "CR"},
	countryInfo{127, "Slovakia", "Bratislava", "", "euro", "SK"},
	countryInfo{128, "Dominican Republic", "Santo Domingo", "", "peso", "DO"},
	countryInfo{129, "Estonia", "Tallinn", "", "kroon", "EE"},
	countryInfo{130, "Denmark", "Copenhagen", "", "krone", "DK"},
	countryInfo{131, "Netherlands", "Amsterdam", "", "euro", "NL"},
	countryInfo{132, "Switzerland", "Bern", "", "Swiss franc", "CH"},
	countryInfo{133, "Bhutan", "Thimphu", "", "ngultrum", "BT"},
	countryInfo{134, "Guinea-Bissau", "Bissau", "Africa", "franc", "GW"},
	countryInfo{135, "Moldova", "Kishinev", "", "leu", "MD"},
	countryInfo{136, "Belgium", "Brussels", "", "euro", "BE"},
	countryInfo{137, "Lesotho", "Maseru", "Africa", "loti", "LS"},
	countryInfo{138, "Armenia", "Yerevan", "", "dram", "AM"},
	countryInfo{139, "Solomon Islands", "Honiara", "", "dollar", "SB"},
	countryInfo{140, "Albania", "Tirana", "", "lek", "AL"},
	countryInfo{141, "Equatorial Guinea", "Malabo", "Africa", "franc", "GQ"},
	countryInfo{142, "Burundi", "Gitega", "Africa", "franc", "BI"},
	countryInfo{143, "Haiti", "Port-au-Prince", "", "gourde", "HT"},
	countryInfo{144, "Rwanda", "Kigali", "Africa", "franc", "RW"},
	countryInfo{145, "North Macedonia", "Skopje", "", "denar", "MK"},
	countryInfo{146, "Djibouti", "Djibouti", "Africa", "franc", "DJ"},
	countryInfo{147, "Belize", "Belmopan", "", "dollar", "BZ"},
	countryInfo{148, "El Salvador", "San Salvador", "", "US dollar", "SV"},
	countryInfo{149, "Israel", "Jerusalem", "", "shekel", "IL"},
	countryInfo{150, "Slovenia", "Ljubljana", "", "euro", "SI"},
	countryInfo{151, "Fiji", "Suva", "", "dollar", "FJ"},
	countryInfo{152, "Kuwait", "Kuwait City", "", "dinar", "KW"},
	countryInfo{153, "Eswatini", "Mbabane", "Africa", "rand", "SZ"},
	countryInfo{154, "East Timor", "Dili", "", "timor-leste", "TL"},
	countryInfo{155, "The Bahamas", "Nassau", "", "dollar", "BS"},
	countryInfo{156, "Montenegro", "Podgorica", "", "euro", "ME"},
	countryInfo{157, "Vanuatu", "Port Vila", "", "vatu", "VU"},
	countryInfo{158, "Qatar", "Doha", "", "rial", "QA"},
	countryInfo{159, "The Gambia", "Banjul", "Africa", "dalasi", "GM"},
	countryInfo{160, "Jamaica", "Kingston", "", "dollar", "JM"},
	countryInfo{161, "Lebanon", "Beirut", "", "pound", "LB"},
	countryInfo{162, "Cyprus", "Nicosia", "", "euro", "CY"},
	countryInfo{163, "Brunei", "Bandar Seri Begawan", "", "dollar", "BN"},
	countryInfo{164, "Trinidad and Tobago", "Port of Spain", "", "dollar", "TT"},
	countryInfo{165, "Cape Verde", "Praia", "Africa", "escudo", "CV"},
	countryInfo{166, "Samoa", "Apia", "", "tala", "WS"},
	countryInfo{167, "Luxembourg", "Luxembourg City", "", "euro", "LU"},
	countryInfo{168, "Mauritius", "Port Louis", "Africa", "rupee", "MU"},
	countryInfo{169, "Comoros", "Moroni", "Africa", "franc", "KM"},
	countryInfo{170, "Sao Tome and Principe", "Sao Tome", "Africa", "dobra", "ST"},
	countryInfo{171, "Kiribati", "South Tarawa", "", "Australian dollar", "KI"},
	countryInfo{172, "Bahrain", "Manama", "", "dinar", "BH"},
	countryInfo{173, "Dominica", "Roseau", "", "Eastern Caribbean dollar", "DM"},
	countryInfo{174, "Tonga", "Nuku'alofa", "", "pa'anga", "TO"},
	countryInfo{175, "Singapore", "Singapore", "", "dollar", "SG"},
	countryInfo{176, "Federated States of Micronesia", "Palikir", "", "US dollar", "FM"},
	countryInfo{177, "Saint Lucia", "Castries", "", "Eastern Caribbean dollar", "LC"},
	countryInfo{178, "Andorra", "Andorra la Vella", "", "euro", "AD"},
	countryInfo{179, "Palau", "Ngerulmud", "", "US dollar", "PW"},
	countryInfo{180, "Seychelles", "Victoria", "Africa", "rupee", "SC"},
	countryInfo{181, "Antigua and Barbuda", "St. John's", "", "Eastern Caribbean dollar", "AG"},
	countryInfo{182, "Barbados", "Bridgetown", "", "dollar", "BB"},
	countryInfo{183, "Saint Vincent and the Grenadines", "Kingstown", "", "Eastern Caribbean dollar", "VC"},
	countryInfo{184, "Grenada", "St. George's", "", "Eastern Caribbean dollar", "GD"},
	countryInfo{185, "Malta", "Valletta", "", "euro", "MT"},
	countryInfo{186, "Maldives", "Male", "", "rufiyaa", "MV"},
	countryInfo{187, "Saint Kitts and Nevis", "Basseterre", "", "Eastern Caribbean dollar", "KN"},
	countryInfo{188, "Marshall Islands", "Majuro", "", "US dollar", "MH"},
	countryInfo{189, "Liechtenstein", "Vaduz", "", "Swiss franc", "LI"},
	countryInfo{190, "San Marino", "San Marino", "", "euro", "SM"},
	countryInfo{191, "Tuvalu", "Funafuti", "", "US dollar", "TV"},
	countryInfo{192, "Nauru", "Yaren", "", "Australian dollar", "NR"},
	countryInfo{193, "Monaco", "Monaco", "", "euro", "MC"},
	countryInfo{194, "Vatican City", "Vatican City", "", "euro", "VA"},
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
		quizCountryFromFlag,
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

func quizCountryFromFlag(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return promptAndResponse{fmt.Sprintf("Which country has this flag: %s", country.flagEmoji()), country.name}
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
