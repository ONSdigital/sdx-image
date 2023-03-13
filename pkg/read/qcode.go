package read

import "strconv"

func getQcodeMap(m map[string]any) map[string]string {
	answerCodes := getListFrom(m, "answer_codes")
	qCodeMap := make(map[string]string, len(answerCodes))
	for _, a := range answerCodes {
		answer := toMap(a)
		code := getStringFrom(answer, "code")
		qCode := getQcode(code)
		qCodeMap[getStringFrom(answer, "answer_id")] = qCode
	}
	return qCodeMap
}

func getQcode(code string) string {
	_, err := strconv.Atoi(code)
	if err != nil {
		return getQcode(code[1:])
	}
	return code
}
