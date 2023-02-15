package tmz

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DisplayInteractiveTable(tableRows [][]string) {
	table := tview.NewTable().
		SetFixed(1, 1).
		SetSelectable(false, true).
		SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorDarkGreen))
	for row := 0; row < len(tableRows); row++ {
		for column := 0; column < len(tableRows[0]); column++ {
			color := tcell.ColorWhite
			align := tview.AlignLeft
			if row%2 == 0 {
				color = tcell.ColorLightGreen
			}
			table.SetCell(row,
				column,
				&tview.TableCell{
					Text:          tableRows[row][column],
					Color:         color,
					NotSelectable: row%2 == 0,
					Attributes:    tcell.AttrBold,
					Align:         align,
					Transparent:   true,
					Expansion:     1,
				})

		}
	}
	tview.NewApplication().
		SetRoot(table, true).
		Run()
}
