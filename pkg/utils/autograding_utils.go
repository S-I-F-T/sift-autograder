package utils

/*
	autograding_utils.go consisting of auxilary functions for autograding, ported from
	autograding_utils.py

https://github.com/Submitty/Submitty/blob/main/autograder/autograder/autograding_utils.py
*/

// Logger struct for file logs
type Logger struct {
	LogDir string
	StackTraceDir string
	CaptureTraces string
	AccumulatedTraces []string
}
