package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "Sa#1234", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "krupin", "the database server")
	user          = flag.String("user", "sa", "the database user")
)

func main() {
	URL := "https://cdn.fishki.net/upload/post/2017/03/19/2245758/tn/02-funny-cat-wallpapercat-wallpaper.jpg"
	id := getAndSaveFile(URL)
	log.Printf("file id: %d", id)

	data := getFile(id)
	log.Printf("file name: %s", data.fileName)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//Параметр URL картинки, картинка сохраняется в БД.
func getAndSaveFile(URL string) int {
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fileName := path.Base(resp.Request.URL.Path)
	id := saveToDB(fileName, body)
	return id
}

func saveToDB(fileName string, body []byte) int {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	check(err)
	defer conn.Close()

	query := "INSERT INTO training.dbo.filesHW04 (fileName, body) VALUES (?, ?); select scope_identity()"
	rows, err := conn.Query(query, fileName, body)
	check(err)
	rows.Next()
	var id int
	err = rows.Scan(&id)
	check(err)
	return id
}

func getFile(id int) struct {
	fileName string
	body     []byte
} {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	check(err)
	defer conn.Close()

	query := "SELECT fileName, body from training.dbo.filesHW04 where id = ?"
	rows, err := conn.Query(query, id)
	check(err)
	rows.Next()
	var fileName string
	var body []byte
	err = rows.Scan(&fileName, &body)
	check(err)

	data := struct {
		fileName string
		body     []byte
	}{
		fileName,
		body}
	return data
}
