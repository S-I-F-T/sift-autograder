package gradables

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
)

const (
	Compilation = "Compilation"
	Execution   = "Execution"
	FileCheck   = "FileCheck"
)

// struct for testcases
type Testcase struct {
	Type            string       `json:"type"`            // type of testcase
	Title           string       `json:"title"`           // name of testcase
	Details         string       `json:"details"`         // details of testcase
	Points          int          `json:"points"`          // points allocated for testcase
	Hidden          bool         `json:"hidden"`          // hidden testcases from user
	ExtraCredit     bool         `json:"extra_credit"`    // extra-credit testcases
	Filenames       []string     `json:"filename"`        // testcase files
	ExecutableNames []string     `json:"executable_name"` // executable for testcase execution
	Commands        []string     `json:"command"`         // linux commands executed during compilation and/or execution
	Containers      []Container  `json:"containers"`      // docker containers and what will be run on each.
	Validations     []Validation `json:"validation"`      // automatic checks for STDOUT.txt, STDERR.txt, and the execution logfile.
	Actions         []string     `json:"actions"`         // actions for testcase
}


/*  Unmarshalling method for the Testcase struct. 
	Allows for initialization of default values and 
	additional validations during the unmarshalling process.
*/
func (tc *Testcase) UnmarshalJSON(data []byte) error {
	// setting the default values for Testcase fields
	tc.Type = "Execution" // default type for testcase
	tc.Details = "" // default empty details
	tc.Points = 0 // setting default points to 0
	tc.Hidden = false // setting default visibility to not hidden
	tc.ExtraCredit = false // default is no extra credit
	tc.Actions = nil // default to actions to null for no actions
	
	// check if input data is empty/null and exit early
	if string(data) == "" || string(data) == "null" {
		return nil
	}

	// TODO: Figure out how Submitty_Count works - placeholder comment

	// create an alias type to avoid recursive unmarshalling
	type JSONTestcase Testcase
	// unmarshal input json data into the testcase using the alias type
	if err := json.Unmarshal(data, (*JSONTestcase)(tc)); err != nil {
		return err
	}

	// perform additional validation on testcase fields after unmarshalling
	// validte() function is described below
	return tc.validate()
}

/*
Validate method for the Testcase struct, which checks if the Testcase fields

	meet specific requirements based on its type. Returns an error if validation fails.
*/
func (tc *Testcase) validate() error {

	// Defines valid types for testcase field 'Type'
	types := []string{"Compilation", "FileCheck", "Execution"}
	// Checks if the slice contains one of the above types
	if !slices.Contains(types, tc.Type) {
		return &ConfigError{fmt.Sprintf("'testcase.type' must be one of: %s", strings.Join(types, ", "))}
	}

	// Ensure that the testcase field 'Title' is not empty
	if tc.Title == "" {
		return &ConfigError{"'testcase.title' is required"}
	}

	// Defines valid types for testcase field 'Filenames'
	requireFiles := []string{"FileCheck", "Execution"}
	/* If Filenames is nil and the slice does not contain the above types,
	   then a config error is outputted to indicate that the Filenames is required. */
	if tc.Filenames == nil && !slices.Contains(requireFiles, tc.Type) {
		return &ConfigError{fmt.Sprintf("`testcase.file_names` required for testcases of types: %s",
			strings.Join(requireFiles, ", "))}
	}

	// Ensures that executable names are provided for testcase
	if tc.ExecutableNames == nil && tc.Type == "Compilation" {
		return &ConfigError{"`testcase.executable_name` required for testcases of type `Compilation`"}
	}

	// Defines types for Commands and Container fields in testcase.
	runTypes := []string{"FileCheck", "Execution"}
	/* Check if tc.type is in runTypes and that the Commands and Containers
	fields of the test case aren't null. A config error is returned to indicate
	that the fields are required for the types. */
	if slices.Contains(runTypes, tc.Type) && tc.Commands == nil && tc.Containers == nil {
		return &ConfigError{fmt.Sprintf("`testcase.commands` or `testcase.containers` required for testcases of types: %s",
			strings.Join(types, ", "))}
	}

	return nil
}
