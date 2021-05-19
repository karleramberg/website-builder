# website_builder

Simple script for converting a folder of markdown files to a mirrored folder of HTML files
- Converts markdown to html using [this](https://github.com/gomarkdown/markdown) library
- Creates a identical file structure in specified folder
- Inserts new html into a specified template at the specified token

Command line usage:
```
$ ./websiteBuilder <markdown folder> <html folder> <html template> <replace token>
```

Example for my personal site found [here](https://github.com/karlramberg/karlramberg.github.io):
```
$ ./websiteBuilder assets/markdown/ ./ assets/template.html CONTENT
```
