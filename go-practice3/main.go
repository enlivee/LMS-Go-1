package main

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	dayOrNightError = errors.New("исправь свой ответ, а лучше ложись поспать")
)

// func TimeNow() time.Time {
// 	t := time.Now()
// 	return t
// }

func currentDayOfTheWeek() string {
	t := TimeNow()
	weekDay := t.Weekday()
	switch weekDay {
	case time.Sunday:
		return "Воскресенье"
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	default:
		return "Суббота"
	}
}

func dayOrNight() string {
	t := TimeNow()
	hour := t.Hour()
	if hour >= 10 && hour <= 22 {
		return "День"
	} 
	return "Ночь"
}

func nextFriday() int {
	today := int(TimeNow().Weekday())
	target := int(time.Friday)
	diff := (7 + target - today) % 7
	return diff
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	weekDay := currentDayOfTheWeek()
	if strings.ToLower(answer) == strings.ToLower(weekDay) {
		return true
	}
	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	day := dayOrNight()
	n := utf8.RuneCountInString(answer)
	if n != 4 {
		err := dayOrNightError
		return false, err
	} else {
		if strings.ToLower(day) == strings.ToLower(answer) {
			return true, nil 
		}
	}
	return false, nil
}