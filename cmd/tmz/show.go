package tmz

import (
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var showAll bool
var customTime string
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays local datetime of all saved timezones",
	Long:  "Get a List of All Locally Saved Timezones defined in the config File. Pass custom time as argument to get the time at different zones at the specified time",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var tableItems = [][]string{}
		var tableHeaders = []string{"Time Zone", "Local Date Time"}

		listOfTimeZones := []string{"UTC", "Local"}
		if showAll {
			listOfTimeZonesFromConfig := tmzUtils.ReadConfigFile()
			listOfTimeZones = append(listOfTimeZones, listOfTimeZonesFromConfig...)
		}
		listOfTimeZones = append(listOfTimeZones, args...)
		if len(args) > 0 {
			customTime = tmzUtils.GetLocalTimeFromTimeZone(args[0], customTime)
		}
		for _, tmZone := range listOfTimeZones {
			displayTime := tmzUtils.GetConvertedTimeAtLocation(tmZone, customTime)
			if tmZone == args[0] {
				tmZone += "*"
			}
			tableItems = append(tableItems, []string{tmZone, displayTime})
		}
		tmzUI.DisplayTable(tableItems, tableHeaders...)
	},
}

func init() {
	showCmd.Flags().BoolVar(&showAll, "all", false, "Load all Timezones from local config")
	showCmd.Flags().StringVarP(&customTime, "time", "t", "", "Show the list at a different time")
	rootCmd.AddCommand(showCmd)
}
