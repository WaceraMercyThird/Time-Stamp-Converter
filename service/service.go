package service

import (
	"fmt"
	"go_praco/factory"
	"time"
)

func UtcToLocal() {
	fmt.Println("..............................")

	utc := time.Now().UTC()

	fmt.Println(utc)

	local := utc
	location, err := time.LoadLocation("Europe/Budapest")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
	local = utc
	location, err = time.LoadLocation("America/Los_Angeles")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))

	local = utc
	location, err = time.LoadLocation("Africa/Nairobi")
	if err == nil {
		local = local.In(location)
	}
	fmt.Println("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
	fmt.Printf("Local Time: %v\n", time.Now().Local())
	fmt.Printf("Time Now: %v", time.Now())

}

func DateTime() {
	fmt.Println("................................")

	utcTime := time.Now().UTC()
	fmt.Println("UTC Time:", utcTime.Format("2006-01-02 15:04:05"))

	// Convert UTC time to Africa/Nairobi time
	localTime, err := factory.ConvertUTCToLocal(utcTime, "Africa/Nairobi")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Africa/Nairobi Time:", localTime.Format("2006-01-02 15:04:05"))

	// Convert UTC time to Europe/Budapest time
	localTime, err = factory.ConvertUTCToLocal(utcTime, "Europe/Budapest")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Europe/Budapest Time:", localTime.Format("2006-01-02 15:04:05"))

	// Convert UTC time to America/Los_Angeles time
	localTime, err = factory.ConvertUTCToLocal(utcTime, "America/Los_Angeles")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("America/Los_Angeles Time:", localTime.Format("2006-01-02 15:04:05"))
}

func AutTime() {
	fmt.Println("............................ ")

	// Example UTC time
	utcTime := time.Now().UTC()
	fmt.Println("Auto UTC Time:", utcTime.Format("2006-01-02 15:04:05"))

	// Convert UTC time to local time
	localTime := factory.ConUTCToLocal(utcTime)
	fmt.Println("Auto Local Time:", localTime.Format("2006-01-02 15:04:05"))
}
