package tmz

import (
	"fmt"
	"log"
	"strings"
	"time"
	tmzUI "tmz/pkg/ui"

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

		if len(countryCode) > 2 {
			res, _ := countryList.From("ALL").GetR()
			conv, _ := res.StringSlice()
			var listOfTimeZones []string
			for _, tmZone := range conv {
				if strings.Contains(strings.ToLower(tmZone), countryCode) {
					listOfTimeZones = append(listOfTimeZones, tmZone)
				}
			}
			selectedTimeZone = tmzUI.SelectTimeZone(listOfTimeZones)
		} else {
			res, _ := countryList.From(strings.ToUpper(countryCode)).GetR()
			conv, _ := res.StringSlice()
			selectedTimeZone = tmzUI.SelectTimeZone(conv)
		}

		location, err := time.LoadLocation(selectedTimeZone)
		if err != nil {
			fmt.Print("error")
		}
		currentTime := time.Now().In(location).Format(time.Stamp)
		tableHeaders := []string{"Time Zone", "Current Time"}
		tableItems := [][]string{{selectedTimeZone, currentTime}}
		tmzUI.DisplayNewTable(tableItems, tableHeaders...)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
