package tmz

import (
	"fmt"
	"log"
	"time"
	tmzUI "tmz/pkg/ui"
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays local datetime of all saved timezones",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		currentTime := true
		out := []string{"Time Zone"}
		lines := tmzUtils.ReadConfigFile()

		if len(args) > 1 {
			log.Fatal("Wrong number of arguments. Expected 0 or 1 Arguments. Recieved ", len(args))
		} else if len(args) == 1 {
			currentTime = false
		}

		if !currentTime {
			out = append(out, "Converted Time")
		} else {
			out = append(out, "Current Time")
		}
		dateToday := time.Now().UTC().Format("2006-01-02")
		fmt.Println(dateToday)
		for _, line := range lines {
			loc, err := time.LoadLocation(line)
			now := time.Now()
			if err != nil {
				fmt.Print("error")
			}
			if !currentTime {
				currentTZ := time.Now().Local().Location()

				now, _ = time.ParseInLocation("2006-01-02 15:04", dateToday+" "+args[0], currentTZ)
			}
			timetoConvert := now.In(loc).Format(time.Stamp)
			out = append(out, line, timetoConvert)
			fmt.Println("ZONE : ", line, "Current Time :", timetoConvert)
		}
		tmzUI.DisplayTable(out, len(out)/2, 2)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
