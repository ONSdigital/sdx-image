package read

import (
	"fmt"
	"sdxImage/pkg/log"
	"sdxImage/pkg/model"
)

var surveyMap = map[string]string{"002": "berd", "073": "blocks", "074": "bricks", "202": "abs"}

func Submission(bytes []byte) (*model.Submission, error) {
	m, err := toCompleteMap(bytes)
	if err != nil {
		log.Error("Failed to convert submission bytes to map", err)
		return nil, &model.SubmissionError{Msg: err.Error()}
	}
	s := toSubmission(m)
	missing := model.MissingFields(s)
	if len(missing) > 0 {
		e := &model.SubmissionError{Msg: fmt.Sprintf("missing required fields: %v", missing)}
		log.Error("Invalid submission", e)
		return nil, e
	}
	return s, nil
}

func toSubmission(m map[string]any) *model.Submission {
	version := getStringFrom(m, "version")
	if version == "v2" {
		return fromV2(m)
	} else {
		return fromV1(m)
	}
}

func fromV1(m map[string]any) *model.Submission {
	submission := &model.Submission{}
	submission.TxId = getStringFrom(m, "tx_id")
	schemaName := locateStringFrom(m, "collection", "schema_name")

	if schemaName == "" {
		surveyId := surveyMap[getStringFrom(m, "survey_id")]
		instrument := locateStringFrom(m, "collection", "instrument_id")
		schemaName = surveyId + "_" + instrument
	}
	submission.SchemaName = schemaName

	submission.SubmittedAt = getStringFrom(m, "submitted_at")

	metadata := getMapFrom(m, "metadata")
	submission.RuRef = getStringFrom(metadata, "ru_ref")
	submission.StartDate = getStringFrom(metadata, "ref_period_start_date")
	submission.EndDate = getStringFrom(metadata, "ref_period_end_date")

	//ru name is not present on v1 submissions
	submission.RuName = "the business"

	submission.DataVersion = getStringFrom(m, "version")
	if submission.DataVersion == "0.0.3" {
		submission.Responses = getResponsesFromList(m)
	} else {
		submission.Responses = getResponsesFromData(m)
	}

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
	if submission.DataVersion == "0.0.3" {
		submission.Responses = getResponsesFromList(m)
	} else {
		submission.Responses = getResponsesFromData(m)
	}

	return submission
}

func getResponsesFromData(m map[string]any) []*model.Response {
	data := getMapFrom(m, "data")
	counter := 0
	responses := make([]*model.Response, len(data))
	for key, value := range data {
		strValue := fmt.Sprintf("%v", value)
		responses[counter] = &model.Response{
			QuestionCode: key,
			Value:        strValue,
			Instance:     0,
		}
		counter++
	}
	return responses
}

func getResponsesFromList(m map[string]any) []*model.Response {
	data := getListFrom(m, "data")
	responses := make([]*model.Response, len(data))

	for index, resp := range data {

		r := toMap(resp)

		responses[index] = &model.Response{
			QuestionCode: getStringFrom(r, "questioncode"),
			Value:        getStringFrom(r, "response"),
			Instance:     int(getFloat64From(r, "instance")),
		}
	}
	return responses
}
