package tmz

import (
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a TimeZone to the Local Saved TimeZones",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tmzone := args[0]
		tmzUtils.AddTimeZoneToConfig(tmzone)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
