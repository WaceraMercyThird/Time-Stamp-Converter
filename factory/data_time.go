package factory

import (
	"fmt"
	"time"
)

func ConvertUTCToLocal(utcTime time.Time, timeZone string) (time.Time, error) {
	// Load the specified time zone location
	location, err := time.LoadLocation(timeZone)

	if err != nil {

		return time.Time{}, fmt.Errorf("error loading location: %w", err)
	}

	// Convert UTC time to local time in the specified time zone
	localTime := utcTime.In(location)
	
	return localTime, nil
}

func ConUTCToLocal(utcTime time.Time) time.Time {
	// Convert UTC time to local time in the system's time zone
	localTime := utcTime.In(time.Local)
	return localTime
}