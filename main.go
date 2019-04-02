package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	connectToAWS()
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	result, err := db.ListTables().All()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "index.gohtml", struct{Tables []string}{result})
		if err != nil {
			log.Fatalln(err)
		}
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}