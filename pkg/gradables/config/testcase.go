package gradables

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
)

type Testcase struct {
	Type            string       `json:"type"`
	Title           *string      `json:"title"`
	Details         string       `json:"details"`
	Points          int          `json:"points"`
	Hidden          bool         `json:"hidden"`
	ExtraCredit     bool         `json:"extra_credit"`
	Filenames       []string     `json:"filename"`
	ExecutableNames []string     `json:"executable_name"`
	Commands        []string     `json:"command"`
	Containers      []Container  `json:"containers"`
	Validations     []Validation `json:"validation"`
	Actions         []string     `json:"actions"`
}

func (tc *Testcase) UnmarshalJSON(data []byte) error {
	tc.Type = "Execution"
	tc.Details = ""
	tc.Points = 0
	tc.Hidden = false
	tc.ExtraCredit = false
	tc.Actions = nil

	if string(data) == "" || string(data) == "null" {
		return nil
	}

	type JSONTestcase Testcase

	if err := json.Unmarshal(data, (*JSONTestcase)(tc)); err != nil {
		return err
	}

	return tc.validate()
}

func (tc *Testcase) validate() error {

	types := []string{"Compilation", "FileCheck", "Execution"}
	if !slices.Contains(types, tc.Type) {
		return &ConfigError{fmt.Sprintf("'testcase.type' must be one of: %s", strings.Join(types, ", "))}
	}

	if tc.Title == nil {
		return &ConfigError{"'testcase.title' is required"}
	}

	requireFiles := []string{"FileCheck", "Execution"}
	if tc.Filenames == nil && !slices.Contains(requireFiles, tc.Type) {
		return &ConfigError{fmt.Sprintf("`testcase.file_names` required for testcases of types: %s",
			strings.Join(requireFiles, ", "))}
	}

	if tc.ExecutableNames == nil && tc.Type == "Compilation" {
		return &ConfigError{"`testcase.executable_name` required for testcases of type `Compilation`"}
	}

	runTypes := []string{"FileCheck", "Execution"}
	if slices.Contains(runTypes, tc.Type) && tc.Commands == nil && tc.Containers == nil {
		return &ConfigError{fmt.Sprintf("`testcase.commands` or `testcase.containers` required for testcases of types: %s",
			strings.Join(types, ", "))}
	}

	return nil
}
