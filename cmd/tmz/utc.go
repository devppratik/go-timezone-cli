package tmz

import (
	"fmt"
	"time"
	tmzUI "tmz/pkg/ui"

	"github.com/spf13/cobra"
	"github.com/thlib/go-timezone-local/tzlocal"
)

var utcCmd = &cobra.Command{
	Use:   "utc",
	Short: "Gives the time in UTC",
	Run: func(cmd *cobra.Command, args []string) {
		loc, err := time.LoadLocation("UTC")
		if err != nil {
			fmt.Print("error")
		}
		now := time.Now().In(loc).Format(time.Stamp)
		localTime := time.Now().Format(time.Stamp)
		tzname, _ := tzlocal.RuntimeTZ()
		fmt.Println("ZONE : ", tzname, "Local Time :", localTime, " UTC Time : ", now)
		out := []string{"Local Time Zone", "Local Time", " UTC Time ", tzname, localTime, now}
		tmzUI.DisplayTable(out, 2, 3)
	},
}

func init() {
	rootCmd.AddCommand(utcCmd)
}
