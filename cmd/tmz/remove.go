package tmz

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a TimeZone from the Local Saved TimeZones",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tmzone := args[0]
		home, _ := os.UserHomeDir()
		file, err := ioutil.ReadFile(home + "/.tmz.list")
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(file), "\n")
		for i, line := range lines {
			if strings.Contains(line, tmzone) {
				lines = slices.Delete(lines, i, i+1)
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(home+"/.tmz.list", []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Print("Removed Timezone ", tmzone, " successfully")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
