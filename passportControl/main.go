package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	male   sex = 1
	female sex = 2

	animalveterinaryPassport animalDocType = 1
	animalOwnershipAct       animalDocType = 2
	animalSafetyСertificate  animalDocType = 3
)

var db_ext_reg map[int]map[string]interface{}
var codeCounter int32

func init() {
	db_ext_reg = make(map[int]map[string]interface{})
}

type passportControl struct {
	personal sex
}

type country string

type passport struct {
	LastName,
	FirstNAme,
	SerialNumber string
	Sex      sex
	Birthday time.Time
	Country  string
}

type ticket struct {
	Departure country `json:"Departure"`
	Arrival   country `json:"Arrival"`
}

type extData struct {
	ticket   `json:"ticket"`
	passport `json:"passport"`
}

type sex int

type animalDocType int

func (p *passport) isAdult() bool {
	return p.Birthday.AddDate(16, 0, 0).Before(time.Now())
}

type animalDoc interface {
	getDocType() animalDocType
	getAge() int
	getName() string
	isSafety() bool
}

type animalDocImpl struct {
	docType animalDocType
	age     int
	name    string
	safety  bool
}

func (d *animalDocImpl) getDocType() animalDocType {
	return d.docType
}

func (d *animalDocImpl) getAge() int {
	return d.age
}

func (d *animalDocImpl) getName() string {
	return d.name
}

func (d *animalDocImpl) isSafety() bool {
	return d.safety
}

func (country country) isArabian() bool {
	arabianCountry := getArabianCountry()
	return arabianCountry[country] == true
}

func getArabianCountry() map[country]bool {
	arabianCountry := make(map[country]bool)
	arabianCountry["Saudi Arabia"] = true
	arabianCountry["United Arab Emirates"] = true
	arabianCountry["Egypt"] = true
	arabianCountry["Iraq"] = true
	arabianCountry["Algeria"] = true
	arabianCountry["Qatar"] = true
	arabianCountry["Kuwait"] = true
	arabianCountry["Morocco"] = true
	arabianCountry["Syria"] = true
	arabianCountry["Oman"] = true
	arabianCountry["Sudan"] = true
	arabianCountry["Lebanon"] = true
	arabianCountry["Jordan"] = true
	arabianCountry["Tunisia"] = true
	arabianCountry["Bahrain"] = true
	arabianCountry["Libya"] = true
	arabianCountry["Yemen"] = true
	arabianCountry["Palestine"] = true
	arabianCountry["Somalia"] = true
	arabianCountry["Mauritania"] = true
	arabianCountry["Djibouti"] = true
	arabianCountry["Comoros"] = true
	return arabianCountry
}

func (pc *passportControl) Check(t *ticket, p *passport, animalDoc animalDoc) (bool, error) {
	if p.Sex == female && p.isAdult() && t.Arrival.isArabian() {
		return false, errors.New("Use check lady to Arabian")
	}
	return pc.checkWithoutCheckArabian(t, p, animalDoc)
}

func (pc *passportControl) CheckLadyToArabian(t *ticket, p *passport, unmarried bool, animalDoc animalDoc) (bool, error) {
	if unmarried {
		return false, nil
	}
	return pc.checkWithoutCheckArabian(t, p, animalDoc)
}
func (pc *passportControl) checkWithoutCheckArabian(t *ticket, p *passport, animalDoc animalDoc) (bool, error) {
	if t == nil || p == nil {
		return false, nil
	}

	if pc.personal == female && animalDoc != nil {
		return false, errors.New("Use dedka passport controller. Babka can't check animal")
	}

	if animalDoc != nil && !animalDoc.isSafety() {
		return false, nil
	}
	return true, nil
}

func getPersonalDataByCode(code int) (*ticket, *passport, error) {
	mapData, ok := db_ext_reg[code]
	if !ok {
		return nil, nil, errors.New("No data")
	}
	return mapData["ticket"].(*ticket), mapData["passport"].(*passport), nil
}

func setPersonalData(ticket *ticket, passport *passport) (code int) {
	if ticket == nil || passport == nil {
		return 0
	}

	code = int(atomic.AddInt32(&codeCounter, 1))
	data := make(map[string]interface{})
	data["ticket"] = ticket
	data["passport"] = passport
	db_ext_reg[code] = data
	return
}

func (pc *passportControl) AutoCheck(code int) (bool, error) {
	t, p, err := getPersonalDataByCode(code)

	if err == nil {
		return false, errors.New("Can't get data by code")
	}

	if t == nil || p == nil {
		return false, errors.New("Data by code is broken")
	}

	return pc.Check(t, p, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var extData extData

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	if err := r.Body.Close(); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	if err := json.Unmarshal(body, &extData); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	code := setPersonalData(&extData.ticket, &extData.passport)

	fmt.Fprintf(w, "Your code registration is: %d", code)
}

func main() {
	http.HandleFunc("/reg", handler)
	//log.Fatal(http.ListenAndServe(":8080", nil))
	go http.ListenAndServe(":8080", nil)

	testServer()

	sec := time.Duration(100)
	log.Printf("You have %d sec to test API: http://localhost:8080/reg", sec)
	time.Sleep(time.Second * sec)

}
func testServer() {
	man_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	man_p := passport{FirstNAme: "Игорь", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: male, LastName: "Столов", SerialNumber: "ER456853", Country: "Russia"}

	extData := extData{ticket: man_t, passport: man_p}

	data, err := json.Marshal(extData)
	if err != nil {
		log.Fatal("%s\n", err)
	}

	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/reg", "application/json", r)
	if err != nil {
		log.Fatal("%s\n", err)
	}

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		log.Fatal("Error: %v", err)
	}
	if err := resp.Body.Close(); err != nil {
		log.Fatal("Error: %v", err)
	}
	code := string(body)
	log.Printf("Test registration. Answer from server '%s'", code)
	t, p, err := getPersonalDataByCode(1)
	if err != nil {
		log.Fatal("Error: can't get data by code 1", err)
	}
	log.Printf("Take data by code 1:\n\tTicket: %v\n\tPassport: %v\n", t, p)
}
