package utils

import (
	"time"
)

func TimeFormatConverter(systemFormatTime time.Time)(databaseFormatTime string) {
	// Set the time like it is stored into CouchDB
	layout := "2006-01-02T15:04:05.000Z"

	formattedTime := systemFormatTime.Format(layout)

	return formattedTime
}
