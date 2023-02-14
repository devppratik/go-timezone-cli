package tmz

import (
	tmzUtils "tmz/pkg/utils"

	"github.com/spf13/cobra"
)

var showLocalConfig bool = false
var visualCmd = &cobra.Command{
	Use:   "visual",
	Short: "Get a visual UI for timezones",
	Run: func(cmd *cobra.Command, args []string) {
		if showLocalConfig {
			lisOfTz := tmzUtils.ReadConfigFile()
			tmzUtils.VisualTimeZone(lisOfTz)
		} else {
			tmzUtils.VisualTimeZone(args)
		}
	},
}

func init() {
	visualCmd.Flags().BoolVar(&showLocalConfig, "all", false, "Load all Timezones from local config")
	rootCmd.AddCommand(visualCmd)
}
