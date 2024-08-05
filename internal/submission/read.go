// Package submission represents the respondents' submission.
package submission

import (
	"encoding/json"
	"fmt"
	"sdxImage/internal/log"
)

// Read transforms a submission's raw bytes into an
// instance that implements interfaces.Submission
func Read(bytes []byte) (*Submission, error) {
	submission := &Submission{}

	//Debugging unmarshalling error
	fmt.Println(bytes)

	err := json.Unmarshal(bytes, submission)
	if err != nil {
		log.Info(fmt.Sprintf("failed to read file with error: %q", err))
		return nil, err
	}
	return submission, nil
}

type Exception struct {
	Msg string
}

func (e *Exception) Error() string {
	return e.Msg
}
