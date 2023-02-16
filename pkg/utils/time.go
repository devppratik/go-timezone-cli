package tmz

import (
	"log"
	"time"
)

func GetCurrentTimeAtLocation(timeZone string) string {
	location, time_error := time.LoadLocation(timeZone)
	if time_error != nil {
		log.Fatalln(time_error)
	}
	currentTime := time.Now().In(location).Format(time.Stamp)
	return currentTime

}

func GetLocalTimeFromTimeZone(timezone string, customTime string) string {
	location, time_err := time.LoadLocation(timezone)
	if time_err != nil {
		log.Fatalln(time_err)
	}
	var dateToday string = time.Now().UTC().Format("2006-01-02")
	currentTZ := time.Now().Local().Location()
	if customTime == "" {
		customTime = GetCurrentTimeAtLocation(timezone)[7:12]
	}
	currentTime, time_err := time.ParseInLocation("2006-01-02 15:04", dateToday+" "+customTime, location)
	if time_err != nil {
		log.Fatalln(err)
	}
	locationTime := currentTime.In(currentTZ).Format(time.Stamp)[7:12]
	return locationTime
}

func GetConvertedTimeAtLocation(timezone string, customTime string) string {
	location, time_err := time.LoadLocation(timezone)
	if time_err != nil {
		log.Fatalln(time_err)
	}
	var dateToday string = time.Now().UTC().Format("2006-01-02")
	currentTZ := time.Now().Local().Location()
	currentTime, time_err := time.ParseInLocation("2006-01-02 15:04", dateToday+" "+customTime, currentTZ)
	if time_err != nil {
		log.Fatalln(err)
	}
	locationTime := currentTime.In(location).Format(time.Stamp)
	return locationTime
}
