package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num := rand.Int() & 255

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"number":%d}`, num)
	})

	log.Fatalf("Error: %+v", http.ListenAndServe("0.0.0.0:8080", nil))
}
