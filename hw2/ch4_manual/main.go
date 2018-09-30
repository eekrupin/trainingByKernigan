package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int `json:"Disk"`
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	//Marshal/Unmarshal
	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)

	data, err := json.Marshal(w)
	if err != nil {
		log.Fatal("%s\n", err)
	}
	fmt.Printf("%s\n\n", data)

	err = ioutil.WriteFile("Wheel.txt", data, 0644)
	check(err)

	circle := Circle{}
	newData, err := ioutil.ReadFile("Wheel.txt")
	err = json.Unmarshal(newData, &circle)
	check(err)
	fmt.Printf("from file: %#v\n", circle)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
