package main

import (
	"os"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	// Exit if 4 arguments are not given
	if len(os.Args) < 5 {
		fmt.Println("USAGE: $ website_builder <input folder> <output folder> <template file> <replace token>")
		return
	}

	inputFolder := os.Args[1]
	outputFolder := os.Args[2]

	// Read the HTML template file into a string
	template, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println("ERROR: Template file not found or is protected")
		return	
	}

	replaceToken := os.Args[4]
	
	// Walk through the root directory file by file
	filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {
		// Create mirrored folder structure
		if info.IsDir() {
			newFolder := outputFolder + path[len(inputFolder):]
			os.MkdirAll(newFolder, 0755)
			fmt.Println("Generating " + newFolder + "...")
			return nil
		}

		// Generate new html file from any html file
		newPath := outputFolder + path[len(inputFolder):len(path)-len(filepath.Ext(path))] + ".html"
		fmt.Println("Generating " + newPath + "...")
		content,_ := ioutil.ReadFile(path)

		// Replace the content token with generated HTML
		output := strings.Replace(string(template), replaceToken, string(content), 1)	
		ioutil.WriteFile(newPath, []byte(output), 0644)
		return nil
	})
}
