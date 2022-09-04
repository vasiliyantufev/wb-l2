package main

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	fileName string
	url      string
)

func main() {

	var url string

	fmt.Println("Enter url")
	fmt.Scan(&url)
	fmt.Println("Enter file name")
	fmt.Scan(&fileName)

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)

	}

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	size, err := io.Copy(file, resp.Body)
	defer file.Close()

	fmt.Printf("File %s size %d", fileName, size)

}
