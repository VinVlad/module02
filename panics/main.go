package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fSource, err := os.OpenFile("./panics/data/in.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fSource.Close()

	fDest, _ := os.OpenFile("./panics/data/out.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer func() {
		rec := recover()
		if rec != nil {
			writtenText, _ := os.ReadFile("./panics/data/out.txt")
			fmt.Println(string(writtenText))
			fmt.Println(rec)
		}
	}()

	defer fDest.Close()

	s := bufio.NewScanner(fSource)
	i := 0

	for s.Scan() {
		i++
		str := s.Text()
		countStr := strings.Count(str, "|")
		if countStr < 2 {
			panic(fmt.Sprintf("parse error: empty field on string %v", i))
		}

		substrings := strings.Split(str, "|")

		writtenString := fmt.Sprintf("Row: %d\nName: %s\nAdress: %s\nCity: %s\n\n\n", i, substrings[0], substrings[1], substrings[2])

		fDest.Write([]byte(writtenString))
	}

}
