package gomoment

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

// Return the regex map
func getRegx() ([]string, map[string]string) {
	order := []string{"date", "today", "before-yesterday", "yesterday", "day", "week", "previousWeek", "currentWeek", "month", "currentMonth", "currentYear", "dateWithYear", "dateWithoutYear"}

	return order, map[string]string{
		"date":             `(?i)([0-9]{1,2})\s*[/-]([0-9]{1,2})\s*[/-]?([0-9]{2,4})?`,
		"today":            `(?i)aujourd'?hui|ce\smoment|ce\sjour|dans\sla\sjourn(é|e|è)e?`,
		"before-yesterday": `(?i)avant-hier|avant\s*hier`,
		"yesterday":        `(?i)hier|veille`,
		"day":              `(?i)([0-9]+)\s*(?:jour|\s?j(?:\s+|$|\?))`,
		"week":             `(?i)([0-9]+)\s*(?:semaine|sem(?:\s+|$|\?))`,
		"previousWeek":     `(?i)\s*semaine.*(derni(e|è|é)re|pas{1,3}(e|è|é))`,
		"currentWeek":      `(?i)\s*semaine`,
		"month":            `(?i)([0-9]+)\s*mois?`,
		"currentMonth":     `(?i)\s*mois?`,
		"currentYear":      `(?i)\s*an{1,3}(é|e)e?`,
		"dateWithYear":     `(?i)(?:\s|[a-zA-Z])*([0-9]+)\s*(?:.*)\s+(?:([0-9]{4})|([0-9]{4}\s*))`,
		"dateWithoutYear":  `(?i).*(?:\s|^)([0-9]+)\s*(?:.*)\s*|.*`,
	}
}

// Find the moment type
func getMoment(moment string) (string, error) {
	order, regx := getRegx()
	var momentFound string

	for _, momentName := range order {
		myRegex, _ := regexp.Compile(regx[momentName])

		if myRegex.MatchString(moment) {
			momentFound = momentName
			break
		}
	}

	if momentFound == "" {
		return momentFound, errors.New("moment not found")
	}

	return momentFound, nil
}

// Return the date or duration
// Example for a date:
//	begin, _, err := GetDate("Donne moi la date d'aujourd'hui", false, nil)
//	// begin is the date of today 0h0min
// Other examples for duration
//	begin, end, err := GetDate("Combien de km j'ai réalisé depuis le mois dernier ?", true, nil)
//	// begin is today - 1 month and end is today 00:00
// Example with another location
//	location, _ := time.LoadLocation("America/New_York")
//	begin, end, err := GetDate("Combien de km j'ai réalisé hier ?", true, location)
func GetDate(moment string, duration bool, location *time.Location) (time.Time, time.Time, error) {
	var locationTime *time.Location
	var err error
	momentName, err := getMoment(moment)
	now := time.Now()

	if location == nil {
		locationTime, err = time.LoadLocation("Europe/Paris")

		if err != nil {
			return time.Time{}, time.Time{}, err
		}
	} else {
		locationTime = location
	}

	now = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, locationTime)

	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	switch momentName {
	case "date":
		return createDate(momentName, moment, locationTime)
	case "today":
		return now, time.Time{}, nil
	case "before-yesterday":
		return now.AddDate(0, 0, -2), now.AddDate(0, 0, -1), nil
	case "yesterday":
		return now.AddDate(0, 0, -1), now, nil
	case "day":
		return subtractDate(momentName, moment, duration)
	case "week":
		return subtractDate(momentName, moment, duration)
	case "previousWeek":
		var number time.Weekday
		previousWeek := now.AddDate(0, 0, -7)
		if previousWeek.Weekday() == 0 {
			number = 6
		} else {
			number = (previousWeek.Weekday() - 1)
		}

		numberInt := int(number)
		previousWeek = previousWeek.AddDate(0, 0, -numberInt)

		if !duration {
			return previousWeek, previousWeek.AddDate(0, 0, 7), nil
		} else {
			return previousWeek, now, nil
		}
	case "currentWeek":
		var number time.Weekday
		if now.Weekday() == 0 {
			number = 6
		} else {
			number = (now.Weekday() - 1)
		}
		numberInt := int(number)
		return now.AddDate(0, 0, -numberInt), time.Time{}, nil
	case "month":
		return subtractDate(momentName, moment, duration)
	case "currentMonth":
		return now.AddDate(0, 0, -(now.Day() - 1)), time.Time{}, nil
	case "currentYear":
		monthNumber := int(now.Month())
		return now.AddDate(0, -(monthNumber - 1), -(now.Day() - 1)), time.Time{}, nil
	case "dateWithYear", "dateWithoutYear":
		return createDateWithText(momentName, moment, locationTime)
	default:
		return time.Time{}, time.Time{}, errors.New("date not found")
	}
}

