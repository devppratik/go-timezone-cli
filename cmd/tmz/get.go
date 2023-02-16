package tmz

import (
	"fmt"
	"log"
	"strings"
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/thedevsaddam/gojsonq/v2"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a timezone using the timezone abbreviation",
	Long:  "Search for a specific timezone using the timezone abbreviation",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var countryCode string = strings.ToUpper(args[0])
		matchedTimeZones := []string{}
		countryList := gojsonq.New().File("pkg/data/abbr.json")
		res := countryList.From("zones").Where("abbr", "=", countryCode).Get()
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
			val := fmt.Sprintf("%s-%s-%s", abbr, name, zones)
			matchedTimeZones = append(matchedTimeZones, val)
		}
		if len(matchedTimeZones) == 0 {
			log.Fatalln("Wrong TimeZone Abbreviation! Please enter correct abbreviation according to IANA List")
		}
		tmZone := strings.Split(tmzUI.SelectTimeZone(matchedTimeZones), "-")[2]
		displayTime := ""
		if len(args) == 2 {
			displayTime = tmzUtils.GetConvertedTimeAtLocation(tmZone, args[1])
		} else {
			displayTime = tmzUtils.GetCurrentTimeAtLocation(tmZone)
		}
		tableHeaders := []string{"Time Zone", "Local Date Time"}
		tableItems := [][]string{{tmZone, displayTime}}
		tmzUI.DisplayNewTable(tableItems, tableHeaders...)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
