package tmz

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
		now := time.Now().In(loc).Format(time.Kitchen)
		localTime := time.Now().Format(time.Kitchen)
		tzname, _ := tzlocal.RuntimeTZ()
		fmt.Println("ZONE : ", tzname, "Local Time :", localTime, " UTC Time : ", now)
		out := []string{"Local Time Zone", "Local Time", " UTC Time ", tzname, localTime, now}
		displayTable(out, 2, 3)
	},
}

func init() {
	rootCmd.AddCommand(utcCmd)
}

func displayTable(out []string, rows int, cols int) {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)
	word := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			if c < 1 || r < 1 {
				color = tcell.ColorYellow
			}
			table.SetCell(r, c,
				tview.NewTableCell(out[word]).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
			word = (word + 1) % len(out)
		}
	}
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
