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
	Short: "Memory quizzes about presidents",
	Run:   quizPresidents,
}

type president struct {
	number         int    `crossquery:"all"`
	name           string `crossquery:"all"`
	startYear      int    `crossquery:"all" crossqueryname:"first year of Presidency"`
	vicePresidents []string
	firstLadies    []string `crossquery:"given" crossqueryname:"First Lady"`
}

func quizPresidents(cmd *cobra.Command, args []string) {
	presidents := []president{
		{1, "George Washington", 1789, []string{"John Adams"}, []string{"Martha Washington"}},
		{2, "John Adams", 1797, []string{"Thomas Jefferson"}, []string{"Abigail Adams"}},
		{3, "Thomas Jefferson", 1801, []string{"Aaron Burr", "George Clinton"}, []string{"Martha Jefferson"}},
		{4, "James Madison", 1809, []string{"George Clinton", "Elbridge Gerry"}, []string{"Dolley Madison"}},
		{5, "James Monroe", 1817, []string{"Daniel Tompkins"}, []string{"Elizabeth Monroe"}},
		{6, "John Quincy Adams", 1825, []string{"John C. Calhoun"}, []string{"Louisa Adams"}},
		{7, "Andrew Jackson", 1829, []string{"John C. Calhoun", "Martin Van Buren"}, []string{"Rachel Jackson", "Emily Donelson"}},
		{8, "Martin Van Buren", 1837, []string{"Richard Mentor Johnson"}, []string{"Hannah Van Buren", "Angelica Van Buren"}},
		{9, "William Henry Harrison", 1841, []string{"John Tyler"}, []string{"Anna Harrison", "Jane Harrison"}},
		{10, "John Tyler", 1841, []string{}, []string{"Letitia Tyler", "Julia Tyler"}},
		{11, "James K. Polk", 1845, []string{"George Dallas"}, []string{"Sarah Polk"}},
		{12, "Zachary Taylor", 1849, []string{"Millard Fillmore"}, []string{"Margaret Taylor"}},
		{13, "Millard Fillmore", 1850, []string{}, []string{"Abigail Powers Fillmore"}},
		{14, "Franklin Pierce", 1853, []string{"William R. King"}, []string{"Jane Pierce"}},
		{15, "James Buchanan", 1857, []string{"John C. Breckinridge"}, []string{"Harriet Lane"}},
		{16, "Abraham Lincoln", 1861, []string{"Hannibal Hamlin", "Andrew Johnson"}, []string{"Mary Lincoln"}},
		{17, "Andrew Johnson", 1865, []string{}, []string{"Eliza Johnson", "Martha Johnson Patterson"}},
		{18, "Ulysses S. Grant", 1869, []string{"Schuyler Colfax", "Henry Wilson"}, []string{"Julia Grant"}},
		{19, "Rutherford B. Hayes", 1877, []string{"William Wheeler"}, []string{"Lucy Hayes"}},
		{20, "James Garfield", 1881, []string{"Chester A. Arthur"}, []string{"Lucretia Garfield"}},
		{21, "Chester A. Arthur", 1881, []string{}, []string{"Ellen Arthur", "Mary Arthur McElroy"}},
		{22, "Grover Cleveland (22)", 1885, []string{"Thomas Hendricks"}, []string{"Rose Cleveland", "Frances Cleveland"}},
		{23, "Benjamin Harrison", 1889, []string{}, []string{"Caroline Harrison"}},
		{24, "Grover Cleveland (24)", 1893, []string{"Adlai Stevenson"}, []string{"Frances Cleveland"}},
		{25, "William McKinley", 1897, []string{"Garret Hobart", "Theodore Roosevelt"}, []string{"Ida McKinley"}},
		{26, "Theodore Roosevelt", 1901, []string{"Charles Fairbanks"}, []string{"Edith Roosevelt"}},
		{27, "William Howard Taft", 1909, []string{"James Sherman"}, []string{"Helen Taft"}},
		{28, "Woodrow Wilson", 1913, []string{"Thomas Marshall"}, []string{"Ellen Wilson", "Edith Wilson"}},
		{29, "Warren G. Harding", 1921, []string{"Calvin Coolidge"}, []string{"Florence Harding"}},
		{30, "Calvin Coolidge", 1923, []string{"Charles Dawes"}, []string{"Grace Coolidge"}},
		{31, "Herbert Hoover", 1929, []string{"Charles Curtis"}, []string{"Lou Hoover"}},
		{32, "Franklin Delano Roosevelt", 1933, []string{"John Garner", "Henry Wallace", "Harry S. Truman"}, []string{"Eleanor Roosevelt"}},
		{33, "Harry S. Truman", 1945, []string{"Alben Barkley"}, []string{"Elizabeth 'Bess' Truman"}},
		{34, "Dwight D. Eisenhower", 1953, []string{"Richard Nixon"}, []string{"Mamie Eisenhower"}},
		{35, "John F. Kennedy", 1961, []string{"Lyndon B. Johnson"}, []string{"Jacqueline Kennedy"}},
		{36, "Lyndon B. Johnson", 1963, []string{"Hubert Humphrey"}, []string{"Claudia 'Ladybird' Johnson"}},
		{37, "Richard M. Nixon", 1969, []string{"Spiro Agnew", "Gerald Ford"}, []string{"Patricia Nixon"}},
		{38, "Gerald Ford", 1974, []string{"Nelson Rockefeller"}, []string{"Betty Ford"}},
		{39, "Jimmy Carter", 1977, []string{"Walter Mondale"}, []string{"Rosalynn Carter"}},
		{40, "Ronald Reagan", 1981, []string{"George H. W. Bush"}, []string{"Nancy Reagan"}},
		{41, "George H. W. Bush", 1989, []string{"Dan Quayle"}, []string{"Barbara Bush"}},
		{42, "Bill Clinton", 1993, []string{"Al Gore"}, []string{"Hillary Clinton"}},
		{43, "George W. Bush", 2001, []string{"Dick Cheney"}, []string{"Laura Bush"}},
		{44, "Barack Obama", 2009, []string{"Joseph R. Biden"}, []string{"Michelle Obama"}},
		{45, "Donald Trump", 2017, []string{"Mike Pence"}, []string{"Melania Trump"}},
		{46, "Joseph R. Biden", 2021, []string{"Kamala Harris"}, []string{"Dr. Jill Biden"}},
		{47, "Donald Trump", 2025, []string{"JD Vance"}, []string{"Melania Trump"}},
	}

	var promptFuncs []presidentQuestion

	if vicePresidentsOnly {
		promptFuncs = []presidentQuestion{
			quizVicePresidents,
			quizPresidentsForVicePresident,
		}
	} else {
		promptFuncs = []presidentQuestion{
			crossQueryPresidentInfo,
			crossQueryPresidentInfo,
			crossQueryPresidentInfo,
			crossQueryPresidentInfo,
			quizBefore,
			quizAfter,
			quizWhenPresidentEnded,
			quizWhoWasPresidentWhen,
			quizVicePresidents,
			quizPresidentsForVicePresident,
			quizFirstLadiesFromPresident,
		}
	}

	function := randomItemFromSlice(promptFuncs)
	promptAndCheckResponse(function(presidents))
}

