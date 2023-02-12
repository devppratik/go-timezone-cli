package tmz

import (
	"fmt"
	"time"
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var selectedLocation string
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Displays local datetime of all saved timezones",
	Run: func(cmd *cobra.Command, args []string) {
		lines := tmzUtils.ReadConfigFile()
		app := tview.NewApplication()
		list := tview.NewList()
		item := 'a'
		for _, line := range lines {
			list.AddItem(line, "", item, func() {
				selectedLocation = lines[list.GetCurrentItem()]
				app.Stop()
			})
			item += 1
		}
		if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
			panic(err)
		}
		loc, err := time.LoadLocation(selectedLocation)
		if err != nil {
			fmt.Print("error")
		}
		now := time.Now().In(loc).Format(time.Stamp)
		out := []string{"Time Zone", "Current Time", selectedLocation, now}
		tmzUI.DisplayTable(out, len(out)/2, 2)
		fmt.Println("ZONE : ", selectedLocation, "Current Time :", now)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
