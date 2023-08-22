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
	countryInfo{1, "Russia", "Moscow", []string{asia, europe}, "ruble", "RU", "RUS"},
	countryInfo{2, "Canada", "Ottawa", []string{north_america}, "dollar", "CA", "CDN"},
	countryInfo{3, "United States", "Washington, D.C.", []string{north_america}, "US dollar", "US", "USA"},
	countryInfo{4, "China", "Beijing", []string{asia}, "renminbi", "CN", "RC"},
	countryInfo{5, "Brazil", "Brasilia", []string{south_america}, "real", "BR", "BR"},
	countryInfo{6, "Australia", "Canberra", []string{oceania}, "Australian dollar", "AU", "AUS"},
	countryInfo{7, "India", "New Delhi", []string{south_asia}, "rupee", "IN", "IND"},
	countryInfo{8, "Argentina", "Buenos Aires", []string{south_america}, "peso", "AR", "RA"},
	countryInfo{9, "Kazakhstan", "Nur-Sultan", []string{asia, europe}, "tenge", "KZ", "KZ"},
	countryInfo{10, "Algeria", "Algiers", []string{africa}, "dinar", "DZ", "DZ"},
	countryInfo{11, "Democratic Republic of Congo", "Kinshasa", []string{africa}, "franc", "CD", "CGO"},
	countryInfo{12, "Saudi Arabia", "Riyadh", []string{middle_east}, "riyal", "SA", "KSA"},
	countryInfo{13, "Mexico", "Mexico City", []string{central_america, north_america}, "peso", "MX", "MEX"},
	countryInfo{14, "Indonesia", "Jakarta", []string{southeast_asia}, "rupiah", "ID", "RI"},
	countryInfo{15, "Sudan", "Khartoum", []string{africa}, "pound", "SD", "SUD"},
	countryInfo{16, "Libya", "Tripoli", []string{africa}, "dinar", "LY", "LAR"},
	countryInfo{17, "Iran", "Tehran", []string{middle_east}, "rial", "IR", "IR"},
	countryInfo{18, "Mongolia", "Ulaanbataar", []string{asia}, "tugrik", "MN", "MGL"},
	countryInfo{19, "Peru", "Lima", []string{south_america}, "sol", "PE", "PE"},
	countryInfo{20, "Chad", "N'Djamena", []string{africa}, "franc", "TD", "TCH"},
	countryInfo{21, "Niger", "Niamey", []string{africa}, "franc", "NE", "RN"},
	countryInfo{22, "Angola", "Luanda", []string{africa}, "kwanza", "AO", "ANG"},
	countryInfo{23, "Mali", "Bamako", []string{africa}, "franc", "ML", "RMM"},
	countryInfo{24, "South Africa", "Pretoria", []string{africa}, "rand", "ZA", "ZA"},
	countryInfo{25, "Colombia", "Bogota", []string{south_america}, "peso", "CO", "CO"},
	countryInfo{26, "Ethiopia", "Addis Ababa", []string{africa, middle_east}, "birr", "ET", "ETH"},
	countryInfo{27, "Bolivia", "Sucre", []string{south_america}, "boliviano", "BO", "BOL"},
	countryInfo{28, "Mauritania", "Nouakchott", []string{africa}, "ouguiya", "MR", "RIM"},
	countryInfo{29, "Egypt", "Cairo", []string{africa}, "pound", "EG", "ET"},
	countryInfo{30, "Tanzania", "Dodoma", []string{africa}, "shilling", "TZ", "EAT"},
	countryInfo{31, "Nigeria", "Abuja", []string{africa}, "naira", "NG", "WAN"},
	countryInfo{32, "Venezuela", "Caracas", []string{south_america}, "bolivar", "VE", "YV"},
	countryInfo{33, "Pakistan", "Islamabad", []string{asia, south_asia}, "rupee", "PK", "PK"},
	countryInfo{34, "Namibia", "Windhoek", []string{africa}, "dollar", "NA", "NAM"},
	countryInfo{35, "Mozambique", "Maputo", []string{africa}, "metical", "MZ", "MOC"},
	countryInfo{36, "Turkey", "Ankara", []string{asia, europe, middle_east}, "lyra", "TR", "TR"},
	countryInfo{37, "Chile", "Santiago", []string{south_america}, "peso", "CL", "RCH"},
	countryInfo{38, "Zambia", "Lusaka", []string{africa}, "kwacha", "ZM", "Z"},
	countryInfo{39, "Myanmar", "Naypyidaw", []string{southeast_asia}, "kyat", "MM", "BUR"},
	countryInfo{40, "Afghanistan", "Kabul", []string{south_asia}, "afghani", "AF", "AFG"},
	countryInfo{41, "South Sudan", "Juba", []string{africa}, "pound", "SS", ""},
	countryInfo{42, "France", "Paris", []string{europe}, "euro", "FR", "F"},
	countryInfo{43, "Somalia", "Mogadishu", []string{africa}, "shilling", "SO", "SO"},
	countryInfo{44, "Central African Republic", "Bangui", []string{africa}, "franc", "CF", "RCA"},
	countryInfo{45, "Ukraine", "Kyiv", []string{europe}, "hryvnia", "UA", "UA"},
	countryInfo{46, "Madagascar", "Antananarivo", []string{africa}, "ariary", "MG", "RM"},
	countryInfo{47, "Botswana", "Gaborone", []string{africa}, "pula", "BW", "BW"},
	countryInfo{48, "Kenya", "Nairobi", []string{africa}, "shilling", "KE", "EAK"},
	countryInfo{49, "Yemen", "Sana'a", []string{middle_east}, "rial", "YE", "YAR"},
	countryInfo{50, "Thailand", "Bangkok", []string{southeast_asia}, "baht", "TH", "T"},
	countryInfo{51, "Spain", "Madrid", []string{europe}, "euro", "ES", "E"},
	countryInfo{52, "Turkmenistan", "Ashgabat", []string{asia}, "manat", "TM", "TM"},
	countryInfo{53, "Cameroon", "Yaounde", []string{africa}, "franc", "CM", "CAM"},
	countryInfo{54, "Papua New Guinea", "Port Moresby", []string{oceania}, "kina", "PG", "PNG"},
	countryInfo{55, "Sweden", "Stockholm", []string{europe}, "krona", "SE", "S"},
	countryInfo{56, "Uzbekistan", "Tashkent", []string{asia}, "som", "UZ", "UZ"},
	countryInfo{57, "Morocco", "Rabat", []string{africa}, "dirham", "MA", "MA"},
	countryInfo{58, "Iraq", "Baghdad", []string{middle_east}, "dinar", "IQ", "IRQ"},
	countryInfo{59, "Paraguay", "Asuncion", []string{south_america}, "guarani", "PY", "PY"},
	countryInfo{60, "Zimbabwe", "Harare", []string{africa}, "US dollar", "ZW", "ZW"},
	countryInfo{61, "Norway", "Oslo", []string{europe}, "krone", "NO", "N"},
	countryInfo{62, "Japan", "Tokyo", []string{asia}, "yen", "JP", "J"},
	countryInfo{63, "Germany", "Berlin", []string{europe}, "euro", "DE", "D"},
	countryInfo{64, "Republic of the Congo", "Brazzaville", []string{africa}, "franc", "CG", "RCB"},
	countryInfo{65, "Finland", "Helsinki", []string{europe}, "euro", "FI", "FIN"},
	countryInfo{66, "Viet Nam", "Hanoi", []string{southeast_asia}, "dong", "VN", "VN"},
	countryInfo{67, "Malaysia", "Kuala Lumpur", []string{southeast_asia}, "ringgit", "MY", "MAL"},
	countryInfo{68, "Ivory Coast", "Yamoussoukro", []string{africa}, "franc", "CI", "CN"},
	countryInfo{69, "Poland", "Warsaw", []string{europe}, "zloty", "PL", "PL"},
	countryInfo{70, "Oman", "Muscat", []string{middle_east}, "rial", "OM", "OM"},
	countryInfo{71, "Italy", "Rome", []string{europe}, "euro", "IT", "I"},
	countryInfo{72, "Phillipines", "Manila", []string{southeast_asia}, "peso", "PH", "RP"},
	countryInfo{73, "Ecuador", "Quito", []string{south_america}, "US dollar", "EC", "EC"},
	countryInfo{74, "Burkina Faso", "Ouagadougou", []string{africa}, "franc", "BF", "BF"},
	countryInfo{75, "New Zealand", "Wellington", []string{oceania}, "dollar", "NZ", "NZ"},
	countryInfo{76, "Gabon", "Libreville", []string{africa}, "franc", "GA", "G"},
	countryInfo{77, "Guinea", "Conakry", []string{africa}, "franc", "GN", "RG"},
	countryInfo{78, "United Kingdom", "London", []string{europe}, "pound", "GB", "UK"},
	countryInfo{79, "Uganda", "Kampala", []string{africa}, "shilling", "UG", "EAU"},
	countryInfo{80, "Ghana", "Accra", []string{africa}, "cedi", "GH", "GH"},
	countryInfo{81, "Romania", "Bucharest", []string{europe}, "leu", "RO", "RO"},
	countryInfo{82, "Laos", "Vientiane", []string{southeast_asia}, "kip", "LA", "LAO"},
	countryInfo{83, "Guyana", "Georgetown", []string{south_america}, "dollar", "GY", "GUY"},
	countryInfo{84, "Belarus", "Minsk", []string{europe}, "ruble", "BY", "BY"},
	countryInfo{85, "Kyrgyzstan", "Bishkek", []string{asia}, "som", "KG", "KG"},
	countryInfo{86, "Senegal", "Dakar", []string{africa}, "franc", "SN", "SN"},
	countryInfo{87, "Syria", "Damascus", []string{middle_east}, "pound", "SY", "SYR"},
	countryInfo{88, "Cambodia", "Phnom Penh", []string{southeast_asia}, "riel", "KH", "K"},
	countryInfo{89, "Uruguay", "Montevideo", []string{south_america}, "peso", "UY", "UY"},
	countryInfo{90, "Suriname", "Paramaribo", []string{south_america}, "dollar", "SR", "SME"},
	countryInfo{91, "Tunisia", "Tunis", []string{africa}, "dinar", "TN", "TN"},
	countryInfo{92, "Bangladesh", "Dhaka", []string{south_asia}, "taka", "BD", "BD"},
	countryInfo{93, "Nepal", "Kathmandu", []string{south_asia}, "rupee", "NP", "NEP"},
	countryInfo{94, "Tajikistan", "Dusharbe", []string{asia}, "somoni", "TJ", "TJ"},
	countryInfo{95, "Greece", "Athens", []string{europe}, "euro", "GR", "GR"},
	countryInfo{96, "Nicaragua", "Managua", []string{central_america}, "cordoba", "NI", "NIC"},
	countryInfo{97, "North Korea", "Pyongyang", []string{asia}, "won", "KP", ""},
	countryInfo{98, "Malawi", "Lilongwe", []string{africa}, "kwacha", "MW", "MW"},
	countryInfo{99, "Eritrea", "Asmara", []string{africa}, "nakfa", "ER", "ER"},
	countryInfo{100, "Benin", "Porto-Novo", []string{africa}, "franc", "BJ", "DY"},
	countryInfo{101, "Honduras", "Tegucigalpa", []string{central_america}, "lempira", "HN", "HN"},
	countryInfo{102, "Liberia", "Monrovia", []string{africa}, "dollar", "LR", "LB"},
	countryInfo{103, "Bulgaria", "Sofia", []string{europe}, "lev", "BG", "BG"},
	countryInfo{104, "Cuba", "Havana", []string{caribbean}, "peso", "CU", "CU"},
	countryInfo{105, "Guatemala", "Guatemala City", []string{central_america}, "quetzal", "GT", "GCA"},
	countryInfo{106, "Iceland", "Reykjavik", []string{europe}, "krona", "IS", "IS"},
	countryInfo{107, "South Korea", "Seoul", []string{asia}, "won", "KR", "ROK"},
	countryInfo{108, "Hungary", "Budapest", []string{europe}, "forint", "HU", "H"},
	countryInfo{109, "Portugal", "Lisbon", []string{europe}, "euro", "PT", "P"},
	countryInfo{110, "Jordan", "Amman", []string{middle_east}, "dinar", "JO", "HKJ"},
	countryInfo{111, "Serbia", "Belgrade", []string{europe}, "dinar", "RS", "SRB"},
	countryInfo{112, "Azerbaijan", "Baku", []string{asia, europe}, "manat", "AZ", "AZ"},
	countryInfo{113, "Austria", "Vienna", []string{europe}, "euro", "AT", "A"},
	countryInfo{114, "United Arab Emirates", "Abu Dhabi", []string{middle_east}, "dirham", "AE", "UAE"},
	countryInfo{115, "Czech Republic", "Prague", []string{europe}, "koruna", "CZ", "CZ"},
	countryInfo{116, "Panama", "Panama City", []string{central_america}, "US dollar", "PA", "PA"},
	countryInfo{117, "Sierra Leone", "Freetown", []string{africa}, "leone", "SL", "WAL"},
	countryInfo{118, "Ireland", "Dublin", []string{europe}, "euro", "IE", "IRL"},
	countryInfo{119, "Georgia", "Tbilisi", []string{asia, europe}, "lari", "GE", "GE"},
	countryInfo{120, "Sri Lanka", "Sri Jayawardenepura Kotte", []string{southeast_asia}, "rupee", "LK", "CL"},
	countryInfo{121, "Lithuania", "Vilnius", []string{europe}, "euro", "LT", "LT"},
	countryInfo{122, "Latvia", "Riga", []string{europe}, "euro", "LV", "LV"},
	countryInfo{123, "Togo", "Lome", []string{africa}, "franc", "TG", "TG"},
	countryInfo{124, "Croatia", "Zagreb", []string{europe}, "kona", "HR", "HR"},
	countryInfo{125, "Bosnia and Herzegovina", "Sarajevo", []string{europe}, "mark", "BA", "BIH"},
	countryInfo{126, "Costa Rica", "San Jose", []string{central_america}, "colon", "CR", "CR"},
	countryInfo{127, "Slovakia", "Bratislava", []string{europe}, "euro", "SK", "SK"},
	countryInfo{128, "Dominican Republic", "Santo Domingo", []string{caribbean}, "peso", "DO", "DOM"},
	countryInfo{129, "Estonia", "Tallinn", []string{europe}, "kroon", "EE", "EST"},
	countryInfo{130, "Denmark", "Copenhagen", []string{europe}, "krone", "DK", "DK"},
	countryInfo{131, "Netherlands", "Amsterdam", []string{europe}, "euro", "NL", "NL"},
	countryInfo{132, "Switzerland", "Bern", []string{europe}, "Swiss franc", "CH", "CH"},
	countryInfo{133, "Bhutan", "Thimphu", []string{south_asia}, "ngultrum", "BT", "BHT"},
	countryInfo{134, "Guinea-Bissau", "Bissau", []string{africa}, "franc", "GW", "RGB"},
	countryInfo{135, "Moldova", "Kishinev", []string{europe}, "leu", "MD", "MD"},
	countryInfo{136, "Belgium", "Brussels", []string{europe}, "euro", "BE", "B"},
	countryInfo{137, "Lesotho", "Maseru", []string{africa}, "loti", "LS", "LS"},
	countryInfo{138, "Armenia", "Yerevan", []string{asia}, "dram", "AM", "AM"},
	countryInfo{139, "Solomon Islands", "Honiara", []string{oceania}, "dollar", "SB", "SOL"},
	countryInfo{140, "Albania", "Tirana", []string{europe}, "lek", "AL", "AL"},
	countryInfo{141, "Equatorial Guinea", "Malabo", []string{africa}, "franc", "GQ", ""},
	countryInfo{142, "Burundi", "Gitega", []string{africa}, "franc", "BI", "RU"},
	countryInfo{143, "Haiti", "Port-au-Prince", []string{caribbean}, "gourde", "HT", "RH"},
	countryInfo{144, "Rwanda", "Kigali", []string{africa}, "franc", "RW", "RWA"},
	countryInfo{145, "North Macedonia", "Skopje", []string{europe}, "denar", "MK", "NMK"},
	countryInfo{146, "Djibouti", "Djibouti", []string{africa}, "franc", "DJ", ""},
	countryInfo{147, "Belize", "Belmopan", []string{central_america}, "dollar", "BZ", "BH"},
	countryInfo{148, "El Salvador", "San Salvador", []string{central_america}, "US dollar", "SV", "ES"},
	countryInfo{149, "Israel", "Jerusalem", []string{middle_east}, "shekel", "IL", "IL"},
	countryInfo{150, "Slovenia", "Ljubljana", []string{europe}, "euro", "SI", "SLO"},
	countryInfo{151, "Fiji", "Suva", []string{oceania}, "dollar", "FJ", "FJI"},
	countryInfo{152, "Kuwait", "Kuwait City", []string{middle_east}, "dinar", "KW", "KWT"},
	countryInfo{153, "Eswatini", "Mbabane", []string{africa}, "rand", "SZ", "SD"},
	countryInfo{154, "East Timor", "Dili", []string{southeast_asia}, "timor-leste", "TL", "TL"},
	countryInfo{155, "The Bahamas", "Nassau", []string{caribbean}, "dollar", "BS", "BS"},
	countryInfo{156, "Montenegro", "Podgorica", []string{europe}, "euro", "ME", "MNE"},
	countryInfo{157, "Vanuatu", "Port Vila", []string{oceania}, "vatu", "VU", "VU"},
	countryInfo{158, "Qatar", "Doha", []string{middle_east}, "rial", "QA", "Q"},
	countryInfo{159, "The Gambia", "Banjul", []string{africa}, "dalasi", "GM", "WAG"},
	countryInfo{160, "Jamaica", "Kingston", []string{caribbean}, "dollar", "JM", "JA"},
	countryInfo{161, "Lebanon", "Beirut", []string{middle_east}, "pound", "LB", "RL"},
	countryInfo{162, "Cyprus", "Nicosia", []string{europe, middle_east}, "euro", "CY", "CY"},
	countryInfo{163, "Brunei", "Bandar Seri Begawan", []string{south_asia}, "dollar", "BN", "BRU"},
	countryInfo{164, "Trinidad and Tobago", "Port of Spain", []string{caribbean}, "dollar", "TT", "TT"},
	countryInfo{165, "Cape Verde", "Praia", []string{africa}, "escudo", "CV", "CV"},
	countryInfo{166, "Samoa", "Apia", []string{oceania}, "tala", "WS", "WS"},
	countryInfo{167, "Luxembourg", "Luxembourg City", []string{europe}, "euro", "LU", "L"},
	countryInfo{168, "Mauritius", "Port Louis", []string{africa}, "rupee", "MU", "MS"},
	countryInfo{169, "Comoros", "Moroni", []string{africa}, "franc", "KM", "COM"},
	countryInfo{170, "Sao Tome and Principe", "Sao Tome", []string{africa}, "dobra", "ST", "STP"},
	countryInfo{171, "Kiribati", "South Tarawa", []string{oceania}, "Australian dollar", "KI", "KIR"},
	countryInfo{172, "Bahrain", "Manama", []string{middle_east}, "dinar", "BH", "BRN"},
	countryInfo{173, "Dominica", "Roseau", []string{caribbean}, "Eastern Caribbean dollar", "DM", "WD"},
	countryInfo{174, "Tonga", "Nuku'alofa", []string{oceania}, "pa'anga", "TO", "TO"},
	countryInfo{175, "Singapore", "Singapore", []string{southeast_asia}, "dollar", "SG", "SGP"},
	countryInfo{176, "Federated States of Micronesia", "Palikir", []string{oceania}, "US dollar", "FM", "FSM"},
	countryInfo{177, "Saint Lucia", "Castries", []string{caribbean}, "Eastern Caribbean dollar", "LC", "WL"},
	countryInfo{178, "Andorra", "Andorra la Vella", []string{europe}, "euro", "AD", "AND"},
	countryInfo{179, "Palau", "Ngerulmud", []string{oceania}, "US dollar", "PW", "PAL"},
	countryInfo{180, "Seychelles", "Victoria", []string{africa}, "rupee", "SC", "SY"},
	countryInfo{181, "Antigua and Barbuda", "St. John's", []string{caribbean}, "Eastern Caribbean dollar", "AG", "AG"},
	countryInfo{182, "Barbados", "Bridgetown", []string{caribbean}, "dollar", "BB", "BDS"},
	countryInfo{183, "Saint Vincent and the Grenadines", "Kingstown", []string{caribbean}, "Eastern Caribbean dollar", "VC", "WV"},
	countryInfo{184, "Grenada", "St. George's", []string{caribbean}, "Eastern Caribbean dollar", "GD", "WG"},
	countryInfo{185, "Malta", "Valletta", []string{europe}, "euro", "MT", "M"},
	countryInfo{186, "Maldives", "Male", []string{south_asia}, "rufiyaa", "MV", "MV"},
	countryInfo{187, "Saint Kitts and Nevis", "Basseterre", []string{caribbean}, "Eastern Caribbean dollar", "KN", "KAN"},
	countryInfo{188, "Marshall Islands", "Majuro", []string{oceania}, "US dollar", "MH", "MH"},
	countryInfo{189, "Liechtenstein", "Vaduz", []string{europe}, "Swiss franc", "LI", "FL"},
	countryInfo{190, "San Marino", "San Marino", []string{europe}, "euro", "SM", "RSM"},
	countryInfo{191, "Tuvalu", "Funafuti", []string{oceania}, "US dollar", "TV", "TUV"},
	countryInfo{192, "Nauru", "Yaren", []string{oceania}, "Australian dollar", "NR", "NAU"},
	countryInfo{193, "Monaco", "Monaco", []string{europe}, "euro", "MC", "MC"},
	countryInfo{194, "Vatican City", "Vatican City", []string{europe}, "euro", "VA", "V"},
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
