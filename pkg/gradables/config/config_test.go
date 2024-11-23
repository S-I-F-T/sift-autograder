// Package gradables contains representations of Submitty gradable state
package gradables

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

/*
	Validating the config.json through the

TestConfigJSON() function and outputting
errors from config structure.
*/
func TestConfigJSON(t *testing.T) {

	// sample title for use
	someTitle := "some title"

	/* define a slice of test cases with different JSON inputs
	   and their corresponding expected Config structs. 
	*/
	tests := []struct {
		Input    string
		Expected Config
	}{
		// testing empty testcases
		{
			`{"testcases":[]}`,
			// expected output: an empty list of testcases and default values for other fields.
			Config{[]Testcase{}, "", 100_000, GradingParameters{}, nil,
				"default", "jailed_sandbox", "jailed_sandbox", Autograding{},
				ContainerOptions{}, nil},
		},
		{
			// a single testcase with a title and a single command.
			`{"testcases":[{"title":"some title","command":["hi"]}]}`,
			Config{[]Testcase{{Title: someTitle, Type: "Execution", Commands: []string{"hi"}}}, "", 100_000, GradingParameters{},
				nil, "default", "jailed_sandbox", "jailed_sandbox",
				Autograding{}, ContainerOptions{}, nil},
		},
		{
			// expected output: a single testcase with the provided title and command, default values for others.
			`{"testcases":[{"title":"some title","command":["g++", "hello.cpp", "-o", "hello.exe"],"extra_credit":true}]}`,
			Config{[]Testcase{{Title: someTitle, Type: "Execution", Commands: []string{"g++", "hello.cpp", "-o", "hello.exe"}, ExtraCredit: true}}, "", 100_000, GradingParameters{},
				nil, "default", "jailed_sandbox", "jailed_sandbox",
				Autograding{}, ContainerOptions{}, nil},
		},
	}

	// define a variable to capture custom configuration errors during unmarshaling.
	var cfgErr *ConfigError
	// iterating over all the test cases to validate JSON unmarshaling behavior.
	for idx, tc := range tests {
		var actual Config
		// attempting to unmarshal the JSON input into Config struct.
		err := json.Unmarshal([]byte(tc.Input), &actual)
		
		// check if the error is of type ConfigError and fail the test if encountered.
		if errors.As(err, &cfgErr) {
			t.Fatalf("Test %d failed. Input %v. Configuration error: %v\n", idx, tc.Input, err)
		}
		
		// check for any other errors during unmarshaling and fail the test if found.
		if err != nil {
			t.Fatalf("Test %d failed. Input %v. Unmarshaling error: %v\n", idx, tc.Input, err)
		}

		// useing reflect.DeepEqual to compare the actual and expected Config structs.
		// if they are not equal, fail the test and output the discrepancy.
		if !reflect.DeepEqual(actual, tc.Expected) {
			t.Fatalf("Test %d failed. Input %v. Expected %v. Actually: %v.\n", idx, tc.Input, tc.Expected, actual)
		}
	}

}

// Testing failures for invalid config.json files.
func TestConfigInvalidJSON(t *testing.T) {

	// define a slice of invalid JSON strings that should fail unmarshaling.
	inputs := []string{
		"null",
		"{}",
	}

	// iterate over each invalid input JSON.
	for idx, input := range inputs {
		// initialize an empty Config object to hold the result of unmarshaling.
		actual := &Config{}
		// attempt to unmarshal the invalid JSON input into the Config object.
		err := json.Unmarshal([]byte(input), actual)

		// check if no error occurred (unexpected behavior for invalid inputs).
		if err == nil {
			t.Errorf("Test %d: JSON unmarshalling should have failed for JSON '%s'", idx, input)
		}

		// define a variable to check if the error is of type ConfigError.
		var cfgErr *ConfigError
		// verify that the error is a ConfigError.
		if !errors.As(err, &cfgErr) {
			t.Errorf("Test %d: Unexpected unmarshalling error for JSON '%s'", idx, input)
		}
	}
}
