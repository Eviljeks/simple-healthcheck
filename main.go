package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("you should specify url, try call \"./simple-healthcheck http://localhost:5000/health\"")
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()

		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatalln(err)
		}

		log.Fatalf(fmt.Sprintf("status code is not 200: %d. Response:\n%s", resp.StatusCode, string(b)))
	}

	fmt.Println("healthy")
	os.Exit(0)
}
