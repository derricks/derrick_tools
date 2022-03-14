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
	"strings"

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
	number         int
	name           string
	startYear      int
	vicePresidents []string
	firstLadies    []string
}

func quizPresidents(cmd *cobra.Command, args []string) {
	presidents := []president{
		president{1, "George Washington", 1789, []string{"John Adams"}, []string{"Martha Washington"}},
		president{2, "John Adams", 1797, []string{"Thomas Jefferson"}, []string{"Abigail Adams"}},
		president{3, "Thomas Jefferson", 1801, []string{"Aaron Burr", "George Clinton"}, []string{"Martha Jefferson"}},
		president{4, "James Madison", 1809, []string{"George Clinton", "Elbridge Gerry"}, []string{"Dolley Madison"}},
		president{5, "James Monroe", 1817, []string{"Daniel Tompkins"}, []string{"Eliizabeth Monroe"}},
		president{6, "John Quincy Adams", 1825, []string{"John C. Calhoun"}, []string{"Louisa Adams"}},
		president{7, "Andrew Jackson", 1829, []string{"John C. Calhoun", "Martin Van Buren"}, []string{"Rachel Jackson", "Emily Donelson"}},
		president{8, "Martin Van Buren", 1837, []string{"Richard Mentor Johnson"}, []string{"Hannah Van Buren", "Angelica Van Buren"}},
		president{9, "William Henry Harrison", 1841, []string{"John Tyler"}, []string{"Anna Harrison", "Jane Harrison"}},
		president{10, "John Tyler", 1841, []string{}, []string{"Letitia Tyler", "Julia Tyler"}},
		president{11, "James K. Polk", 1845, []string{"George Dallas"}, []string{"Sarah Polk"}},
		president{12, "Zachary Taylor", 1849, []string{"Millard Fillmore"}, []string{"Margaret Taylor"}},
		president{13, "Millard Fillmore", 1850, []string{}, []string{"Abigail Powers Fillmore"}},
		president{14, "Franklin Pierce", 1853, []string{"William R. King"}, []string{"Jane Pierce"}},
		president{15, "James Buchanan", 1857, []string{"John C. Breckinridge"}, []string{"Harriet Lane"}},
		president{16, "Abraham Lincoln", 1861, []string{"Hannibal Hamlin", "Andrew Johnson"}, []string{"Mary Lincoln"}},
		president{17, "Andrew Johnson", 1865, []string{}, []string{"Eliza Johnson", "Martha Johnson Patterson"}},
		president{18, "Ulysses S. Grant", 1869, []string{"Schuyler Colfax", ""}, []string{"Julia Grant"}},
		president{19, "Rutherford B. Hayes", 1877, []string{"William Wheeler"}, []string{"Lucy Hayes"}},
		president{20, "James Garfield", 1881, []string{"Chester A. Arthur"}, []string{"Lucretia Garfield"}},
		president{21, "Chester A. Arthur", 1881, []string{}, []string{"Ellen Arthur", "Mary Arthur McElroy"}},
		president{22, "Grover Cleveland (22)", 1885, []string{"Thomas Hendricks"}, []string{"Rose Cleveland", "Frances Cleveland"}},
		president{23, "Benjamin Harrison", 1889, []string{}, []string{"Caroline Harrison"}},
		president{24, "Grover Cleveland (24)", 1893, []string{"Adlai Stevenson"}, []string{"Frances Cleveland"}},
		president{25, "William McKinley", 1897, []string{"Garret Hobart", "Theodore Roosevelt"}, []string{"Ida McKinley"}},
		president{26, "Theodore Roosevelt", 1901, []string{"Charles Fairbanks"}, []string{"Edith Roosevelt"}},
		president{27, "William Howard Taft", 1909, []string{"James Sherman"}, []string{"Helen Taft"}},
		president{28, "Woodrow Wilson", 1913, []string{"Thomas Marshall"}, []string{"Ellen Wilson", "Edith Wilson"}},
		president{29, "Warren G. Harding", 1921, []string{"Calvin Coolidge"}, []string{"Florence Harding"}},
		president{30, "Calvin Coolidge", 1923, []string{"Charles Dawes"}, []string{"Grace Coolidge"}},
		president{31, "Herbert Hoover", 1929, []string{"Charles Curtis"}, []string{"Lou Hoover"}},
		president{32, "Franklin Delano Roosevelt", 1933, []string{"John Garner", "Henry Wallace", "Harry S. Truman"}, []string{"Eleanor Roosevelt"}},
		president{33, "Harry S. Truman", 1945, []string{"Alben Barkley"}, []string{"Elizabeth 'Bess' Truman"}},
		president{34, "Dwight D. Eisenhower", 1953, []string{"Richard Nixon"}, []string{"Mamie Eisenhower"}},
		president{35, "John F. Kennedy", 1961, []string{"Lyndon B. Johnson"}, []string{"Jacqueline Kennedy"}},
		president{36, "Lyndon B. Johnson", 1963, []string{"Hubert Humphrey"}, []string{"Claudia 'Ladybird' Johnson"}},
		president{37, "Richard M. Nixon", 1969, []string{"Spiro Agnew", "Gerald Ford"}, []string{"Patricia Nixon"}},
		president{38, "Gerald Ford", 1974, []string{"Nelson Rockefeller"}, []string{"Betty Ford"}},
		president{39, "Jimmy Carter", 1977, []string{"Walter Mondale"}, []string{"Rosalynn Carter"}},
		president{40, "Ronald Reagan", 1981, []string{"George H. W. Bush"}, []string{"Nancy Reagan"}},
		president{41, "George H. W. Bush", 1989, []string{"Dan Quayle"}, []string{"Barbara Bush"}},
		president{42, "Bill Clinton", 1993, []string{"Al Gore"}, []string{"Hillary Clinton"}},
		president{43, "George W. Bush", 2001, []string{"Dick Cheney"}, []string{"Laura Bush"}},
		president{44, "Barack Obama", 2009, []string{"Joseph R. Biden"}, []string{"Michelle Obama"}},
		president{45, "Donald Trump", 2017, []string{"Mike Pence"}, []string{"Melania Trump"}},
		president{46, "Joseph R. Biden", 2021, []string{"Kamala Harris"}, []string{"Dr. Jill Biden"}},
	}

	var promptFuncs []presidentQuestion

	if vicePresidentsOnly {
		promptFuncs = []presidentQuestion{
			quizVicePresidents,
			quizPresidentsForVicePresident,
		}
	} else {
		promptFuncs = []presidentQuestion{
			quizIndex,
			quizBefore,
			quizAfter,
			quizWhichNumber,
			quizWhenPresidentStarted,
			quizWhenPresidentEnded,
			quizWhoWasPresidentWhen,
			quizVicePresidents,
			quizPresidentsForVicePresident,
			quizFirstLadiesFromPresident,
			quizPresidentFromFirstLady,
		}
	}

	function := promptFuncs[rand.Intn(len(promptFuncs))]
	promptAndCheckResponse(function(presidents))
}

