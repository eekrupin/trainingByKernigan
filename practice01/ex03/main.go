package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

const imagePattern = `^\.|\.jpg$|\.gif$|.png$`

//Получет список картинок в переданном каталоге
func main() {
	root := os.Args[1:][0]
	files := List(root)
	for _, file := range files {
		fmt.Println(file)
	}
}

func List(dir string) []string {
	var files []string
	filesInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range filesInfo {
		if !fileInfo.IsDir() {
			r, err := regexp.MatchString(imagePattern, fileInfo.Name())
			if err == nil && r {
				files = append(files, fileInfo.Name())
			}
		}
	}
	return files
}
