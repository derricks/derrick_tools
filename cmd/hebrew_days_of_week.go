/*
Copyright © 2024 Derrick Schneider derrick.schneider@gmail.com
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

var hebrewWeekCmd = &cobra.Command{
	Use:   "hebrew-week",
	Short: "Quiz Hebrew days of week",
	Run:   quizHebrewWeek,
}

type hebrewDayOfWeek struct {
	index        int    `crossquery:"all" crossqueryname:"index"`
	hebrewName   string `crossquery:"all" crossqueryname:"Hebrew name"`
	englishMonth string `crossquery:"all" crossqueryname:"English name"`
}

var hebrewWeek = []hebrewDayOfWeek{
	{1, "Rishon", "Sunday"},
	{2, "Sheni", "Monday"},
	{3, "Shlishi", "Tuesday"},
	{4, "Revi'i", "Wednesday"},
	{5, "Chamishi", "Thursday"},
	{6, "Shishi", "Friday"},
	{7, "Shabbat", "Saturday"},
}

func quizHebrewWeek(cmd *cobra.Command, args []string) {
	promptAndCheckResponse(crossQueryHebrewWeek(hebrewWeek))
}

func crossQueryHebrewWeek(daysOfWeek []hebrewDayOfWeek) promptAndResponse {
	foundDay := randomItemFromSlice(hebrewWeek)
	return constructCrossQuery("Hebrew day", foundDay)
}

func init() {
	memoryquizCmd.AddCommand(hebrewWeekCmd)
}
