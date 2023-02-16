package tmz

import (
	tmzUI "tmz/pkg/ui"

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
		// var dateToday string = time.Now().UTC().Format("2006-01-02")
		// var listOfTimeZones []string = tmzUtils.ReadConfigFile()
		// displayTime := ""
		// if()
		// for _, tmZone := range listOfTimeZones {
		// 	location, err := time.LoadLocation(tmZone)
		// 	currentTime := time.Now()
		// 	if err != nil {
		// 		log.Fatalln(err)
		// 	}
		// 	if !cntTime {
		// 		currentTZ := time.Now().Local().Location()
		// 		currentTime, err = time.ParseInLocation("2006-01-02 15:04", dateToday+" "+args[0], currentTZ)
		// 		if err != nil {
		// 			log.Fatalln(err)
		// 		}
		// 	}
		// 	locationTime := currentTime.In(location).Format(time.Stamp)
		// tableItems = append(tableItems, []string{tmZone, displayTime})
		// }
		tmzUI.DisplayNewTable(tableItems, tableHeaders...)
	},
}

func init() {
	showCmd.Flags().BoolVar(&showAll, "all", false, "Load all Timezones from local config")
	showCmd.Flags().StringVarP(&customTime, "time", "t", "", "Show the list at a different time")
	rootCmd.AddCommand(showCmd)
}
