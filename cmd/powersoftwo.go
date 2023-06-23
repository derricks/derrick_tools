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
	"math"
	"math/rand"
	"strconv"

	"github.com/spf13/cobra"
)

// powersoftwoCmd represents the powersoftwo command
var powersoftwoCmd = &cobra.Command{
	Use:   "powersoftwo",
	Short: "Memory quiz on powers of two (up to 2^32)",
	Run:   quizPowersOfTwo,
}

func quizPowersOfTwo(cmd *cobra.Command, args []string) {
	exponent := rand.Intn(33)
	funcs := []func(int) promptAndResponse{
		quizExponentForPowerOfTwo,
		quizPowerOfTwoFromExponent,
		quizPowerOfTwoOrderOfMagnitude,
	}
	quizFunc := funcs[rand.Intn(len(funcs))]
	promptAndCheckResponse(quizFunc(exponent))
}

func quizExponentForPowerOfTwo(exponent int) promptAndResponse {
	twoToExponent := int(math.Exp2(float64(exponent)))
	return promptAndResponse{fmt.Sprintf("What exponent for 2 gives you %d?", twoToExponent), strconv.Itoa(exponent)}
}

func quizPowerOfTwoFromExponent(exponent int) promptAndResponse {
	twoToExponent := powerOfTwoFromExponent(exponent)
	return promptAndResponse{fmt.Sprintf("What is 2^%d?", exponent), strconv.Itoa(twoToExponent)}
}

// quiz the order of magnitude (1, 10, 10000, etc) for a given power of two
func quizPowerOfTwoOrderOfMagnitude(exponent int) promptAndResponse {

	twoToExponent := powerOfTwoFromExponent(exponent)
	log := int(math.Log10(float64(twoToExponent)))
	return promptAndResponse{fmt.Sprintf("What is the order of magnitude of 2^%d", exponent), strconv.Itoa(int(math.Pow10(log)))}
}

func powerOfTwoFromExponent(exponent int) int {
	return int(math.Exp2(float64(exponent)))
}

func init() {
	memoryquizCmd.AddCommand(powersoftwoCmd)
}
