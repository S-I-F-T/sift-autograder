package gradables

// Set up homework.go Homework and Shipper structs

//
//	Constructing Homework struct from homework_form.json

// https://github.com/Submitty/Submitty/tree/main/autograder/tests/data
// 
type Homework struct {
	GradableID   string `json:"gradable_id"`  // id for hw gradeable
	ConfigPath   string `json:"config_path"`  // configuration path for homework
	DateDue      string `json:"date_due"`     // due date for assignment
	UploadType   string `json:"upload_type"`  // upload types for the homework
	Subdirectory string `json:"subdirectory"` // hw subdirectories
}

//
// 	Constructing Shipper struct from shipper_config.json (adjusted for SIFT)

// https://github.com/Submitty/Submitty/blob/main/autograder/tests/data/shipper_config.json
//
type Shipper struct {
	CheckoutIncludedSymLinks bool `json:"checkout_included_symlinks"`
	CheckoutTotalSize int `json:"checkout_total_size"`
	GradeResult string `json:"grade_result"`
	Gradeable string `json:"gradeable"`
	GradingTime float64 `json:"gradingtime"`
	IsTeam bool `json:"is_team"` // curious how this would work with SIFT, will ask...
	MaxPossibleGradingTime int `json:"max_possible_grading_time"`
	User string `json:"user"`
	Who string `json:"test_student"`
}
