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
	orderInUnion int
	name         string
	capital      string
	yearJoined   int
	nicknames    []string
}

var states = []state{
	state{1, "Delaware", "Dover", 1787, []string{"First State"}},
	state{2, "Pennsylvania", "Harrisburg", 1787, []string{"Keystone State"}},
	state{3, "New Jersey", "Trenton", 1787, []string{"Garden State"}},
	state{4, "Georgia", "Atlanta", 1788, []string{"Peach State"}},
	state{5, "Connecticut", "Hartford", 1788, []string{"Constitution State"}},
	state{6, "Massachusetts", "Boston", 1788, []string{"Bay State"}},
	state{7, "Maryland", "Annapolis", 1788, []string{"Free State", "Old Line State"}},
	state{8, "South Carolina", "Columbia", 1788, []string{"Palmetto State"}},
	state{9, "New Hampshire", "Concord", 1788, []string{"Granite State"}},
	state{10, "Virginia", "Richmond", 1788, []string{"Old Dominion"}},
	state{11, "New York", "Albany", 1788, []string{"Empire State"}},
	state{12, "North Carolina", "Raleigh", 1789, []string{"Tarheel State"}},
	state{13, "Rhode Island", "Providence", 1790, []string{"Ocean State"}},
	state{14, "Vermont", "Montpelier", 1791, []string{"Green Mountain State"}},
	state{15, "Kentucky", "Frankfort", 1792, []string{"Bluegrass State"}},
	state{16, "Tennessee", "Nashville", 1796, []string{"Volunteer State"}},
	state{17, "Ohio", "Columbus", 1803, []string{"Buckeye State"}},
	state{18, "Louisiana", "Baton Rouge", 1812, []string{"Pelican State"}},
	state{19, "Indiana", "Indianapolis", 1816, []string{"Hoosier State"}},
	state{20, "Mississippi", "Jackson", 1817, []string{"Magnolia State"}},
	state{21, "Illinois", "Springfield", 1818, []string{"Prairie State"}},
	state{22, "Alabama", "Montgomery", 1819, []string{"Heart of Dixie"}},
	state{23, "Maine", "Augusta", 1820, []string{"Pine Tree State", "Vacationland"}},
	state{24, "Missouri", "Jefferson City", 1821, []string{"Show Me State"}},
	state{25, "Arkansas", "Little Rock", 1836, []string{"Natural State"}},
	state{26, "Michigan", "Lansing", 1837, []string{"Wolverine State", "Great Lakes State"}},
	state{27, "Florida", "Tallahassee", 1845, []string{"Sunshine State"}},
	state{28, "Texas", "Austin", 1845, []string{"Lone Star State"}},
	state{29, "Iowa", "Des Moines", 1846, []string{"Hawkeye State"}},
	state{30, "Wisconsin", "Madison", 1848, []string{"America's Dairyland"}},
	state{31, "California", "Sacramento", 1850, []string{"Golden State"}},
	state{32, "Minnesota", "Saint Paul", 1853, []string{"Land of 10,000 Lakes"}},
	state{33, "Oregon", "Salem", 1859, []string{"Beaver State"}},
	state{34, "Kansas", "Topeka", 1861, []string{"Sunflower State"}},
	state{35, "West Virginia", "Charleston", 1863, []string{"Mountain State"}},
	state{36, "Nevada", "Carson City", 1864, []string{"Silver State"}},
	state{37, "Nebraska", "Lincoln", 1867, []string{"Cornhusker State"}},
	state{38, "Colorado", "Denver", 1876, []string{"Centennial State"}},
	state{39, "North Dakota", "Bismarck", 1889, []string{"Peace Garden State"}},
	state{40, "South Dakota", "Pierre", 1889, []string{"Mount Rushmore State"}},
	state{41, "Montana", "Helena", 1889, []string{"Treasure State"}},
	state{42, "Washington", "Olympia", 1889, []string{"Evergreen State"}},
	state{43, "Idaho", "Boise", 1890, []string{"Gem State"}},
	state{44, "Wyoming", "Cheyenne", 1890, []string{"Equality State"}},
	state{45, "Utah", "Salt Lake City", 1896, []string{"Beehive State"}},
	state{46, "Oklahoma", "Oklahoma City", 1907, []string{"Sooner State"}},
	state{47, "New Mexico", "Santa Fe", 1912, []string{"Land of Enchantment"}},
	state{48, "Arizona", "Phoenix", 1912, []string{"Grand Canyon State"}},
	state{49, "Alaska", "Juneau", 1959, []string{"Last Frontier"}},
	state{50, "Hawaii", "Honolulu", 1959, []string{"Aloha State"}},
}

type statesQuestion func([]state) promptAndResponse

func quizStates(cmd *cobra.Command, args []string) {

	var promptFuncs = []statesQuestion{
		quizOrderToName,
		quizNameToOrder,
		quizStateToCapital,
		quizCapitalToState,
		quizStateJoinedEarliest,
		quizWhenStateJoined,
		quizHowManyStatesInYear,
		quizStateByNickname,
		quizNicknamesForState,
	}

	function := promptFuncs[rand.Intn(len(promptFuncs))]
	promptAndCheckResponse(function(states))
}

func quizOrderToName(states []state) promptAndResponse {
	foundState := randomState(states)
	return promptAndResponse{fmt.Sprintf("Which state was added to the United States at position %d?", foundState.orderInUnion), foundState.name}
}

func quizNameToOrder(states []state) promptAndResponse {
	foundState := randomState(states)
	return promptAndResponse{fmt.Sprintf("What order was %s added to the Union?", foundState.name), strconv.Itoa(foundState.orderInUnion)}
}

func quizStateToCapital(states []state) promptAndResponse {
	foundState := randomState(states)
	return promptAndResponse{fmt.Sprintf("What is the capital of %s?", foundState.name), foundState.capital}
}

func quizCapitalToState(states []state) promptAndResponse {
	foundState := randomState(states)
	return promptAndResponse{fmt.Sprintf("%s is the capital of which state?", foundState.capital), foundState.name}
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

func quizWhenStateJoined(states []state) promptAndResponse {
	state := randomState(states)
	return promptAndResponse{fmt.Sprintf("When did %s join the United States?", state.name), strconv.Itoa(state.yearJoined)}
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

func quizStateByNickname(states []state) promptAndResponse {
	state := randomState(states)
	nickname := state.nicknames[rand.Intn(len(state.nicknames))]
	return promptAndResponse{fmt.Sprintf("What state has the nickname %s?", nickname), state.name}
}

func quizNicknamesForState(states []state) promptAndResponse {
	state := randomState(states)
	sort.Strings(state.nicknames)
	nicknames := strings.Join(state.nicknames, ",")
	return promptAndResponse{fmt.Sprintf("What are the nicknames of %s?", state.name), nicknames}
}

func randomState(states []state) state {
	return states[rand.Intn(len(states))]
}

func init() {
	memoryquizCmd.AddCommand(statesCmd)
}
