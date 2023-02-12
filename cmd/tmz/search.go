package tmz

import (
	"fmt"
	"log"
	"strings"
	"time"
	tmzUI "tmz/pkg/ui"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/thedevsaddam/gojsonq/v2"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a timezones (Interactive)",
	Long:  "Search for timezone using the country abbrevaition or timezone name & get the current time for the timezone",
	Run: func(cmd *cobra.Command, args []string) {
		var selectedTimeZone string
		countryCode := args[0]
		if len(countryCode) < 2 {
			log.Fatalln("Wrong or Invalid Country Code or Name. Enter 2 or more characters")
		}
		// Read the Country List JSON
		countryList := gojsonq.New().File("pkg/data/country.json")
		app := tview.NewApplication()
		list := tview.NewList()

		if len(countryCode) > 2 {
			res, _ := countryList.From("ALL").GetR()
			conv, _ := res.StringSlice()
			var listOfTimeZones []string
			for _, tmZone := range conv {
				if strings.Contains(strings.ToLower(tmZone), countryCode) {
					listOfTimeZones = append(listOfTimeZones, tmZone)
				}
			}
			for _, tmZone := range listOfTimeZones {
				list.AddItem(tmZone, "", 'a', func() {
					selectedTimeZone = listOfTimeZones[list.GetCurrentItem()]
					app.Stop()
				})
			}
		} else {
			res, _ := countryList.From(strings.ToUpper(countryCode)).GetR()
			conv, _ := res.StringSlice()
			for _, tmZone := range conv {
				list.AddItem(tmZone, "", 'a', func() {
					selectedTimeZone = conv[list.GetCurrentItem()]
					app.Stop()
				})
			}
		}

		location, err := time.LoadLocation(selectedTimeZone)
		if err != nil {
			fmt.Print("error")
		}
		// pages := tview.NewPages()
		// table := tmzUI.GetTableWidget(tableItems, len(tableItems)/2, 2)
		// fmt.Println("ZONE : ", foundSearchItem, "Current Time :", currentTime)
		// pages.AddPage("List", list, true, true)
		// pages.AddPage("Display", table, true, false)
		currentTime := time.Now().In(location).Format(time.Stamp)
		tableItems := []string{"Time Zone", "Current Time", selectedTimeZone, currentTime}
		if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
			panic(err)
		}
		tmzUI.DisplayTable(tableItems, len(tableItems)/2, 2)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
