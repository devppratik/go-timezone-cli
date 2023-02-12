package tmz

import (
	"log"
	"time"
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays local datetime of all saved timezones",
	Long:  "Get a List of All Locally Saved Timezones defined in the config File. Pass custom time as argument to get the time at different zones at the specified time",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Variable to determine current time to display or custom time
		var cntTime bool = true
		if len(args) > 1 {
			log.Fatalln("Wrong number of arguments. Expected 1 or No Arguments but Recieved ", len(args))
		} else if len(args) == 1 {
			cntTime = false
		}

		var tableItems = []string{"Time Zone"}
		var dateToday string = time.Now().UTC().Format("2006-01-02")
		var listOfTimeZones []string = tmzUtils.ReadConfigFile()
		if !cntTime {
			tableItems = append(tableItems, "Converted Time")
		} else {
			tableItems = append(tableItems, "Current Time")
		}

		for _, tmZone := range listOfTimeZones {
			location, err := time.LoadLocation(tmZone)
			currentTime := time.Now()
			if err != nil {
				log.Fatalln(err)
			}
			if !cntTime {
				currentTZ := time.Now().Local().Location()
				currentTime, err = time.ParseInLocation("2006-01-02 15:04", dateToday+" "+args[0], currentTZ)
				if err != nil {
					log.Fatalln(err)
				}
			}
			locationTime := currentTime.In(location).Format(time.Stamp)
			tableItems = append(tableItems, tmZone, locationTime)
		}
		tmzUI.DisplayTable(tableItems, len(tableItems)/2, 2)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
