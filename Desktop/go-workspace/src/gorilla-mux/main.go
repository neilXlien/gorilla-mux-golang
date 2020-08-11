package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = ":3000"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func abo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello this is ABOUT page</h1>")
}

func con(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>hello this is CONTACT page</h1>")
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", hello)
	router.HandleFunc("/about", abo)
	router.HandleFunc("/contact", con)
	log.Println("listening on localhost", port)
	http.ListenAndServe(port, router)

}
