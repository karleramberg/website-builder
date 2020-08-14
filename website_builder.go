package main

import (
	"github.com/gomarkdown/markdown"
	"os"
	"path/filepath"
	"io/ioutil"
	"fmt"
)

const ROOT string = "pages/"
const HEADER string = "assets/header.html"
const FOOTER string = "assets/footer.html"
func main() {
	// Read the surrounding HTML into bytes
	header,_ := ioutil.ReadFile(HEADER)
	footer,_ := ioutil.ReadFile(FOOTER)

	// Walk through the ROOT directory file by file
	filepath.Walk(ROOT, func(path string, info os.FileInfo, err error) error {
		// Create mirrored folder structure
		if info.IsDir() {
			newFolder := path[len(ROOT):]
			os.MkdirAll(newFolder, 0755)
			fmt.Println("Generating " + newFolder + "...")
			return nil
		}

		// Generate html file from any markdown file
		newPath := path[len(ROOT):len(path)-len(filepath.Ext(path))] + ".html"
		fmt.Println("Generating " + newPath + "...")
		md,_ := ioutil.ReadFile(path)
		html := append(header, markdown.ToHTML(md, nil, nil)[:]...)
		html = append(html, footer[:]...)
		ioutil.WriteFile(newPath, html, 0644)
		return nil
	})
}
