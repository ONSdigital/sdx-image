package submission

import (
	"encoding/json"
	"sdxImage/pkg/interfaces"
	"sdxImage/pkg/log"
)

func Read(bytes []byte) (interfaces.Submission, error) {
	submission := &V2Submission{}
	err := json.Unmarshal(bytes, submission)
	if err != nil {
		v1 := &V1Submission{}
		err2 := json.Unmarshal(bytes, v1)
		if err2 != nil {
			log.Error("Failed to convert submission bytes to map", err)
			return nil, &Exception{Msg: err.Error()}
		}
		return v1, nil
	}
	return submission, nil
}

type Exception struct {
	Msg string
}

func (e *Exception) Error() string {
	return e.Msg
}
