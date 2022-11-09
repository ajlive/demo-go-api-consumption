package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func run() error {
	req, err := http.NewRequest("GET", "http://localhost:8989/api/pets?id=05cc936a-6068-11ed-8427-32bc7acc9df8", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
