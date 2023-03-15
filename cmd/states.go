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
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var statesCmd = &cobra.Command{
	Use:   "states",
	Short: "Quiz state information",
	Run:   quizStates,
}

type state struct {
	orderInUnion int      `crossquery:"all" crossqueryname:"order"`
	name         string   `crossquery:"all"`
	capital      string   `crossquery:"all"`
	yearJoined   int      `crossquery:"guess" crossqueryname:"year of joining"`
	nicknames    []string `crossquery:"given"`
	flowers      []string `crossquery:"guess" crossqueryname:"flower"`
}

var states = []state{
	state{1, "Delaware", "Dover", 1787, []string{"First State"}, []string{"Peach Blossom"}},
	state{2, "Pennsylvania", "Harrisburg", 1787, []string{"Keystone State"}, []string{"Mountain Laurel"}},
	state{3, "New Jersey", "Trenton", 1787, []string{"Garden State"}, []string{"Violet"}},
	state{4, "Georgia", "Atlanta", 1788, []string{"Peach State"}, []string{"Cherokee Rose"}},
	state{5, "Connecticut", "Hartford", 1788, []string{"Constitution State"}, []string{"Mountain Laurel"}},
	state{6, "Massachusetts", "Boston", 1788, []string{"Bay State"}, []string{"Mayflower"}},
	state{7, "Maryland", "Annapolis", 1788, []string{"Free State", "Old Line State"}, []string{"Black-eyed Susan"}},
	state{8, "South Carolina", "Columbia", 1788, []string{"Palmetto State"}, []string{"Yellow Jessamine"}},
	state{9, "New Hampshire", "Concord", 1788, []string{"Granite State"}, []string{"Purple Lilac"}},
	state{10, "Virginia", "Richmond", 1788, []string{"Old Dominion"}, []string{"American Dogwood"}},
	state{11, "New York", "Albany", 1788, []string{"Empire State"}, []string{"Rose"}},
	state{12, "North Carolina", "Raleigh", 1789, []string{"Tarheel State"}, []string{"Flowering Dogwood"}},
	state{13, "Rhode Island", "Providence", 1790, []string{"Ocean State"}, []string{"Violet"}},
	state{14, "Vermont", "Montpelier", 1791, []string{"Green Mountain State"}, []string{"Red Clover"}},
	state{15, "Kentucky", "Frankfort", 1792, []string{"Bluegrass State"}, []string{"Gooldenrod"}},
	state{16, "Tennessee", "Nashville", 1796, []string{"Volunteer State"}, []string{"Iris"}},
	state{17, "Ohio", "Columbus", 1803, []string{"Buckeye State"}, []string{"Scarlet Carnation"}},
	state{18, "Louisiana", "Baton Rouge", 1812, []string{"Pelican State"}, []string{"Magnolia"}},
	state{19, "Indiana", "Indianapolis", 1816, []string{"Hoosier State"}, []string{"Peony"}},
	state{20, "Mississippi", "Jackson", 1817, []string{"Magnolia State"}, []string{"Magnolia"}},
	state{21, "Illinois", "Springfield", 1818, []string{"Prairie State"}, []string{"Violet"}},
	state{22, "Alabama", "Montgomery", 1819, []string{"Heart of Dixie"}, []string{"Camellia"}},
	state{23, "Maine", "Augusta", 1820, []string{"Pine Tree State", "Vacationland"}, []string{"White Pine Cone and Tassel"}},
	state{24, "Missouri", "Jefferson City", 1821, []string{"Show Me State"}, []string{"Hawthorn"}},
	state{25, "Arkansas", "Little Rock", 1836, []string{"Natural State"}, []string{"Apple Blossom"}},
	state{26, "Michigan", "Lansing", 1837, []string{"Wolverine State", "Great Lakes State"}, []string{"Apple Blossom"}},
	state{27, "Florida", "Tallahassee", 1845, []string{"Sunshine State"}, []string{"Orange Blossom"}},
	state{28, "Texas", "Austin", 1845, []string{"Lone Star State"}, []string{"Bluebonnet"}},
	state{29, "Iowa", "Des Moines", 1846, []string{"Hawkeye State"}, []string{"Wild Rose"}},
	state{30, "Wisconsin", "Madison", 1848, []string{"America's Dairyland"}, []string{"Wood Violet"}},
	state{31, "California", "Sacramento", 1850, []string{"Golden State"}, []string{"California Poppy"}},
	state{32, "Minnesota", "Saint Paul", 1853, []string{"Land of 10,000 Lakes"}, []string{"Pink and White Lady's Slipper"}},
	state{33, "Oregon", "Salem", 1859, []string{"Beaver State"}, []string{"Oregon Rose"}},
	state{34, "Kansas", "Topeka", 1861, []string{"Sunflower State"}, []string{"Sunflower"}},
	state{35, "West Virginia", "Charleston", 1863, []string{"Mountain State"}, []string{"Rhododendron"}},
	state{36, "Nevada", "Carson City", 1864, []string{"Silver State"}, []string{"Sagebrush"}},
	state{37, "Nebraska", "Lincoln", 1867, []string{"Cornhusker State"}, []string{"Goldenrod"}},
	state{38, "Colorado", "Denver", 1876, []string{"Centennial State"}, []string{"Colorado blue columbine"}},
	state{39, "North Dakota", "Bismarck", 1889, []string{"Peace Garden State"}, []string{"Wild Prairie Rose"}},
	state{40, "South Dakota", "Pierre", 1889, []string{"Mount Rushmore State"}, []string{"Pasque Flower"}},
	state{41, "Montana", "Helena", 1889, []string{"Treasure State"}, []string{"Bitterroot"}},
	state{42, "Washington", "Olympia", 1889, []string{"Evergreen State"}, []string{"Coast rhododendron"}},
	state{43, "Idaho", "Boise", 1890, []string{"Gem State"}, []string{"Syringa"}},
	state{44, "Wyoming", "Cheyenne", 1890, []string{"Equality State"}, []string{"Indian Paintbrush"}},
	state{45, "Utah", "Salt Lake City", 1896, []string{"Beehive State"}, []string{"Sego Lily"}},
	state{46, "Oklahoma", "Oklahoma City", 1907, []string{"Sooner State"}, []string{"Oklahoma Rose"}},
	state{47, "New Mexico", "Santa Fe", 1912, []string{"Land of Enchantment"}, []string{"Yucca Flower"}},
	state{48, "Arizona", "Phoenix", 1912, []string{"Grand Canyon State"}, []string{"Saguaro Cactus Blossom"}},
	state{49, "Alaska", "Juneau", 1959, []string{"Last Frontier"}, []string{"Forget-me-not"}},
	state{50, "Hawaii", "Honolulu", 1959, []string{"Aloha State"}, []string{"Hawaiian Hibiscus"}},
}

