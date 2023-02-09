package tmz

import (
	"fmt"
	"time"

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
		// displayTable()
	},
}

func init() {
	rootCmd.AddCommand(utcCmd)
}

// func displayTable() {
// 	app := tview.NewApplication()
// 	table := tview.NewTable().
// 		SetBorders(true)
// 	lorem := strings.Split("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.", " ")
// 	cols, rows := 2, 2
// 	word := 0
// 	for r := 0; r < rows; r++ {
// 		for c := 0; c < cols; c++ {
// 			color := tcell.ColorWhite
// 			if c < 1 || r < 1 {
// 				color = tcell.ColorYellow
// 			}
// 			table.SetCell(r, c,
// 				tview.NewTableCell(lorem[word]).
// 					SetTextColor(color).
// 					SetAlign(tview.AlignCenter))
// 			word = (word + 1) % len(lorem)
// 		}
// 	}
// 	if err := app.SetRoot(table, true).SetFocus(table).Run(); err != nil {
// 		panic(err)
// 	}
// }
