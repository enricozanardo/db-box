package utils

import (
	"testing"
	"time"
	"fmt"
)


func Test_TimeFormatConverter(t *testing.T) {
	// Set the time like it is stored into the DB
	// Aspected format "2017-11-16T11:40:04.723Z"
	mytime := time.Now()

	formattedTime := TimeFormatConverter(mytime)

	if len(formattedTime) != 24 {
		t.Error("It was not possible to convert the time like in the CouchDB")
	}

	fmt.Printf("Ok, input time %s, formatted format %s", mytime, formattedTime)
}


