package tmz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

var home, _ = os.UserHomeDir()
var file, err = os.OpenFile(home+"/.tmz.list", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)

func ReadConfigFile() []string {
	if err != nil {
		log.Fatal(err)
	}
	sc := bufio.NewScanner(file)
	lines := make([]string, 0)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func AddTimeZoneToConfig(tmzone string) {
	lines := ReadConfigFile()
	if slices.Contains(lines, tmzone) {
		fmt.Printf("TimeZone Already Exists")
		return
	}
	if _, err := file.WriteString(tmzone + "\n"); err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Added Timezone ", tmzone, " successfully")
}

func RemoveTimeZoneFromConfig(tmzone string) {
	lines := ReadConfigFile()
	for i, line := range lines {
		if strings.Contains(line, tmzone) {
			lines = slices.Delete(lines, i, i+1)
		}
	}
	output := strings.Join(lines, "\n")
	err := os.WriteFile(home+"/.tmz.list", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = file.WriteString("\n"); err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Removed Timezone ", tmzone, " successfully")
}
