package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func getAllItemsFromCurrentDir() {

	excludeStrings := []string{".DS_Store", ".idea", "dockerfile", "go.mod", "readme.md", ".", ".gitignore", "modules.xml", "tree.iml", "workspace.xml"}
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !slices.Contains(excludeStrings, info.Name()) {
				splitedPath := strings.Split(path, "/")
				lenPath := len(splitedPath) // ex: 2
				result := ""
				for i := 1; i < lenPath; i++ {
					result += "│\t"
				}
				if !info.IsDir() {
					if len(splitedPath) > 1 {
						fmt.Printf("%s└───%s (%d)\n", result, info.Name(), info.Size())
					} else {
						fmt.Printf("├───%s (%d)\n", info.Name(), info.Size())
					}
				} else {
					if len(splitedPath) > 1 {
						fmt.Printf("%s│───%s \n", result, info.Name())
					} else {
						fmt.Printf("└───%s \n", info.Name())
					}

				}

			}
			return nil

		})
	if err != nil {
		log.Println(err)
	}

}
func main() {
	getAllItemsFromCurrentDir() // get all files, dirs and exclude trash files

}
/*
my resolution output
├───main.go (1180)
├───main_test.go (1865)
└───testdata 
│       │───project 
│       │       └───file.txt (19)
│       │       └───gopher.png (70372)
│       │───static 
│       │       │───a_lorem 
│       │       │       └───dolor.txt (0)
│       │       │       └───gopher.png (70372)
│       │       │       │───ipsum 
│       │       │       │       └───gopher.png (70372)
│       │       │───css 
│       │       │       └───body.css (28)
│       │       └───empty.txt (0)
│       │       │───html 
│       │       │       └───index.html (57)
│       │       │───js 
│       │       │       └───site.js (10)
│       │       │───z_lorem 
│       │       │       └───dolor.txt (0)
│       │       │       └───gopher.png (70372)
│       │       │       │───ipsum 
│       │       │       │       └───gopher.png (70372)
│       │───zline 
│       │       └───empty.txt (0)
│       │       │───lorem 
│       │       │       └───dolor.txt (0)
│       │       │       └───gopher.png (70372)
│       │       │       │───ipsum 
│       │       │       │       └───gopher.png (70372)
│       └───zzfile.txt (0)

how it should be look 
go run main.go . -f
├───main.go (1881b)
├───main_test.go (1318b)
└───testdata
	├───project
	│	├───file.txt (19b)
	│	└───gopher.png (70372b)
	├───static
	│	├───css
	│	│	└───body.css (28b)
	│	├───html
	│	│	└───index.html (57b)
	│	└───js
	│		└───site.js (10b)
	├───zline
	│	└───empty.txt (empty)
	└───zzfile.txt (empty)
*/
