package main

import (
	"testing"
	"time"
)

var pc_babka passportControl
var pc_dedka passportControl

func init() {
	pc_babka = passportControl{personal: female}
	pc_dedka = passportControl{personal: male}
}

func Test_man_emir_alon_to_babka(t *testing.T) {
	man_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	man_p := passport{FirstNAme: "Игорь", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: male, LastName: "Столов", SerialNumber: "ER456853", Country: "Russia"}
	res, err := pc_babka.Check(&man_t, &man_p, nil)
	if err != nil || !res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, true)
	}
}

func Test_man_emir_alon_to_dedka(t *testing.T) {
	man_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	man_p := passport{FirstNAme: "Игорь", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: male, LastName: "Столов", SerialNumber: "ER456853", Country: "Russia"}
	res, err := pc_dedka.Check(&man_t, &man_p, nil)
	if !res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, true)
	}
}

func Test_man_emir_to_dedka_with_safetyAnimal(t *testing.T) {
	man_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	man_p := passport{FirstNAme: "Игорь", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: male, LastName: "Столов", SerialNumber: "ER456853", Country: "Russia"}
	animalDoc := animalDocImpl{docType: animalveterinaryPassport, name: "Гоша", age: 3, safety: true}
	res, err := pc_dedka.Check(&man_t, &man_p, &animalDoc)
	if !res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, true)
	}
}

func Test_man_emir_to_babka_with_safetyAnimal(t *testing.T) {
	man_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	man_p := passport{FirstNAme: "Игорь", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: male, LastName: "Столов", SerialNumber: "ER456853", Country: "Russia"}
	animalDoc := animalDocImpl{docType: animalveterinaryPassport, name: "Гоша", age: 3, safety: true}
	res, err := pc_babka.Check(&man_t, &man_p, &animalDoc)
	if res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, false)
	}
}

func Test_man_emir_to_dedka_with_unsafetyAnimal(t *testing.T) {
	man_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	man_p := passport{FirstNAme: "Игорь", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: male, LastName: "Столов", SerialNumber: "ER456853", Country: "Russia"}
	animalDoc := animalDocImpl{docType: animalveterinaryPassport, name: "Гоша", age: 3, safety: false}
	res, err := pc_dedka.Check(&man_t, &man_p, &animalDoc)
	if res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, false)
	}
}

func Test_man_emir_to_dedka_with_animal_withEmptyDoc(t *testing.T) {
	man_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	man_p := passport{FirstNAme: "Игорь", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: male, LastName: "Столов", SerialNumber: "ER456853", Country: "Russia"}
	animalDoc := animalDocImpl{}
	res, err := pc_dedka.Check(&man_t, &man_p, &animalDoc)
	if res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, false)
	}
}

func Test_check_woman_married_emir_to_dedka(t *testing.T) {
	woman_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	woman_p := passport{FirstNAme: "Анна", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: female, LastName: "Столова", SerialNumber: "ER456853", Country: "Russia"}
	res, err := pc_dedka.Check(&woman_t, &woman_p, nil)
	if err == nil {
		t.Errorf("Result: %b, error: %s. Must be error", res, err)
	}
}

func Test_woman_married_emir_to_dedka(t *testing.T) {
	woman_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	woman_p := passport{FirstNAme: "Анна", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: female, LastName: "Столова", SerialNumber: "ER456853", Country: "Russia"}
	res, err := pc_dedka.CheckLadyToArabian(&woman_t, &woman_p, false, nil)
	if !res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, true)
	}
}

func Test_woman_unmarried_emir_to_dedka(t *testing.T) {
	woman_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	woman_p := passport{FirstNAme: "Анна", Birthday: time.Date(1990, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: female, LastName: "Столова", SerialNumber: "ER456853", Country: "Russia"}
	res, err := pc_dedka.CheckLadyToArabian(&woman_t, &woman_p, true, nil)
	if res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, false)
	}
}

func Test_girl_unmarried_emir_to_dedka(t *testing.T) {
	girl_t := ticket{Departure: "Moscow", Arrival: "Iraq"}
	girl_p := passport{FirstNAme: "Анна", Birthday: time.Date(time.Now().Year()-10, 3, 15, 0, 0, 0, 0, time.UTC),
		Sex: female, LastName: "Столова", SerialNumber: "ER456853", Country: "Russia"}
	res, err := pc_dedka.Check(&girl_t, &girl_p, nil)
	if !res {
		t.Errorf("Result: %b, error: %s. Must %b", res, err, true)
	}
}
