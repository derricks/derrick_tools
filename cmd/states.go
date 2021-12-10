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
}

var states = []state{
	state{1, "Delaware"},
	state{2, "Pennsylvania"},
	state{3, "New Jersey"},
	state{4, "Georgia"},
	state{5, "Connecticut"},
	state{6, "Massachusetts"},
	state{7, "Maryland"},
	state{8, "South Carolina"},
	state{9, "New Hampshire"},
	state{10, "Virginia"},
	state{11, "New York"},
	state{12, "North Carolina"},
	state{13, "Rhode Island"},
	state{14, "Vermont"},
	state{15, "Kentucky"},
	state{16, "Tennessee"},
	state{17, "Ohio"},
	state{18, "Louisiana"},
	state{19, "Indiana"},
	state{20, "Mississippi"},
	state{21, "Illinois"},
	state{22, "Alabama"},
	state{23, "Maine"},
	state{24, "Missouri"},
	state{25, "Arkansas"},
	state{26, "Michigan"},
	state{27, "Florida"},
	state{28, "Texas"},
	state{29, "Iowa"},
	state{30, "Wisconson"},
	state{31, "California"},
	state{32, "Minnesota"},
	state{33, "Oregon"},
	state{34, "Kansas"},
	state{35, "West Virginia"},
	state{36, "Nevada"},
	state{37, "Nebraska"},
	state{38, "Colorado"},
	state{39, "North Dakota"},
	state{40, "South Dakota"},
	state{41, "Montana"},
	state{42, "Washington"},
	state{43, "Idaho"},
	state{44, "Wyoming"},
	state{45, "Utah"},
	state{46, "Oklahoma"},
	state{47, "New Mexico"},
	state{48, "Arizona"},
	state{49, "Alaska"},
	state{50, "Hawaii"},
}

type statesQuestion func([]state) promptAndResponse

func quizStates(cmd *cobra.Command, args []string) {

	var promptFuncs = []statesQuestion{
		quizOrderToName,
		quizNameToOrder,
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

func randomState(states []state) state {
	return states[rand.Intn(len(states))]
}

func init() {
	memoryquizCmd.AddCommand(statesCmd)
}
