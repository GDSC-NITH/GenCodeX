package handler

import (
	models "GenCodeX/models"
	"log"
	"os"
	"sync"

	"github.com/adhocore/chin"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

//	func check(e error) {
//		if e != nil {
//			panic(e)
//		}
//	}
func GenerateCode(template models.BoilerPlate) {
	var wg sync.WaitGroup
	s := chin.New().WithWait(&wg)

	color.Blue("\n \n ⚡⚡ Generating code for %s  ⚡⚡ \n", template.Title)
	go s.Start()
	for _, file := range template.Files {
		iterateFile(file)
	}
	s.Stop()
	wg.Wait()
	color.Green("\n  \n ⚡⚡ Code generated successfully  ⚡⚡ \n\n ")

}

func iterateFile(file models.File) bool {
	dirPath := file.PathToGo
	fileName := file.Name + "." + file.Extension

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatalf("\n Error creating directory %s: %v", dirPath, err)
			return false
		}
	}

	// Check if the file exists
	if _, err := os.Stat(dirPath + "/" + fileName); err == nil {
		// File already exists
		color.Yellow("\n File %s already exists", dirPath)

		prompt := promptui.Prompt{
			Label:     "Do you want to overwrite it?",
			IsConfirm: true,
		}

		result, err := prompt.Run()

		if err != nil {
			color.Magenta("Prompt failed %v\n", err)
			return false
		}

		if result == "y" || result == "Y" {
			// User wants to overwrite
			err := os.WriteFile(dirPath+"/"+fileName, []byte(file.Content), 0644)
			if err != nil {
				log.Fatal(err)
			}
			color.Green("File %s Overwritten", dirPath)
		} else {
			color.Blue("File %s not overwritten", dirPath)
		}
	} else if os.IsNotExist(err) {
		// File does not exist, create it
		err := os.WriteFile(dirPath+"/"+fileName, []byte(file.Content), 0644)
		if err != nil {
			log.Fatal(err)
			return false
		}

		color.Green("File %s Created", dirPath)
	} else {
		// Other error while checking path
		log.Fatalf("Error while checking path %s", err)
		return false
	}

	return true
}
