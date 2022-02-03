/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
}

var states = []state{
	state{1, "Delaware", "Dover", 1787},
	state{2, "Pennsylvania", "Harrisburg", 1787},
	state{3, "New Jersey", "Trenton", 1787},
	state{4, "Georgia", "Atlanta", 1788},
	state{5, "Connecticut", "Hartford", 1788},
	state{6, "Massachusetts", "Boston", 1788},
	state{7, "Maryland", "Annapolis", 1788},
	state{8, "South Carolina", "Columbia", 1788},
	state{9, "New Hampshire", "Concord", 1788},
	state{10, "Virginia", "Richmond", 1788},
	state{11, "New York", "Albany", 1788},
	state{12, "North Carolina", "Raleigh", 1789},
	state{13, "Rhode Island", "Providence", 1790},
	state{14, "Vermont", "Montpelier", 1791},
	state{15, "Kentucky", "Frankfort", 1792},
	state{16, "Tennessee", "Nashville", 1796},
	state{17, "Ohio", "Columbus", 1803},
	state{18, "Louisiana", "Baton Rouge", 1812},
	state{19, "Indiana", "Indianapolis", 1816},
	state{20, "Mississippi", "Jackson", 1817},
	state{21, "Illinois", "Springfield", 1818},
	state{22, "Alabama", "Montgomery", 1819},
	state{23, "Maine", "Augusta", 1820},
	state{24, "Missouri", "Jefferson City", 1821},
	state{25, "Arkansas", "Little Rock", 1836},
	state{26, "Michigan", "Lansing", 1837},
	state{27, "Florida", "Tallahassee", 1845},
	state{28, "Texas", "Austin", 1845},
	state{29, "Iowa", "Des Moines", 1846},
	state{30, "Wisconsin", "Madison", 1848},
	state{31, "California", "Sacramento", 1850},
	state{32, "Minnesota", "Saint Paul", 1853},
	state{33, "Oregon", "Salem", 1859},
	state{34, "Kansas", "Topeka", 1861},
	state{35, "West Virginia", "Charleston", 1863},
	state{36, "Nevada", "Carson City", 1864},
	state{37, "Nebraska", "Lincoln", 1867},
	state{38, "Colorado", "Denver", 1876},
	state{39, "North Dakota", "Bismarck", 1889},
	state{40, "South Dakota", "Pierre", 1889},
	state{41, "Montana", "Helena", 1889},
	state{42, "Washington", "Ølympia", 1889},
	state{43, "Idaho", "Boise", 1890},
	state{44, "Wyoming", "Cheyenne", 1890},
	state{45, "Utah", "Salt Lake City", 1896},
	state{46, "Oklahoma", "Oklahoma City", 1907},
	state{47, "New Mexico", "Santa Fe", 1912},
	state{48, "Arizona", "Phoenix", 1912},
	state{49, "Alaska", "Juneau", 1959},
	state{50, "Hawaii", "Honolulu", 1959},
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
	for stateIndex := 0; states[stateIndex].yearJoined <= targetYear && stateIndex < len(states); stateIndex++ {
		countOfStates++
	}
	return promptAndResponse{fmt.Sprintf("How many states were in the Union by the end of %d?", targetYear), strconv.Itoa(countOfStates)}
}

func randomState(states []state) state {
	return states[rand.Intn(len(states))]
}

func init() {
	memoryquizCmd.AddCommand(statesCmd)
}
