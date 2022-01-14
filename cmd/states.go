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
}

var states = []state{
	state{1, "Delaware", "Dover"},
	state{2, "Pennsylvania", "Harrisburg"},
	state{3, "New Jersey", "Trenton"},
	state{4, "Georgia", "Atlanta"},
	state{5, "Connecticut", "Hartford"},
	state{6, "Massachusetts", "Boston"},
	state{7, "Maryland", "Annapolis"},
	state{8, "South Carolina", "Columbia"},
	state{9, "New Hampshire", "Concord"},
	state{10, "Virginia", "Richmond"},
	state{11, "New York", "Albany"},
	state{12, "North Carolina", "Raleigh"},
	state{13, "Rhode Island", "Providence"},
	state{14, "Vermont", "Montpelier"},
	state{15, "Kentucky", "Frankfort"},
	state{16, "Tennessee", "Nashville"},
	state{17, "Ohio", "Columbus"},
	state{18, "Louisiana", "Baton Rouge"},
	state{19, "Indiana", "Indianapolis"},
	state{20, "Mississippi", "Jackson"},
	state{21, "Illinois", "Springfield"},
	state{22, "Alabama", "Montgomery"},
	state{23, "Maine", "Augusta"},
	state{24, "Missouri", "Jefferson City"},
	state{25, "Arkansas", "Little Rock"},
	state{26, "Michigan", "Lansing"},
	state{27, "Florida", "Tallahassee"},
	state{28, "Texas", "Austin"},
	state{29, "Iowa", "Des Moines"},
	state{30, "Wisconsin", "Madison"},
	state{31, "California", "Sacramento"},
	state{32, "Minnesota", "Saint Paul"},
	state{33, "Oregon", "Salem"},
	state{34, "Kansas", "Topeka"},
	state{35, "West Virginia", "Charleston"},
	state{36, "Nevada", "Carson City"},
	state{37, "Nebraska", "Lincoln"},
	state{38, "Colorado", "Denver"},
	state{39, "North Dakota", "Bismarck"},
	state{40, "South Dakota", "Pierre"},
	state{41, "Montana", "Helena"},
	state{42, "Washington", "Ølympia"},
	state{43, "Idaho", "Boise"},
	state{44, "Wyoming", "Cheyenne"},
	state{45, "Utah", "Salt Lake City"},
	state{46, "Oklahoma", "Oklahoma City"},
	state{47, "New Mexico", "Santa Fe"},
	state{48, "Arizona", "Phoenix"},
	state{49, "Alaska", "Juneau"},
	state{50, "Hawaii", "Honolulu"},
}

type statesQuestion func([]state) promptAndResponse

func quizStates(cmd *cobra.Command, args []string) {

	var promptFuncs = []statesQuestion{
		quizOrderToName,
		quizNameToOrder,
		quizStateToCapital,
		quizCapitalToState,
		quizStateJoinedEarliest,
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

func randomState(states []state) state {
	return states[rand.Intn(len(states))]
}

func init() {
	memoryquizCmd.AddCommand(statesCmd)
}
