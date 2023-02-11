package tmz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	tmz "tmz/pkg/ui"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays local datetime of all saved timezones",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		currentTime := true
		if len(args) > 1 {
			log.Fatal("Wrong number of arguments")
		} else if len(args) == 1 {
			currentTime = false
		}
		home, _ := os.UserHomeDir()
		file, err := os.OpenFile(home+"/.tmz.list", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		sc := bufio.NewScanner(file)
		lines := make([]string, 0)
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
		out := []string{"Time Zone"}
		if !currentTime {
			out = append(out, "Converted Time")
		} else {
			out = append(out, "Current Time")
		}

		for _, line := range lines {
			loc, err := time.LoadLocation(line)
			if err != nil {
				fmt.Print("error")
			}
			now := time.Now()
			if !currentTime {
				currentTZ := time.Now().Local().Location()
				now, _ = time.ParseInLocation("2006-01-02 15:04", "2023-01-01 "+args[0], currentTZ)

			}
			timetoConvert := now.In(loc).Format(time.Stamp)
			out = append(out, line, timetoConvert)
			fmt.Println("ZONE : ", line, "Current Time :", timetoConvert)
		}
		tmz.DisplayTable(out, len(out)/2, 2)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
