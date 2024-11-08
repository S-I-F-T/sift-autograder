package autograder

import (
	"context"
	"log"

	gradables "github.com/S-I-F-T/sift-autograder/pkg/gradables/config"
)

func Run(submissionFolder string, cfg *gradables.Config, ctx *context.Context, logger *log.Logger) (map[string]int, error) {

	if sizeOk, err := submissionSizeMaintained(submissionFolder, cfg.MaxSubmissionSize); err != nil {
		logger.Printf("Couldn't determine size of directory %s%n", submissionFolder)
		return nil, &ErrorForUser{"Trouble reading from submission folder"}
	} else if !sizeOk {
		logger.Println("WARNING: submission too big")
	}

	// 1. Set up environment
	// 2. Run file checks
	// 3. Run compilation
	// 4. Run execution (in parallel)
	panic("implement me")
}

func submissionSizeMaintained(submissionPath string, maxSize int) (bool, error) {

}

type ErrorForUser struct {
	Message string
}

func (err *ErrorForUser) Error() string {
	return err.Message
}
