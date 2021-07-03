Script for building a static site
- Grabs HTML files from the input folder
- Sticks their content into a template file at the specified token
- Spits the finished files into the output folder

### Usage
```
$ website_builder <input folder> <output folder> <template file> <replace token>
```

Example for my personal site found [here](https://github.com/karlramberg/karlramberg.github.io):
```
$ website_builder site/ docs/ template.html REPLACE_TOKEN
```
