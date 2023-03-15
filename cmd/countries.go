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
)

// countriesCmd represents the countries command
var countriesCmd = &cobra.Command{
	Use:   "countries",
	Short: "Memory quizzes about countries, including capitals and rank in area",
	Run:   quizCountries,
}

type countryInfo struct {
	rankInArea  int    `crossquery:"all" crossqueryname:"size rank"`
	name        string `crossquery:"all"`
	capital     string `crossquery:"all"`
	region      []string
	currency    string `crossquery:"guess"`
	countryCode string
}

const (
	asia            string = "Asia"
	europe          string = "Europe"
	north_america   string = "North America"
	south_america   string = "South America"
	central_america string = "Central America"
	africa          string = "Africa"
	oceania         string = "Oceania"
	middle_east     string = "Middle East"
	caribbean       string = "Caribbean"
	south_asia      string = "South Asia"
	southeast_asia  string = "Southeast Asia"
)

func (c countryInfo) flagEmoji() string {
	base := rune(0x1f1e6)
	return string(base+rune(c.countryCode[0])-'A') + string(base+rune(c.countryCode[1])-'A')
}

var countries = []countryInfo{
	countryInfo{1, "Russia", "Moscow", []string{asia, europe}, "ruble", "RU"},
	countryInfo{2, "Canada", "Ottawa", []string{north_america}, "dollar", "CA"},
	countryInfo{3, "United States", "Washington, D.C.", []string{north_america}, "US dollar", "US"},
	countryInfo{4, "China", "Beijing", []string{asia}, "renminbi", "CN"},
	countryInfo{5, "Brazil", "Brasilia", []string{south_america}, "real", "BR"},
	countryInfo{6, "Australia", "Canberra", []string{oceania}, "Australian dollar", "AU"},
	countryInfo{7, "India", "New Delhi", []string{south_asia}, "rupee", "IN"},
	countryInfo{8, "Argentina", "Buenos Aires", []string{south_america}, "peso", "AR"},
	countryInfo{9, "Kazakhstan", "Nur-Sultan", []string{asia, europe}, "tenge", "KZ"},
	countryInfo{10, "Algeria", "Algiers", []string{africa}, "dinar", "DZ"},
	countryInfo{11, "Democratic Republic of Congo", "Kinshasa", []string{africa}, "franc", "CD"},
	countryInfo{12, "Saudi Arabia", "Riyadh", []string{middle_east}, "riyal", "SA"},
	countryInfo{13, "Mexico", "Mexico City", []string{central_america, north_america}, "peso", "MX"},
	countryInfo{14, "Indonesia", "Jakarta", []string{southeast_asia}, "rupiah", "ID"},
	countryInfo{15, "Sudan", "Khartoum", []string{africa}, "pound", "SD"},
	countryInfo{16, "Libya", "Tripoli", []string{africa}, "dinar", "LY"},
	countryInfo{17, "Iran", "Tehran", []string{middle_east}, "rial", "IR"},
	countryInfo{18, "Mongolia", "Ulaanbataar", []string{asia}, "tugrik", "MN"},
	countryInfo{19, "Peru", "Lima", []string{south_america}, "sol", "PE"},
	countryInfo{20, "Chad", "N'Djamena", []string{africa}, "franc", "TD"},
	countryInfo{21, "Niger", "Niamey", []string{africa}, "franc", "NE"},
	countryInfo{22, "Angola", "Luanda", []string{africa}, "kwanza", "AO"},
	countryInfo{23, "Mali", "Bamako", []string{africa}, "franc", "ML"},
	countryInfo{24, "South Africa", "Pretoria", []string{africa}, "rand", "ZA"},
	countryInfo{25, "Colombia", "Bogota", []string{south_america}, "peso", "CO"},
	countryInfo{26, "Ethiopia", "Addis Ababa", []string{africa, middle_east}, "birr", "ET"},
	countryInfo{27, "Bolivia", "Sucre", []string{south_america}, "boliviano", "BO"},
	countryInfo{28, "Mauritania", "Nouakchott", []string{africa}, "ouguiya", "MR"},
	countryInfo{29, "Egypt", "Cairo", []string{africa}, "pound", "EG"},
	countryInfo{30, "Tanzania", "Dodoma", []string{africa}, "shilling", "TZ"},
	countryInfo{31, "Nigeria", "Abuja", []string{africa}, "naira", "NG"},
	countryInfo{32, "Venezuela", "Caracas", []string{south_america}, "bolivar", "VE"},
	countryInfo{33, "Pakistan", "Islamabad", []string{asia, south_asia}, "rupee", "PK"},
	countryInfo{34, "Namibia", "Windhoek", []string{africa}, "dollar", "NA"},
	countryInfo{35, "Mozambique", "Maputo", []string{africa}, "metical", "MZ"},
	countryInfo{36, "Turkey", "Ankara", []string{asia, europe, middle_east}, "lyra", "TR"},
	countryInfo{37, "Chile", "Santiago", []string{south_america}, "peso", "CL"},
	countryInfo{38, "Zambia", "Lusaka", []string{africa}, "kwacha", "ZM"},
	countryInfo{39, "Myanmar", "Naypyidaw", []string{southeast_asia}, "kyat", "MM"},
	countryInfo{40, "Afghanistan", "Kabul", []string{south_asia}, "afghani", "AF"},
	countryInfo{41, "South Sudan", "Juba", []string{africa}, "pound", "SS"},
	countryInfo{42, "France", "Paris", []string{europe}, "euro", "FR"},
	countryInfo{43, "Somalia", "Mogadishu", []string{africa}, "shilling", "SO"},
	countryInfo{44, "Central African Republic", "Bangui", []string{africa}, "franc", "CF"},
	countryInfo{45, "Ukraine", "Kyiv", []string{europe}, "hryvnia", "UA"},
	countryInfo{46, "Madagascar", "Antananarivo", []string{africa}, "ariary", "MG"},
	countryInfo{47, "Botswana", "Gaborone", []string{africa}, "pula", "BW"},
	countryInfo{48, "Kenya", "Nairobi", []string{africa}, "shilling", "KE"},
	countryInfo{49, "Yemen", "Sana'a", []string{middle_east}, "rial", "YE"},
	countryInfo{50, "Thailand", "Bangkok", []string{southeast_asia}, "baht", "TH"},
	countryInfo{51, "Spain", "Madrid", []string{europe}, "euro", "ES"},
	countryInfo{52, "Turkmenistan", "Ashgabat", []string{asia}, "manat", "TM"},
	countryInfo{53, "Cameroon", "Yaounde", []string{africa}, "franc", "CM"},
	countryInfo{54, "Papua New Guinea", "Port Moresby", []string{oceania}, "kina", "PG"},
	countryInfo{55, "Sweden", "Stockholm", []string{europe}, "krona", "SE"},
	countryInfo{56, "Uzbekistan", "Tashkent", []string{asia}, "som", "UZ"},
	countryInfo{57, "Morocco", "Rabat", []string{africa}, "dirham", "MA"},
	countryInfo{58, "Iraq", "Baghdad", []string{middle_east}, "dinar", "IQ"},
	countryInfo{59, "Paraguay", "Asuncion", []string{south_america}, "guarani", "PY"},
	countryInfo{60, "Zimbabwe", "Harare", []string{africa}, "US dollar", "ZW"},
	countryInfo{61, "Norway", "Oslo", []string{europe}, "krone", "NO"},
	countryInfo{62, "Japan", "Tokyo", []string{asia}, "yen", "JP"},
	countryInfo{63, "Germany", "Berlin", []string{europe}, "euro", "DE"},
	countryInfo{64, "Republic of the Congo", "Brazzaville", []string{africa}, "franc", "CG"},
	countryInfo{65, "Finland", "Helsinki", []string{europe}, "euro", "FI"},
	countryInfo{66, "Viet Nam", "Hanoi", []string{southeast_asia}, "dong", "VN"},
	countryInfo{67, "Malaysia", "Kuala Lumpur", []string{southeast_asia}, "ringgit", "MY"},
	countryInfo{68, "Ivory Coast", "Yamoussoukro", []string{africa}, "franc", "CI"},
	countryInfo{69, "Poland", "Warsaw", []string{europe}, "zloty", "PL"},
	countryInfo{70, "Oman", "Muscat", []string{middle_east}, "rial", "OM"},
	countryInfo{71, "Italy", "Rome", []string{europe}, "euro", "IT"},
	countryInfo{72, "Phillipines", "Manila", []string{southeast_asia}, "peso", "PH"},
	countryInfo{73, "Ecuador", "Quito", []string{south_america}, "US dollar", "EC"},
	countryInfo{74, "Burkina Faso", "Ouagadougou", []string{africa}, "franc", "BF"},
	countryInfo{75, "New Zealand", "Wellington", []string{oceania}, "dollar", "NZ"},
	countryInfo{76, "Gabon", "Libreville", []string{africa}, "franc", "GA"},
	countryInfo{77, "Guinea", "Conakry", []string{africa}, "franc", "GN"},
	countryInfo{78, "United Kingdom", "London", []string{europe}, "pound", "GB"},
	countryInfo{79, "Uganda", "Kampala", []string{africa}, "shilling", "UG"},
	countryInfo{80, "Ghana", "Accra", []string{africa}, "cedi", "GH"},
	countryInfo{81, "Romania", "Bucharest", []string{europe}, "leu", "RO"},
	countryInfo{82, "Laos", "Vientiane", []string{southeast_asia}, "kip", "LA"},
	countryInfo{83, "Guyana", "Georgetown", []string{south_america}, "dollar", "GY"},
	countryInfo{84, "Belarus", "Minsk", []string{europe}, "ruble", "BY"},
	countryInfo{85, "Kyrgyzstan", "Bishkek", []string{asia}, "som", "KG"},
	countryInfo{86, "Senegal", "Dakar", []string{africa}, "franc", "SN"},
	countryInfo{87, "Syria", "Damascus", []string{middle_east}, "pound", "SY"},
	countryInfo{88, "Cambodia", "Phnom Penh", []string{southeast_asia}, "riel", "KH"},
	countryInfo{89, "Uruguay", "Montevideo", []string{south_america}, "peso", "UY"},
	countryInfo{90, "Suriname", "Paramaribo", []string{south_america}, "dollar", "SR"},
	countryInfo{91, "Tunisia", "Tunis", []string{africa}, "dinar", "TN"},
	countryInfo{92, "Bangladesh", "Dhaka", []string{south_asia}, "taka", "BD"},
	countryInfo{93, "Nepal", "Kathmandu", []string{south_asia}, "rupee", "NP"},
	countryInfo{94, "Tajikistan", "Dusharbe", []string{asia}, "somoni", "TJ"},
	countryInfo{95, "Greece", "Athens", []string{europe}, "euro", "GR"},
	countryInfo{96, "Nicaragua", "Managua", []string{central_america}, "cordoba", "NI"},
	countryInfo{97, "North Korea", "Pyongyang", []string{asia}, "won", "KP"},
	countryInfo{98, "Malawi", "Lilongwe", []string{africa}, "kwacha", "MW"},
	countryInfo{99, "Eritrea", "Asmara", []string{africa}, "nakfa", "ER"},
	countryInfo{100, "Benin", "Porto-Novo", []string{africa}, "franc", "BJ"},
	countryInfo{101, "Honduras", "Tegucigalpa", []string{central_america}, "lempira", "HN"},
	countryInfo{102, "Liberia", "Monrovia", []string{africa}, "dollar", "LR"},
	countryInfo{103, "Bulgaria", "Sofia", []string{europe}, "lev", "BG"},
	countryInfo{104, "Cuba", "Havana", []string{caribbean}, "peso", "CU"},
	countryInfo{105, "Guatemala", "Guatemala City", []string{central_america}, "quetzal", "GT"},
	countryInfo{106, "Iceland", "Reykjavik", []string{europe}, "krona", "IS"},
	countryInfo{107, "South Korea", "Seoul", []string{asia}, "won", "KR"},
	countryInfo{108, "Hungary", "Budapest", []string{europe}, "forint", "HU"},
	countryInfo{109, "Portugal", "Lisbon", []string{europe}, "euro", "PT"},
	countryInfo{110, "Jordan", "Amman", []string{middle_east}, "dinar", "JO"},
	countryInfo{111, "Serbia", "Belgrade", []string{europe}, "dinar", "RS"},
	countryInfo{112, "Azerbaijan", "Baku", []string{asia, europe}, "manat", "AZ"},
	countryInfo{113, "Austria", "Vienna", []string{europe}, "euro", "AT"},
	countryInfo{114, "United Arab Emirates", "Abu Dhabi", []string{middle_east}, "dirham", "AE"},
	countryInfo{115, "Czech Republic", "Prague", []string{europe}, "koruna", "CZ"},
	countryInfo{116, "Panama", "Panama City", []string{central_america}, "US dollar", "PA"},
	countryInfo{117, "Sierra Leone", "Freetown", []string{africa}, "leone", "SL"},
	countryInfo{118, "Ireland", "Dublin", []string{europe}, "euro", "IE"},
	countryInfo{119, "Georgia", "Tbilisi", []string{asia, europe}, "lari", "GE"},
	countryInfo{120, "Sri Lanka", "Sri Jayawardenepura Kotte", []string{southeast_asia}, "rupee", "LK"},
	countryInfo{121, "Lithuania", "Vilnius", []string{europe}, "euro", "LT"},
	countryInfo{122, "Latvia", "Riga", []string{europe}, "euro", "LV"},
	countryInfo{123, "Togo", "Lome", []string{africa}, "franc", "TG"},
	countryInfo{124, "Croatia", "Zagreb", []string{europe}, "kona", "HR"},
	countryInfo{125, "Bosnia and Herzegovina", "Sarajevo", []string{europe}, "mark", "BA"},
	countryInfo{126, "Costa Rica", "San Jose", []string{central_america}, "colon", "CR"},
	countryInfo{127, "Slovakia", "Bratislava", []string{europe}, "euro", "SK"},
	countryInfo{128, "Dominican Republic", "Santo Domingo", []string{caribbean}, "peso", "DO"},
	countryInfo{129, "Estonia", "Tallinn", []string{europe}, "kroon", "EE"},
	countryInfo{130, "Denmark", "Copenhagen", []string{europe}, "krone", "DK"},
	countryInfo{131, "Netherlands", "Amsterdam", []string{europe}, "euro", "NL"},
	countryInfo{132, "Switzerland", "Bern", []string{europe}, "Swiss franc", "CH"},
	countryInfo{133, "Bhutan", "Thimphu", []string{south_asia}, "ngultrum", "BT"},
	countryInfo{134, "Guinea-Bissau", "Bissau", []string{africa}, "franc", "GW"},
	countryInfo{135, "Moldova", "Kishinev", []string{europe}, "leu", "MD"},
	countryInfo{136, "Belgium", "Brussels", []string{europe}, "euro", "BE"},
	countryInfo{137, "Lesotho", "Maseru", []string{africa}, "loti", "LS"},
	countryInfo{138, "Armenia", "Yerevan", []string{asia}, "dram", "AM"},
	countryInfo{139, "Solomon Islands", "Honiara", []string{oceania}, "dollar", "SB"},
	countryInfo{140, "Albania", "Tirana", []string{europe}, "lek", "AL"},
	countryInfo{141, "Equatorial Guinea", "Malabo", []string{africa}, "franc", "GQ"},
	countryInfo{142, "Burundi", "Gitega", []string{africa}, "franc", "BI"},
	countryInfo{143, "Haiti", "Port-au-Prince", []string{caribbean}, "gourde", "HT"},
	countryInfo{144, "Rwanda", "Kigali", []string{africa}, "franc", "RW"},
	countryInfo{145, "North Macedonia", "Skopje", []string{europe}, "denar", "MK"},
	countryInfo{146, "Djibouti", "Djibouti", []string{africa}, "franc", "DJ"},
	countryInfo{147, "Belize", "Belmopan", []string{central_america}, "dollar", "BZ"},
	countryInfo{148, "El Salvador", "San Salvador", []string{central_america}, "US dollar", "SV"},
	countryInfo{149, "Israel", "Jerusalem", []string{middle_east}, "shekel", "IL"},
	countryInfo{150, "Slovenia", "Ljubljana", []string{europe}, "euro", "SI"},
	countryInfo{151, "Fiji", "Suva", []string{oceania}, "dollar", "FJ"},
	countryInfo{152, "Kuwait", "Kuwait City", []string{middle_east}, "dinar", "KW"},
	countryInfo{153, "Eswatini", "Mbabane", []string{africa}, "rand", "SZ"},
	countryInfo{154, "East Timor", "Dili", []string{southeast_asia}, "timor-leste", "TL"},
	countryInfo{155, "The Bahamas", "Nassau", []string{caribbean}, "dollar", "BS"},
	countryInfo{156, "Montenegro", "Podgorica", []string{europe}, "euro", "ME"},
	countryInfo{157, "Vanuatu", "Port Vila", []string{oceania}, "vatu", "VU"},
	countryInfo{158, "Qatar", "Doha", []string{middle_east}, "rial", "QA"},
	countryInfo{159, "The Gambia", "Banjul", []string{africa}, "dalasi", "GM"},
	countryInfo{160, "Jamaica", "Kingston", []string{caribbean}, "dollar", "JM"},
	countryInfo{161, "Lebanon", "Beirut", []string{middle_east}, "pound", "LB"},
	countryInfo{162, "Cyprus", "Nicosia", []string{europe, middle_east}, "euro", "CY"},
	countryInfo{163, "Brunei", "Bandar Seri Begawan", []string{south_asia}, "dollar", "BN"},
	countryInfo{164, "Trinidad and Tobago", "Port of Spain", []string{caribbean}, "dollar", "TT"},
	countryInfo{165, "Cape Verde", "Praia", []string{africa}, "escudo", "CV"},
	countryInfo{166, "Samoa", "Apia", []string{oceania}, "tala", "WS"},
	countryInfo{167, "Luxembourg", "Luxembourg City", []string{europe}, "euro", "LU"},
	countryInfo{168, "Mauritius", "Port Louis", []string{africa}, "rupee", "MU"},
	countryInfo{169, "Comoros", "Moroni", []string{africa}, "franc", "KM"},
	countryInfo{170, "Sao Tome and Principe", "Sao Tome", []string{africa}, "dobra", "ST"},
	countryInfo{171, "Kiribati", "South Tarawa", []string{oceania}, "Australian dollar", "KI"},
	countryInfo{172, "Bahrain", "Manama", []string{middle_east}, "dinar", "BH"},
	countryInfo{173, "Dominica", "Roseau", []string{caribbean}, "Eastern Caribbean dollar", "DM"},
	countryInfo{174, "Tonga", "Nuku'alofa", []string{oceania}, "pa'anga", "TO"},
	countryInfo{175, "Singapore", "Singapore", []string{southeast_asia}, "dollar", "SG"},
	countryInfo{176, "Federated States of Micronesia", "Palikir", []string{oceania}, "US dollar", "FM"},
	countryInfo{177, "Saint Lucia", "Castries", []string{caribbean}, "Eastern Caribbean dollar", "LC"},
	countryInfo{178, "Andorra", "Andorra la Vella", []string{europe}, "euro", "AD"},
	countryInfo{179, "Palau", "Ngerulmud", []string{oceania}, "US dollar", "PW"},
	countryInfo{180, "Seychelles", "Victoria", []string{africa}, "rupee", "SC"},
	countryInfo{181, "Antigua and Barbuda", "St. John's", []string{caribbean}, "Eastern Caribbean dollar", "AG"},
	countryInfo{182, "Barbados", "Bridgetown", []string{caribbean}, "dollar", "BB"},
	countryInfo{183, "Saint Vincent and the Grenadines", "Kingstown", []string{caribbean}, "Eastern Caribbean dollar", "VC"},
	countryInfo{184, "Grenada", "St. George's", []string{caribbean}, "Eastern Caribbean dollar", "GD"},
	countryInfo{185, "Malta", "Valletta", []string{europe}, "euro", "MT"},
	countryInfo{186, "Maldives", "Male", []string{south_asia}, "rufiyaa", "MV"},
	countryInfo{187, "Saint Kitts and Nevis", "Basseterre", []string{caribbean}, "Eastern Caribbean dollar", "KN"},
	countryInfo{188, "Marshall Islands", "Majuro", []string{oceania}, "US dollar", "MH"},
	countryInfo{189, "Liechtenstein", "Vaduz", []string{europe}, "Swiss franc", "LI"},
	countryInfo{190, "San Marino", "San Marino", []string{europe}, "euro", "SM"},
	countryInfo{191, "Tuvalu", "Funafuti", []string{oceania}, "US dollar", "TV"},
	countryInfo{192, "Nauru", "Yaren", []string{oceania}, "Australian dollar", "NR"},
	countryInfo{193, "Monaco", "Monaco", []string{europe}, "euro", "MC"},
	countryInfo{194, "Vatican City", "Vatican City", []string{europe}, "euro", "VA"},
}

type countryQuery func([]countryInfo) promptAndResponse

func quizCountries(cmd *cobra.Command, args []string) {
	quizFuncs := []countryQuery{
		// add a bunch of these to bias the randomizer
		crossQueryCountryInfo,
		crossQueryCountryInfo,
		crossQueryCountryInfo,
		crossQueryCountryInfo,
		crossQueryCountryInfo,
		crossQueryCountryInfo,
		crossQueryCountryInfo,
		quizWhichIsBigger,
		quizWhichIsSmaller,
		quizCountryFromFlag,
	}
	function := quizFuncs[rand.Intn(len(quizFuncs))]
	promptAndCheckResponse(function(countries))

}

func randomCountry(countries []countryInfo) countryInfo {
	return countries[rand.Intn(len(countries))]
}

func crossQueryCountryInfo(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return constructCrossQuery("country", country)
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
