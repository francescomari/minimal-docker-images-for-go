package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatalf("unable to execute request: %v\n", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("unable to read the response body: %v\n", err)
	}

	fmt.Print(string(body))
}
