package main

import "fmt"

func main() {
	s := []string{"Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота", "Воскресенье"}

	workingDays := make([]string, 5)

	p := copy(workingDays, s)
	fmt.Println(p)
	fmt.Println(workingDays)

	weekend := make([]string, 2)
	t := copy(weekend, s[5:])

	fmt.Println(t)
	fmt.Println(weekend)

	weekDay := append(workingDays, weekend...)

	fmt.Println(weekDay)

}
