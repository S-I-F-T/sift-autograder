package gradables

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"
	"syscall"
)

const (
	Compilation = "Compilation"
	Execution   = "Execution"
	FileCheck   = "FileCheck"
)

// struct for testcases
type Testcase struct {
	Type            string         `json:"type"`            // type of testcase
	Title           string         `json:"title"`           // name of testcase
	Details         string         `json:"details"`         // details of testcase
	Points          int            `json:"points"`          // points allocated for testcase
	Hidden          bool           `json:"hidden"`          // hidden testcases from user
	ExtraCredit     bool           `json:"extra_credit"`    // extra-credit testcases
	Filenames       []string       `json:"filename"`        // testcase files
	ExecutableNames []string       `json:"executable_name"` // executable for testcase execution
	Commands        []string       `json:"command"`         // linux commands executed during compilation and/or execution
	Containers      []Container    `json:"containers"`      // docker containers and what will be run on each.
	Validations     []Validation   `json:"validation"`      // automatic checks for STDOUT.txt, STDERR.txt, and the execution logfile.
	Actions         []string       `json:"actions"`         // actions for testcase
	ResourceLimits  ResourceLimits `json:"resource_limits"` // resource limits for testcase to prevent overuse
}

func (tc *Testcase) UnmarshalJSON(data []byte) error {
	tc.Type = "Execution" // Could also
	tc.Details = ""
	tc.Points = 0
	tc.Hidden = false
	tc.ExtraCredit = false
	tc.Actions = nil

	if string(data) == "" || string(data) == "null" {
		return nil
	}

	var Typed struct {
		Type string `json:"type"`
	}

	if err := json.Unmarshal(data, &Typed); err != nil {
		return err
	}

	if Typed.Type == Compilation {
		tc.ResourceLimits.MaxCPUTime = 60
		tc.ResourceLimits.MaxFileSize = 10000000
	} else {
		tc.ResourceLimits.MaxCPUTime = 10
		tc.ResourceLimits.MaxFileSize = 100_1000
	}

	if false { // FIXME: Replace once we figure out how Submitty_Count works
		panic("implement me!")
	} else {
		// FIXME: Whatever Submitty_Count changes from the default, change back here
	}

	tc.ResourceLimits.MaxData = 500_000_000
	tc.ResourceLimits.MaxStack = 500_000_000
	tc.ResourceLimits.MaxCoreFile = 0
	tc.ResourceLimits.MaxResidentSet = 1_000_000_000
	tc.ResourceLimits.MaxFileNumber = 100
	tc.ResourceLimits.MaxLockedMemory = 500_000_000
	tc.ResourceLimits.MaxAddressSpace = syscall.RLIM_INFINITY
	tc.ResourceLimits.MaxLocks = 100
	tc.ResourceLimits.MaxSignalQueue = 0
	tc.ResourceLimits.MaxMessageQueue = 0
	tc.ResourceLimits.MaxNice = 1_000_000_000 // Infinity
	tc.ResourceLimits.MaxRTPriority = 0
	tc.ResourceLimits.MaxRTTime = 0

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

	if tc.Title == "" {
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
