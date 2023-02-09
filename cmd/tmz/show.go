package tmz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
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
		out := []string{"Time Zone", "Current Time"}
		for _, line := range lines {
			loc, err := time.LoadLocation(line)
			if err != nil {
				fmt.Print("error")
			}
			now := time.Now().In(loc).Format(time.Kitchen)
			out = append(out, line, now)
			fmt.Println("ZONE : ", line, "Current Time :", now)
		}
		localTime := time.Now().Format(time.Kitchen)
		tzname, _ := tzlocal.RuntimeTZ()
		fmt.Println("ZONE : ", tzname, "Local Time :", localTime)
		displayTableShow(out)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func displayTableShow(out []string) {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)
	cols, rows := 2, len(out)/2
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
