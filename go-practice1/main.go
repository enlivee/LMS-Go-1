package main

import "fmt"
import "time"
import "math"

func main(){
	solve1()
}

func solve1() {
    var dateStr string
	fmt.Scan(&dateStr)

	layout := "02.01.2006"
	date, _ := time.Parse(layout, dateStr)

	new_date := date.AddDate(0, 0, 15)
	new_date_str := new_date.Format(layout)

	var name, surname, patronymic string
	fmt.Scan(&name)
	fmt.Scan(&surname)
	fmt.Scan(&patronymic)

	var number1, number2, number3 float64
	fmt.Scan(&number1)
	fmt.Scan(&number2)
	fmt.Scan(&number3)

	total := number1 + number2 + number3
	rubles := math.Floor(total)
	kopecks := math.Round((total - rubles) * 100)

	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\nДата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %d руб. %d коп. \n\nС уважением,\nГл. бух. Иванов А.Е.", surname, name, patronymic, new_date_str, int(rubles), int(kopecks))
}