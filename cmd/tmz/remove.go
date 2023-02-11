package tmz

import (
	tmz "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a TimeZone from the Local Saved TimeZones",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tmzone := args[0]
		tmz.RemoveTimeZoneFromConfig(tmzone)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
