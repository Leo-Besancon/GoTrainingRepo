package main

import (
	"io"
	"log"
	"os"
)

func main() {

	myFileName := os.Args[1]

	myFile, err := os.Open(myFileName)

	if err != nil {
		log.Fatal("Fatal error : ", err)
	}

	_, err2 := io.Copy(os.Stdout, myFile)

	if err2 != nil {
		log.Fatal("Fatal error : ", err2)
	}
}
