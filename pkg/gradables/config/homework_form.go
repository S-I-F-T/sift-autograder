package gradables

/* Constructing Homework struct from homework_form.json
https://github.com/Submitty/Submitty/tree/main/autograder/tests/data */
type Homework struct {
	GradableID string `json:"gradable_id"` // id for hw gradeable
	ConfigPath string `json:"config_path"` // configuration path for homework
	DateDue string `json:"date_due"` // due date for assignment
	UploadType string `json:upload_type` // upload types for the homework
	Subdirectory string `json:subdirectory` // hw subdirectories
}