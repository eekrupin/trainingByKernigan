package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "", "the database server")
	user          = flag.String("user", "", "the database user")
)

type rowData struct {
	name string
	age  int
	date time.Time
}

type data struct {
	item []rowData
}

func main() {

	fileName := os.Args[1:][0]

	data := data{}
	byteData, err := ioutil.ReadFile(fileName)
	err = json.Unmarshal(byteData, &data)
	check(err)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	check(err)
	defer conn.Close()

	for _, item := range data.item {
		_, err := conn.Exec("INSERT INTO users (name, age, date) VALUES (@1, @2, @3)", item.name, item.age, item.date)
		check(err)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
