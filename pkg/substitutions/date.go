package substitutions

import (
	"sdxImage/pkg/log"
	"time"
)

// The reference data is always for "Mon, 02 Jan 2006 15:04:05 MST"
func convertDate(date string) string {
	dt, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Error("Unable to convert date: "+date, err)
		return date
	}
	return dt.Format("02/01/2006")
}

func convertSubmittedAt(dateTime string) string {
	dt, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		log.Error("Unable to convert date: "+dateTime, err)
		return dateTime
	}
	return dt.Format("02 January 2006 15:04:05")
}
