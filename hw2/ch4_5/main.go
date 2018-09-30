package main

import (
	"log"
	"strings"
)

func main() {
	str := "some some text with with with with with some text text !"

	text := strings.Fields(str)

	i := 0
	var prev string
	for _, s := range text {
		if s != prev {
			text[i] = s
			i++
		}
		prev = s
	}

	log.Println(text[:i])

}
