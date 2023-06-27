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
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var format string
var startTime string
var endTime string
var interval string

var chronoSeqCmd = &cobra.Command{
	Use:   "chrono-sequence",
	Short: "Print out an inclusive list of dates/times from one moment to another by some interval",
	Long: `
  The simplest way to use this is to pass YYYY-MM-DD strings to --start and --end, and the tool will give you each day from start to end, inclusively.

  However, you can also speciy an input format, and output format, and a skip amount (done in go duration notation)
  `,
	Run: generateChronoSequence,
}

func generateChronoSequence(cmd *cobra.Command, args []string) {

	translatedFormat := translateFormat(format)
	start, err := time.Parse(translatedFormat, startTime)
	if err != nil {
		log.Fatalf("Could not parse start time: %v", err)
	}
	current := start

	end, err := time.Parse(translatedFormat, endTime)
	if err != nil {
		log.Fatalf("Could not parse end time: %v", err)
	}

	actualDuration, err := translateDuration(interval)
	if err != nil {
		log.Fatalf("Could not translate interval to go syntax: %v", err)
	}

	skip, err := time.ParseDuration(actualDuration)
	if err != nil {
		log.Fatalf("Could not parse interval: %v", err)
	}

	// go until current time surpasses end time
	for current.Compare(end) < 1 {
		fmt.Println(current.Format(translatedFormat))
		current = current.Add(skip)
	}

}

// Do any necessary translations on the duration flag.
// The command-line arguments allow for Xd durations, but ParseDuration does not, so this code does the translation
func translateDuration(curValue string) (string, error) {
	if strings.HasSuffix(curValue, "d") {
		numDays, err := strconv.Atoi(strings.Replace(curValue, "d", "", -1))
		if err != nil {
			return "", fmt.Errorf("Could not convert day amount to number: %v", err)
		}
		return fmt.Sprintf("%dh", numDays*24), nil
	}
	return curValue, nil
}

// convert strings into the constants named in time for time formats
func translateFormat(curValue string) string {
	switch curValue {
	case "ANSIC":
		return time.ANSIC
	case "UnixDate":
		return time.UnixDate
	case "RubyDate":
		return time.RubyDate
	case "RFC822":
		return time.RFC822
	case "RFC822Z":
		return time.RFC822Z
	case "RFC850":
		return time.RFC850
	case "RFC1123":
		return time.RFC1123
	case "RFC1123Z":
		return time.RFC1123Z
	case "RFC3339":
		return time.RFC3339
	case "RFC3339Nano":
		return time.RFC3339Nano
	case "Kitchen":
		return time.Kitchen
	case "Stamp":
		return time.Stamp
	case "StampMilli":
		return time.StampMilli
	case "StampMicro":
		return time.StampMicro
	case "StampNano":
		return time.StampNano
	case "DateTime":
		return time.DateTime
	case "DateOnly":
		return time.DateOnly
	case "TimeOnly":
		return time.TimeOnly
	default:
		return curValue
	}
}

func init() {
	// todo: allow user to specify format by constant name if possible (e.g., DateOnly)
	chronoSeqCmd.Flags().StringVarP(&format, "format", "f", time.DateOnly, "The format of the time string. See https://pkg.go.dev/time#pkg-constants for options")
	chronoSeqCmd.Flags().StringVarP(&startTime, "start", "s", "", "The starting date or time to use")
	chronoSeqCmd.Flags().StringVarP(&endTime, "end", "e", "", "The ending date or time to use")
	chronoSeqCmd.Flags().StringVarP(&interval, "interval", "i", "24h", "The interval to add in each step")

	rootCmd.AddCommand(chronoSeqCmd)
}
