/*
	Karl Ramberg
	Walker v4.0
	walker.go
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var iFlag = flag.String("i", "src", "Input folder of incomplete HTML pages")
var oFlag = flag.String("o", "docs", "Output folder to place complete HTML pages. This folder is created if it doesn't exist. Existing *.html files, *.htm files, and empty subdirectories in here are deleted!")
var tFlag = flag.String("t", "template.html", "Template HTML page to inject HTML into")

func main() {
	flag.Parse()

	fmt.Println("> Building site...")
	fmt.Println("> Input folder is " + *iFlag)
	fmt.Println("> Output folder is " + *oFlag)
	fmt.Println("> Template file is " + *tFlag)

	// Setup the input folder, checking that it exists
	_, err := os.Stat(*iFlag)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("> ERROR: Input folder " + *iFlag + " not found")
		return
	}

	// Read template into a string, checking that it exists
	template, err := os.ReadFile(*tFlag)
	if err != nil {
		fmt.Println("> ERROR: Template not found or is protected")
		return
	}

	// Setup the output folder, checking that it exists, create it if user wants
	_, err = os.Stat(*oFlag)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("> Generating " + *oFlag + "...")
		os.MkdirAll(*oFlag, 0755)

	} else { // Clean-up the output folder
		fmt.Println("> Cleaning " + *oFlag + "...")

		// Delete old HTML files
		filepath.Walk(*oFlag, func(path string, info os.FileInfo, err error) error {
			extension := filepath.Ext(path)
			if extension == ".html" || extension == ".htm" {
				os.Remove(path)
			}
			return nil
		})

		// Remove any empty folders
		filepath.Walk(*oFlag, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && path != *oFlag {
				contents, _ := os.ReadDir(path)
				if len(contents) == 0 {
					os.Remove(path)
				}
			}
			return nil
		})
	}

	// Walk through the input folder file by file
	filepath.Walk(*iFlag, func(path string, info os.FileInfo, err error) error {
		// If a folder is found, mirror it in the output folder
		if info.IsDir() {
			newFolder := *oFlag + path[len(*iFlag):]
			if newFolder != *oFlag {
				os.MkdirAll(newFolder, 0755)
				fmt.Println("> Generating " + newFolder + "...")
			}
		}

		// If an HTML file is found, generate a completed one
		// and write it to the same location in the output folder
		extension := filepath.Ext(path)
		if extension == ".html" || extension == ".htm" {
			// Read in contents of input file
			input, err := os.ReadFile(path)
			if err != nil {
				fmt.Println("> ERROR: Cannot access " + path)
				return nil
			}

			// Grab the title from the first line and the rest is the content
			title, content, _ := strings.Cut(string(input), "\n")

			// Replace [TITLE] with page title, [CONTENT] with the content
			output := strings.Replace(string(template), "[TITLE]", string(title), 1)
			output = strings.Replace(output, "[CONTENT]", string(content), 1)

			// Genearate the output file's path and write to disk
			newFile := *oFlag + path[len(*iFlag):]
			fmt.Println("> Generating " + newFile + "...")
			os.WriteFile(newFile, []byte(output), 0644)
		}

		return nil
	})

	fmt.Println("> Site built successfully! :)")
}
