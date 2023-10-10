// main.go
package main

import (
	controller "GenCodeX/controller"
	db "GenCodeX/db"
	"GenCodeX/handler"
	"GenCodeX/models"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"

	"github.com/urfave/cli/v2"
)

func main() {
	var language string = "javascript"

	color.Cyan("CodegenX : Generate boilerplate code for your project in no time")

	client := db.Dbconnect()

	app := &cli.App{
		Name:  "CodeGenX",
		Usage: "Generate boilerplate code for your project in no time",
		Authors: []*cli.Author{{
			Name:  "Kanak Kholwal",
			Email: "kanakkholwal@gmail.com",
		}},
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "javascript",
				Usage:       "language for the code to be generated",
				Destination: &language,
			},
		},
		Action: func(cCtx *cli.Context) error {

			langcontroller := controller.GetLangs(client)
			if len(langcontroller) == 0 {
				color.Magenta("No languages found")
				return nil
			}
			// Sort the language controller alphabetically
			sort.SliceStable(langcontroller, func(i, j int) bool {
				return langcontroller[i].ProgLang < langcontroller[j].ProgLang
			})

			// Create a prompt for user selection
			langPrompt := promptui.Select{
				Label: "Select a language",
				Items: Map(langcontroller, func(lang controller.Options) string {
					return lang.ProgLang
				}),
			}

			_, selectedLang, err := langPrompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				panic(err)
			}

			// Perform a search to find the selected language
			var langIndex = 0
			for idx, lang := range langcontroller {
				if lang.ProgLang == selectedLang {
					langIndex = idx
					break
				}
			}
			topicPrompt := promptui.Select{
				Label: "Select a Topic",
				Items: langcontroller[langIndex].Topics,
			}

			_, selectedTopic, err := topicPrompt.Run()

			if err != nil {
				color.Red("Prompt failed %v\n", err)
				panic(err)
			}
			// fmt.Println("Selected : ", langIndex, langcontroller[langIndex].ProgLang, selectedTopic)

			codeSuggestions := controller.GetTemplateSuggestions(client, langcontroller[langIndex].ProgLang, selectedTopic)
			sort.SliceStable(codeSuggestions, func(i, j int) bool {
				return codeSuggestions[i].Title < codeSuggestions[j].Title
			})
			if len(codeSuggestions) == 0 {
				color.Magenta("No code templates found")
				return nil
			}
			codePrompt := promptui.Select{
				Label: "Select a Code Template",
				Items: Map(codeSuggestions, func(template models.BoilerPlate) string {
					return template.Title
				}),
			}
			_, selectedCodeTemplate, err := codePrompt.Run()
			if err != nil {
				color.Red("Prompt failed %v\n", err)
				panic(err)
			}
			// Perform a search to find the selected template
			var tempIndex = 0
			for idx, lang := range codeSuggestions {
				if lang.Title == selectedCodeTemplate {
					tempIndex = idx
					break
				}
			}
			var template models.BoilerPlate = codeSuggestions[tempIndex]

			handler.GenerateCode(template)

			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}
