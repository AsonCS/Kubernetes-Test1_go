package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	port     = os.Getenv("PORT")
	name     = os.Getenv("NAME")
	age      = os.Getenv("AGE")
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
)

var startedTime = time.Now()

func main() {
	// for {
	// 	println("Hello World")
	// 	time.Sleep(20 * time.Second)
	// }
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/", Hello)
	// get port from env
	if port == "" {
		port = "80"
	}
	println("Listening on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, I'm %s, I'm %s years old</h1>", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
		return
	}
	fmt.Fprintf(w, "<h1>My family:</h1><br /><pre>%s</pre>", string(data))
}

func Healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedTime)
	if duration.Seconds() > 120 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Not healthy, uptime: %s", duration.String())
	} else if duration.Seconds() < 15 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Not ready, uptime: %s", duration.String())
	} else {
		fmt.Fprintf(w, "Healthy, uptime: %s", duration.String())
	}
}

func Secret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, I'm %s, my password is %s</h1>", user, password)
}
