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

	"github.com/spf13/cobra"
)

// presidentsCmd represents the presidents command
var presidentsCmd = &cobra.Command{
	Use:   "presidents",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: quizPresidents,
}

type president struct {
	number int
	name   string
}

func quizPresidents(cmd *cobra.Command, args []string) {
	presidents := []president{
		president{1, "George Washington"},
		president{2, "John Adams"},
		president{3, "Thomas Jefferson"},
		president{4, "James Madison"},
		president{5, "James Monroe"},
		president{6, "John Quincy Adams"},
		president{7, "Andrew Jackson"},
		president{8, "Martin Van Buren"},
		president{9, "William Henry Harrison"},
		president{10, "John Tyler"},
		president{11, "James K. Polk"},
		president{12, "Zachary Taylor"},
		president{13, "Millard Fillmore"},
		president{14, "Franklin Pierce"},
		president{15, "James Buchanan"},
		president{16, "Abraham Lincoln"},
		president{17, "Andrew Johnson"},
		president{18, "Ulysses S. Grant"},
		president{19, "Rutherford B. Hayes"},
		president{20, "James Garfield"},
		president{21, "Chester A. Arthur"},
		president{22, "Grover Cleveland (22)"},
		president{23, "Benjamin Harrison"},
		president{24, "Grover Cleveland (24)"},
		president{25, "William McKinley"},
		president{26, "Theodore Roosevelt"},
		president{27, "William Howard Taft"},
		president{28, "Woodrow Wilson"},
		president{29, "Warren G. Harding"},
		president{30, "Calvin Coolidge"},
		president{31, "Herbert Hoover"},
		president{32, "Franklin Delano Roosevelt"},
		president{33, "Harry S. Truman"},
		president{34, "Dwight D. Eisenhower"},
		president{35, "John F. Kennedy"},
		president{36, "Lyndon B. Johnson"},
		president{37, "Richard M. Nixon"},
		president{38, "Gerald Ford"},
		president{39, "Jimmy Carter"},
		president{40, "Ronald Reagan"},
		president{41, "George H. W. Bush"},
		president{42, "Bill Clinton"},
		president{43, "George W. Bush"},
		president{44, "Barack Obama"},
		president{45, "Donald Trump"},
		president{46, "Joseph R. Biden"},
	}

	promptFuncs := []presidentQuestion{
		quizIndex,
		quizBefore,
		quizAfter,
		quizWhichNumber,
	}

	function := promptFuncs[rand.Intn(len(promptFuncs))]
	promptAndCheckResponse(function(presidents))
}

type presidentQuestion func([]president) promptAndResponse

func quizIndex(presidents []president) promptAndResponse {
	president := presidents[rand.Intn(len(presidents))]
	return promptAndResponse{fmt.Sprintf("Who was president %d?", president.number), president.name}
}

func quizBefore(presidents []president) promptAndResponse {
	index := 0
	for index == 0 {
		index = rand.Intn(len(presidents))
	}
	return promptAndResponse{fmt.Sprintf("Who was President before %s?", presidents[index].name), presidents[index-1].name}
}

func quizAfter(presidents []president) promptAndResponse {
	index := len(presidents) - 1
	for index == len(presidents)-1 {
		index = rand.Intn(len(presidents))
	}
	return promptAndResponse{fmt.Sprintf("Who was President after %s?", presidents[index].name), presidents[index+1].name}
}

func quizWhichNumber(presidents []president) promptAndResponse {
	president := presidents[rand.Intn(len(presidents))]
	return promptAndResponse{fmt.Sprintf("What number president was %s?", president.name), strconv.Itoa(president.number)}
}

func init() {
	memoryquizCmd.AddCommand(presidentsCmd)
}
