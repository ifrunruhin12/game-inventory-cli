// Package models keep the custom types and models for the CLI
package models

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
