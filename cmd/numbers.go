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
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var numberLength int

// numbersCmd represents the numbers command
var numbersCmd = &cobra.Command{
	Use:   "numbers",
	Short: "Test the ability to remember large numbers",
	Run:   quizLargeNumbers,
}

func generateNumberStringOfLength(length int) string {
	numbers := make([]string, length, length)
	for i := 0; i < length; i++ {
		numbers[i] = strconv.Itoa(rand.Intn(10))
	}
	return strings.Join(numbers, "")
}

func quizLargeNumbers(cmd *cobra.Command, args []string) {

	// generate a large number
	// display and start a timer
	stringToMemorize := generateNumberStringOfLength(numberLength)
	fmt.Println(stringToMemorize)

	startTime := time.Now()
	// ask user to hit a key when ready.
	var pressedKey string
	fmt.Println("Press enter when you've memorized the number")
	fmt.Scanln(&pressedKey)

	fmt.Print("\u001b[3A")
	fmt.Print("\u001b[1000D")
	endTime := time.Now()
	fmt.Printf("You took %.2f seconds to memorize\n", endTime.Sub(startTime).Seconds())
	guess := responseFromPrompt(promptAndResponse{"Enter the number and press the Enter key when you're done", stringToMemorize})

	if guess == stringToMemorize {
		fmt.Println("Awesome! You memorized it!")
	} else {
		fmt.Println("Original   : " + stringToMemorize)
		fmt.Println("Your guess : " + guess)

		// find score, which is the last correct character
		score := 0
		for i := 0; i < len(stringToMemorize) && i < len(guess); i++ {
			if []byte(stringToMemorize)[i] == []byte(guess)[i] {
				score = score + 1
			} else {
				break
			}
		}
		fmt.Printf("Score: %d\n", score)
	}
}

func init() {
	memoryquizCmd.AddCommand(numbersCmd)
	numbersCmd.Flags().IntVarP(&numberLength, "length", "l", 20, "length of number to present")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// numbersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// numbersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
