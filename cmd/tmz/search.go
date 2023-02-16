package tmz

import (
	"log"
	"strings"
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

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

		countryList := gojsonq.New().File("pkg/data/country.json")

		if len(countryCode) > 2 {
			res, err := countryList.From("ALL").GetR()
			if err != nil {
				log.Fatalln("No Country with the given search term found")
			}
			conv, err := res.StringSlice()
			if err != nil {
				log.Fatalln("No Country with the given search term found")
			}
			var listOfTimeZones []string
			for _, tmZone := range conv {
				if strings.Contains(strings.ToLower(tmZone), countryCode) {
					listOfTimeZones = append(listOfTimeZones, tmZone)
				}
			}
			if len(listOfTimeZones) == 0 {
				log.Fatalln("No Country with the given search term found")
			}
			selectedTimeZone = tmzUI.SelectTimeZone(listOfTimeZones)
		} else {
			res, err := countryList.From(strings.ToUpper(countryCode)).GetR()
			if err != nil {
				log.Fatalln("No Country with the given search term found")
			}
			conv, _ := res.StringSlice()
			selectedTimeZone = tmzUI.SelectTimeZone(conv)
		}

		currentTime := tmzUtils.GetCurrentTimeAtLocation(selectedTimeZone)
		tableHeaders := []string{"Time Zone", "Current Time"}
		tableItems := [][]string{{selectedTimeZone, currentTime}}
		tmzUI.DisplayNewTable(tableItems, tableHeaders...)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
