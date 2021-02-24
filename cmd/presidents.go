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
	number    int
	name      string
	startYear int
}

func quizPresidents(cmd *cobra.Command, args []string) {
	presidents := []president{
		president{1, "George Washington", 1789},
		president{2, "John Adams", 1797},
		president{3, "Thomas Jefferson", 1801},
		president{4, "James Madison", 1809},
		president{5, "James Monroe", 1817},
		president{6, "John Quincy Adams", 1825},
		president{7, "Andrew Jackson", 1829},
		president{8, "Martin Van Buren", 1837},
		president{9, "William Henry Harrison", 1841},
		president{10, "John Tyler", 1841},
		president{11, "James K. Polk", 1845},
		president{12, "Zachary Taylor", 1849},
		president{13, "Millard Fillmore", 1850},
		president{14, "Franklin Pierce", 1853},
		president{15, "James Buchanan", 1857},
		president{16, "Abraham Lincoln", 1861},
		president{17, "Andrew Johnson", 1865},
		president{18, "Ulysses S. Grant", 1869},
		president{19, "Rutherford B. Hayes", 1877},
		president{20, "James Garfield", 1881},
		president{21, "Chester A. Arthur", 1881},
		president{22, "Grover Cleveland (22)", 1885},
		president{23, "Benjamin Harrison", 1889},
		president{24, "Grover Cleveland (24)", 1893},
		president{25, "William McKinley", 1897},
		president{26, "Theodore Roosevelt", 1901},
		president{27, "William Howard Taft", 1909},
		president{28, "Woodrow Wilson", 1913},
		president{29, "Warren G. Harding", 1921},
		president{30, "Calvin Coolidge", 1923},
		president{31, "Herbert Hoover", 1929},
		president{32, "Franklin Delano Roosevelt", 1933},
		president{33, "Harry S. Truman", 1945},
		president{34, "Dwight D. Eisenhower", 1953},
		president{35, "John F. Kennedy", 1961},
		president{36, "Lyndon B. Johnson", 1963},
		president{37, "Richard M. Nixon", 1969},
		president{38, "Gerald Ford", 1974},
		president{39, "Jimmy Carter", 1977},
		president{40, "Ronald Reagan", 1981},
		president{41, "George H. W. Bush", 1989},
		president{42, "Bill Clinton", 1993},
		president{43, "George W. Bush", 2001},
		president{44, "Barack Obama", 2009},
		president{45, "Donald Trump", 2017},
		president{46, "Joseph R. Biden", 2021},
	}

	promptFuncs := []presidentQuestion{
		quizIndex,
		quizBefore,
		quizAfter,
		quizWhichNumber,
		quizWhenPresidentStarted,
		quizWhenPresidentEnded,
		quizWhoWasPresidentWhen,
	}

	function := promptFuncs[rand.Intn(len(promptFuncs))]
	promptAndCheckResponse(function(presidents))
}

type presidentQuestion func([]president) promptAndResponse

func quizIndex(presidents []president) promptAndResponse {
	president := randomPresident(presidents)
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
	president := randomPresident(presidents)
	return promptAndResponse{fmt.Sprintf("What number president was %s?", president.name), strconv.Itoa(president.number)}
}

func quizWhenPresidentStarted(presidents []president) promptAndResponse {
	president := randomPresident(presidents)
	return promptAndResponse{fmt.Sprintf("What was the first year of %s's presidency?", president.name), strconv.Itoa(president.startYear)}
}

func quizWhenPresidentEnded(presidents []president) promptAndResponse {
	presidentIndex := rand.Intn(len(presidents) - 1)
	president := presidents[presidentIndex]
	nextPresident := presidents[presidentIndex+1]
	return promptAndResponse{fmt.Sprintf("What was the last year of %s's presidency?", president.name), strconv.Itoa(nextPresident.startYear)}
}

// ask who was president in a given year
func quizWhoWasPresidentWhen(presidents []president) promptAndResponse {
	// figure out two presidents more than two years apart so you can have an unambiguous year
	// asking who was president in 1841, for instance, has three answers, depending on what part of the year
	var president1, president2 president
	distanceBetweenStarts := 0
	for distanceBetweenStarts < 2 {
		// it doesn't make sense to ask for the next president for the one currently in office
		presidentIndex := rand.Intn(len(presidents) - 1)
		president1 = presidents[presidentIndex]
		president2 = presidents[presidentIndex+1]
		distanceBetweenStarts = president2.startYear - president1.startYear
	}

	// now we want to figure out an offset from president1's start year to figure out the actual year
	// we'll query about.
	offsetFromCurrentPresident := 0
	for offsetFromCurrentPresident == 0 {
		offsetFromCurrentPresident = rand.Intn(president2.startYear - president1.startYear)
	}
	return promptAndResponse{fmt.Sprintf("Who was president in %d?", president1.startYear+offsetFromCurrentPresident), president1.name}
}

func randomPresident(presidents []president) president {
	return presidents[rand.Intn(len(presidents))]
}

func init() {
	memoryquizCmd.AddCommand(presidentsCmd)
}
