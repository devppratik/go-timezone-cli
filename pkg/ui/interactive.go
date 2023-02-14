package tmz

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DisplayInteractiveTable(tableRows [][]string) {
	table := tview.NewTable().
		SetFixed(1, 1).
		SetSelectable(false, true).
		SetSeparator(tview.Borders.Vertical)
	for row := 0; row < len(tableRows); row++ {
		for column := 0; column < len(tableRows[0]); column++ {
			val := string(tableRows[row][column])
			color := tcell.ColorWhite
			align := tview.AlignCenter
			if row%2 == 0 {
				color = tcell.ColorLightGreen
				align = tview.AlignLeft
			}
			if row%2 == 1 {
				i, _ := strconv.Atoi(val[0:2])
				if i <= 5 || i >= 20 {
					color = tcell.ColorDarkSlateBlue
				}
				if i >= 6 && i <= 8 {
					color = tcell.ColorLightCyan
				}
				if i >= 9 && i <= 16 {
					color = tcell.ColorYellow
				}
				if i >= 17 && i <= 19 {
					color = tcell.ColorOrange
				}
			}
			table.SetCell(row,
				column,
				&tview.TableCell{
					Text:          val,
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
