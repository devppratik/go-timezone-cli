package tmz

import (
	"fmt"
	"strings"
	"time"
	tmz "tmz/pkg/ui"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/thedevsaddam/gojsonq/v2"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a timezone using the timezone abbreviation",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var selectedTimeZone string
		countryCode := strings.ToUpper(args[0])
		countryList := gojsonq.New().File("pkg/abbr.json")
		res := countryList.From("zones").Where("abbr", "=", countryCode).Get()
		out := []string{}
		app := tview.NewApplication()
		list := tview.NewList()
		item := 'a'

		for _, value := range res.([]interface{}) {
			a, b, c := "", "", ""
			for key, value := range value.(map[string]interface{}) {
				if key == "utc" {
					utcname, _ := value.([]interface{})[0].(string)
					a = utcname
				} else if key == "value" {
					b = value.(string)
				} else if key == "abbr" {
					c = value.(string)
				}
			}
			out = append(out, a, b, c)
		}

		for i := 0; i < len(out); i += 3 {
			list.AddItem(out[i+2]+" "+out[i+1], out[i], item, func() {
				selectedTimeZone = out[list.GetCurrentItem()*3]
				app.Stop()
			})
			item += 1
		}
		if err := app.SetRoot(list, true).EnableMouse(false).Run(); err != nil {
			panic(err)
		}

		loc, err := time.LoadLocation(selectedTimeZone)
		if err != nil {
			fmt.Print("error")
		}
		now := time.Now().In(loc).Format(time.Stamp)
		items := []string{"Time Zone", "Current Time", selectedTimeZone, now}
		tmz.DisplayTable(items, len(items)/2, 2)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
