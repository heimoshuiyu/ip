package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", IPHandler)

	log.Println("Listening on port", *port)
	http.ListenAndServe(":"+strconv.Itoa(*port), mux)
}

func IPHandler(w http.ResponseWriter, r *http.Request) {
	addr := r.RemoteAddr

	w.Write([]byte(addr))
	w.Write([]byte("\n"))

	if r.Header.Get("X-Real-IP") != "" {
		w.Write([]byte("X-Real-IP: "))
		w.Write([]byte(r.Header.Get("X-Real-IP")))
		w.Write([]byte("\n"))
	}

	if r.Header.Get("X-Forwarded-For") != "" {
		w.Write([]byte("X-Forwarded-For: "))
		w.Write([]byte(r.Header.Get("X-Forwarded-For")))
		w.Write([]byte("\n"))
	}
}
