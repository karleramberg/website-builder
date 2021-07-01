Simple script for converting a folder of markdown files to a mirrored folder of HTML files
- Converts markdown to html using [this](https://github.com/gomarkdown/markdown) library
- Creates an identical file structure in specified folder
- Leaves other files in html folder alone, such as 'style.css'
- Inserts new html into a specified template at the specified token

## Usage:
```
$ website_builder <markdown folder> <html folder> <html template> <replace token>
```

Example for my personal site found [here](https://github.com/karlramberg/karlramberg.github.io):
```
$ website_builder markdown/ docs/ template.html REPLACE_TOKEN
```
