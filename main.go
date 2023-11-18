package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"astuart.co/goq"
)

type Example struct {
	AHrefList []map[string]string `goquery:"a,text,[href]"`
}

func main() {
	// res, err := http.Get("file:///")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	fileInfo, err := os.Stat("bookmarks_11_18_23.html")
	if err != nil {
		log.Fatal(err)
	}

	a := fileInfo.ModTime().Format(http.TimeFormat)
	log.Println(a)

	data, err := os.ReadFile("bookmarks_11_18_23.html")
	if err != nil {
		log.Fatal(err)
	}

	var ex Example
	err = goq.NewDecoder(bytes.NewReader(data)).Decode(&ex)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range ex.AHrefList {
		for key, value := range item {
			fmt.Println("Key:", key, "Value:", value)
		}
	}
}
