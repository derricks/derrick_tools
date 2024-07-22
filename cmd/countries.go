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
	rankInArea         int    `crossquery:"all" crossqueryname:"size rank"`
	name               string `crossquery:"all"`
	capital            string `crossquery:"all"`
	region             []string
	currency           string `crossquery:"guess"`
	countryCode        string
	ivrCode            string `crossquery:"all" crossqueryname:"IVR code"`
	landlocked         bool
	fractionalCurrency string `crossquery:"guess" crossqueryname:"fractional currency"`
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
	{1, "Russia", "Moscow", []string{asia, europe}, "ruble", "RU", "RUS", false, "kopeck"},
	{2, "Canada", "Ottawa", []string{north_america}, "dollar", "CA", "CDN", false, "cent"},
	{3, "United States", "Washington, D.C.", []string{north_america}, "US dollar", "US", "USA", false, "cent"},
	{4, "China", "Beijing", []string{asia}, "renminbi", "CN", "RC", false, "jiao"},
	{5, "Brazil", "Brasilia", []string{south_america}, "real", "BR", "BR", false, "centavo"},
	{6, "Australia", "Canberra", []string{oceania}, "Australian dollar", "AU", "AUS", false, ""},
	{7, "India", "New Delhi", []string{south_asia}, "rupee", "IN", "IND", false, "paisa"},
	{8, "Argentina", "Buenos Aires", []string{south_america}, "peso", "AR", "RA", false, "centavo"},
	{9, "Kazakhstan", "Nur-Sultan", []string{asia, europe}, "tenge", "KZ", "KZ", true, "tiyin"},
	{10, "Algeria", "Algiers", []string{africa}, "dinar", "DZ", "DZ", false, "centime"},
	{11, "Democratic Republic of Congo", "Kinshasa", []string{africa}, "franc", "CD", "CGO", false, "centime"},
	{12, "Saudi Arabia", "Riyadh", []string{middle_east}, "riyal", "SA", "KSA", false, "halala"},
	{13, "Mexico", "Mexico City", []string{central_america, north_america}, "peso", "MX", "MEX", false, "centavo"},
	{14, "Indonesia", "Jakarta", []string{southeast_asia}, "rupiah", "ID", "RI", false, "sen"},
	{15, "Sudan", "Khartoum", []string{africa}, "pound", "SD", "SUD", false, "piastre"},
	{16, "Libya", "Tripoli", []string{africa}, "dinar", "LY", "LAR", false, "dirham"},
	{17, "Iran", "Tehran", []string{middle_east}, "rial", "IR", "IR", false, "rial"},
	{18, "Mongolia", "Ulaanbataar", []string{asia}, "tugrik", "MN", "MGL", true, "mongo"},
	{19, "Peru", "Lima", []string{south_america}, "sol", "PE", "PE", false, "centimo"},
	{20, "Chad", "N'Djamena", []string{africa}, "franc", "TD", "TCH", false, "centime"},
	{21, "Niger", "Niamey", []string{africa}, "franc", "NE", "RN", true, "centime"},
	{22, "Angola", "Luanda", []string{africa}, "kwanza", "AO", "ANG", true, "centimo"},
	{23, "Mali", "Bamako", []string{africa}, "franc", "ML", "RMM", false, "centime"},
	{24, "South Africa", "Pretoria", []string{africa}, "rand", "ZA", "ZA", false, "cent"},
	{25, "Colombia", "Bogota", []string{south_america}, "peso", "CO", "CO", false, "centavo"},
	{26, "Ethiopia", "Addis Ababa", []string{africa, middle_east}, "birr", "ET", "ETH", true, "santim"},
	{27, "Bolivia", "Sucre", []string{south_america}, "boliviano", "BO", "BOL", true, "centavo"},
	{28, "Mauritania", "Nouakchott", []string{africa}, "ouguiya", "MR", "RIM", false, "khoums"},
	{29, "Egypt", "Cairo", []string{africa}, "pound", "EG", "ET", false, "piastre"},
	{30, "Tanzania", "Dodoma", []string{africa}, "shilling", "TZ", "EAT", false, "cent"},
	{31, "Nigeria", "Abuja", []string{africa}, "naira", "NG", "WAN", false, "kobo"},
	{32, "Venezuela", "Caracas", []string{south_america}, "bolivar", "VE", "YV", false, "cent"},
	{33, "Pakistan", "Islamabad", []string{asia, south_asia}, "rupee", "PK", "PK", false, "paisa"},
	{34, "Namibia", "Windhoek", []string{africa}, "dollar", "NA", "NAM", false, "cent"},
	{35, "Mozambique", "Maputo", []string{africa}, "metical", "MZ", "MOC", false, "centavo"},
	{36, "Turkey", "Ankara", []string{asia, europe, middle_east}, "lyra", "TR", "TR", false, "kurush"},
	{37, "Chile", "Santiago", []string{south_america}, "peso", "CL", "RCH", false, "centavo"},
	{38, "Zambia", "Lusaka", []string{africa}, "kwacha", "ZM", "Z", true, "ngwee"},
	{39, "Myanmar", "Naypyidaw", []string{southeast_asia}, "kyat", "MM", "BUR", false, "pya"},
	{40, "Afghanistan", "Kabul", []string{south_asia}, "afghani", "AF", "AFG", true, "pul"},
	{41, "South Sudan", "Juba", []string{africa}, "pound", "SS", "", true, "piaster"},
	{42, "France", "Paris", []string{europe}, "euro", "FR", "F", false, "cent"},
	{43, "Somalia", "Mogadishu", []string{africa}, "shilling", "SO", "SO", false, "cent"},
	{44, "Central African Republic", "Bangui", []string{africa}, "franc", "CF", "RCA", true, "centime"},
	{45, "Ukraine", "Kyiv", []string{europe}, "hryvnia", "UA", "UA", false, "kopek"},
	{46, "Madagascar", "Antananarivo", []string{africa}, "ariary", "MG", "RM", false, "iraimbilanja"},
	{47, "Botswana", "Gaborone", []string{africa}, "pula", "BW", "BW", true, "thebe"},
	{48, "Kenya", "Nairobi", []string{africa}, "shilling", "KE", "EAK", false, "cent"},
	{49, "Yemen", "Sana'a", []string{middle_east}, "rial", "YE", "YAR", false, "fils"},
	{50, "Thailand", "Bangkok", []string{southeast_asia}, "baht", "TH", "T", false, "satang"},
	{51, "Spain", "Madrid", []string{europe}, "euro", "ES", "E", false, "cent"},
	{52, "Turkmenistan", "Ashgabat", []string{asia}, "manat", "TM", "TM", true, "tenge"},
	{53, "Cameroon", "Yaounde", []string{africa}, "franc", "CM", "CAM", false, "centime"},
	{54, "Papua New Guinea", "Port Moresby", []string{oceania}, "kina", "PG", "PNG", false, "toea"},
	{55, "Sweden", "Stockholm", []string{europe}, "krona", "SE", "S", false, "ore"},
	{56, "Uzbekistan", "Tashkent", []string{asia}, "som", "UZ", "UZ", true, "tiyin"},
	{57, "Morocco", "Rabat", []string{africa}, "dirham", "MA", "MA", false, "centime"},
	{58, "Iraq", "Baghdad", []string{middle_east}, "dinar", "IQ", "IRQ", false, "fils"},
	{59, "Paraguay", "Asuncion", []string{south_america}, "guarani", "PY", "PY", true, "centimo"},
	{60, "Zimbabwe", "Harare", []string{africa}, "US dollar", "ZW", "ZW", true, "cent"},
	{61, "Norway", "Oslo", []string{europe}, "krone", "NO", "N", false, "ore"},
	{62, "Japan", "Tokyo", []string{asia}, "yen", "JP", "J", false, "sen"},
	{63, "Germany", "Berlin", []string{europe}, "euro", "DE", "D", false, "cent"},
	{64, "Republic of the Congo", "Brazzaville", []string{africa}, "franc", "CG", "RCB", false, "centime"},
	{65, "Finland", "Helsinki", []string{europe}, "euro", "FI", "FIN", false, "cent"},
	{66, "Viet Nam", "Hanoi", []string{southeast_asia}, "dong", "VN", "VN", false, "hao"},
	{67, "Malaysia", "Kuala Lumpur", []string{southeast_asia}, "ringgit", "MY", "MAL", false, "sen"},
	{68, "Ivory Coast", "Yamoussoukro", []string{africa}, "franc", "CI", "CN", false, "centime"},
	{69, "Poland", "Warsaw", []string{europe}, "zloty", "PL", "PL", false, "grosz"},
	{70, "Oman", "Muscat", []string{middle_east}, "rial", "OM", "OM", false, "baisa"},
	{71, "Italy", "Rome", []string{europe}, "euro", "IT", "I", false, "cent"},
	{72, "Phillipines", "Manila", []string{southeast_asia}, "peso", "PH", "RP", false, "sentimo"},
	{73, "Ecuador", "Quito", []string{south_america}, "US dollar", "EC", "EC", false, "centavo"},
	{74, "Burkina Faso", "Ouagadougou", []string{africa}, "franc", "BF", "BF", true, "centime"},
	{75, "New Zealand", "Wellington", []string{oceania}, "dollar", "NZ", "NZ", false, "cent"},
	{76, "Gabon", "Libreville", []string{africa}, "franc", "GA", "G", false, "centime"},
	{77, "Guinea", "Conakry", []string{africa}, "franc", "GN", "RG", false, "centime"},
	{78, "United Kingdom", "London", []string{europe}, "pound", "GB", "UK", false, "penny"},
	{79, "Uganda", "Kampala", []string{africa}, "shilling", "UG", "EAU", true, ""},
	{80, "Ghana", "Accra", []string{africa}, "cedi", "GH", "GH", false, "pesewa"},
	{81, "Romania", "Bucharest", []string{europe}, "leu", "RO", "RO", false, ""},
	{82, "Laos", "Vientiane", []string{southeast_asia}, "kip", "LA", "LAO", true, ""},
	{83, "Guyana", "Georgetown", []string{south_america}, "dollar", "GY", "GUY", false, ""},
	{84, "Belarus", "Minsk", []string{europe}, "ruble", "BY", "BY", true, ""},
	{85, "Kyrgyzstan", "Bishkek", []string{asia}, "som", "KG", "KG", true, ""},
	{86, "Senegal", "Dakar", []string{africa}, "franc", "SN", "SN", false, ""},
	{87, "Syria", "Damascus", []string{middle_east}, "pound", "SY", "SYR", false, ""},
	{88, "Cambodia", "Phnom Penh", []string{southeast_asia}, "riel", "KH", "K", false, ""},
	{89, "Uruguay", "Montevideo", []string{south_america}, "peso", "UY", "UY", false, ""},
	{90, "Suriname", "Paramaribo", []string{south_america}, "dollar", "SR", "SME", false, ""},
	{91, "Tunisia", "Tunis", []string{africa}, "dinar", "TN", "TN", false, ""},
	{92, "Bangladesh", "Dhaka", []string{south_asia}, "taka", "BD", "BD", false, ""},
	{93, "Nepal", "Kathmandu", []string{south_asia}, "rupee", "NP", "NEP", true, ""},
	{94, "Tajikistan", "Dusharbe", []string{asia}, "somoni", "TJ", "TJ", true, ""},
	{95, "Greece", "Athens", []string{europe}, "euro", "GR", "GR", false, ""},
	{96, "Nicaragua", "Managua", []string{central_america}, "cordoba", "NI", "NIC", false, ""},
	{97, "North Korea", "Pyongyang", []string{asia}, "won", "KP", "", false, ""},
	{98, "Malawi", "Lilongwe", []string{africa}, "kwacha", "MW", "MW", true, ""},
	{99, "Eritrea", "Asmara", []string{africa}, "nakfa", "ER", "ER", false, ""},
	{100, "Benin", "Porto-Novo", []string{africa}, "franc", "BJ", "DY", false, ""},
	{101, "Honduras", "Tegucigalpa", []string{central_america}, "lempira", "HN", "HN", false, ""},
	{102, "Liberia", "Monrovia", []string{africa}, "dollar", "LR", "LB", false, ""},
	{103, "Bulgaria", "Sofia", []string{europe}, "lev", "BG", "BG", false, ""},
	{104, "Cuba", "Havana", []string{caribbean}, "peso", "CU", "CU", false, ""},
	{105, "Guatemala", "Guatemala City", []string{central_america}, "quetzal", "GT", "GCA", false, ""},
	{106, "Iceland", "Reykjavik", []string{europe}, "krona", "IS", "IS", false, ""},
	{107, "South Korea", "Seoul", []string{asia}, "won", "KR", "ROK", false, ""},
	{108, "Hungary", "Budapest", []string{europe}, "forint", "HU", "H", true, ""},
	{109, "Portugal", "Lisbon", []string{europe}, "euro", "PT", "P", false, ""},
	{110, "Jordan", "Amman", []string{middle_east}, "dinar", "JO", "HKJ", false, ""},
	{111, "Serbia", "Belgrade", []string{europe}, "dinar", "RS", "SRB", true, ""},
	{112, "Azerbaijan", "Baku", []string{asia, europe}, "manat", "AZ", "AZ", true, ""},
	{113, "Austria", "Vienna", []string{europe}, "euro", "AT", "A", true, ""},
	{114, "United Arab Emirates", "Abu Dhabi", []string{middle_east}, "dirham", "AE", "UAE", false, ""},
	{115, "Czech Republic", "Prague", []string{europe}, "koruna", "CZ", "CZ", true, ""},
	{116, "Panama", "Panama City", []string{central_america}, "US dollar", "PA", "PA", false, ""},
	{117, "Sierra Leone", "Freetown", []string{africa}, "leone", "SL", "WAL", false, ""},
	{118, "Ireland", "Dublin", []string{europe}, "euro", "IE", "IRL", false, ""},
	{119, "Georgia", "Tbilisi", []string{asia, europe}, "lari", "GE", "GE", false, ""},
	{120, "Sri Lanka", "Sri Jayawardenepura Kotte", []string{southeast_asia}, "rupee", "LK", "CL", false, ""},
	{121, "Lithuania", "Vilnius", []string{europe}, "euro", "LT", "LT", false, ""},
	{122, "Latvia", "Riga", []string{europe}, "euro", "LV", "LV", false, ""},
	{123, "Togo", "Lome", []string{africa}, "franc", "TG", "TG", false, ""},
	{124, "Croatia", "Zagreb", []string{europe}, "kona", "HR", "HR", false, ""},
	{125, "Bosnia and Herzegovina", "Sarajevo", []string{europe}, "mark", "BA", "BIH", false, ""},
	{126, "Costa Rica", "San Jose", []string{central_america}, "colon", "CR", "CR", false, ""},
	{127, "Slovakia", "Bratislava", []string{europe}, "euro", "SK", "SK", true, ""},
	{128, "Dominican Republic", "Santo Domingo", []string{caribbean}, "peso", "DO", "DOM", false, ""},
	{129, "Estonia", "Tallinn", []string{europe}, "kroon", "EE", "EST", false, ""},
	{130, "Denmark", "Copenhagen", []string{europe}, "krone", "DK", "DK", false, ""},
	{131, "Netherlands", "Amsterdam", []string{europe}, "euro", "NL", "NL", false, ""},
	{132, "Switzerland", "Bern", []string{europe}, "Swiss franc", "CH", "CH", true, ""},
	{133, "Bhutan", "Thimphu", []string{south_asia}, "ngultrum", "BT", "BHT", true, ""},
	{134, "Guinea-Bissau", "Bissau", []string{africa}, "franc", "GW", "RGB", false, ""},
	{135, "Moldova", "Kishinev", []string{europe}, "leu", "MD", "MD", true, ""},
	{136, "Belgium", "Brussels", []string{europe}, "euro", "BE", "B", false, ""},
	{137, "Lesotho", "Maseru", []string{africa}, "loti", "LS", "LS", true, ""},
	{138, "Armenia", "Yerevan", []string{asia}, "dram", "AM", "AM", true, ""},
	{139, "Solomon Islands", "Honiara", []string{oceania}, "dollar", "SB", "SOL", false, ""},
	{140, "Albania", "Tirana", []string{europe}, "lek", "AL", "AL", false, ""},
	{141, "Equatorial Guinea", "Malabo", []string{africa}, "franc", "GQ", "", false, ""},
	{142, "Burundi", "Gitega", []string{africa}, "franc", "BI", "RU", true, ""},
	{143, "Haiti", "Port-au-Prince", []string{caribbean}, "gourde", "HT", "RH", false, ""},
	{144, "Rwanda", "Kigali", []string{africa}, "franc", "RW", "RWA", true, ""},
	{145, "North Macedonia", "Skopje", []string{europe}, "denar", "MK", "NMK", true, ""},
	{146, "Djibouti", "Djibouti", []string{africa}, "franc", "DJ", "", false, ""},
	{147, "Belize", "Belmopan", []string{central_america}, "dollar", "BZ", "BH", false, ""},
	{148, "El Salvador", "San Salvador", []string{central_america}, "US dollar", "SV", "ES", false, ""},
	{149, "Israel", "Jerusalem", []string{middle_east}, "shekel", "IL", "IL", false, ""},
	{150, "Slovenia", "Ljubljana", []string{europe}, "euro", "SI", "SLO", false, ""},
	{151, "Fiji", "Suva", []string{oceania}, "dollar", "FJ", "FJI", false, ""},
	{152, "Kuwait", "Kuwait City", []string{middle_east}, "dinar", "KW", "KWT", false, ""},
	{153, "Eswatini", "Mbabane", []string{africa}, "rand", "SZ", "SD", true, ""},
	{154, "East Timor", "Dili", []string{southeast_asia}, "timor-leste", "TL", "TL", false, ""},
	{155, "The Bahamas", "Nassau", []string{caribbean}, "dollar", "BS", "BS", false, ""},
	{156, "Montenegro", "Podgorica", []string{europe}, "euro", "ME", "MNE", false, ""},
	{157, "Vanuatu", "Port Vila", []string{oceania}, "vatu", "VU", "VU", false, ""},
	{158, "Qatar", "Doha", []string{middle_east}, "rial", "QA", "Q", false, ""},
	{159, "The Gambia", "Banjul", []string{africa}, "dalasi", "GM", "WAG", false, ""},
	{160, "Jamaica", "Kingston", []string{caribbean}, "dollar", "JM", "JA", false, ""},
	{161, "Lebanon", "Beirut", []string{middle_east}, "pound", "LB", "RL", false, ""},
	{162, "Cyprus", "Nicosia", []string{europe, middle_east}, "euro", "CY", "CY", false, ""},
	{163, "Brunei", "Bandar Seri Begawan", []string{south_asia}, "dollar", "BN", "BRU", false, ""},
	{164, "Trinidad and Tobago", "Port of Spain", []string{caribbean}, "dollar", "TT", "TT", false, ""},
	{165, "Cape Verde", "Praia", []string{africa}, "escudo", "CV", "CV", false, ""},
	{166, "Samoa", "Apia", []string{oceania}, "tala", "WS", "WS", false, ""},
	{167, "Luxembourg", "Luxembourg City", []string{europe}, "euro", "LU", "L", true, ""},
	{168, "Mauritius", "Port Louis", []string{africa}, "rupee", "MU", "MS", false, ""},
	{169, "Comoros", "Moroni", []string{africa}, "franc", "KM", "COM", false, ""},
	{170, "Sao Tome and Principe", "Sao Tome", []string{africa}, "dobra", "ST", "STP", false, ""},
	{171, "Kiribati", "South Tarawa", []string{oceania}, "Australian dollar", "KI", "KIR", false, ""},
	{172, "Bahrain", "Manama", []string{middle_east}, "dinar", "BH", "BRN", false, ""},
	{173, "Dominica", "Roseau", []string{caribbean}, "Eastern Caribbean dollar", "DM", "WD", false, ""},
	{174, "Tonga", "Nuku'alofa", []string{oceania}, "pa'anga", "TO", "TO", false, ""},
	{175, "Singapore", "Singapore", []string{southeast_asia}, "dollar", "SG", "SGP", false, ""},
	{176, "Federated States of Micronesia", "Palikir", []string{oceania}, "US dollar", "FM", "FSM", false, ""},
	{177, "Saint Lucia", "Castries", []string{caribbean}, "Eastern Caribbean dollar", "LC", "WL", false, ""},
	{178, "Andorra", "Andorra la Vella", []string{europe}, "euro", "AD", "AND", true, ""},
	{179, "Palau", "Ngerulmud", []string{oceania}, "US dollar", "PW", "PAL", false, ""},
	{180, "Seychelles", "Victoria", []string{africa}, "rupee", "SC", "SY", false, ""},
	{181, "Antigua and Barbuda", "St. John's", []string{caribbean}, "Eastern Caribbean dollar", "AG", "AG", false, ""},
	{182, "Barbados", "Bridgetown", []string{caribbean}, "dollar", "BB", "BDS", false, ""},
	{183, "Saint Vincent and the Grenadines", "Kingstown", []string{caribbean}, "Eastern Caribbean dollar", "VC", "WV", false, ""},
	{184, "Grenada", "St. George's", []string{caribbean}, "Eastern Caribbean dollar", "GD", "WG", false, ""},
	{185, "Malta", "Valletta", []string{europe}, "euro", "MT", "M", false, ""},
	{186, "Maldives", "Male", []string{south_asia}, "rufiyaa", "MV", "MV", false, ""},
	{187, "Saint Kitts and Nevis", "Basseterre", []string{caribbean}, "Eastern Caribbean dollar", "KN", "KAN", false, ""},
	{188, "Marshall Islands", "Majuro", []string{oceania}, "US dollar", "MH", "MH", false, ""},
	{189, "Liechtenstein", "Vaduz", []string{europe}, "Swiss franc", "LI", "FL", true, ""},
	{190, "San Marino", "San Marino", []string{europe}, "euro", "SM", "RSM", true, ""},
	{191, "Tuvalu", "Funafuti", []string{oceania}, "US dollar", "TV", "TUV", false, ""},
	{192, "Nauru", "Yaren", []string{oceania}, "Australian dollar", "NR", "NAU", false, ""},
	{193, "Monaco", "Monaco", []string{europe}, "euro", "MC", "MC", false, ""},
	{194, "Vatican City", "Vatican City", []string{europe}, "euro", "VA", "V", true, ""},
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
