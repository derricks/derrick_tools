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

var hebrewCalendarCmd = &cobra.Command{
	Use:   "hebrew-calendar",
	Short: "Quiz Hebrew Calendar",
	Run:   quizHebrewCalendar,
}

type hebrewCalendar struct {
	hebrewIndex    int    `crossquery:"all" crossqueryname:"index in the Hebrew Calendar"`
	hebrewMonth    string `crossquery:"all" crossqueryname:"Hebrew month"`
	gregorianMonth string `crossquery:"all" crossqueryname:"Gregorian month"`
}

var hebrewMonths = []hebrewCalendar{
	{1, "Nisan", "March"},
	{2, "Iyar", "April"},
	{3, "Sivan", "May"},
	{4, "Tammuz", "June"},
	{5, "Av", "July"},
	{6, "Elul", "August"},
	{7, "Tishrei", "September"},
	{8, "Heshvan", "October"},
	{9, "Kislev", "November"},
	{10, "Tevet", "December"},
	{11, "Shivat", "January"},
	{12, "Adar", "February"},
}

type hebrewCalendarQuestion func([]hebrewCalendar) promptAndResponse

func quizHebrewCalendar(cmd *cobra.Command, args []string) {
	promptAndCheckResponse(crossQueryHebrewCalendar(hebrewMonths))
}

func crossQueryHebrewCalendar(months []hebrewCalendar) promptAndResponse {
	foundMonth := randomItemFromSlice(months)
	return constructCrossQuery("Hebrew calendar", foundMonth)
}

func init() {
	memoryquizCmd.AddCommand(hebrewCalendarCmd)
}
