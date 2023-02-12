package tmz

import (
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a TimeZone",
	Long:  "Removes a Given TimeZone from the locally saved timezones in the config file. The config file is set at ~/.tmz.list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tmzone := args[0]
		tmzUtils.RemoveTimeZoneFromConfig(tmzone)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
