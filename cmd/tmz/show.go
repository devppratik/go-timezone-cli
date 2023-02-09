package tmz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/thlib/go-timezone-local/tzlocal"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Displays local datetime of all saved timezones",
	Run: func(cmd *cobra.Command, args []string) {
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

		for _, line := range lines {
			loc, err := time.LoadLocation(line)
			if err != nil {
				fmt.Print("error")
			}
			now := time.Now().In(loc).Format(time.Kitchen)
			fmt.Println("ZONE : ", line, "Current Time :", now)
		}
		localTime := time.Now().Format(time.Kitchen)
		tzname, _ := tzlocal.RuntimeTZ()
		fmt.Println("ZONE : ", tzname, "Local Time :", localTime)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
