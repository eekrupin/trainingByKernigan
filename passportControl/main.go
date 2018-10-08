package main

import (
	"github.com/pkg/errors"
	"time"
)

type passportControl struct {
	personal sex
}

type country string

type ticket struct {
	departure,
	arrival country
}

type sex int

type animalDocType int

const (
	male   sex = 1
	female sex = 2

	animalveterinaryPassport animalDocType = 1
	animalOwnershipAct       animalDocType = 2
	animalOSafetyСertificate animalDocType = 3
)

func (p *passport) isAdult() bool {
	return p.birthday.AddDate(16, 0, 0).Before(time.Now())
}

type passport struct {
	lastName,
	firstNAme,
	serialNumber string
	sex      sex
	birthday time.Time
	country  string
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
	//кря... кря... обратился к DB
	return &ticket{departure: "Moscow", arrival: "Moscow"}, &passport{firstNAme: "user"}, nil
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

func main() {

}
