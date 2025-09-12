package main

import (
	"os"
	"text/template"
)

func main() {
	const tmplText = "Hello, {{.}}! Welcome to your inventory report. \n"

	tmpl, err := template.New("report").Parse(tmplText)
	if err != nil {
		panic(err)
	}

	playerName := "Popcycle"

	err = tmpl.Execute(os.Stdout, playerName)
	if err != nil {
		panic(err)
	}
}
