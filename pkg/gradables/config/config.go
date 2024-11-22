// Package gradables implements gradables.
package gradables

import "encoding/json"

// A Config represents a gradable's configuration.
//
// It's specification matches the Submitty gradable gradables:
// https://submitty.org/instructor/autograding/specification
type Config struct {
	Testcases            []Testcase        `json:"testcases"`
	AssignmentMessage    string            `json:"assignment_message"`
	MaxSubmissionSize    int               `json:"max_submission_size"`
	GradingParameters    GradingParameters `json:"grading_parameters"`
	PartNames            []string          `json:"part_names"`
	RequiredCapabilities string            `json:"required_capabilities"`
	AutogradingMethod    string            `json:"autograding_method"`
	ContainerType        string            `json:"container options"`
	Autograding          Autograding       `json:"autograding"`
	ContainerOptions     ContainerOptions  `json:"container_options"`
	AllowedSyscalls      []string          `json:"allow_system_calls"`
}

// UnmarshalJSON Grab data from JSON file, parse them into their correct data structure
func (c *Config) UnmarshalJSON(data []byte) error {
	c.Testcases = nil
	c.AssignmentMessage = ""
	c.MaxSubmissionSize = 100_000
	c.GradingParameters = GradingParameters{} // TODO
	c.PartNames = nil
	c.RequiredCapabilities = "default"
	c.AutogradingMethod = "jailed_sandbox"
	c.ContainerType = "jailed_sandbox"
	c.Autograding = Autograding{}           // TODO
	c.ContainerOptions = ContainerOptions{} // TODO
	c.AllowedSyscalls = nil

	type JSONConfig Config
	if err := json.Unmarshal(data, (*JSONConfig)(c)); err != nil {
		return err
	}

	if c.Testcases == nil {
		return &ConfigError{"'testcase' field cannot be left blank"}
	}

	return nil
}

// Autograding configuration structure
type Autograding struct {
	CompilationsToRunner     []string `json:"compilation_to_runner"`
	CompilationsToValidation []string `json:"compilation_to_validation"`
	SubmissionsToCompilation []string `json:"submission_to_compilation"`
	SubmissionsToRunner      []string `json:"submission_to_runner"`
	SubmissionsToValidation  []string `json:"submission_to_validation"`
	WorkToDetails            []string `json:"work_to_details"`
	UseCheckoutSubdirectory  string   `json:"use_checkout_subdirectory"`
}

// FileMovement define how files are moved through the autograding process
type FileMovement struct {
	CompilationsToRunner     []string `json:"compilation_to_runner"`
	CompilationsToValidation []string `json:"compilation_to_validation"`
	SubmissionsToCompilation []string `json:"submission_to_compilation"`
	SubmissionsToRunner      []string `json:"submission_to_runner"`
	SubmissionsToValidation  []string `json:"submission_to_validation"`
	WorkToDetails            []string `json:"work_to_details"`
	UseCheckoutSubdirectory  string   `json:"use_checkout_subdirectory"`
}

// ContainerOptions the different options the container of the autograder uses
type ContainerOptions struct {
	ContainerImage         string `json:"container_image"`
	NumberOfPorts          uint16 `json:"number_of_ports"`
	SinglePortPerContainer bool   `json:"single_port_per_container"`
	UseRouter              bool   `json:"use_router"`
}

func (co *ContainerOptions) validate() error {
	if len(co.ContainerImage) <= 0 {
		return &ConfigError{"'container_image' field cannot be left blank"}
	}

	return nil
}

/*
	Grading parameters structure to keep track of

auto-grader and extra-credit points.
*/
type GradingParameters struct {
	AutoPoints        uint16 `json:"AUTO_POINTS"`
	ExtraCreditPoints uint16 `json:"EXTRA_CREDIT_POINTS"`
}

/* Validates the grading parameters */
func (gp *GradingParameters) validate() error {
	return nil
}

// Container Specify Docker containers for this testcase and what will be run in each of them.
// Each container is specified by an object with no key that contains the following fields.
type Container struct {
	Commands            []string `json:"commands"`
	ContainerName       string   `json:"container_name"`
	ContainerImage      string   `json:"container_image"`
	NumberOfPorts       int      `json:"number_of_ports"`
	OutgoingConnections []string `json:"outgoing_connections"`
	Server              bool     `json:"server"`
}

type ConfigError struct {
	Message string
}

func (err *ConfigError) Error() string {
	return err.Message
}
