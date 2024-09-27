/*
Copyright © 2024 Derrick Schneider derrick.schneider@gmail.com
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
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http-codes",
	Short: "Quiz HTTP Error Codes",
	Run:   quizHttpCodes,
}

type httpCode struct {
	code    int    `crossquery:"all"`
	message string `crossquery:"all"`
}

var httpCodes = []httpCode{
	{100, "Continue"},
	{101, "Switching protocols"},
	{102, "Processing"},
	{103, "Early hints"},
	{200, "OK"},
	{201, "Created"},
	{202, "Accepted"},
	{203, "Non-authoritative information"},
	{204, "No content"},
	{205, "Reset content"},
	{206, "Partial content"},
	{207, "Multi-Status"},
	{208, "Already Reported"},
	{226, "IM Used"},
	{300, "Multiple choices"},
	{301, "Moved permanently"},
	{302, "Found"},
	{303, "See other"},
	{304, "Not modified"},
	{305, "Use proxy"},
	{306, "Switch proxy"},
	{307, "Temporary redirect"},
	{308, "Permanent redirect"},
}

type httpCodeQuestion func([]httpCode) promptAndResponse

func quizHttpCodes(cmd *cobra.Command, args []string) {

	var promptFuncs = []httpCodeQuestion{
		crossQueryHttpCodeInfo,
	}

	function := randomItemFromSlice(promptFuncs)
	promptAndCheckResponse(function(httpCodes))
}

func crossQueryHttpCodeInfo(codes []httpCode) promptAndResponse {
	foundCode := randomItemFromSlice(codes)
	return constructCrossQuery("HTTP", foundCode)
}

func init() {
	memoryquizCmd.AddCommand(httpCmd)
}
