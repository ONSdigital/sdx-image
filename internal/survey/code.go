package survey

import "strconv"

func getQCode(code, surveyId string) string {
	if code == "" {
		return ""
	}
	switch surveyId {
	case "002":
		return removeLetters(code)
	case "068":
		return removeLetters(code)
	case "071":
		return removeLetters(code)
	case "066":
		return removeLetters(code)
	case "076":
		return removeLetters(code)
	case "221":
		return getBresCode(code)
	case "024":
		return getFuelsCode(code)
	case "194":
		return getRailwaysCode(code)
	default:
		return code
	}
}

// removeLetters removes letters from the beginning
// of the code leaving the rightmost digits.
// E.G. "a4b123" -> "123"
func removeLetters(code string) string {
	if len(code) == 1 {
		return code
	}
	_, err := strconv.Atoi(code)
	if err != nil {
		return removeLetters(code[1:])
	}

	return code
}

// getBresCode temporary fix for bad qcode for bres.
func getBresCode(code string) string {
	if code == "156c" {
		return "156"
	}
	return removeLetters(code)
}

func getFuelsCode(code string) string {
	mapping := map[string]string{
		"10":  "0a",
		"11":  "0b",
		"12":  "0c",
		"13":  "0d",
		"14":  "0e",
		"110": "1",
		"120": "2a",
		"121": "2b",
		"122": "2c",
		"130": "3",
		"140": "4a",
		"141": "4b",
		"142": "4c",
		"150": "5",
		"160": "6",
		"180": "8",
		"190": "9",
		"200": "11",
		"210": "12",
		"211": "12a",
		"220": "13",
		"230": "15",
		"240": "16",
		"250": "18",
		"260": "19",
		"270": "20",
		"271": "20a",
		"280": "21",
		"290": "23",
		"300": "24",
		"310": "26",
		"320": "27",
		"330": "28",
		"340": "29",
		"350": "31",
		"360": "32",
		"370": "34",
		"146": "146",
		"12a": "17",
		"20a": "25",
		"28":  "33",
	}

	c, found := mapping[code]
	if found {
		return c
	}
	return code
}

func getRailwaysCode(code string) string {
	mapping := map[string]string{
		"2":   "1.1",
		"3":   "1.2",
		"4":   "2.1",
		"5":   "2.2",
		"6":   "3.1",
		"7":   "3.2",
		"8":   "3.3",
		"9":   "3.4",
		"10":  "4.1",
		"13":  "4.2",
		"146": "146",
	}

	c, found := mapping[code]
	if found {
		return c
	}
	return code
}
