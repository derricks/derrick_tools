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
	ivrCode     string `crossquery:"all" crossqueryname:"IVR code"`
	landlocked  bool
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
	countryInfo{1, "Russia", "Moscow", []string{asia, europe}, "ruble", "RU", "RUS", false},
	countryInfo{2, "Canada", "Ottawa", []string{north_america}, "dollar", "CA", "CDN", false},
	countryInfo{3, "United States", "Washington, D.C.", []string{north_america}, "US dollar", "US", "USA", false},
	countryInfo{4, "China", "Beijing", []string{asia}, "renminbi", "CN", "RC", false},
	countryInfo{5, "Brazil", "Brasilia", []string{south_america}, "real", "BR", "BR", false},
	countryInfo{6, "Australia", "Canberra", []string{oceania}, "Australian dollar", "AU", "AUS", false},
	countryInfo{7, "India", "New Delhi", []string{south_asia}, "rupee", "IN", "IND", false},
	countryInfo{8, "Argentina", "Buenos Aires", []string{south_america}, "peso", "AR", "RA", false},
	countryInfo{9, "Kazakhstan", "Nur-Sultan", []string{asia, europe}, "tenge", "KZ", "KZ", true},
	countryInfo{10, "Algeria", "Algiers", []string{africa}, "dinar", "DZ", "DZ", false},
	countryInfo{11, "Democratic Republic of Congo", "Kinshasa", []string{africa}, "franc", "CD", "CGO", false},
	countryInfo{12, "Saudi Arabia", "Riyadh", []string{middle_east}, "riyal", "SA", "KSA", false},
	countryInfo{13, "Mexico", "Mexico City", []string{central_america, north_america}, "peso", "MX", "MEX", false},
	countryInfo{14, "Indonesia", "Jakarta", []string{southeast_asia}, "rupiah", "ID", "RI", false},
	countryInfo{15, "Sudan", "Khartoum", []string{africa}, "pound", "SD", "SUD", false},
	countryInfo{16, "Libya", "Tripoli", []string{africa}, "dinar", "LY", "LAR", false},
	countryInfo{17, "Iran", "Tehran", []string{middle_east}, "rial", "IR", "IR", false},
	countryInfo{18, "Mongolia", "Ulaanbataar", []string{asia}, "tugrik", "MN", "MGL", true},
	countryInfo{19, "Peru", "Lima", []string{south_america}, "sol", "PE", "PE", false},
	countryInfo{20, "Chad", "N'Djamena", []string{africa}, "franc", "TD", "TCH", false},
	countryInfo{21, "Niger", "Niamey", []string{africa}, "franc", "NE", "RN", true},
	countryInfo{22, "Angola", "Luanda", []string{africa}, "kwanza", "AO", "ANG", true},
	countryInfo{23, "Mali", "Bamako", []string{africa}, "franc", "ML", "RMM", false},
	countryInfo{24, "South Africa", "Pretoria", []string{africa}, "rand", "ZA", "ZA", false},
	countryInfo{25, "Colombia", "Bogota", []string{south_america}, "peso", "CO", "CO", false},
	countryInfo{26, "Ethiopia", "Addis Ababa", []string{africa, middle_east}, "birr", "ET", "ETH", true},
	countryInfo{27, "Bolivia", "Sucre", []string{south_america}, "boliviano", "BO", "BOL", true},
	countryInfo{28, "Mauritania", "Nouakchott", []string{africa}, "ouguiya", "MR", "RIM", false},
	countryInfo{29, "Egypt", "Cairo", []string{africa}, "pound", "EG", "ET", false},
	countryInfo{30, "Tanzania", "Dodoma", []string{africa}, "shilling", "TZ", "EAT", false},
	countryInfo{31, "Nigeria", "Abuja", []string{africa}, "naira", "NG", "WAN", false},
	countryInfo{32, "Venezuela", "Caracas", []string{south_america}, "bolivar", "VE", "YV", false},
	countryInfo{33, "Pakistan", "Islamabad", []string{asia, south_asia}, "rupee", "PK", "PK", false},
	countryInfo{34, "Namibia", "Windhoek", []string{africa}, "dollar", "NA", "NAM", false},
	countryInfo{35, "Mozambique", "Maputo", []string{africa}, "metical", "MZ", "MOC", false},
	countryInfo{36, "Turkey", "Ankara", []string{asia, europe, middle_east}, "lyra", "TR", "TR", false},
	countryInfo{37, "Chile", "Santiago", []string{south_america}, "peso", "CL", "RCH", false},
	countryInfo{38, "Zambia", "Lusaka", []string{africa}, "kwacha", "ZM", "Z", true},
	countryInfo{39, "Myanmar", "Naypyidaw", []string{southeast_asia}, "kyat", "MM", "BUR", false},
	countryInfo{40, "Afghanistan", "Kabul", []string{south_asia}, "afghani", "AF", "AFG", true},
	countryInfo{41, "South Sudan", "Juba", []string{africa}, "pound", "SS", "", true},
	countryInfo{42, "France", "Paris", []string{europe}, "euro", "FR", "F", false},
	countryInfo{43, "Somalia", "Mogadishu", []string{africa}, "shilling", "SO", "SO", false},
	countryInfo{44, "Central African Republic", "Bangui", []string{africa}, "franc", "CF", "RCA", true},
	countryInfo{45, "Ukraine", "Kyiv", []string{europe}, "hryvnia", "UA", "UA", false},
	countryInfo{46, "Madagascar", "Antananarivo", []string{africa}, "ariary", "MG", "RM", false},
	countryInfo{47, "Botswana", "Gaborone", []string{africa}, "pula", "BW", "BW", true},
	countryInfo{48, "Kenya", "Nairobi", []string{africa}, "shilling", "KE", "EAK", false},
	countryInfo{49, "Yemen", "Sana'a", []string{middle_east}, "rial", "YE", "YAR", false},
	countryInfo{50, "Thailand", "Bangkok", []string{southeast_asia}, "baht", "TH", "T", false},
	countryInfo{51, "Spain", "Madrid", []string{europe}, "euro", "ES", "E", false},
	countryInfo{52, "Turkmenistan", "Ashgabat", []string{asia}, "manat", "TM", "TM", true},
	countryInfo{53, "Cameroon", "Yaounde", []string{africa}, "franc", "CM", "CAM", false},
	countryInfo{54, "Papua New Guinea", "Port Moresby", []string{oceania}, "kina", "PG", "PNG", false},
	countryInfo{55, "Sweden", "Stockholm", []string{europe}, "krona", "SE", "S", false},
	countryInfo{56, "Uzbekistan", "Tashkent", []string{asia}, "som", "UZ", "UZ", true},
	countryInfo{57, "Morocco", "Rabat", []string{africa}, "dirham", "MA", "MA", false},
	countryInfo{58, "Iraq", "Baghdad", []string{middle_east}, "dinar", "IQ", "IRQ", false},
	countryInfo{59, "Paraguay", "Asuncion", []string{south_america}, "guarani", "PY", "PY", true},
	countryInfo{60, "Zimbabwe", "Harare", []string{africa}, "US dollar", "ZW", "ZW", true},
	countryInfo{61, "Norway", "Oslo", []string{europe}, "krone", "NO", "N", false},
	countryInfo{62, "Japan", "Tokyo", []string{asia}, "yen", "JP", "J", false},
	countryInfo{63, "Germany", "Berlin", []string{europe}, "euro", "DE", "D", false},
	countryInfo{64, "Republic of the Congo", "Brazzaville", []string{africa}, "franc", "CG", "RCB", false},
	countryInfo{65, "Finland", "Helsinki", []string{europe}, "euro", "FI", "FIN", false},
	countryInfo{66, "Viet Nam", "Hanoi", []string{southeast_asia}, "dong", "VN", "VN", false},
	countryInfo{67, "Malaysia", "Kuala Lumpur", []string{southeast_asia}, "ringgit", "MY", "MAL", false},
	countryInfo{68, "Ivory Coast", "Yamoussoukro", []string{africa}, "franc", "CI", "CN", false},
	countryInfo{69, "Poland", "Warsaw", []string{europe}, "zloty", "PL", "PL", false},
	countryInfo{70, "Oman", "Muscat", []string{middle_east}, "rial", "OM", "OM", false},
	countryInfo{71, "Italy", "Rome", []string{europe}, "euro", "IT", "I", false},
	countryInfo{72, "Phillipines", "Manila", []string{southeast_asia}, "peso", "PH", "RP", false},
	countryInfo{73, "Ecuador", "Quito", []string{south_america}, "US dollar", "EC", "EC", false},
	countryInfo{74, "Burkina Faso", "Ouagadougou", []string{africa}, "franc", "BF", "BF", true},
	countryInfo{75, "New Zealand", "Wellington", []string{oceania}, "dollar", "NZ", "NZ", false},
	countryInfo{76, "Gabon", "Libreville", []string{africa}, "franc", "GA", "G", false},
	countryInfo{77, "Guinea", "Conakry", []string{africa}, "franc", "GN", "RG", false},
	countryInfo{78, "United Kingdom", "London", []string{europe}, "pound", "GB", "UK", false},
	countryInfo{79, "Uganda", "Kampala", []string{africa}, "shilling", "UG", "EAU", true},
	countryInfo{80, "Ghana", "Accra", []string{africa}, "cedi", "GH", "GH", false},
	countryInfo{81, "Romania", "Bucharest", []string{europe}, "leu", "RO", "RO", false},
	countryInfo{82, "Laos", "Vientiane", []string{southeast_asia}, "kip", "LA", "LAO", true},
	countryInfo{83, "Guyana", "Georgetown", []string{south_america}, "dollar", "GY", "GUY", false},
	countryInfo{84, "Belarus", "Minsk", []string{europe}, "ruble", "BY", "BY", true},
	countryInfo{85, "Kyrgyzstan", "Bishkek", []string{asia}, "som", "KG", "KG", true},
	countryInfo{86, "Senegal", "Dakar", []string{africa}, "franc", "SN", "SN", false},
	countryInfo{87, "Syria", "Damascus", []string{middle_east}, "pound", "SY", "SYR", false},
	countryInfo{88, "Cambodia", "Phnom Penh", []string{southeast_asia}, "riel", "KH", "K", false},
	countryInfo{89, "Uruguay", "Montevideo", []string{south_america}, "peso", "UY", "UY", false},
	countryInfo{90, "Suriname", "Paramaribo", []string{south_america}, "dollar", "SR", "SME", false},
	countryInfo{91, "Tunisia", "Tunis", []string{africa}, "dinar", "TN", "TN", false},
	countryInfo{92, "Bangladesh", "Dhaka", []string{south_asia}, "taka", "BD", "BD", false},
	countryInfo{93, "Nepal", "Kathmandu", []string{south_asia}, "rupee", "NP", "NEP", true},
	countryInfo{94, "Tajikistan", "Dusharbe", []string{asia}, "somoni", "TJ", "TJ", true},
	countryInfo{95, "Greece", "Athens", []string{europe}, "euro", "GR", "GR", false},
	countryInfo{96, "Nicaragua", "Managua", []string{central_america}, "cordoba", "NI", "NIC", false},
	countryInfo{97, "North Korea", "Pyongyang", []string{asia}, "won", "KP", "", false},
	countryInfo{98, "Malawi", "Lilongwe", []string{africa}, "kwacha", "MW", "MW", true},
	countryInfo{99, "Eritrea", "Asmara", []string{africa}, "nakfa", "ER", "ER", false},
	countryInfo{100, "Benin", "Porto-Novo", []string{africa}, "franc", "BJ", "DY", false},
	countryInfo{101, "Honduras", "Tegucigalpa", []string{central_america}, "lempira", "HN", "HN", false},
	countryInfo{102, "Liberia", "Monrovia", []string{africa}, "dollar", "LR", "LB", false},
	countryInfo{103, "Bulgaria", "Sofia", []string{europe}, "lev", "BG", "BG", false},
	countryInfo{104, "Cuba", "Havana", []string{caribbean}, "peso", "CU", "CU", false},
	countryInfo{105, "Guatemala", "Guatemala City", []string{central_america}, "quetzal", "GT", "GCA", false},
	countryInfo{106, "Iceland", "Reykjavik", []string{europe}, "krona", "IS", "IS", false},
	countryInfo{107, "South Korea", "Seoul", []string{asia}, "won", "KR", "ROK", false},
	countryInfo{108, "Hungary", "Budapest", []string{europe}, "forint", "HU", "H", true},
	countryInfo{109, "Portugal", "Lisbon", []string{europe}, "euro", "PT", "P", false},
	countryInfo{110, "Jordan", "Amman", []string{middle_east}, "dinar", "JO", "HKJ", false},
	countryInfo{111, "Serbia", "Belgrade", []string{europe}, "dinar", "RS", "SRB", true},
	countryInfo{112, "Azerbaijan", "Baku", []string{asia, europe}, "manat", "AZ", "AZ", true},
	countryInfo{113, "Austria", "Vienna", []string{europe}, "euro", "AT", "A", true},
	countryInfo{114, "United Arab Emirates", "Abu Dhabi", []string{middle_east}, "dirham", "AE", "UAE", false},
	countryInfo{115, "Czech Republic", "Prague", []string{europe}, "koruna", "CZ", "CZ", true},
	countryInfo{116, "Panama", "Panama City", []string{central_america}, "US dollar", "PA", "PA", false},
	countryInfo{117, "Sierra Leone", "Freetown", []string{africa}, "leone", "SL", "WAL", false},
	countryInfo{118, "Ireland", "Dublin", []string{europe}, "euro", "IE", "IRL", false},
	countryInfo{119, "Georgia", "Tbilisi", []string{asia, europe}, "lari", "GE", "GE", false},
	countryInfo{120, "Sri Lanka", "Sri Jayawardenepura Kotte", []string{southeast_asia}, "rupee", "LK", "CL", false},
	countryInfo{121, "Lithuania", "Vilnius", []string{europe}, "euro", "LT", "LT", false},
	countryInfo{122, "Latvia", "Riga", []string{europe}, "euro", "LV", "LV", false},
	countryInfo{123, "Togo", "Lome", []string{africa}, "franc", "TG", "TG", false},
	countryInfo{124, "Croatia", "Zagreb", []string{europe}, "kona", "HR", "HR", false},
	countryInfo{125, "Bosnia and Herzegovina", "Sarajevo", []string{europe}, "mark", "BA", "BIH", false},
	countryInfo{126, "Costa Rica", "San Jose", []string{central_america}, "colon", "CR", "CR", false},
	countryInfo{127, "Slovakia", "Bratislava", []string{europe}, "euro", "SK", "SK", true},
	countryInfo{128, "Dominican Republic", "Santo Domingo", []string{caribbean}, "peso", "DO", "DOM", false},
	countryInfo{129, "Estonia", "Tallinn", []string{europe}, "kroon", "EE", "EST", false},
	countryInfo{130, "Denmark", "Copenhagen", []string{europe}, "krone", "DK", "DK", false},
	countryInfo{131, "Netherlands", "Amsterdam", []string{europe}, "euro", "NL", "NL", false},
	countryInfo{132, "Switzerland", "Bern", []string{europe}, "Swiss franc", "CH", "CH", true},
	countryInfo{133, "Bhutan", "Thimphu", []string{south_asia}, "ngultrum", "BT", "BHT", true},
	countryInfo{134, "Guinea-Bissau", "Bissau", []string{africa}, "franc", "GW", "RGB", false},
	countryInfo{135, "Moldova", "Kishinev", []string{europe}, "leu", "MD", "MD", true},
	countryInfo{136, "Belgium", "Brussels", []string{europe}, "euro", "BE", "B", false},
	countryInfo{137, "Lesotho", "Maseru", []string{africa}, "loti", "LS", "LS", true},
	countryInfo{138, "Armenia", "Yerevan", []string{asia}, "dram", "AM", "AM", true},
	countryInfo{139, "Solomon Islands", "Honiara", []string{oceania}, "dollar", "SB", "SOL", false},
	countryInfo{140, "Albania", "Tirana", []string{europe}, "lek", "AL", "AL", false},
	countryInfo{141, "Equatorial Guinea", "Malabo", []string{africa}, "franc", "GQ", "", false},
	countryInfo{142, "Burundi", "Gitega", []string{africa}, "franc", "BI", "RU", true},
	countryInfo{143, "Haiti", "Port-au-Prince", []string{caribbean}, "gourde", "HT", "RH", false},
	countryInfo{144, "Rwanda", "Kigali", []string{africa}, "franc", "RW", "RWA", true},
	countryInfo{145, "North Macedonia", "Skopje", []string{europe}, "denar", "MK", "NMK", true},
	countryInfo{146, "Djibouti", "Djibouti", []string{africa}, "franc", "DJ", "", false},
	countryInfo{147, "Belize", "Belmopan", []string{central_america}, "dollar", "BZ", "BH", false},
	countryInfo{148, "El Salvador", "San Salvador", []string{central_america}, "US dollar", "SV", "ES", false},
	countryInfo{149, "Israel", "Jerusalem", []string{middle_east}, "shekel", "IL", "IL", false},
	countryInfo{150, "Slovenia", "Ljubljana", []string{europe}, "euro", "SI", "SLO", false},
	countryInfo{151, "Fiji", "Suva", []string{oceania}, "dollar", "FJ", "FJI", false},
	countryInfo{152, "Kuwait", "Kuwait City", []string{middle_east}, "dinar", "KW", "KWT", false},
	countryInfo{153, "Eswatini", "Mbabane", []string{africa}, "rand", "SZ", "SD", true},
	countryInfo{154, "East Timor", "Dili", []string{southeast_asia}, "timor-leste", "TL", "TL", false},
	countryInfo{155, "The Bahamas", "Nassau", []string{caribbean}, "dollar", "BS", "BS", false},
	countryInfo{156, "Montenegro", "Podgorica", []string{europe}, "euro", "ME", "MNE", false},
	countryInfo{157, "Vanuatu", "Port Vila", []string{oceania}, "vatu", "VU", "VU", false},
	countryInfo{158, "Qatar", "Doha", []string{middle_east}, "rial", "QA", "Q", false},
	countryInfo{159, "The Gambia", "Banjul", []string{africa}, "dalasi", "GM", "WAG", false},
	countryInfo{160, "Jamaica", "Kingston", []string{caribbean}, "dollar", "JM", "JA", false},
	countryInfo{161, "Lebanon", "Beirut", []string{middle_east}, "pound", "LB", "RL", false},
	countryInfo{162, "Cyprus", "Nicosia", []string{europe, middle_east}, "euro", "CY", "CY", false},
	countryInfo{163, "Brunei", "Bandar Seri Begawan", []string{south_asia}, "dollar", "BN", "BRU", false},
	countryInfo{164, "Trinidad and Tobago", "Port of Spain", []string{caribbean}, "dollar", "TT", "TT", false},
	countryInfo{165, "Cape Verde", "Praia", []string{africa}, "escudo", "CV", "CV", false},
	countryInfo{166, "Samoa", "Apia", []string{oceania}, "tala", "WS", "WS", false},
	countryInfo{167, "Luxembourg", "Luxembourg City", []string{europe}, "euro", "LU", "L", true},
	countryInfo{168, "Mauritius", "Port Louis", []string{africa}, "rupee", "MU", "MS", false},
	countryInfo{169, "Comoros", "Moroni", []string{africa}, "franc", "KM", "COM", false},
	countryInfo{170, "Sao Tome and Principe", "Sao Tome", []string{africa}, "dobra", "ST", "STP", false},
	countryInfo{171, "Kiribati", "South Tarawa", []string{oceania}, "Australian dollar", "KI", "KIR", false},
	countryInfo{172, "Bahrain", "Manama", []string{middle_east}, "dinar", "BH", "BRN", false},
	countryInfo{173, "Dominica", "Roseau", []string{caribbean}, "Eastern Caribbean dollar", "DM", "WD", false},
	countryInfo{174, "Tonga", "Nuku'alofa", []string{oceania}, "pa'anga", "TO", "TO", false},
	countryInfo{175, "Singapore", "Singapore", []string{southeast_asia}, "dollar", "SG", "SGP", false},
	countryInfo{176, "Federated States of Micronesia", "Palikir", []string{oceania}, "US dollar", "FM", "FSM", false},
	countryInfo{177, "Saint Lucia", "Castries", []string{caribbean}, "Eastern Caribbean dollar", "LC", "WL", false},
	countryInfo{178, "Andorra", "Andorra la Vella", []string{europe}, "euro", "AD", "AND", true},
	countryInfo{179, "Palau", "Ngerulmud", []string{oceania}, "US dollar", "PW", "PAL", false},
	countryInfo{180, "Seychelles", "Victoria", []string{africa}, "rupee", "SC", "SY", false},
	countryInfo{181, "Antigua and Barbuda", "St. John's", []string{caribbean}, "Eastern Caribbean dollar", "AG", "AG", false},
	countryInfo{182, "Barbados", "Bridgetown", []string{caribbean}, "dollar", "BB", "BDS", false},
	countryInfo{183, "Saint Vincent and the Grenadines", "Kingstown", []string{caribbean}, "Eastern Caribbean dollar", "VC", "WV", false},
	countryInfo{184, "Grenada", "St. George's", []string{caribbean}, "Eastern Caribbean dollar", "GD", "WG", false},
	countryInfo{185, "Malta", "Valletta", []string{europe}, "euro", "MT", "M", false},
	countryInfo{186, "Maldives", "Male", []string{south_asia}, "rufiyaa", "MV", "MV", false},
	countryInfo{187, "Saint Kitts and Nevis", "Basseterre", []string{caribbean}, "Eastern Caribbean dollar", "KN", "KAN", false},
	countryInfo{188, "Marshall Islands", "Majuro", []string{oceania}, "US dollar", "MH", "MH", false},
	countryInfo{189, "Liechtenstein", "Vaduz", []string{europe}, "Swiss franc", "LI", "FL", true},
	countryInfo{190, "San Marino", "San Marino", []string{europe}, "euro", "SM", "RSM", true},
	countryInfo{191, "Tuvalu", "Funafuti", []string{oceania}, "US dollar", "TV", "TUV", false},
	countryInfo{192, "Nauru", "Yaren", []string{oceania}, "Australian dollar", "NR", "NAU", false},
	countryInfo{193, "Monaco", "Monaco", []string{europe}, "euro", "MC", "MC", false},
	countryInfo{194, "Vatican City", "Vatican City", []string{europe}, "euro", "VA", "V", true},
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
		quizCountryLandlocked,
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

func quizCountryLandlocked(countries []countryInfo) promptAndResponse {
	country := randomCountry(countries)
	return promptAndResponse{fmt.Sprintf("%s is landlocked: true or false?", country.name),
		strconv.FormatBool(country.landlocked)}
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
