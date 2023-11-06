/*
Copyright © 2023 Derrick Schneider derrick.schneider@gmail.com

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
var caCountiesCmd = &cobra.Command{
	Use:   "ca-counties",
	Short: "Memory quizzes about California, including county seats",
	Run:   quizCaCounties,
}

// note that this struct and most of the methods below can be reused if I ever add
// other counties
type countyInfo struct {
	sizeRank   int    `crossquery:"all" crossqueryname:"size rank"`
	name       string `crossquery:"all"`
	countySeat string `crossquery:"all" crossqueryname:"county seat"`
}

var caCounties = []countyInfo{
	countyInfo{1, "San Bernadino", "San Bernadino"},
	countyInfo{2, "Inyo", "Independence"},
	countyInfo{3, "Kern", "Bakersfield"},
	countyInfo{4, "Riverside", "Riverside"},
	countyInfo{5, "Siskiyou", "Yreka"},
	countyInfo{6, "Fresno", "Fresno"},
	countyInfo{7, "Tulare", "Visalia"},
	countyInfo{8, "Lassen", "Susanville"},
	countyInfo{9, "San Diego", "San Diego"},
	countyInfo{10, "Imperial", "El Centro"},
	countyInfo{11, "Los Angeles", "Los Angeles"},
	countyInfo{12, "Modoc", "Alturas"},
	countyInfo{13, "Shasta", "Redding"},
	countyInfo{14, "Humboldt", "Eureka"},
	countyInfo{15, "Mendocino", "Ukiah"},
	countyInfo{16, "Monterey", "Salinas"},
	countyInfo{17, "San Luis Obispo", "San Luis Obispo"},
	countyInfo{18, "Trinity", "Weavervill"},
	countyInfo{19, "Mono", "Bridgeport"},
	countyInfo{20, "Tehama", "Red Bluff"},
	countyInfo{21, "Santa Barbara", "Santa Barbara"},
	countyInfo{22, "Plumas", "Quincy"},
	countyInfo{23, "Tuolumne", "Sonora"},
	countyInfo{24, "Madera", "Madera"},
	countyInfo{25, "Merced", "Merced"},
	countyInfo{26, "Ventura", "Ventura"},
	countyInfo{27, "El Dorado", "Placerville"},
	countyInfo{28, "Butte", "Oroville"},
	countyInfo{29, "Sonoma", "Santa Rosa"},
	countyInfo{30, "Stanislaus", "Modesto"},
	countyInfo{31, "Mariposa", "Mariposa"},
	countyInfo{32, "Placer", "Auburn"},
	countyInfo{33, "San Joaquin", "Stockton"},
	countyInfo{34, "Kings", "Hanford"},
	countyInfo{35, "San Benito", "Hollister"},
	countyInfo{36, "Glenn", "Willows"},
	countyInfo{37, "Santa Clara", "San Jose"},
	countyInfo{38, "Lake", "Lakeport"},
	countyInfo{39, "Colusa", "Colusa"},
	countyInfo{40, "Calaveras", "San Andreas"},
	countyInfo{41, "Yolo", "Woodland"},
	countyInfo{42, "Del Norte", "Crescent City"},
	countyInfo{43, "Sacramento", "Sacramento"},
	countyInfo{44, "Nevada", "Nevada City"},
	countyInfo{45, "Sierra", "Downieville"},
	countyInfo{46, "Orange", "Santa Ana"},
	countyInfo{47, "Solano", "Fairfield"},
	countyInfo{48, "Napa", "Napa"},
	countyInfo{49, "Alpine", "Markleeville"},
	countyInfo{50, "Alameda", "Oakland"},
	countyInfo{51, "Contra Costa", "Martinez"},
	countyInfo{52, "Yuba", "Marysville"},
	countyInfo{53, "Amador", "Jackson"},
	countyInfo{54, "Sutter", "Yuba City"},
	countyInfo{55, "Marin", "San Rafael"},
	countyInfo{56, "San Mateo", "Redwood City"},
	countyInfo{57, "Santa Cruz", "Santa Cruz"},
	countyInfo{58, "San Francisco", "San Francisco"},
}

type countyQuery func([]countyInfo) promptAndResponse

func quizCaCounties(cmd *cobra.Command, args []string) {
	quizFuncs := []countyQuery{
		crossQueryCaCountyInfo,
		crossQueryCaCountyInfo,
		crossQueryCaCountyInfo,
		crossQueryCaCountyInfo,
		crossQueryCaCountyInfo,
		crossQueryCaCountyInfo,
		crossQueryCaCountyInfo,
		quizWhichCountyIsBigger,
		quizWhichCountyIsSmaller,
	}
	function := quizFuncs[rand.Intn(len(quizFuncs))]
	promptAndCheckResponse(function(caCounties))

}

func randomCounty(counties []countyInfo) countyInfo {
	return counties[rand.Intn(len(counties))]
}

func crossQueryCaCountyInfo(counties []countyInfo) promptAndResponse {
	county := randomCounty(counties)
	return constructCrossQuery("CA county", county)
}

func quizWhichCountyIsBigger(counties []countyInfo) promptAndResponse {
	county1 := randomCounty(counties)
	county2 := randomCounty(counties)

	for county1.name == county2.name {
		county2 = randomCounty(counties)
	}

	response := promptAndResponse{fmt.Sprintf("Which county is bigger: %s or %s?", county1.name, county2.name), ""}
	// rank is inverse to size: 1 means the biggest country, and 50 is bigger than 70
	if county1.sizeRank < county2.sizeRank {
		response.response = county1.name
	} else {
		response.response = county2.name
	}
	return response
}

func quizWhichCountyIsSmaller(counties []countyInfo) promptAndResponse {
	county1 := randomCounty(counties)
	county2 := randomCounty(counties)

	for county1.name == county2.name {
		county2 = randomCounty(counties)
	}

	response := promptAndResponse{fmt.Sprintf("Which county is smaller: %s or %s?", county1.name, county2.name), ""}
	// rank is inverse to size: 1 means the biggest country, and 50 is bigger than 70
	if county1.sizeRank > county2.sizeRank {
		response.response = county1.name
	} else {
		response.response = county2.name
	}
	return response
}

func init() {
	memoryquizCmd.AddCommand(caCountiesCmd)
}
