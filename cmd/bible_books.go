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

	"github.com/spf13/cobra"
)

// greekCmd represents the greek command
var bibleCmd = &cobra.Command{
	Use:   "bible",
	Short: "Test memory of the books of the King James Bible",
	Run:   quizKingJames,
}

var bibleBooks = []string{
	"Genesis",
	"Exodus",
	"Leviticus",
	"Numbers",
	"Deuteronomy",
	"Joshua",
	"Judges",
	"Ruth",
	"1 Samuel",
	"2 Samuel",
	"1 Kings",
	"2 Kings",
	"1 Chronicles",
	"2 Chronicles",
	"Ezra",
	"Nehemiah",
	"Esther",
	"Job",
	"Psalms",
	"Proverbs",
	"Ecclesiastes",
	"Song of Solomon",
	"Isaiah",
	"Jeremiah",
	"Lamentations",
	"Ezekiel",
	"Daniel",
	"Hosea",
	"Joel",
	"Amos",
	"Obadiah",
	"Jonah",
	"Micah",
	"Nahum",
	"Habakkuk",
	"Zephaniah",
	"Haggai",
	"Zachariah",
	"Malachi",
	"Matthew",
	"Mark",
	"Luke",
	"John",
	"Acts",
	"Romans",
	"1 Corinthians",
	"2 Corinthians",
	"Galatians",
	"Ephesians",
	"Philippians",
	"Colossians",
	"1 Thessolonians",
	"2 Thessolonians",
	"1 Timothy",
	"2 Timothy",
	"Titus",
	"Philemon",
	"Hebrews",
	"James",
	"1 Peter",
	"2 Peter",
	"1 John",
	"2 John",
	"3 John",
	"Jude",
	"Revelations",
}

type quizBibleFunc func([]string) promptAndResponse

func quizKingJames(cmd *cobra.Command, args []string) {
	funcs := []quizGreekFunc{
		quizBibleBookFromPosition,
		quizPositionFromBibleBook,
		quizBibleBookBefore,
		quizBibleBookAfter,
	}

	function := funcs[rand.Intn(len(funcs))]
	promptAndCheckResponse(function(bibleBooks))
}

func quizPositionFromBibleBook(books []string) promptAndResponse {
	return quizIndexOfStringInList(books)
}

func quizBibleBookFromPosition(books []string) promptAndResponse {
	return quizStringAtIndexInList("book of the Bible", books)
}

func quizBibleBookBefore(books []string) promptAndResponse {
	index := 0
	for index == 0 {
		index = rand.Intn(len(books))
	}
	return promptAndResponse{fmt.Sprintf("What book comes before %s?", books[index]), books[index-1]}
}

func quizBibleBookAfter(books []string) promptAndResponse {
	index := rand.Intn(len(books) - 1) // -1 to ensure we don't get the last item
	return promptAndResponse{fmt.Sprintf("What book comes after %s?", books[index]), books[index+1]}
}

func init() {
	memoryquizCmd.AddCommand(bibleCmd)
}
