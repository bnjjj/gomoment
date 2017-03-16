package gomoment

import (
	"fmt"
	"testing"
	"time"
)

func TestOrder(test *testing.T) {
	order, regx := getRegx()

	if len(order) != len(regx) {
		test.Error("ERROR !!!! OrderArray and length of map keys must be equals")
	}
}

func TestGetDateOfDay(test *testing.T) {
	moments := []string{" aujourd'hui ?"}

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		now := time.Now()

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != now.Day() || date.Month() != now.Month() || date.Year() != now.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", now.Day(), now.Month(), now.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfBeforeYesterday(test *testing.T) {
	moments := []string{" avant-hier ?", " avant  hier "}

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		beforeYesterday := time.Now().AddDate(0, 0, -2)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != beforeYesterday.Day() || date.Month() != beforeYesterday.Month() || date.Year() != beforeYesterday.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", beforeYesterday.Day(), beforeYesterday.Month(), beforeYesterday.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfYesterday(test *testing.T) {
	moments := []string{" hier ", " la veille "}

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		yesterday := time.Now().AddDate(0, 0, -1)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != yesterday.Day() || date.Month() != yesterday.Month() || date.Year() != yesterday.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", yesterday.Day(), yesterday.Month(), yesterday.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfDays(test *testing.T) {
	moments := []string{" 5 jours ", " 5jours ", "5j", " 5j ?", "5j?"}

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		daysBefore := time.Now().AddDate(0, 0, -5)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfWeeks(test *testing.T) {
	moments := []string{" 2 semaines ", " 2semaines ", "2sem", " 2semaines ?", "2sem?"}

	for _, moment := range moments {
		date, _, err := GetDate(moment, true, nil)
		daysBefore := time.Now().AddDate(0, 0, -14)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfMonths(test *testing.T) {
	moments := []string{" 5 mois ", " 5mois ", "5mois?"}

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		daysBefore := time.Now().AddDate(0, -5, -(time.Now().Day() - 1))

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfCurrentMonth(test *testing.T) {
	moments := []string{" ce mois ", " ce mois-ci "}

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		daysBefore := time.Now().AddDate(0, 0, -(time.Now().Day() - 1))

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfPreviousWeek(test *testing.T) {
	test.Skip()
	moments := []string{" la semaine passée ", " semaine dernière ? "}

	for _, moment := range moments {
		date, end, err := GetDate(moment, true, nil)
		fmt.Println(date)
		fmt.Println(end)
		daysBefore := time.Now().AddDate(0, -5, -(time.Now().Day() - 1))

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfCurrentWeek(test *testing.T) {
	test.Skip()
	moments := []string{" cette semaine ", " semaine-ci "}

	for _, moment := range moments {
		date, end, err := GetDate(moment, false, nil)
		daysBefore := time.Now().AddDate(0, -5, -(time.Now().Day() - 1))
		fmt.Println(date)
		fmt.Println(end)
		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfCurrentYear(test *testing.T) {
	moments := []string{" cette annéee ", " début de l'annee "}

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		monthNumber := int(time.Now().Month())
		daysBefore := time.Now().AddDate(0, -(monthNumber - 1), -(time.Now().Day() - 1))

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfWeeksBis(test *testing.T) {
	moments := []string{" 5 semaines ", " 5semaines ", "5sem", "5sem?"}

	for _, moment := range moments {
		date, _, err := GetDate(moment, true, nil)
		daysBefore := time.Now().AddDate(0, 0, -5*7)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfDate(test *testing.T) {
	moments := []string{" 5 décembre 2015 ", " 5 decembre 2015 "}
	location, _ := time.LoadLocation("Europe/Paris")

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		daysBefore := time.Date(2015, time.Month(12), 5, 0, 0, 0, 0, location)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfDateWithoutYear(test *testing.T) {
	moments := []string{" le 5 janvier ?", "5janvier?", " j'ai fait depuis le 5 janvier ?"}
	location, _ := time.LoadLocation("Europe/Paris")

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		daysBefore := time.Date(2017, time.Month(1), 5, 0, 0, 0, 0, location)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateOfDateWithoutYearBis(test *testing.T) {
	moments := []string{" le 30 janvier ?", "30janvier?", " j'ai fait depuis le 30 janvier ?"}
	location, _ := time.LoadLocation("Europe/Paris")

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, nil)
		daysBefore := time.Date(2017, 1, 30, 0, 0, 0, 0, location)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}

func TestGetDateParsed(test *testing.T) {
	moments := []string{"le 5/01/2017 ?", "05/1/17?", " 5-01 ", "05/01"}
	location, _ := time.LoadLocation("Europe/Paris")

	for _, moment := range moments {
		date, _, err := GetDate(moment, false, location)
		daysBefore := time.Date(2017, time.Month(1), 5, 0, 0, 0, 0, location)

		if err != nil {
			test.Errorf("Got an error for [%s] -> %s", moment, err.Error())
			return
		}

		if date.Day() != daysBefore.Day() || date.Month() != daysBefore.Month() || date.Year() != daysBefore.Year() {
			test.Errorf("Date is not the right date: expected [%d/%d/%d] got [%d/%d/%d]", daysBefore.Day(), daysBefore.Month(), daysBefore.Year(), date.Day(), date.Month(), date.Year())
		}
	}
}
