package main

import (
	"fmt"
)

func main() {
	var p1 string = "-"
	var p2 string = "-"
	var p3 string = "-"
	var p4 string = "-"
	var p5 string = "-"
	
	var x string
	var count int
	var y int

	for {

		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}

		if x == "количество" {
			fmt.Printf("Осталось свободных мест: %d\n", 5-count)
			fmt.Printf("Всего человек в очереди: %d\n", count)
		} else if x == "конец" {
			fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", p1, p2, p3, p4, p5)
			break
		} else if x == "очередь" {
			fmt.Printf("1. %s\n2. %s\n3. %s\n4. %s\n5. %s\n", p1, p2, p3, p4, p5)
		} else if x == "" {
			continue
		} else {
			_, errNum := fmt.Scan(&y)
			if errNum != nil {
				var trash string
				fmt.Scan(&trash)
				fmt.Printf("Запись на место номер %s невозможна: некорректный ввод\n", x)
				continue
			}
			if y <= 0 || y > 5 {
				fmt.Printf("Запись на место номер %d невозможна: очередь переполнена\n", y)
				continue
			}
			isBusy := false
			switch y {
			case 1: if p1 != "-" { isBusy = true }
			case 2: if p2 != "-" { isBusy = true }
			case 3: if p3 != "-" { isBusy = true }
			case 4: if p4 != "-" { isBusy = true }
			case 5: if p5 != "-" { isBusy = true }
			}
			if isBusy {
				fmt.Printf("Запись на место номер %d невозможна: место уже занято\n", y)
				continue
			}
			switch y {
			case 1: p1 = x
			case 2: p2 = x
			case 3: p3 = x
			case 4: p4 = x
			case 5: p5 = x
			}
			count++
		}
	}
}
