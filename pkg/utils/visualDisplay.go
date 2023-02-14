package tmz

import (
	"fmt"
	"log"
	"time"
	tmzUI "tmz/pkg/ui"

	"github.com/kyokomi/emoji/v2"
)

func VisualTimeZone(zones []string) {
	zones = append([]string{"UTC"}, zones...)
	var dateToday string = time.Now().UTC().Format("2006-01-02")
	if err != nil {
		log.Fatalln(err)
	}
	tableRows := [][]string{}
	localemptyRow := make([]string, 17)
	localemptyRow[0] = emoji.Sprint(":clock8:Local")
	tableRows = append(tableRows, localemptyRow)
	for i, zone := range zones {
		location, err := time.LoadLocation(zone)
		if err != nil {
			log.Fatalln(err)
		}

		currentTZ := time.Now().Local().Location()
		locationTimeRow := []string{}
		localTimeRow := []string{}
		emptyRow := make([]string, 17)
		emptyRow[0] = emoji.Sprintf(":clock%d:%s", i+1, zone)
		for i := 5; i < 22; i++ {
			timeVal := ""
			if i < 10 {
				timeVal += "0" + fmt.Sprint(i) + ":00"
			} else {
				timeVal = fmt.Sprint(i) + ":00"
			}
			currentTime, err := time.ParseInLocation("2006-01-02 15:00", dateToday+" "+timeVal, currentTZ)
			if err != nil {
				log.Fatalln(err)
			}
			locationTime := currentTime.In(location).Format(time.Stamp)
			timeFormat := locationTime[7:12] + " "
			locationTimeRow = append(locationTimeRow, timeFormat)
			if len(tableRows) == 1 {
				localTimeRow = append(localTimeRow, timeVal)
			}
		}
		if len(tableRows) == 1 {
			tableRows = append(tableRows, localTimeRow)
		}
		tableRows = append(tableRows, emptyRow, locationTimeRow)
	}
	tmzUI.DisplayInteractiveTable(tableRows)
}
