package tmz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var selectedLocation string
var selectCmd = &cobra.Command{
	Use:   "select",
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
		app := tview.NewApplication()
		list := tview.NewList()
		item := 'a'
		for _, line := range lines {
			list.AddItem(line, "", item, func() {
				selectedLocation = lines[list.GetCurrentItem()]
				app.Stop()
			})
			item += 1
		}
		if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
			panic(err)
		}
		loc, err := time.LoadLocation(selectedLocation)
		if err != nil {
			fmt.Print("error")
		}
		now := time.Now().In(loc).Format(time.Kitchen)
		out := []string{"Time Zone", "Current Time", selectedLocation, now}
		displayTableShow(out)
		fmt.Println("ZONE : ", selectedLocation, "Current Time :", now)
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)
}
