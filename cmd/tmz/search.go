package tmz

import (
	"fmt"
	"log"
	"strings"
	"time"

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
		countryList := gojsonq.New().File("pkg/country.json")

		if len(countryCode) > 2 {
			res, _ := countryList.From("ALL").GetR()
			conv, _ := res.StringSlice()
			var lines []string
			app := tview.NewApplication()
			list := tview.NewList()
			var foundSearchItem string
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
			if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
				panic(err)
			}
			loc, err := time.LoadLocation(foundSearchItem)
			if err != nil {
				fmt.Print("error")
			}
			now := time.Now().In(loc).Format(time.Kitchen)
			fmt.Println("ZONE : ", foundSearchItem, "Current Time :", now)
			return
		}
		res, _ := countryList.From(strings.ToUpper(countryCode)).GetR()
		conv, _ := res.StringSlice()
		app := tview.NewApplication()
		list := tview.NewList()
		var foundSearchItem string
		for _, line := range conv {
			list.AddItem(line, "", 'a', func() {
				foundSearchItem = conv[list.GetCurrentItem()]
				app.Stop()
			})
		}
		if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
			panic(err)
		}
		loc, err := time.LoadLocation(foundSearchItem)
		if err != nil {
			fmt.Print("error")
		}
		now := time.Now().In(loc).Format(time.Kitchen)
		out := []string{"Time Zone", "Current Time", foundSearchItem, now}
		displayTableShow(out)
		fmt.Println("ZONE : ", foundSearchItem, "Current Time :", now)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
