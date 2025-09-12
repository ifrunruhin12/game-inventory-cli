package main

import (
	"os"
	"path/filepath"
	"text/template"
)

type Item struct {
	Name  string
	Count int
}

type Player struct {
	Name  string
	Level int
}

type ReportData struct {
	Player    Player
	Inventory []Item
}

func main() {
	tmplPath := filepath.Join("templates", "report.tmpl")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		panic(err)
	}

	data := ReportData{
		Player: Player{
			Name:  "Popcycle",
			Level: 7,
		},
		Inventory: []Item{
			{Name: "Sword", Count: 2},
			{Name: "Potion", Count: 3},
			{Name: "Bow", Count: 1},
		},
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
