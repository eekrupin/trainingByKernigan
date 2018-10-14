package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

//Параметр URL картинки, картинка сохраняется в текущий путь программы по собственному имени.
func main() {

	//URL := https://cdn.fishki.net/upload/post/2017/03/19/2245758/tn/02-funny-cat-wallpapercat-wallpaper.jpg
	URL := os.Args[1:][0]

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	fileName := path.Base(resp.Request.URL.Path)
	err = ioutil.WriteFile(fileName, body, 0644)
}
