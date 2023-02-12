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
var configFile, err = os.OpenFile(home+"/.tmz.list", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)

func ReadConfigFile() []string {
	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(configFile)
	listOfTimeZones := make([]string, 0)
	for sc.Scan() {
		listOfTimeZones = append(listOfTimeZones, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalln(err)
	}
	return listOfTimeZones
}

func AddTimeZoneToConfig(tmzone string) {
	listOfTimeZones := ReadConfigFile()
	if slices.Contains(listOfTimeZones, tmzone) {
		log.Fatalln("TimeZone Already Exists")
		return
	}
	if _, err := configFile.WriteString(tmzone + "\n"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Added Timezone ", tmzone, " successfully")
}

func RemoveTimeZoneFromConfig(tmzone string) {
	listOfTimeZones := ReadConfigFile()
	for i, line := range listOfTimeZones {
		if strings.Contains(line, tmzone) {
			listOfTimeZones = slices.Delete(listOfTimeZones, i, i+1)
		}
	}
	output := strings.Join(listOfTimeZones, "\n")
	err := os.WriteFile(home+"/.tmz.list", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = configFile.WriteString("\n"); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Removed Timezone ", tmzone, " successfully")
}
