package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/form.html")
		fmt.Fprint(w, "Broo")
		return
	} else if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// parse form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm Error: %v\n", err)
	}

	w.Write([]byte("POST request successfull"))

	// form values
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Not supperted method 404", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello!"))
}

func main() {
	fs := http.FileServer(http.Dir("./static/"))
	// this will serve the static directory by default
	// and look for index.html
	http.Handle("/", fs)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	// formHandler)
	fmt.Println("Sever on prot 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
