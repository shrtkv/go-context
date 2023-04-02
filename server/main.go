package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Server started")
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

// func home(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	log.Println("request started")
// 	defer log.Println("request ended")

// 	select {
// 	case <-time.After(time.Second * 5):
// 		log.Println("-- request succeeded")
// 		w.Write([]byte("-- request succeeded :D\n"))
// 	case <-ctx.Done():
// 		http.Error(w, "-- request aborted :(\n", http.StatusRequestTimeout)
// 	}
// }

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ip := r.RemoteAddr
	log.Printf("request started from IP address: %s", ip)
	defer log.Println("request ended")

	select {
	case <-time.After(time.Second * 5):
		log.Println("-- request succeeded")
		fmt.Fprintf(w, " \nYour IP address is: %s\n", ip) // include IP address in response

		w.Write([]byte("-- request succeeded :D\n"))
	case <-ctx.Done():
		http.Error(w, "-- request aborted :(\n", http.StatusRequestTimeout)
	}
}
