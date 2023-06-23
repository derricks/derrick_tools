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

var chronoSeqCmd = &cobra.Command{
	Use:   "chrono-sequence",
	Short: "Print out a list of dates/times from one moment to another by some interval",
	Long: `
  The simplest way to use this is to pass YYYY-MM-DD strings to --start and --end, and the tool will give you each day from start to end, inclusively.

  However, you can also speciy an input format, and output format, and a skip amount (done in go duration notation)
  `,
	Run: generateChronoSequence,
}

func generateChronoSequence(cmd *cobra.Command, args []string) {

}

func init() {
	rootCmd.AddCommand(chronoSeqCmd)
}
