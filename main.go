package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

func init() {
	connectToAWS()
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := db.ListTables().All()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(result)
		err = tpl.ExecuteTemplate(w, "index.gohtml", struct{Tables []string}{result})
		if err != nil {
			log.Fatalln(err)
		}
	})

	//manejador del nombre para enseñar los elementos
	http.HandleFunc("/table/", func(w http.ResponseWriter, r *http.Request) {
		table := db.Table(strings.TrimPrefix(r.URL.Path, "/table/"))
		//fmt.Println(table.Describe().Run())

		table.Put(struct{b_id int}{2}).Run()
		q := table.Get("b_id", 2)
		fmt.Println(q)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}