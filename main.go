package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	Name        string
	Description string
	Price       float32
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %.2f", c.Name, c.Description, c.Price)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	course := Course{
		Name:        "Golang",
		Description: "Golang course",
		Price:       150.90,
	}

	json.NewEncoder(w).Encode(course)
}
