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
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type promptAndResponse struct {
	prompt   string
	response string
}

// promptAndCheckResponse will use promot to pose a question to the user and wait for
// a response. If correct, it will print Correct! and return true. Otherwise it will
// print the user's answer and the right answer and return false
func promptAndCheckResponse(prompt promptAndResponse) bool {
	userResponse := responseFromPrompt(prompt)
	if userResponse != "" {
		if strings.TrimSpace(userResponse) == prompt.response {
			fmt.Println("Correct!")
			return true
		} else {
			fmt.Printf("Incorrect. The right answer was %s\n", prompt.response)
			return false
		}
	}
	return false
}

func responseFromPrompt(prompt promptAndResponse) string {
	fmt.Println(prompt.prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""

}

// memoryquizCmd represents the memoryquiz command
var memoryquizCmd = &cobra.Command{
	Use:   "memoryquiz",
	Short: "Fire up various memory quizzes",
}

func init() {
	rootCmd.AddCommand(memoryquizCmd)

}
