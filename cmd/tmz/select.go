package tmz

import (
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
		customTime := tmzUI.PromptForCustomTime()
		displayTime := tmzUtils.GetCurrentTimeAtLocation(selectedLocation)
		if customTime != "" {
			displayTime = tmzUtils.GetConvertedTimeAtLocation(selectedLocation, customTime)
		}
		tableHeaders := []string{"Time Zone", "Local Date Time"}
		tableItems := [][]string{{selectedLocation, displayTime}}
		tmzUI.DisplayTable(tableItems, tableHeaders...)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
