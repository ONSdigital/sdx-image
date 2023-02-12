package read

import (
	"fmt"
	"sdxImage/pkg/log"
	"sdxImage/pkg/model"
)

func Submission(bytes []byte) (*model.Submission, error) {
	m, err := toCompleteMap(bytes)
	if err != nil {
		log.Error("Failed to convert submission bytes to map", err)
		return nil, &model.SubmissionError{Msg: err.Error()}
	}
	s, e := toSubmission(m)
	if e != nil {
		return nil, e
	}
	return s, nil
}

func toSubmission(m map[string]any) (*model.Submission, error) {
	version := getStringFrom(m, "version")
	var submission *model.Submission
	if version == "v2" {
		submission = fromV2(m)
	} else {
		submission = fromV1(m)
	}
	missing := model.MissingFields(submission)
	if len(missing) > 0 {
		err := &model.SubmissionError{Msg: fmt.Sprintf("missing required fields: %v", missing)}
		log.Error("Invalid submission", err)
		return nil, err
	}
	return submission, nil
}

func fromV1(m map[string]any) *model.Submission {
	submission := &model.Submission{}
	submission.TxId = getStringFrom(m, "tx_id")
	submission.SchemaName = locateStringFrom(m, "collection", "schema_name")
	submission.SubmittedAt = getStringFrom(m, "submitted_at")

	metadata := getMapFrom(m, "metadata")
	submission.RuRef = getStringFrom(metadata, "ru_ref")
	submission.StartDate = getStringFrom(metadata, "ref_period_start_date")
	submission.EndDate = getStringFrom(metadata, "ref_period_end_date")

	//ru name is not present on v1 submissions
	submission.RuName = "the business"

	submission.DataVersion = getStringFrom(m, "version")
	submission.Data = getData(m)
	return submission
}

func fromV2(m map[string]any) *model.Submission {
	submission := &model.Submission{}
	submission.TxId = getStringFrom(m, "tx_id")
	submission.SchemaName = getStringFrom(m, "schema_name")
	submission.SubmittedAt = getStringFrom(m, "submitted_at")

	metadata := getMapFrom(m, "survey_metadata")
	submission.RuRef = getStringFrom(metadata, "ru_ref")
	submission.RuName = getStringFrom(metadata, "ru_name")
	submission.StartDate = getStringFrom(metadata, "ref_p_start_date")
	submission.EndDate = getStringFrom(metadata, "ref_p_end_date")

	submission.DataVersion = getStringFrom(m, "data_version")
	submission.Data = getData(m)
	return submission
}

func getData(m map[string]any) map[string]string {
	data := getMapFrom(m, "data")
	dataMap := make(map[string]string)
	for key, value := range data {
		strValue := fmt.Sprintf("%v", value)
		dataMap[key] = strValue
	}
	return dataMap
}