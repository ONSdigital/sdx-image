// Package submission represents the respondents' submission.
package submission

import (
	"encoding/json"
	"sdxImage/pkg/interfaces"
	"sdxImage/pkg/log"
)

// Read transforms a submission's raw bytes into an
// instance that implements interfaces.Submission
func Read(bytes []byte) (interfaces.Submission, error) {
	submission := &V2Submission{}
	err := json.Unmarshal(bytes, submission)
	if err != nil {
		return readV1(bytes)
	}
	if submission.Version != "v2" {
		return readV1(bytes)
	}
	return submission, nil
}

func readV1(bytes []byte) (interfaces.Submission, error) {
	v1 := &V1Submission{}
	err := json.Unmarshal(bytes, v1)
	if err != nil {
		log.Error("Failed to convert submission bytes to map", err)
		return nil, &Exception{Msg: err.Error()}
	}
	return v1, nil
}

type Exception struct {
	Msg string
}

func (e *Exception) Error() string {
	return e.Msg
}
