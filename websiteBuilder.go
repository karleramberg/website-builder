/*
   Website Builder
   Author: Karl Ramberg
   Last modified: 2021-07-11
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Exit if 5 arguments are not given
	if len(os.Args) != 5 {
		fmt.Println("USAGE: $ website_builder <input folder> <output folder> <template> <token>")
		return
	}

	inputFolder := strings.TrimLeft(os.Args[1], "./\\")
	outputFolder := strings.TrimLeft(os.Args[2], "./\\")

	// Read template into a string
	template, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println("ERROR: Template not found or is protected")
		return
	}

	token := os.Args[4]

	// Walk through the output folder, deleting any html files
	fmt.Println("Clearing " + outputFolder + " of HTML files...")
	filepath.Walk(outputFolder, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".html" {
			os.Remove(path)
		}
		return nil
	})

	// Remove any empty folders in the output folder
	fmt.Println("Clearing " + outputFolder + " of empty folders...")
	filepath.Walk(outputFolder, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			contents, _ := ioutil.ReadDir(path)
			if len(contents) == 0 {
				os.Remove(path)
			}
		}
		return nil
	})

	// Walk through the input folder file by file
	filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {
		// If a folder is found, mirror it in the output folder
		if info.IsDir() {
			newFolder := outputFolder + path[len(inputFolder):]
			os.MkdirAll(newFolder, 0755)
			fmt.Println("Generating " + newFolder + "...")
		}

		// If an HTML file is found, generate a completed one
		// and write it to the same location in the output folder
		if filepath.Ext(path) == ".html" {
			// Read in contents of input file
			content, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("ERROR: Cannot access " + path)
				return nil
			}

			// Replace the token with content
			output := strings.Replace(string(template), token, string(content), 1)

			// Genearate the output file's path and write to disk
			newFile := outputFolder + path[len(inputFolder):]
			fmt.Println("Generating " + newFile + "...")
			ioutil.WriteFile(newFile, []byte(output), 0644)
		}

		return nil
	})
}
