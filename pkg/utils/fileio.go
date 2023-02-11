package tmz

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

func AddTimeZoneToConfig(tmzone string) {
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
	if _, err := file.WriteString(tmzone + "\n"); err != nil {
		panic(err)
	}
	fmt.Print("Added Timezone ", tmzone, " successfully")
}

func RemoveTimeZoneFromConfig(tmzone string) {
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
	for i, line := range lines {
		if strings.Contains(line, tmzone) {
			lines = slices.Delete(lines, i, i+1)
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(home+"/.tmz.list", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = file.WriteString("\n"); err != nil {
		panic(err)
	}
	fmt.Print("Removed Timezone ", tmzone, " successfully")
}
