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
		return nil, err
	}
	s := toSubmission(m)
	fmt.Println(s.Data)
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
	submission.SchemaName = locateStringFrom(m, "collection", "schema_name")
	submission.SubmittedAt = getStringFrom(m, "submitted_at")

	metadata := getMapFrom(m, "metadata")
	submission.RuRef = getStringFrom(metadata, "ru_ref")
	submission.StartDate = getStringFrom(metadata, "ref_period_start_date")
	submission.EndDate = getStringFrom(metadata, "ref_period_end_date")

	submission.DataVersion = getStringFrom(m, "version")
	submission.Data = getData(m)
	return submission
}

func fromV2(m map[string]any) *model.Submission {
	fmt.Println(m)
	submission := &model.Submission{}
	submission.TxId = getStringFrom(m, "tx_id")
	submission.SchemaName = getStringFrom(m, "schema_name")
	submission.SubmittedAt = getStringFrom(m, "submitted_at")

	metadata := getMapFrom(m, "survey_metadata")
	submission.RuRef = getStringFrom(metadata, "ru_ref")
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
