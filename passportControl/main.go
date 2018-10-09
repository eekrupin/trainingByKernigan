package main

import (
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
	animalOSafetyСertificate animalDocType = 3
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
	lastName,
	firstNAme,
	serialNumber string
	sex      sex
	birthday time.Time
	country  string
}

type ticket struct {
	departure,
	arrival country
}

type extData struct {
	ticket
	passport
}

type sex int

type animalDocType int

func (p *passport) isAdult() bool {
	return p.birthday.AddDate(16, 0, 0).Before(time.Now())
}

type animalDoc interface {
	getDocType() animalDocType
	getAge() int
	getName() string
	isSafety() bool
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
	if p.sex == female && p.isAdult() && t.arrival.isArabian() {
		return false, errors.New("Use check lady to Arabian")
	}
	return pc.checkWithoutCheckArabian(t, p, animalDoc)
}

func (pc *passportControl) CheckLadyToArabian(t *ticket, p *passport, unmarried bool, animalDoc animalDoc) (bool, error) {
	if !unmarried {
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
	//time.Sleep(time.Second*5)
	//fmt.Println("dd")

	var res bool
	var err error

	pc_babka := passportControl{personal: female}

	man_t := ticket{departure: "Moscow", arrival: "Iraq"}
	man_p := passport{firstNAme: "Игорь", birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC), sex: male, lastName: "Столов", serialNumber: "ER456853", country: "Russia"}
	res, err = pc_babka.Check(&man_t, &man_p, nil)
	log.Printf("Result: %b, error: %s", res, err)

	girl_t := ticket{departure: "Moscow", arrival: "Iraq"}
	girl_p := passport{firstNAme: "Анна", birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC), sex: female, lastName: "Столова", serialNumber: "ER456853", country: "Russia"}
	res, err = pc_babka.Check(&girl_t, &girl_p, nil)
	log.Printf("Result: %b, error: %s", res, err)

	res, err = pc_babka.CheckLadyToArabian(&girl_t, &girl_p, true, nil)
	log.Printf("Result: %b, error: %s", res, err)

	sec := time.Duration(100)
	log.Printf("You have %d sec to test API: http://localhost:8080/reg", sec)
	time.Sleep(time.Second * sec)

}
