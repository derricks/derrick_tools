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
	"time"

	"github.com/spf13/cobra"
)

// powersoftwoCmd represents the powersoftwo command
var doomsdayCmd = &cobra.Command{
	Use:   "doomsday",
	Short: "Memory quiz on doomsdays for given years",
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
`,
	Run: quizDoomsday,
}

func quizDoomsday(cmd *cobra.Command, args []string) {
	yearsAfter1800 := rand.Intn(400)
	year := 1800 + yearsAfter1800
	doomsdayDate := time.Date(year, 12, 12, 0, 0, 0, 0, time.UTC)
	dayOfWeek := doomsdayDate.Weekday().String()
	promptAndCheckResponse(promptAndResponse{fmt.Sprintf("What day of the week is the doomsday for %d?", year), dayOfWeek})
}

func init() {
	memoryquizCmd.AddCommand(doomsdayCmd)
}
