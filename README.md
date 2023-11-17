# Website Builder

Website Builder is an application to build complete HTML pages in an output folder using a template and an input folder of incomplete HTML pages.

The first line of an input file is used as a page title, replacing the [TITLE] token in the template. The rest of the file replaces the [CONTENT] token.

## Install
Executable files are located in the releases section. Download the .exe file if you are on Windows and the un-extended file
if you are on Linux or macOS. Run as you would a normal comamnd in your terminal. You may have to give the file
executable permissions (Using `chmod +x` on Linux/macOS or the "Run as Administrator" function on Windows).
  
If neither file runs on your system, you can build a new one using Go. Install the Go langauge on your system from
https://go.dev/. Clone this repository and then run `go build websiteBuilder.go` in your terminal to make a new file
for your system's architecture.

## CLI Usage
  ```$ wb input_folder/ output_folder/ template.html```

### Example website(s)
  https://github.com/karlramberg/karlramberg.com
