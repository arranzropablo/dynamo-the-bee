package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/underarmour/dynago.v2"
	"html/template"
	"log"
	"net/http"
)

type Table struct {
	Fields []string
	Items []dynago.Document
}

var tpl *template.Template

func init() {
	connectToAWS()
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	router := mux.NewRouter()

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		err := tpl.ExecuteTemplate(w, "index.gohtml", home())
		if err != nil {
			log.Fatalln(err)
		}
	}).Methods("GET")

	tablemux := router.PathPrefix("/table/{name}").Subrouter()

	tablemux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]

		//mostrar el boton de buscar y los campos y eso
		err := tpl.ExecuteTemplate(w, "index.gohtml", tableDetail(name))

		if err != nil {
			log.Fatalln(err)
		}
	}).Methods("GET")

	tablemux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {

	}).Methods("POST")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}