// Package gradables contains representations of Submitty gradable state
package gradables

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

/* Validating the config.json through the 
TestConfigJSON() function and outputting
errors from config structure. */
func TestConfigJSON(t *testing.T) {

	inputs := []string{
		`{"testcases":[]}`,
		`{"testcases":[{"title":"some title","command":["hi"]}]}`,
		`{"testcases":[{"title":"some title","command":["g++", "hello.cpp", "-o", "hello.exe"],"extra_credit":true}]}`,
	}

	someTitle := "some title"

	expected := []*Config{
		{[]Testcase{}, "", 100_000, GradingParameters{}, nil,
			"default", "jailed_sandbox", "jailed_sandbox", Autograding{},
			ContainerOptions{}, nil, ResourceLimits{}},

		{[]Testcase{{Title: someTitle, Type: "Execution", Commands: []string{"hi"}}}, "", 100_000, GradingParameters{},
			nil, "default", "jailed_sandbox", "jailed_sandbox",
			Autograding{}, ContainerOptions{}, nil, ResourceLimits{}},

		{[]Testcase{{Title: someTitle, Type: "Execution", Commands: []string{"g++", "hello.cpp", "-o", "hello.exe"}, ExtraCredit: true}}, "", 100_000, GradingParameters{},
			nil, "default", "jailed_sandbox", "jailed_sandbox",
			Autograding{}, ContainerOptions{}, nil, ResourceLimits{}},
	}

	if len(inputs) != len(expected) {
		t.Fatal("Input-Expected length mismatch")
	}

	var cfgErr *ConfigError
	for idx, input := range inputs {
		var actual Config
		err := json.Unmarshal([]byte(input), &actual)

		if errors.As(err, &cfgErr) {
			t.Fatalf("Config error for JSON '%s': %v", input, err)
		}

		if err != nil {
			t.Fatalf("Unmarshalling error for JSON '%s': %v", input, err)
		}

		if !reflect.DeepEqual(&actual, expected[idx]) {
			t.Fatalf("\nExpected: %v\nActually: %v", expected[idx], &actual)
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
