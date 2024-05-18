package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
  //parses the form
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successfull\n")

  // reads the values from the request
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  //checks the path
	if r.URL.Path != "/hello" {
		http.Error(w, "Not found bro 404", http.StatusNotFound)
		return
	}

  //checks the method
	if r.Method != "GET" {
		http.Error(w, "Not supported method bro 404", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!\n")
}

func main() {
	//create a servemux
	mux := http.NewServeMux()

  //FileServer serves the whole file
  //Dir is used to get the directory
	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/", fileServer)
  // root is handled by fileServer it looks for index.html
  // and any other files in the dir can be used 
  // for eg: apipath/from.html will serve the form.html as the response

	mux.HandleFunc("/form", formHandler)
	mux.HandleFunc("/hello", helloHandler)

	fmt.Println("Listening to server on port 8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}
