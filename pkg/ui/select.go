package tmz

import (
	"log"

	"github.com/manifoldco/promptui"
)

func SelectTimeZone(listOfTimeZones []string) string {
	prompt := promptui.Select{
		Label: "Select the TimeZone",
		Items: listOfTimeZones,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalln("Prompt failed ", err)
	}
	return result
}
