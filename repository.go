package main

import (
	"gopkg.in/underarmour/dynago.v2"
	"log"
)

var client *dynago.Client

func connectToAWS() {

	client = dynago.NewClient(dynago.NewAwsExecutor(dynago.ExecutorConfig{
		Region:    "us-east-1",
		AccessKey: "",
		SecretKey: "",
		SessionToken: "",
	}))

}

func getTables() []string {
	result, err := client.ListTables().Execute()
	if err != nil {
		log.Fatalln(err)
	}
	return result.TableNames
}

func getTableItems(tableName string, limit uint) (fields []string, items []dynago.Document)  {
	table, _ := client.DescribeTable(tableName)

	rows, _ := client.Scan(tableName).Limit(limit).Execute()
	set := make(map[string]bool)
	fields = make([]string, len(table.Table.KeySchema))

	for _, i := range table.Table.KeySchema {
		if v, ok := set[i.AttributeName]; !ok && !v {
			fields = append(fields, i.AttributeName)
		}
		set[i.AttributeName] = true
	}

	for _, i := range rows.Items {
		for k := range i {
			if v, ok := set[k]; !ok && !v {
				fields = append(fields, k)
			}
			set[k] = true
		}
	}

	return fields, rows.Items
}