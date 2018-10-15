package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"io/ioutil"
	"os"
	"time"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "Sa#1234", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "krupin", "the database server")
	user          = flag.String("user", "sa", "the database user")
)

type rowData struct {
	Name string    `json:ame`
	Age  int       `json:age`
	Date time.Time `json:date`
}

type data struct {
	Items []rowData
}

//парсим json, складываем в БД
func main() {

	//.\practice01\ex01\users.json
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

	for _, item := range data.Items {
		dateString := item.Date.Format(time.RFC3339)
		_, err := conn.Exec("INSERT INTO training.dbo.usersPractice01_ex01 (name, age, date) VALUES (?, ?, ?)", item.Name, item.Age, dateString)
		check(err)
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
