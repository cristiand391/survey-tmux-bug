package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// the question to ask
	qs := []*survey.Question{
		{
			Name: "color",
			Prompt: &survey.Select{
				Message: "Choose a color:",
				Options: []string{"red", "blue", "green"},
				Default: "red",
			},
		},
	}

	// the answer will be written to this struct
	answers := struct {
		FavoriteColor string `survey:"color"`
	}{}

	// perform the question
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s", answers.FavoriteColor)
}
