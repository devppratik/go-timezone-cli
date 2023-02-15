package tmz

import (
	"log"
	"strings"
	"time"
	tmzUI "tmz/pkg/ui"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/thedevsaddam/gojsonq/v2"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a timezone using the timezone abbreviation",
	Long:  "Search for a specific timezone using the timezone abbreviation",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var selectedTimeZone string
		var countryCode string = strings.ToUpper(args[0])
		matchedTimeZones := []string{}
		itemOption := 'a'

		// Read List From abbreviations JSON File
		countryList := gojsonq.New().File("pkg/data/abbr.json")
		res := countryList.From("zones").Where("abbr", "=", countryCode).Get()

		app := tview.NewApplication()
		list := tview.NewList()

		for _, item := range res.([]interface{}) {
			var zones, name, abbr string
			for key, value := range item.(map[string]interface{}) {
				if key == "utc" {
					utcname, _ := value.([]interface{})[0].(string)
					zones = utcname
				} else if key == "value" {
					name = value.(string)
				} else if key == "abbr" {
					abbr = value.(string)
				}
			}
			matchedTimeZones = append(matchedTimeZones, zones, name, abbr)
		}
		if len(matchedTimeZones) == 0 {
			log.Fatalln("Wrong TimeZone Abbreviation! Please enter correct abbreviation according to IANA List")
		}

		for i := 0; i < len(matchedTimeZones); i += 3 {
			list.AddItem(matchedTimeZones[i+2]+" "+matchedTimeZones[i+1], matchedTimeZones[i], itemOption, func() {
				selectedTimeZone = matchedTimeZones[list.GetCurrentItem()*3]
				app.Stop()
			})
			itemOption += 1
		}
		if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
			log.Fatalln(err)
		}

		location, err := time.LoadLocation(selectedTimeZone)
		if err != nil {
			log.Fatalln(err)
		}
		currentTZTime := time.Now().In(location).Format(time.Stamp)
		tableHeaders := []string{"Time Zone", "Current Time"}
		tableItems := [][]string{{selectedTimeZone, currentTZTime}}
		tmzUI.DisplayNewTable(tableItems, tableHeaders...)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