func createDate(momentName string, moment string, location *time.Location) (time.Time, time.Time, error) {
	var date time.Time
	now := time.Now()
	_, regx := getRegx()
	myRegx, _ := regexp.Compile(regx[momentName])
	submatch := myRegx.FindStringSubmatch(moment)

	if len(submatch) < 4 {
		return time.Time{}, time.Time{}, errors.New("Date introuvable")
	}

	day, _ := strconv.Atoi(submatch[1])
	month, _ := strconv.Atoi(submatch[2])

	if submatch[3] == "" {
		date = time.Date(now.Year(), time.Month(month), day, 0, 0, 0, 0, location)
		return date, date.AddDate(0, 0, 1), nil
	} else {
		year := 0
		if len(submatch[3]) == 2 {
			year, _ = strconv.Atoi(submatch[3])
			year += 2000
		} else {
			year, _ = strconv.Atoi(submatch[3])
		}
		date = time.Date(year, time.Month(month), day, 0, 0, 0, 0, location)

		return date, date.AddDate(0, 0, 1), nil
	}
}

// Create date from text with day, month and year into text
func createDateWithText(momentName string, moment string, location *time.Location) (time.Time, time.Time, error) {
	var date time.Time
	now := time.Now()
	_, regx := getRegx()
	monthId, err := getMonthId(moment)
	myRegx, _ := regexp.Compile(regx[momentName])
	submatch := myRegx.FindStringSubmatch(moment)

	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	if len(submatch) <= 1 || submatch[1] == "" {
		return time.Time{}, time.Time{}, errors.New("Date incorrecte")
	}
	day, _ := strconv.Atoi(submatch[1])

	if len(submatch) <= 2 || submatch[2] == "" {
		date = time.Date(now.Year(), time.Month(monthId+1), day, 0, 0, 0, 0, location)

		return date, date.AddDate(0, 0, 1), nil
	}
	year, _ := strconv.Atoi(submatch[2])

	date = time.Date(year, time.Month(monthId+1), day, 0, 0, 0, 0, location)

	return date, date.AddDate(0, 0, 1), nil
}

// Get month id from text
func getMonthId(moment string) (int, error) {
	months := []string{`(?i)janvier`, `(?i)f(?:é|e|è)vrier`, `(?i)mars`, `(?i)avril`, `(?i)mais?(?:\s|$)`, `(?i)juin`, `(?i)juillet`, `(?i)ao(?:û|u|ù)t`, `(?i)septembre`, `(?i)octobre`, `(?i)novembre`, `(?i)d(?:é|e|è)cembre`}
	id := -1

	for monthId, monthRx := range months {
		myRegx, _ := regexp.Compile(monthRx)

		if myRegx.MatchString(moment) {
			id = monthId
			break
		}
	}
	if id == -1 {
		return -1, errors.New("Mois introuvable")
	}

	return id, nil
}

// Subtract good time for specific moment
func subtractDate(momentName string, moment string, duration bool) (time.Time, time.Time, error) {
	now := time.Now()

	switch momentName {
	case "day":
		number, err := getNumber(momentName, moment)
		return now.AddDate(0, 0, -number), now.AddDate(0, 0, -(number - 1)), err
	case "month":
		number, err := getNumber(momentName, moment)
		if duration {
			return now.AddDate(0, -number, 0), time.Time{}, err
		}
		return now.AddDate(0, -number, -(now.Day() - 1)), time.Time{}, err
	case "week":
		number, err := getNumber(momentName, moment)
		if duration {
			return now.AddDate(0, 0, -(number * 7)), now.AddDate(0, 0, -((number - 1) * 7)), err
		}

		previousWeek := now.AddDate(0, 0, -(number * 7))
		var weekday time.Weekday
		if previousWeek.Weekday() == 0 {
			weekday = 6
		} else {
			weekday = (now.Weekday() - 1)
		}
		numberInt := int(weekday)

		return previousWeek.AddDate(0, 0, -numberInt), previousWeek.AddDate(0, 0, 7-numberInt), err
	default:
		return time.Time{}, time.Time{}, errors.New("cannot subtract date")
	}
}

// Get number from moment
func getNumber(momentName string, moment string) (int, error) {
	_, regx := getRegx()
	myRegx, _ := regexp.Compile(regx[momentName])
	submatch := myRegx.FindStringSubmatch(moment)

	if len(submatch) <= 1 {
		return -1, errors.New("number not found")
	}
	number, err := strconv.Atoi(submatch[1])

	return number, err
}
