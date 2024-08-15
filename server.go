package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Greetings struct {
	Message string `json:"message"`
}

func greeter(w http.ResponseWriter, my_greet string, ch chan string) {
	greet := Greetings{
		Message: my_greet,
	}
	json.NewEncoder(w).Encode(greet)
	ch <- "Done"
}

func main() {
	r := mux.NewRouter()
	c := make(chan string)
	r.HandleFunc("/good_greet", func(w http.ResponseWriter, r *http.Request) {
		go greeter(w, "Hello stranger", c)
	})

	r.HandleFunc("/bad_greet", func(w http.ResponseWriter, r *http.Request) {
		go greeter(w, "Piss of will ya", c)
	})

	http.ListenAndServe(":8080", r)
}
