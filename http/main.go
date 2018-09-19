package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type MyWriter struct{}

func main() {

	// bs := make([]byte, 100000)

	res, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	// res.Body.Read(bs)

	mw := MyWriter{}

	io.Copy(mw, res.Body)
	// fmt.Println(string(bs))

}

func (MyWriter) Write(bs []byte) (int, error) {

	fmt.Println(string(bs))
	return 0, nil
}