var vicePresidentsOnly bool

type presidentQuestion func([]president) promptAndResponse

func crossQueryPresidentInfo(presidents []president) promptAndResponse {
	president := randomItemFromSlice(presidents)
	return constructCrossQuery("president", president)
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
	president := randomItemFromSlice(presidents)
	for len(president.vicePresidents) == 0 {
		president = randomItemFromSlice(presidents)
	}
	return promptAndResponse{fmt.Sprintf("Who served as vice president under %s (separate names with commas)?", president.name), strings.Join(president.vicePresidents, ",")}
}

// the complicated logic here is because some vice presidents served under more than one president
func quizPresidentsForVicePresident(presidents []president) promptAndResponse {
	p := randomItemFromSlice(presidents)
	// not all presidents had vice presidents
	for len(p.vicePresidents) == 0 {
		p = randomItemFromSlice(presidents)
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
	p := randomItemFromSlice(presidents)
	for len(p.firstLadies) == 0 {
		// not a necessity now, but future-proofing
		p = randomItemFromSlice(presidents)
	}
	return promptAndResponse{fmt.Sprintf("Who were %s's First Ladies (join with commas)?", p.name), strings.Join(p.firstLadies, ",")}
}

func vpServedUnderPres(vp string, pres president) bool {
	for _, curVp := range pres.vicePresidents {
		if curVp == vp {
			return true
		}
	}
	return false
}

func init() {
	presidentsCmd.Flags().BoolVarP(&vicePresidentsOnly, "vicepresidents", "", false, "If set, only ask questions about vice presidents")
	memoryquizCmd.AddCommand(presidentsCmd)
}
