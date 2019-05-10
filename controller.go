package main

type Response map[string]interface{}

func home() Response {
	r := Response{}
	r["Tables"] = getTables()
	return r
}

func tableDetail(name string) Response {
	r := home()

	fields, items := getTableItems(name, 0)
	r["Table"] = Table{fields, items[:100]}
	return r
}

