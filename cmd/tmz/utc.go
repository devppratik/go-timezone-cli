package tmz

import (
	"log"
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/thlib/go-timezone-local/tzlocal"
)

var utcCmd = &cobra.Command{
	Use:   "utc",
	Short: "Gives the time in UTC",
	Long:  "Outputs the current UTC Time along with the current time at the local timezone",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		currentUTCTime := tmzUtils.GetCurrentTimeAtLocation("UTC")
		currentLocalTime := tmzUtils.GetCurrentTimeAtLocation("Local")
		localTZName, err := tzlocal.RuntimeTZ()
		if err != nil {
			log.Fatalln(err)
		}
		tableHeaders := []string{"Local Time Zone", "Local Date Time", " UTC Time "}
		tableItems := [][]string{{localTZName, currentLocalTime, currentUTCTime}}
		tmzUI.DisplayTable(tableItems, tableHeaders...)
	},
}

func init() {
	rootCmd.AddCommand(utcCmd)
}
