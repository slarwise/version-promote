package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world\n"))
	})
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong\n"))
	})

	http.ListenAndServe(":8080", nil)
}
