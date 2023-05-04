package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main() {

	type NotificationStructure struct {
		Instance string
		NodeId   string
		Summary  string
	}

	test := NotificationStructure{"testA", "id_1234d6vr", "Disk is full"}

	tmpl, err2 := template.New("temporary").Parse("instance: <b>{{ .Instance }}</b>\nnode id: <b>{{ .NodeId }}</b>\nsummary: <b>{{ .Summary }}</b>")
	if err2 != nil {
		panic(err2)
	}

	var tpl bytes.Buffer
	err2 = tmpl.Execute(&tpl, test)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(tpl.String(), "\n")

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	tmpl2, err2 := template.New("temporary").Parse("map value {{.k1}} and {{.k2}}")
	if err2 != nil {
		panic(err2)
	}

	var tpl2 bytes.Buffer
	err2 = tmpl2.Execute(&tpl2, m)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(tpl2.String(), "\n")

	/*
		Create := func(name, t string) *template.Template {
			return template.Must(template.New(name).Parse(t))
		}

		t2 := Create("t2", "Name: {{.Name}}\n")

		t2.Execute(os.Stdout, struct {
			Name string
		}{"Jane Doe"})

		t2.Execute(os.Stdout, map[string]string{
			"Name": "Mickey Mouse",
		})

		t3 := Create("t3",
			"{{if . -}} yes {{else -}} no {{end}}\n")
		t3.Execute(os.Stdout, "not empty")
		t3.Execute(os.Stdout, "")

		t4 := Create("t4",
			"Range: {{range .}}{{.}} {{end}}\n")
		t4.Execute(os.Stdout,
			[]string{
				"Go",
				"Rust",
				"C++",
				"C#",
			})
	*/
}
