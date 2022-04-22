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
	"strconv"

	"github.com/spf13/cobra"
)

// speedmathCmd represents the speedmath command
var speedmathCmd = &cobra.Command{
	Use:   "speedmath",
	Short: "Quizzes to test speed math abilities",
	Run:   speedMathTesting,
}

type speedMathFunc func() promptAndResponse

func speedMathTesting(cmd *cobra.Command, args []string) {
	speedMathFuncs := []speedMathFunc{
		speedMathAddition,
		speedMathSubtraction,
		speedMath1xNMultiplication,
		speedMathSquareTwoDigits,
		speedMath2x2Multiplication,
		speedMathSquareThreeDigits,
		speedMathCubeTwoDigits,
	}

	mathFunc := speedMathFuncs[rand.Intn(len(speedMathFuncs))]
	promptAndCheckResponse(mathFunc())
}

func speedMathAddition() promptAndResponse {
	addend1 := rand.Intn(10000)
	addend2 := rand.Intn(10000)
	return promptAndResponse{fmt.Sprintf("%d + %d = ", addend1, addend2), strconv.Itoa(addend1 + addend2)}
}

func speedMathSubtraction() promptAndResponse {
	minuend := rand.Intn(9900)
	minuend += 100                   // ensure that minuend is always a reasonably sized number
	subtrahend := rand.Intn(minuend) // ensure that subtrahend is always smaller
	return promptAndResponse{fmt.Sprintf("%d - %d = ", minuend, subtrahend), strconv.Itoa(minuend - subtrahend)}
}

func speedMath1xNMultiplication() promptAndResponse {
	factor1 := rand.Intn(1000)
	factor2 := rand.Intn(10)
	return promptAndResponse{fmt.Sprintf("%d * %d = ", factor1, factor2), strconv.Itoa(factor1 * factor2)}
}

func speedMathSquareTwoDigits() promptAndResponse {
	base := twoDigitNumber()
	return promptAndResponse{fmt.Sprintf("%d^2 = ", base), strconv.Itoa(base * base)}
}

func speedMath2x2Multiplication() promptAndResponse {
	factor1 := twoDigitNumber()
	factor2 := twoDigitNumber()
	return promptAndResponse{fmt.Sprintf("%d * %d =", factor1, factor2), strconv.Itoa(factor1 * factor2)}
}

func speedMathSquareThreeDigits() promptAndResponse {
	base := randNumberBetween(100, 1000)
	return promptAndResponse{fmt.Sprintf("%d^2 = ", base), strconv.Itoa(base * base)}
}

func speedMathCubeTwoDigits() promptAndResponse {
	base := twoDigitNumber()
	return promptAndResponse{fmt.Sprintf("%d^3 = ", base), strconv.Itoa(base * base * base)}
}

func twoDigitNumber() int {
	return randNumberBetween(10, 100)
}

// return a random number between the lower number (inclusive) and the upper nuumber (exclusive)
func randNumberBetween(lower, upper int) int {
	if lower == upper {
		return lower
	}

	if lower > upper {
		return rand.Intn(lower-upper) + upper
	} else {
		return rand.Intn(upper-lower) + lower
	}
}

func init() {
	rootCmd.AddCommand(speedmathCmd)
}
