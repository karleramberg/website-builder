/*
	Karl Ramberg
	Website Builder v4.0
	websiteBuilder.go
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
	// Check that all 4 arguments were given
	if len(os.Args) != 4 {
		fmt.Println("USAGE: $ website_builder <input folder> <output folder> <template>")
		return
	}

	// Setup the input folder, checking that it exists
	inputFolder := strings.TrimLeft(os.Args[1], "./\\")
	_, err := os.Stat(inputFolder)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("ERROR: Input folder " + inputFolder + " not found")
		return
	}

	// Read template into a string, checking that it exists
	template, err := ioutil.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println("ERROR: Template not found or is protected")
		return
	}

	// Setup the output folder, checking that it exists, create it if it does not
	outputFolder := strings.TrimLeft(os.Args[2], "./\\")
	_, err = os.Stat(outputFolder)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("Generating " + outputFolder + "...")
		os.MkdirAll(outputFolder, 0755)
	} else { // Clean-up the output folder
		fmt.Println("Cleaning " + outputFolder + "...")

		// Delete old HTML files
		filepath.Walk(outputFolder, func(path string, info os.FileInfo, err error) error {
			extension := filepath.Ext(path)
			if extension == ".html" || extension == ".htm" {
				os.Remove(path)
			}
			return nil
		})

		// Remove any empty folders
		filepath.Walk(outputFolder, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && path != outputFolder {
				contents, _ := ioutil.ReadDir(path)
				if len(contents) == 0 {
					os.Remove(path)
				}
			}
			return nil
		})
	}

	// Walk through the input folder file by file
	filepath.Walk(inputFolder, func(path string, info os.FileInfo, err error) error {
		// If a folder is found, mirror it in the output folder
		if info.IsDir() {
			newFolder := outputFolder + path[len(inputFolder):]
			if newFolder != outputFolder {
				os.MkdirAll(newFolder, 0755)
				fmt.Println("Generating " + newFolder + "...")
			}
		}

		// If an HTML file is found, generate a completed one
		// and write it to the same location in the output folder
		extension := filepath.Ext(path)
		if extension == ".html" || extension == ".htm" {
			// Read in contents of input file
			input, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Println("ERROR: Cannot access " + path)
				return nil
			}

			// Grab the title from the first line and the rest is the content
			title, content, _ := strings.Cut(string(input), "\n")

			// Replace [TITLE] with page title, [CONTENT] with the content
			output := strings.Replace(string(template), "[TITLE]", string(title), 1)
			output = strings.Replace(output, "[CONTENT]", string(content), 1)

			// Genearate the output file's path and write to disk
			newFile := outputFolder + path[len(inputFolder):]
			fmt.Println("Generating " + newFile + "...")
			ioutil.WriteFile(newFile, []byte(output), 0644)
		}

		return nil
	})

	fmt.Println("Done.")
}
