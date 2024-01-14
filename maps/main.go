package main

import "fmt"

func main() {

	var borrowers map[string]map[string][]string = map[string]map[string][]string{
		"Винокуров": {
			"Газеты": []string{"123", "456", "789"},
			"Книги":  []string{"Война и Мира", "Как стать ниже", "Незнайка на луне"},
		},
		"Киреев": {
			"Газеты": []string{"222", "555", "888"},
			"Книги":  []string{"C# для маленьких", "Кубернетис в картинках", "Секс в большом городе"},
		},
	}

	fmt.Println(borrowers)

	fmt.Println(fmt.Sprintf("Читателей в базе: %v", len(borrowers)))

	for index, value := range borrowers {
		fmt.Println(fmt.Sprintf("Читатель %v взял в билиотеке:", index))
		m := value
		for ind, val := range m {
			fmt.Println(fmt.Sprintf("%v:", ind))
			for _, in := range val {
				fmt.Println(fmt.Sprintf("\"%v\"", in))
			}

		}
	}

}
