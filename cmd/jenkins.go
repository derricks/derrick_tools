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
	//	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var jenkinsJobUrl string
var jenkinsFollowLog bool
var jenkinsJobFlags []string
var jenkinsUser string
var jenkinsToken string

// jenkinsCmd represents the jenkins command
var jenkinsCmd = &cobra.Command{
	Use:   "jenkins",
	Short: "Utilities for working with Jenkins installations",
	Long:  "",
}

// jobCmd holds the functionality for working with Jenkins jobs
var jobCmd = &cobra.Command{
	Use:   "job",
	Short: "Tools for working with Jenkins jobs",
	Long:  "",
}

// jobStartCmd will start a given jenkins job
var jobStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a Jenkins job",
	Long:  "",
	Run:   startJenkinsJob,
}

func init() {
	jobStartCmd.Flags().BoolVarP(&jenkinsFollowLog, "follow", "f", false, "Specify whether or not to follow the log once it's kicked off")

	jobStartCmd.Flags().StringVarP(&jenkinsJobUrl, "url", "u", "", "The URL of the job to kick off")
	jobStartCmd.MarkFlagRequired("url")

	jobStartCmd.Flags().StringSliceVarP(&jenkinsJobFlags, "param", "", []string{}, "List of params to pass to Jenkins. Format is key=value.")

	jenkinsCmd.PersistentFlags().StringVarP(&jenkinsUser, "username", "", "", "Your Jenkins username")
	jenkinsCmd.PersistentFlags().StringVarP(&jenkinsToken, "token", "", "", "Your Jenkins API token")

	jobCmd.AddCommand(jobStartCmd)
	jenkinsCmd.AddCommand(jobCmd)
	rootCmd.AddCommand(jenkinsCmd)
}

// jenkins offers two kick-off URLs depending on whether
// the job has parameters or not
const unparamaterizedBuildPath = "build"
const parameterizedBuildPath = "buildWithParameters"

// startJenkinsJob determines the URL to use for a given Jenkins job and kicks it off
func startJenkinsJob(cmd *cobra.Command, args []string) {

	params, err := createJenkinsPostParams(jenkinsJobFlags)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var buildPath string
	if len(params) != 0 {
		buildPath = parameterizedBuildPath
	} else {
		buildPath = unparamaterizedBuildPath
	}

	// jenkins offers two kick-off URLs depending on whether
	// the job has parameters or not
	normalizedUrl := normalizeJenkinsUrl(jenkinsJobUrl, buildPath)

	response, err := http.PostForm(normalizedUrl, params)

	if err != nil {
		fmt.Printf("Error retrieving url %s: %v\n", normalizedUrl, err)
		os.Exit(1)
	}

	// if you kick off a parameterized build using the "build" path (even if you're using the defaults for the parameters)
	// Jenkins will give a 400. So track that particular response
	if buildPath == unparamaterizedBuildPath && response.StatusCode == http.StatusBadRequest {
		response.Body.Close()
		normalizedUrl = normalizeJenkinsUrl(jenkinsJobUrl, parameterizedBuildPath)
		response, err = http.PostForm(normalizedUrl, url.Values{})
	}

	defer response.Body.Close()
	if err != nil {
		fmt.Printf("Error retrieving url %s: %v\n", normalizedUrl, err)
		os.Exit(1)
	}
}

var lastSlashRegex = regexp.MustCompile("/$")

// normalizeJenkinsUrl takes the URL the user might paste and normalizes it with
// username and password if those are set, appending the passed-in buildPath
func normalizeJenkinsUrl(currentUrl string, buildPath string) string {
	// input urls could be of the formats:
	//  http://jenkinsurl
	//  https://jenkinsurl
	//  http://user:password@jenkinsurl
	//  https://user:password@jenkinsurl
	//  jenkinsUrl
	//  user:pass@jenkinsUrl

	parsedUrl, err := url.Parse(currentUrl)
	if err != nil {
		fmt.Printf("Invalid URL: %v", err)
		os.Exit(1)
	}

	// add auth info if present
	if jenkinsUser != "" {
		if jenkinsToken != "" {
			parsedUrl.User = url.UserPassword(jenkinsUser, jenkinsToken)
		} else {
			parsedUrl.User = url.User(jenkinsUser)
		}
	}

	if lastSlashRegex.Match([]byte(parsedUrl.Path)) {
		parsedUrl.Path = parsedUrl.Path + buildPath
	} else {
		parsedUrl.Path = parsedUrl.Path + "/" + buildPath
	}

	return parsedUrl.String()
}

var hasEqualsRegex = regexp.MustCompile("=")

func createJenkinsPostParams(inputPairs []string) (url.Values, error) {
	values := url.Values{}
	if len(inputPairs) == 0 {
		return values, nil
	}

	for _, keyValue := range inputPairs {
		if !hasEqualsRegex.Match([]byte(keyValue)) {
			return nil, fmt.Errorf("Key/values must be of the form key=value. Invalid format: %s", keyValue)
		}

		keyValueParts := strings.Split(keyValue, "=")
		var value string
		if len(keyValueParts) == 1 {
			// the key= case
			value = ""
		} else {
			value = keyValueParts[1]
		}
		values[keyValueParts[0]] = []string{value}
	}
	fmt.Println(values)
	return values, nil
}
