/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var fileName string
var delimiter string
var title string
var screenWidth int

// asciichartCmd represents the asciichart command
var asciichartCmd = &cobra.Command{
	Use:   "asciichart",
	Short: "A variety of tools for presenting data in ASCII-art charts",
	Long: `The commands here work with either stdin or a file specified with -f.

  Data should be separated as follows:
	   item<delimiter>quantity

	When "item" is repeated, the "quantity" fields are added together.
	When "quantity" is empty, it will be treated as "1". If any data causes a problem, the
	command will error out. If there's more than one "quantity" on a line, any beyond
	the first will be ignored.

	The command will update the graph as long as data is coming in, allowing this to dynamically
  show data from a stream.

	`,
}

var barChartCmd = &cobra.Command{
	Use:   "bar",
	Short: "Create an ASCII-art bar chart.",
	Long:  "See the long description under asciichart for more information",
	Run:   generateBarChart,
}

type dataPoint struct {
	label    string
	quantity int
}

type dataSet map[string]*dataPoint

func (data dataSet) addDataPoint(label string, value int) {
	currentDataPoint, exists := data[label]
	if !exists {
		data[label] = &dataPoint{label, value}
	} else {
		currentDataPoint.quantity += value
	}
}

// longestLabel returns the longest label in the data set
func (data dataSet) longestLabel() string {
	currentLabel := ""
	for key, _ := range data {
		if len(key) > len(currentLabel) {
			currentLabel = key
		}
	}
	return currentLabel
}

// largestValue returns the largest value in the data set.
func (data dataSet) largestValue() int {
	value := math.MinInt32
	for _, dataPoint := range data {
		if dataPoint.quantity > value {
			value = dataPoint.quantity
		}
	}
	return value
}

func generateBarChart(command *cobra.Command, args []string) {
	if title != "" {
		fmt.Println(centerTextInSpace(title, screenWidth))
		fmt.Println(centerTextInSpace(strings.Repeat("-", len(title)), screenWidth))
	} else {
		fmt.Println("")
	}
	// ensure that labels show up in a consistent order. otherwise they bounce around per invocation
	labelOrder := make([]string, 0)

	collectData(func(data dataSet) {
		// move the cursor back the size of dataSet rows (to the top of where the bars will draw)
		fmt.Printf("\u001b[%dA", len(data))
		fmt.Printf("\u001b[1000D")
	}, func(data dataSet) {
		// make sure the format codes are sent
		// and short circuit any extra work if there's nothing to do
		if len(data) == 0 {
			fmt.Printf("\n")
		}

		// verify all keys in data are in labelOrder. Add if not.
		for label, _ := range data {
			if !isStringInSlice(label, labelOrder) {
				labelOrder = append(labelOrder, label)
			}
		}

		longestLabel := data.longestLabel()
		largestValue := data.largestValue()
		// the area you have to draw a bar (and the maximum bar width you'll have)
		// is the total width of the screen
		// minus the length of the longest label,
		// minus the length of the largest value (printed at the end of the bar),
		// minus the length of " | " which is between the label and the chart
		// minus the length of " " printed between the bar and the number
		barAreaWidth := screenWidth - len(longestLabel) - len(strconv.Itoa(largestValue)) - len(" | ") - len(" ")

		for _, label := range labelOrder {
			point, exists := data[label]
			if !exists {
				// very unlikely, but just in case
				fmt.Printf("Unexpected error: %s should have been in data set but was not\n", label)
				os.Exit(1)
			}

			barWidth := scaledBarWidth(point.quantity, largestValue, barAreaWidth)
			fmt.Printf("\u001b[2K%*s | %s %d\n", len(longestLabel), point.label, strings.Repeat("=", barWidth), point.quantity)
		}
	}, nil)
}

// scaledBarWidth maps dataValue to maxBarWidth by comparing it to maxBarValue and returns the new length of the bar
// Examples:
//   dataValue == maxBarValue, maxBarWidth is returned
//   dataValue = 4, maxBarValue = 8, maxBarWidth is 20, the return is 10 (4/8 * 20)
func scaledBarWidth(dataValue, maxBarValue, maxBarWidth int) int {
	return int(math.Round((float64(dataValue) / float64(maxBarValue)) * float64(maxBarWidth)))
}

// centerTextInSpace takes an input string, and a width, and returns a string padded with enough space to appear centered in that width
func centerTextInSpace(text string, width int) string {
	// truncate if necessary
	if len(text) > width {
		return text[0:width]
	}

	center := width / 2
	leftIndent := center - (len(text) / 2)
	return fmt.Sprintf("%s%s", strings.Repeat(" ", leftIndent), text)
}

// collectData reads from the specified stream and calls the preData and postData functions before and after the data is added, respectively
// optional reader can be used for controlling what reader is used.
func collectData(preData func(data dataSet), postData func(data dataSet), reader *bufio.Reader) {
	dataPoints := dataSet(make(map[string]*dataPoint))
	channel := make(chan string)
	if reader != nil {
		go pushReaderToChannel(reader, channel)
	} else if fileName == "" {
		go pushStdinToChannel(channel)
	} else {
		go pushFileToChannel(fileName, channel)
	}

	for currentLine := range channel {
		// skip blank lines
		if currentLine == "" {
			continue
		}

		fields := strings.Split(currentLine, delimiter)
		value := 1
		var err error
		if len(fields) > 1 {
			value, err = strconv.Atoi(fields[1])
			if err != nil {
				fmt.Printf("Bad value in data: %v\n", fields[1])
				os.Exit(1)
			}
		}
		// we know the data is good, so now invoke the callbackFunctions
		preData(dataPoints)
		dataPoints.addDataPoint(fields[0], value)
		postData(dataPoints)
	}
}

func init() {

	asciichartCmd.AddCommand(barChartCmd)
	asciichartCmd.PersistentFlags().StringVarP(&fileName, "filename", "f", "", "A file to use as the source of data. If not specified, stdin will be used")
	asciichartCmd.PersistentFlags().StringVarP(&delimiter, "delimiter", "d", " ", "The delimiter to use for separating fields.")
	asciichartCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "The title of the chart")
	asciichartCmd.PersistentFlags().IntVarP(&screenWidth, "width", "w", 120, "The maximum width of the chart")
	rootCmd.AddCommand(asciichartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// asciichartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// asciichartCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
