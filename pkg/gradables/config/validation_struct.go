package gradables

import (
	"encoding/json"
	"slices"
	"strings"

	"github.com/S-I-F-T/sift-autograder/pkg/utils"
)

const (
	customMethod = "custom_method"
)

// Validation represents the configuration settings for a Submitty validation.
//
// Adheres to: https://submitty.org/instructor/autograding/validation
type Validation struct {
	Method              *string  `json:"method"`
	Description         string   `json:"description"`
	ActualFiles         []string `json:"actual_file"`
	ExpectedFiles       []string `json:"expected_file"`
	Deduction           float64  `json:"deduction"`
	ShowMessage         string   `json:"show_message"`
	ShowActual          string   `json:"show_actual"`
	ShowExpected        string   `json:"show_expected"`
	ShowDifferenceImage string   `json:"show_difference_image"`
	AcceptableThreshold float64  `json:"acceptable_threshold"`
	FailureMessage      string   `json:"failure_message"`
	Command             *string  `json:"command"`
}

func (v *Validation) UnmarshalJSON(data []byte) error {
	v.Deduction = 1.0
	v.ShowActual = "always"
	v.ShowDifferenceImage = "always"
	v.ShowExpected = "always"
	v.ShowMessage = "always"

	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	type JSONValidation Validation
	jsonValidation := (*JSONValidation)(v)

	if err := json.Unmarshal(data, jsonValidation); err != nil {
		return nil
	}

	return v.validate()
}

// Validate validates a Validation structure.
//
// Only an object adhering to https://submitty.org/instructor/autograding/validation
// should pass this validation.
func (v *Validation) validate() error {
	// TODO: check Method against files provided

	if *v.Method == customMethod && v.Command == nil {
		return &ConfigError{"Command cannot be empty for custom method in a validation"}
	}

	if v.Method == nil {
		return &ConfigError{"Method cannot be empty for validation"}
	}

	if !utils.IsUnit(v.Deduction, 0.001) {
		return &ConfigError{"A validation deduction must be between 0 and 1"}
	}

	if !utils.IsUnit(v.AcceptableThreshold, 0.001) {
		return &ConfigError{"A validation's acceptable threshold must be between 0 and 1"}
	}

	levels := []string{"always", "never", "on_success", "on_failure"}
	levelsDescription := strings.Join(levels, ", ")

	if !slices.Contains(levels, v.ShowActual) {
		return &ConfigError{"A validation's showActual must be one of " + levelsDescription}
	}

	if !slices.Contains(levels, v.ShowDifferenceImage) {
		return &ConfigError{"A validation's showDifferenceImage must be one of " + levelsDescription}
	}

	if !slices.Contains(levels, v.ShowExpected) {
		return &ConfigError{"A validation's showExpected must be one of " + levelsDescription}
	}

	if !slices.Contains(levels, v.ShowMessage) {
		return &ConfigError{"A validation's showMessage must be one of " + levelsDescription}
	}

	return nil
}
