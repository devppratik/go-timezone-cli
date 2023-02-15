package tmz

import (
	"os"

	"github.com/aquasecurity/table"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DisplayTable(tableItems []string, rows int, cols int) {
	app := tview.NewApplication()
	table := GetTableWidget(tableItems, rows, cols)
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}

func GetTableWidget(tableItems []string, rows int, cols int) *tview.Table {
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
				tview.NewTableCell(tableItems[word]).
					SetTextColor(color).
					SetAlign(tview.AlignCenter))
			word = (word + 1) % len(tableItems)
		}
	}
	return table
}

func DisplayNewTable(rowItems [][]string, rowHeaders ...string) {
	t := table.New(os.Stdout)
	t.SetLineStyle(table.StyleWhite)
	t.SetHeaders(rowHeaders...)
	t.SetHeaderStyle(table.StyleBold)
	t.SetHeaderStyle(table.StyleBrightYellow)
	for _, row := range rowItems {
		t.AddRows(row)
	}
	t.Render()
}
