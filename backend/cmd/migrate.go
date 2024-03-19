package main

import (
	db "backend/adapters/persistence"
	"log"
)

func main() {
	db.ConnectGorm()

	g := db.GetDbGorm()
	// Migrate the schema

	models := db.GetModels()
	err := g.AutoMigrate(models...)
	if err != nil {
		log.Fatal(err.Error())
	}
}
