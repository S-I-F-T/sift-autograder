package gradables

/*
	Constructing Course struct from config.json

https://github.com/Submitty/Submitty/blob/main/autograder/tests/data/config_files
*/
type Course struct {
	CourseName        string `json:"course_name"`          // name of course
	CourseHomeURL     string `json:"course_home_url"`      // course home url
	DefaultHWLateDays int    `json:"default_hw_late_days"` /* default number of late days allocated
	that can be used for hw assignments. */
	DefaultStudentLateDays int `json:"default_student_late_days"` /* default number
	of late days students an use for an assignment */
	UploadMessage string `json:"upload_message"` /* default upload message
	when submitting an assignment. */
	CourseEmail string `json:"course_email"` // course email
}
