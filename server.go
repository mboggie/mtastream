package main

import (
	//	"encoding/json"
	// "fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var c chan string

func getMtaStatus(outchan chan string) {
	interval := 1

	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	for {
		select {
		case <-ticker.C:
			//get mta status here
			outchan <- "hello"
		}
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	for {
		msg := <-c
		w.Write([]byte(msg))
		w.Write([]byte("\r\n"))
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

func main() {

	c = make(chan string)

	go getMtaStatus(c)

	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)

	http.Handle("/", r)

	err := http.ListenAndServe(":8765", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
