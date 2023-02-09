package tmz

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a TimeZone to the Local Saved TimeZones",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tmzone := args[0]
		home, _ := os.UserHomeDir()
		file, err := os.OpenFile(home+"/.tmz.list", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		sc := bufio.NewScanner(file)
		lines := make([]string, 0)
		for sc.Scan() {
			lines = append(lines, sc.Text())
		}
		if err := sc.Err(); err != nil {
			log.Fatal(err)
		}
		if slices.Contains(lines, tmzone) {
			fmt.Printf("TimeZone Already Exists")
			return
		}
		if _, err = file.WriteString(tmzone + "\n"); err != nil {
			panic(err)
		}
		fmt.Print("Added Timezone ", tmzone, " successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
