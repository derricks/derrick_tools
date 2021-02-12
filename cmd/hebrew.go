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
  "math/rand"

	"github.com/spf13/cobra"
)

// hebrewCmd represents the hebrew command
var hebrewCmd = &cobra.Command{
	Use:   "hebrew",
	Short: "Quiz command of hebrew alphabet",
	Run:   quizHebrewAlphabet,
}

var hebrewAlphabet = []string{
	"aleph",
	"bet",
	"gimel",
	"dalet",
	"he",
	"vav",
	"zayin",
	"het",
	"tet",
	"yod",
	"kaf",
	"lamed",
	"mem",
	"nun",
	"samekh",
	"ayin",
	"pe",
	"tsadi",
	"qof",
	"resh",
	"shin",
	"tav",
}

type quizHebrewFunc func([]string) promptAndResponse

func quizHebrewAlphabet(cmd *cobra.Command, args []string) {
	funcs := []quizHebrewFunc{
		quizPositionFromLetter,
		quizLetterFromPosition,
		quizLetterBefore,
		quizLetterAfter,
	}

	function := funcs[rand.Intn(len(funcs))]
	promptAndCheckResponse(function(hebrewAlphabet))
}

func init() {
	memoryquizCmd.AddCommand(hebrewCmd)
}
