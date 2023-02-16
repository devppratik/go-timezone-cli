package tmz

import (
	"os"

	"github.com/aquasecurity/table"
)

func DisplayTable(rowItems [][]string, rowHeaders ...string) {
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
