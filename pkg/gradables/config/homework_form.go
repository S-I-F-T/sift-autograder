package gradables

/* Constructing Homework struct from homework_form.json
https://github.com/Submitty/Submitty/tree/main/autograder/tests/data */
type Homework struct {
	GradableID string `json:"gradable_id"`
	ConfigPath string `json:"config_path"`
	DateDue string `json:"date_due"`
	UploadType string `json:upload_type`
	Subdirectory string `json:subdirectory`
	VCSPartialPath string `json:vcs_partial_path`
}