package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// for {
	// 	println("Hello World")
	// 	time.Sleep(20 * time.Second)
	// }
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/", Hello)
	// get port from env
	p := os.Getenv("PORT")
	if p == "" {
		p = "80"
	}
	println("Listening on http://localhost:" + p)
	http.ListenAndServe(":"+p, nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "<h1>Hello, I'm %s, I'm %s years old</h1>", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
		return
	}
	fmt.Fprintf(w, "<h1>My family:</h1><br /><pre>%s</pre>", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "<h1>Hello, I'm %s, my password is %s</h1>", user, password)
}
