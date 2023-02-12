package tmz

import (
	"log"
	"time"
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

	"github.com/rivo/tview"
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
		var itemOption rune = 'a'
		app := tview.NewApplication()
		list := tview.NewList()

		for _, tmZone := range listOfTimeZones {
			list.AddItem(tmZone, "", itemOption, func() {
				selectedLocation = listOfTimeZones[list.GetCurrentItem()]
				app.Stop()
			})
			itemOption += 1
		}
		if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
			log.Fatalln(err)
		}

		location, err := time.LoadLocation(selectedLocation)
		if err != nil {
			log.Fatalln(err)
		}
		currentTime := time.Now().In(location).Format(time.Stamp)
		tableItems := []string{"Time Zone", "Current Time", selectedLocation, currentTime}
		tmzUI.DisplayTable(tableItems, len(tableItems)/2, 2)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
