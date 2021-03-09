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
	Run:   quizDayOfWeekCalculation,
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
