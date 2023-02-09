package tmz

import (
	"fmt"
	"strings"
	"time"

	"github.com/rivo/tview"
	"github.com/spf13/cobra"
	"github.com/thedevsaddam/gojsonq/v2"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a timezone using the timezone abbreviation",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		countryCode := strings.ToUpper(args[0])
		countryList := gojsonq.New().File("pkg/abbr.json")
		res := countryList.From("zones").Where("abbr", "=", countryCode).Get()
		// result, _ := res.StringSlice()
		// fmt.Println(res.([]interface{}))
		out := []string{}
		for _, value := range res.([]interface{}) {
			a, b, c := "", "", ""
			for key, value := range value.(map[string]interface{}) {
				if key == "utc" {
					// fmt.Println(key, value.([]interface{})[0])
					utcname, _ := value.([]interface{})[0].(string)
					a = utcname
				} else if key == "value" {
					b = value.(string)
					// fmt.Println(key, value)
				} else if key == "abbr" {
					c = value.(string)
					// fmt.Println(key, value)
				}
			}
			out = append(out, a, b, c)
		}
		app := tview.NewApplication()
		list := tview.NewList()
		item := 'a'
		var selectedTimeZone string
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
		fmt.Println(selectedTimeZone)
		loc, err := time.LoadLocation(selectedTimeZone)
		if err != nil {
			fmt.Print("error")
		}
		now := time.Now().In(loc).Format(time.Kitchen)
		items := []string{"Time Zone", "Current Time", selectedTimeZone, now}
		displayTableShow(items)
		// print(out)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
