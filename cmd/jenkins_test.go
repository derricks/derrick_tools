package cmd

import (
	"testing"
)

type normalizeJenkinsUrlTest struct {
	baseUrl   string
	username  string
	password  string
	buildPath string
	expected  string
}

func TestNormalizeUrl(test *testing.T) {
	testCases := []normalizeJenkinsUrlTest{
		normalizeJenkinsUrlTest{"https://localhost:8080", "", "", "build", "https://localhost:8080/build"},
		normalizeJenkinsUrlTest{"http://localhost:8080/123/", "", "", "buildWithParameters", "http://localhost:8080/123/buildWithParameters"},
		normalizeJenkinsUrlTest{"http://localhost:8080/123", "testuser", "", "buildWithParameters", "http://testuser@localhost:8080/123/buildWithParameters"},
		normalizeJenkinsUrlTest{"https://localhost:8080/123/", "testuser", "testpass", "build", "https://testuser:testpass@localhost:8080/123/build"},
	}

	for index, testCase := range testCases {
		jenkinsUser = testCase.username
		jenkinsToken = testCase.password

		actualUrl := normalizeJenkinsUrl(testCase.baseUrl, testCase.buildPath)

		if actualUrl != testCase.expected {
			test.Errorf("Test case %d: Expected url of %s but got %s", index, testCase.expected, actualUrl)
		}
	}
}

type createJenkinsPostParamsTest struct {
	pairs          []string
	expectedResult map[string]string

	errorExpected bool
}

func TestCreateJenkinsPostParams(test *testing.T) {
	testCases := []createJenkinsPostParamsTest{
		createJenkinsPostParamsTest{[]string{"key1=value1", "key2="}, map[string]string{"key1": "value1", "key2": ""}, false},
		createJenkinsPostParamsTest{[]string{"key1=value1", "key2"}, nil, true},
		createJenkinsPostParamsTest{[]string{}, make(map[string]string), false},
	}

	for index, testCase := range testCases {
		actualParams, err := createJenkinsPostParams(testCase.pairs)

		if err == nil && testCase.errorExpected {
			test.Errorf("Test case %d: expected error but none was produced", index)
		}

		if err != nil && !testCase.errorExpected {
			test.Errorf("Test case %d: unexpected error %v", index, err)
		}

		if err != nil && testCase.errorExpected {
			if actualParams != nil {
				test.Errorf("Test case %d: Expected nil return but it was not %v", index, actualParams)
			}
			// there was an expected error, no reason to continue with these tests
			continue
		}

		for key, value := range testCase.expectedResult {
			actualValue := actualParams.Get(key)
			if actualValue != value {
				test.Errorf("Test case %d: Expected value of %s for key %s but was %s", index, value, key, actualValue)
			}
		}
	}
}
