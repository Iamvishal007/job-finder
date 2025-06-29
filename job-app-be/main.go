package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello form the Go backend! Current time: %s", time.Now().Format("2006-01-02 15:04:05"))
	})

	log.Println("Go server starting on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
