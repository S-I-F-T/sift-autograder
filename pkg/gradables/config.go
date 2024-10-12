// Package gradables implements gradables.
package gradables

// A Config represents a gradable's configuration.
//
// It's specification matches the submitty gradable gradables:
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

// ConfigFromRequired constructs a configuration out of the minimal configuration possible.
//
// All non-required fields are given default values.
func ConfigFromRequired(testcases []Testcase) Config {
	return Config{
		Testcases:            testcases,
		AssignmentMessage:    "",
		MaxSubmissionSize:    100_000,
		GradingParameters:    GradingParameters{}, // TODO
		PartNames:            nil,
		RequiredCapabilities: "default",
		AutogradingMethod:    "jailed_sandbox",
		ContainerType:        "jailed_sandbox",
		Autograding:          Autograding{},      // TODO
		ContainerOptions:     ContainerOptions{}, // TODO
		AllowedSyscalls:      nil,
	}
}

func (c *Config) Validate() bool {
	for _, tc := range c.Testcases {
		if !tc.validate() {
			return false
		}
	}

	if !(c.GradingParameters.validate() &&
		c.Autograding.validate() &&
		c.ContainerOptions.validate()) {
		return false
	}

	panic("Implement me")
}

type Autograding struct {
	CompilationsToRunner     []string `json:"compilation_to_runner"`
	CompilationsToValidation []string `json:"compilation_to_validation"`
	SubmissionsToCompilation []string `json:"submission_to_compilation"`
	SubmissionsToRunner      []string `json:"submission_to_runner"`
	SubmissionsToValidation  []string `json:"submission_to_validation"`
	WorkToDetails            []string `json:"work_to_details"`
	UseCheckoutSubdirectory  string   `json:"use_checkout_subdirectory"`
}

func (ag *Autograding) validate() bool {
	panic("Implement me!")
}

type ContainerOptions struct {
	ContainerImage         string `json:"container_image"`
	NumberOfPorts          int    `json:"number_of_ports"`
	SinglePortPerContainer bool   `json:"single_port_per_container"`
	UseRouter              bool   `json:"use_router"`
}

func (co *ContainerOptions) validate() bool {
	panic("Implement me!")
}

type GradingParameters struct{} // TODO: implement

func (gp *GradingParameters) validate() bool {
	panic("Implement me!")
}

type Testcase struct {
	Type            string       `json:"type"`
	Title           string       `json:"title"`
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

func (tc *Testcase) validate() bool {
	for _, c := range tc.Containers {
		if !c.validate() {
			return false
		}
	}

	for _, v := range tc.Validations {
		if !v.validate() {
			return false
		}
	}

	panic("Implement me!")
}

type Container struct {
	Commands            []string `json:"commands"`
	ContainerName       string   `json:"container_name"`
	ContainerImage      string   `json:"container_image"`
	NumberOfPorts       int      `json:"number_of_ports"`
	OutgoingConnections []string `json:"outgoing_connections"`
	Server              bool     `json:"server"`
}

func (c *Container) validate() bool {
	panic("Implement me!")
}

type Validation struct {
	Method              string   `json:"method"`
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
}

func (v *Validation) validate() bool {
	panic("Implement me!")
}
