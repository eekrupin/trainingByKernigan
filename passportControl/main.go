package main

import "time"

type passportControl struct {
}

type ticket struct {
}

type sex int

const (
	male   = 1
	female = 2
)

type passport struct {
	lastName,
	firstNAme,
	serialNumber string
	sex      sex
	birthday time.Time
	country  string
}

func isArabian(country string) bool {
	arabianCountry := getArabianCountry()
	return arabianCountry[country] == true
}

func getArabianCountry() map[string]bool {
	arabianCountry := make(map[string]bool)
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

func (pc *passportControl) check(t *ticket, p *passport, checkLadyToArabian bool) bool {
	return false
}

func (pc *passportControl) checkLadyToArabian(t *ticket, p *passport, unmarried bool) bool {
	if !unmarried {
		return false
	}
	return pc.check(t, p, false)
}

func main() {

}
