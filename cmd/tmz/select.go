package tmz

import (
	"log"
	"time"
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a timezone from all saved timezones",
	Long:  "Select a timezone from all saved timezones to view the current time in that timezone",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var selectedLocation string
		var listOfTimeZones []string = tmzUtils.ReadConfigFile()
		selectedLocation = tmzUI.SelectTimeZone(listOfTimeZones)
		location, err := time.LoadLocation(selectedLocation)
		if err != nil {
			log.Fatalln(err)
		}
		currentTime := time.Now().In(location).Format(time.Stamp)
		tableHeaders := []string{"Time Zone", "Current Time"}
		tableItems := []string{selectedLocation, currentTime}
		tmzUI.DisplayNewTable(tableItems, tableHeaders...)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
