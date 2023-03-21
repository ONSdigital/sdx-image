package read

func getQcodeMap(m map[string]any) map[string]string {
	answerCodes := getListFrom(m, "answer_codes")
	qCodeMap := make(map[string]string, len(answerCodes))
	for _, a := range answerCodes {
		answer := toMap(a)
		code := getStringFrom(answer, "code")
		qCode := code
		qCodeMap[getStringFrom(answer, "answer_id")] = qCode
	}
	return qCodeMap
}
