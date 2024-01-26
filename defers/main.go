package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var start time.Time

func LogTime() {
	duration := time.Since(start)
	fmt.Printf("Программа была выполнена за %v секунд", duration.Seconds())
}

func main() {
	start = time.Now()

	defer LogTime()

	fSource, err := os.OpenFile("./defers/data/in.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fSource.Close()

	fDest, _ := os.OpenFile("./defers/data/out.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer fDest.Close()

	s := bufio.NewScanner(fSource)

	t := 0

	for s.Scan() {
		t++
		fDest.WriteString(strconv.Itoa(t) + ". " + s.Text() + "\n")

	}
	g, _ := fDest.Stat()

	defer fmt.Printf("В файл записано %v строк и %v байт\n", t, g.Size())

}
