package main

import (
	"fmt"
	"gopkg.in/underarmour/dynago.v2"
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
		result, err := client.ListTables().Execute()
		if err != nil {
			log.Fatalln(err)
		}

		err = tpl.ExecuteTemplate(w, "index.gohtml", struct{Tables []string}{result.TableNames})
		if err != nil {
			log.Fatalln(err)
		}
	})

	//manejador del nombre para enseñar los elementos
	http.HandleFunc("/table/", func(w http.ResponseWriter, r *http.Request) {
		tableName := strings.TrimPrefix(r.URL.Path, "/table/")
		table, _ := client.DescribeTable(tableName)

		fmt.Println(table.Table.AttributeDefinitions)
		fmt.Println(table.Table.KeySchema)
		fmt.Println(table.Table.GlobalSecondaryIndexes)
		fmt.Println(table.Table.TableStatus)
		a, _ := client.Query(tableName).KeyConditionExpression("b_id != :id", dynago.P(":id", 2)).Desc().Execute()

		fmt.Println(a)
		//
		//table.Put(struct{b_id int}{2}).Run()
		//q := table.Get("b_id", 2)
		//fmt.Println(q)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}