package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const banjoURL = "http://localhost:8989/api/pets?id=05cc936a-6068-11ed-8427-32bc7acc9df8"

type location struct {
	City  string
	State string
}
type pet struct {
	ID       string
	Name     string
	Species  string
	Color    string
	Age      int64
	Weight   float64
	Location location
}

func run() error {
	req, err := http.NewRequest("GET", banjoURL, nil)
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var banjo pet
	if err := json.Unmarshal(b, &banjo); err != nil {
		return err
	}

	fmt.Printf("requested URL:\n%v\n\nresponse:\n%v\n\nunmarshalled to Go struct:\n%#v\n", banjoURL, string(b), banjo)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
