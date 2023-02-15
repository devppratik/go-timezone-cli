package tmz

import (
	"log"
	"time"
	tmzUI "tmz/pkg/ui"

	"github.com/spf13/cobra"
	"github.com/thlib/go-timezone-local/tzlocal"
)

var utcCmd = &cobra.Command{
	Use:   "utc",
	Short: "Gives the time in UTC",
	Long:  "Outputs the current UTC Time along with the current time at the local timezone",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		utcLocation, err := time.LoadLocation("UTC")
		if err != nil {
			log.Fatalln(err)
		}
		currentUTCTime := time.Now().In(utcLocation).Format(time.Stamp)
		currentLocalTime := time.Now().Format(time.Stamp)
		localTZName, err := tzlocal.RuntimeTZ()
		if err != nil {
			log.Fatalln(err)
		}
		tableHeaders := []string{"Local Time Zone", "Local Time", " UTC Time "}
		tableItems := [][]string{{localTZName, currentLocalTime, currentUTCTime}}
		tmzUI.DisplayNewTable(tableItems, tableHeaders...)
	},
}

func init() {
	rootCmd.AddCommand(utcCmd)
}
