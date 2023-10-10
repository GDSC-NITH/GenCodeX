// main.go
package main

import (
	options "GenCodeX/controller"
	db "GenCodeX/db"
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

			langOptions := options.GetLangs(client)
			// Sort the language options alphabetically
			sort.SliceStable(langOptions, func(i, j int) bool {
				return langOptions[i].ProgLang < langOptions[j].ProgLang
			})

			// Create a prompt for user selection
			prompt := promptui.Select{
				Label: "Select a language",
				Items: Map(langOptions, func(lang options.Options) string {
					return lang.ProgLang
				}),
			}

			_, selectedLang, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				panic(err)
			}
			fmt.Printf("Selected Language %q\n", selectedLang)

			// Perform a search to find the selected language
			var langIndex = 0
			for idx, lang := range langOptions {
				if lang.ProgLang == selectedLang {
					langIndex = idx
					break
				}
			}
			prompt2 := promptui.Select{
				Label: "Select a Topic",
				Items: langOptions[langIndex].Topics,
			}

			_, selectedTopic, err := prompt2.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				panic(err)
			}
			fmt.Printf("Selected Topic %q\n", selectedTopic)

			fmt.Println("Selected : ", langIndex, langOptions[langIndex].ProgLang, selectedTopic)

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
