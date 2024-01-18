package main

import (
	"bufio"
	"os"
)

func main() {

	fSource, err := os.OpenFile("./defers/data/in.txt", os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fSource.Close()

	fDest, _ := os.OpenFile("./defers/data/out.txt", os.O_RDWR|os.O_CREATE, 0666)
	defer fDest.Close()

	s := bufio.NewScanner(fSource)

	for s.Scan() {
		fDest.Write(s.Bytes())
	}

}
