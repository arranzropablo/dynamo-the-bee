package main

import (
	"gopkg.in/underarmour/dynago.v2"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Table struct {
	Fields []string
	Items []dynago.Document
}

var tpl *template.Template
var templateArgs = make(map[string]interface{})

func init() {
	connectToAWS()
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templateArgs["Tables"] = getTables()

		err := tpl.ExecuteTemplate(w, "index.gohtml", templateArgs)
		if err != nil {
			log.Fatalln(err)
		}
	})

	http.HandleFunc("/table/", func(w http.ResponseWriter, r *http.Request) {
		if _, ok := templateArgs["Tables"]; !ok {
			templateArgs["Tables"] = getTables()
		}

		tableName := strings.TrimPrefix(r.URL.Path, "/table/")

		fields, items := getTableItems(tableName, 100)

		templateArgs["Table"] = Table{fields, items}
		err := tpl.ExecuteTemplate(w, "index.gohtml", templateArgs)
		if err != nil {
			log.Fatalln(err)
		}

		delete(templateArgs, "Table")

		///como tengo todos los fields en un set solo tengo q hacer un for sobre esos fields y cogerlos para cada elem asi estaran en el mismo orden
		//cuando los vaya a mostrar por pantalla q la primary key sea la primera y luego las dema


	})

	//Aqui meter para hacer la busqueda, q te muestre los fields
	//http.HandleFunc("/table/", func(w http.ResponseWriter, r *http.Request) {
	//	tableName := strings.TrimPrefix(r.URL.Path, "/table/")
	//
	//	//realmente esto seria para hacer la busqueda, para q me salgan todos los campos, para mostrar por pantalla la tabla hay q hacer el limit y no hacen falta todos los fields
	//	//maybe this takes too much time in large tables and I should limit the results to retrieve the fields
	//	b, _ := client.Scan(tableName).Execute()
	//	set := make(map[string]bool)
	//	for _, i := range b.Items {
	//		for k := range i {
	//			set[k] = true
	//		}
	//	}
	//	//cuando los vaya a mostrar por pantalla q la primary key sea la primera y luego las dema
	//	println(b)
	//
	//})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}