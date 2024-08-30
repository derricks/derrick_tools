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
	"github.com/spf13/cobra"
)

var canadaCmd = &cobra.Command{
	Use:   "canada",
	Short: "Quiz territories and provinces of Canada",
	Run:   quizCanada,
}

type canadaRegion struct {
	orderBySize int    `crossquery:"all" crossqueryname:"size rank"`
	name        string `crossquery:"all"`
	capital     string `crossquery:"all"`
}

var canadianRegions = []canadaRegion{
	{1, "Nunavut", "Iqaluit"},
	{2, "Quebec", "Quebec City"},
	{3, "Northwest Territory", "Yellowknife"},
	{4, "Ontario", "Toronto"},
	{5, "British Columbia", "Victoria"},
	{6, "Alberta", "Edmonton"},
	{7, "Saskatchewan", "Regina"},
	{8, "Manitoba", "Winnipeg"},
	{9, "Yukon", "Whitehorse"},
	{10, "New Foundland and Labrador", "St. John's"},
	{11, "New Brunswick", "Fredericton"},
	{12, "Nova Scotia", "Halifax"},
	{13, "Prince Edward Island", "Charlottetown"},
}

type canadaQuestion func([]canadaRegion) promptAndResponse

func quizCanada(cmd *cobra.Command, args []string) {

	var promptFuncs = []canadaQuestion{
		crossQueryCanadaInfo,
	}

	function := randomItemFromSlice(promptFuncs)
	promptAndCheckResponse(function(canadianRegions))
}

func crossQueryCanadaInfo(regions []canadaRegion) promptAndResponse {
	region := randomItemFromSlice(regions)
	return constructCrossQuery("Canadian region", region)
}

func init() {
	memoryquizCmd.AddCommand(canadaCmd)
}
