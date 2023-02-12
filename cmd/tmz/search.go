package tmz

import (
	"fmt"
	"log"
	"strings"
	"time"
	tmzUI "tmz/pkg/ui"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/thedevsaddam/gojsonq/v2"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for local date time from the timezones (Interactive)",
	Run: func(cmd *cobra.Command, args []string) {
		countryCode := args[0]
		if len(countryCode) < 2 {
			log.Fatal("Wrong or Invalid Country Code or Name. Enter 2 or more characters")
		}
		countryList := gojsonq.New().File("pkg/data/country.json")
		app := tview.NewApplication()
		// pages := tview.NewPages()
		list := tview.NewList()
		var foundSearchItem string

		if len(countryCode) > 2 {
			res, _ := countryList.From("ALL").GetR()
			conv, _ := res.StringSlice()
			var lines []string
			for _, line := range conv {
				if strings.Contains(strings.ToLower(line), countryCode) {
					lines = append(lines, line)
				}
			}
			for _, line := range lines {
				list.AddItem(line, "", 'a', func() {
					foundSearchItem = lines[list.GetCurrentItem()]
					app.Stop()
				})
			}
		} else {
			res, _ := countryList.From(strings.ToUpper(countryCode)).GetR()
			conv, _ := res.StringSlice()
			for _, line := range conv {
				list.AddItem(line, "", 'a', func() {
					foundSearchItem = conv[list.GetCurrentItem()]
					app.Stop()
				})
			}
		}
		loc, err := time.LoadLocation(foundSearchItem)
		if err != nil {
			fmt.Print("error")
		}

		now := time.Now().In(loc).Format(time.Stamp)
		out := []string{"Time Zone", "Current Time", foundSearchItem, now}

		// table := tmzUI.GetTableWidget(out, len(out)/2, 2)
		// fmt.Println("ZONE : ", foundSearchItem, "Current Time :", now)
		// pages.AddPage("List", list, true, true)
		// pages.AddPage("Display", table, true, false)
		if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
			panic(err)
		}
		tmzUI.DisplayTable(out, len(out)/2, 2)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
