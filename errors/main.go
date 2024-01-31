package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	fSource, _ := os.OpenFile("./errors/data/in.txt", os.O_RDONLY, 0666)
	defer func() {
		err := fSource.Close()
		if err != nil {
			log.Fatalf("Не удалось закрыть файл in.txt: %v", err)
		}
	}()

	fDest, _ := os.OpenFile("./errors/data/out.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer func() {
		err := fDest.Close()
		if err != nil {
			log.Fatalf("Не удалось закрыть файл out.txt: %v", err)
		}
	}()

	s := bufio.NewScanner(fSource)

	var i int
	defer func() {
		fmt.Printf("Total string: %d", i)
	}()
	for s.Scan() {
		i++
	}
	err := s.Err()
	if !errors.Is(err, io.EOF) && err != nil {
		log.Fatalf("Не удалось завершить чтение файла, %v", err)
	}

}
