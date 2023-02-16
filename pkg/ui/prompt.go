package tmz

import (
	"log"

	"github.com/manifoldco/promptui"
)

func PromptForCustomTime() string {
	prompt := promptui.Prompt{Label: "Enter Local Time to Convert(Press Enter for current time)"}
	result, err := prompt.Run()

	if err != nil {
		log.Fatalln("Prompt failed ", err)
	}
	return result
}
