package tmz

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DisplayTable(out []string, rows int, cols int) {
	app := tview.NewApplication()
	table := GetTableWidget(out, rows, cols)
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}

func GetTableWidget(out []string, rows int, cols int) *tview.Table {
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
	return table
}
