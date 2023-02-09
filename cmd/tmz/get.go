package tmz

import (
	"fmt"
	"strings"

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
		res := countryList.From("zones").Where("Abbreviation", "=", countryCode).Only("Name")
		fmt.Print(res)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
