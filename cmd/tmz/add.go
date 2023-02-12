package tmz

import (
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a New TimeZone",
	Long:  "Adds a New TimeZone to the locally saved timezones in the config file. The config file is set at ~/.tmz.list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tmzone := args[0]
		tmzUtils.AddTimeZoneToConfig(tmzone)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
