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
	{1, "San Bernadino", "San Bernadino"},
	{2, "Inyo", "Independence"},
	{3, "Kern", "Bakersfield"},
	{4, "Riverside", "Riverside"},
	{5, "Siskiyou", "Yreka"},
	{6, "Fresno", "Fresno"},
	{7, "Tulare", "Visalia"},
	{8, "Lassen", "Susanville"},
	{9, "San Diego", "San Diego"},
	{10, "Imperial", "El Centro"},
	{11, "Los Angeles", "Los Angeles"},
	{12, "Modoc", "Alturas"},
	{13, "Shasta", "Redding"},
	{14, "Humboldt", "Eureka"},
	{15, "Mendocino", "Ukiah"},
	{16, "Monterey", "Salinas"},
	{17, "San Luis Obispo", "San Luis Obispo"},
	{18, "Trinity", "Weavervill"},
	{19, "Mono", "Bridgeport"},
	{20, "Tehama", "Red Bluff"},
	{21, "Santa Barbara", "Santa Barbara"},
	{22, "Plumas", "Quincy"},
	{23, "Tuolumne", "Sonora"},
	{24, "Madera", "Madera"},
	{25, "Merced", "Merced"},
	{26, "Ventura", "Ventura"},
	{27, "El Dorado", "Placerville"},
	{28, "Butte", "Oroville"},
	{29, "Sonoma", "Santa Rosa"},
	{30, "Stanislaus", "Modesto"},
	{31, "Mariposa", "Mariposa"},
	{32, "Placer", "Auburn"},
	{33, "San Joaquin", "Stockton"},
	{34, "Kings", "Hanford"},
	{35, "San Benito", "Hollister"},
	{36, "Glenn", "Willows"},
	{37, "Santa Clara", "San Jose"},
	{38, "Lake", "Lakeport"},
	{39, "Colusa", "Colusa"},
	{40, "Calaveras", "San Andreas"},
	{41, "Yolo", "Woodland"},
	{42, "Del Norte", "Crescent City"},
	{43, "Sacramento", "Sacramento"},
	{44, "Nevada", "Nevada City"},
	{45, "Sierra", "Downieville"},
	{46, "Orange", "Santa Ana"},
	{47, "Solano", "Fairfield"},
	{48, "Napa", "Napa"},
	{49, "Alpine", "Markleeville"},
	{50, "Alameda", "Oakland"},
	{51, "Contra Costa", "Martinez"},
	{52, "Yuba", "Marysville"},
	{53, "Amador", "Jackson"},
	{54, "Sutter", "Yuba City"},
	{55, "Marin", "San Rafael"},
	{56, "San Mateo", "Redwood City"},
	{57, "Santa Cruz", "Santa Cruz"},
	{58, "San Francisco", "San Francisco"},
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
	function := randomItemFromSlice(quizFuncs)
	promptAndCheckResponse(function(caCounties))

}

func crossQueryCaCountyInfo(counties []countyInfo) promptAndResponse {
	county := randomItemFromSlice(counties)
	return constructCrossQuery("CA county", county)
}

func quizWhichCountyIsBigger(counties []countyInfo) promptAndResponse {
	county1 := randomItemFromSlice(counties)
	county2 := randomItemFromSlice(counties)

	for county1.name == county2.name {
		county2 = randomItemFromSlice(counties)
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
	county1 := randomItemFromSlice(counties)
	county2 := randomItemFromSlice(counties)

	for county1.name == county2.name {
		county2 = randomItemFromSlice(counties)
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
