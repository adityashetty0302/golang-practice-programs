package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.in")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("No. of bytes", len(bs))
	return len(bs), nil
}
