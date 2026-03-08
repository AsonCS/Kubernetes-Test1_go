package main

import (
	"net/http"
	"os"
)

func main() {
	// for {
	// 	println("Hello World")
	// 	time.Sleep(20 * time.Second)
	// }
	http.HandleFunc("/", Hello)
	// get port from env
	p := os.Getenv("PORT")
	println("Listening on http://localhost:" + p)
	http.ListenAndServe(":"+p, nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	println("Hello World")
	w.Write([]byte("<h1>Hello World</h1>"))
}
