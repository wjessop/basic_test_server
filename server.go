package main

import (
	"log"
	"net/http"
	"time"
)

func logRequest(req *http.Request) {
	log.Println(req.RemoteAddr, req.Method, req.RequestURI, req.Proto, req.ContentLength, req.TransferEncoding)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	logRequest(req)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("You got in!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	s := &http.Server{
		Addr:           "127.0.0.1:8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