type statesQuestion func([]state) promptAndResponse

func quizStates(cmd *cobra.Command, args []string) {

	var promptFuncs = []statesQuestion{
		crossQueryStateInfo,
		crossQueryStateInfo,
		crossQueryStateInfo,
		crossQueryStateInfo,
		crossQueryStateInfo,
		crossQueryStateInfo,
		crossQueryStateInfo,
		crossQueryStateInfo,
		quizStateJoinedEarliest,
		quizHowManyStatesInYear,
		quizNicknamesForState,
		quizStatesThatJoinedInAYear,
	}

	function := promptFuncs[rand.Intn(len(promptFuncs))]
	promptAndCheckResponse(function(states))
}

func crossQueryStateInfo(states []state) promptAndResponse {
	foundState := randomState(states)
	return constructCrossQuery("state", foundState)
}

func quizStateJoinedEarliest(states []state) promptAndResponse {
	state1 := randomState(states)
	state2 := randomState(states)
	for state1.name == state2.name {
		state2 = randomState(states)
	}

	prompt := fmt.Sprintf("Which state joined first? %s or %s?", state1.name, state2.name)
	if state1.orderInUnion < state2.orderInUnion {
		return promptAndResponse{prompt, state1.name}
	} else {
		return promptAndResponse{prompt, state2.name}
	}
}

func quizHowManyStatesInYear(states []state) promptAndResponse {
	today := time.Now()
	thisYear := today.Year()
	firstYear := states[0].yearJoined
	possibleDelta := rand.Intn(thisYear - states[0].yearJoined)
	targetYear := firstYear + possibleDelta

	countOfStates := 0
	for stateIndex := 0; stateIndex < len(states) && states[stateIndex].yearJoined <= targetYear; stateIndex++ {
		countOfStates++
	}
	return promptAndResponse{fmt.Sprintf("How many states were in the Union by the end of %d?", targetYear), strconv.Itoa(countOfStates)}
}

func quizNicknamesForState(states []state) promptAndResponse {
	state := randomState(states)
	sort.Strings(state.nicknames)
	nicknames := strings.Join(state.nicknames, ",")
	return promptAndResponse{fmt.Sprintf("What are the nicknames of %s?", state.name), nicknames}
}

func quizStatesThatJoinedInAYear(states []state) promptAndResponse {
	state := randomState(states)
	statesThatJoinedThatYear := 0
	for _, stateToCheck := range states {
		if stateToCheck.yearJoined == state.yearJoined {
			statesThatJoinedThatYear++
		}
	}
	return promptAndResponse{fmt.Sprintf("How many states joined in %d?", state.yearJoined), strconv.Itoa(statesThatJoinedThatYear)}
}

func randomState(states []state) state {
	return states[rand.Intn(len(states))]
}

func init() {
	memoryquizCmd.AddCommand(statesCmd)
}
