package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Andreffelipe/git-flow/greeting"
)

func GreetingHandle(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()["g"]
	if !ok {
		log.Println("Url Param 'key' is missing")
		return
	}
	g, err := greeting.Greeting(key[0])

	if err != nil {
		msg := err.Error()
		json.NewEncoder(w).Encode(fmt.Sprintf("error: %v", msg))
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(fmt.Sprintf("massage: %v", g))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", GreetingHandle)
	log.Fatal(http.ListenAndServe(":3000", mux))
}