var vicePresidentsOnly bool

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

func quizVicePresidents(presidents []president) promptAndResponse {
	president := randomPresident(presidents)
	for len(president.vicePresidents) == 0 {
		president = randomPresident(presidents)
	}
	return promptAndResponse{fmt.Sprintf("Who served as vice president under %s (separate names with commas)?", president.name), strings.Join(president.vicePresidents, ",")}
}

// the complicated logic here is because some vice presidents served under more than one president
func quizPresidentsForVicePresident(presidents []president) promptAndResponse {
	p := randomPresident(presidents)
	// not all presidents had vice presidents
	for len(p.vicePresidents) == 0 {
		p = randomPresident(presidents)
	}

	vp := p.vicePresidents[rand.Intn(len(p.vicePresidents))]
	presList := make([]string, 0)

	for _, president := range presidents {
		if vpServedUnderPres(vp, president) {
			presList = append(presList, president.name)
		}
	}
	return promptAndResponse{fmt.Sprintf("Which Presidents did %s serve under as Vice President? (Separate names with commas)", vp), strings.Join(presList, ",")}
}

func quizFirstLadiesFromPresident(presidents []president) promptAndResponse {
	p := randomPresident(presidents)
	for len(p.firstLadies) == 0 {
		// not a necessity now, but future-proofing
		p = randomPresident(presidents)
	}
	return promptAndResponse{fmt.Sprintf("Who were %s's First Ladies (join with commas)?", p.name), strings.Join(p.firstLadies, ",")}
}

func quizPresidentFromFirstLady(presidents []president) promptAndResponse {
	p := randomPresident(presidents)
	fl := p.firstLadies[rand.Intn(len(p.firstLadies))]
	return promptAndResponse{fmt.Sprintf("Who did %s serve as First Lady?", fl), p.name}
}

func vpServedUnderPres(vp string, pres president) bool {
	for _, curVp := range pres.vicePresidents {
		if curVp == vp {
			return true
		}
	}
	return false
}

func randomPresident(presidents []president) president {
	return presidents[rand.Intn(len(presidents))]
}

func init() {
	presidentsCmd.Flags().BoolVarP(&vicePresidentsOnly, "vicepresidents", "", false, "If set, only ask questions about vice presidents")
	memoryquizCmd.AddCommand(presidentsCmd)
}
