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

	someTitle := "some title"

	tests := []struct {
		Input    string
		Expected Config
	}{
		{
			`{"testcases":[]}`,
			Config{[]Testcase{}, "", 100_000, GradingParameters{}, nil,
				"default", "jailed_sandbox", "jailed_sandbox", Autograding{},
				ContainerOptions{}, nil},
		},
		{
			`{"testcases":[{"title":"some title","command":["hi"]}]}`,
			Config{[]Testcase{{Title: someTitle, Type: "Execution", Commands: []string{"hi"}}}, "", 100_000, GradingParameters{},
				nil, "default", "jailed_sandbox", "jailed_sandbox",
				Autograding{}, ContainerOptions{}, nil},
		},
		{
			`{"testcases":[{"title":"some title","command":["g++", "hello.cpp", "-o", "hello.exe"],"extra_credit":true}]}`,
			Config{[]Testcase{{Title: someTitle, Type: "Execution", Commands: []string{"g++", "hello.cpp", "-o", "hello.exe"}, ExtraCredit: true}}, "", 100_000, GradingParameters{},
				nil, "default", "jailed_sandbox", "jailed_sandbox",
				Autograding{}, ContainerOptions{}, nil},
		},
	}

	var cfgErr *ConfigError
	for idx, tc := range tests {
		var actual Config
		err := json.Unmarshal([]byte(tc.Input), &actual)

		if errors.As(err, &cfgErr) {
			t.Fatalf("Test %d failed. Input %v. Configuration error: %v\n", idx, tc.Input, err)
		}

		if err != nil {
			t.Fatalf("Test %d failed. Input %v. Unmarshaling error: %v\n", idx, tc.Input, err)
		}

		if !reflect.DeepEqual(actual, tc.Expected) {
			t.Fatalf("Test %d failed. Input %v. Expected %v. Actually: %v.\n", idx, tc.Input, tc.Expected, actual)
		}
	}

}

// Testing failures for invalid config.json files.
func TestConfigInvalidJSON(t *testing.T) {
	inputs := []string{
		"null",
		"{}",
	}

	for idx, input := range inputs {
		actual := &Config{}
		err := json.Unmarshal([]byte(input), actual)

		if err == nil {
			t.Errorf("Test %d: JSON unmarshalling should have failed for JSON '%s'", idx, input)
		}

		var cfgErr *ConfigError
		if !errors.As(err, &cfgErr) {
			t.Errorf("Test %d: Unexpected unmarshalling error for JSON '%s'", idx, input)
		}
	}
}
