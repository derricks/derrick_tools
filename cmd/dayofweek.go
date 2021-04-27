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
	"time"

	"github.com/spf13/cobra"
)

// dayofweekCmd represents the dayofweek command
var dayofweekCmd = &cobra.Command{
	Use:   "dayofweek",
	Short: "Tests the ability to figure out the day of week for a given date",
	Long: `
Conway's method for finding the day of week for any given year

1. Figure out the modifier for the given century (the four centuries repeat, so 1700-1799 = 2100-2199):
	- 1800 - 1899: 5
	- 1900 - 1999: 3
	- 2000 - 2099: 2
	- 2100 - 2199: 0

2. Divide the year (without the century) by 12
3. Figure out the remainder from step 2
4. Divide the answer from step 3 by 4 and ignore the remainder
5. Add the results from steps 1, 2, 3, 4.
6. Mod the result of step 5 by 7
7. The answer from step 6 gives the day of the week that Doomsday falls on
	- 0 = Sunday
	- 1 = Monday
	- 2 = Tuesday
	- 3 = Wednesday
	- 4 = Thursday
	- 5 = Friday
	- 6 = Saturday

 8. Find the date you want relative to the doomsday in a given month
 	- January = 3 or 4 (leap year)
 	- February = the last day of February
	- March 7th
 	- April 4th
 	- May 9th
 	- June 6th
 	- July 11th
 	- August 8th
 	- September 5th
 	- October 10th
 	- November 7th
 	- December 12th`,
	Run: quizDayOfWeekCalculation,
}

func quizDayOfWeekCalculation(cmd *cobra.Command, args []string) {
	century := rand.Intn(4) + 18
	twoDigitYear := rand.Intn(100)
	year := (century * 100) + twoDigitYear
	month := rand.Intn(12) + 1
	day := rand.Intn(31) + 1

	// note Date will do the right thing if, for instance, you pass September 31; it will set it to October 1.
	// so we can just give it the date and let it figure it out
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	promptAndCheckResponse(promptAndResponse{fmt.Sprintf("What day of the week does %s fall on?", fmt.Sprintf("%d/%d/%d", date.Month(), date.Day(), date.Year())), fmt.Sprintf("%v", date.Weekday())})
}

func init() {
	speedmathCmd.AddCommand(dayofweekCmd)
}
