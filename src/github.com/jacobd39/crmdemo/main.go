package main

import (
	"log"

	db "github.com/jacobd39/crmdemo/bd"
	hnd "github.com/jacobd39/crmdemo/handlers"
)

func main() {
	if db.CheckDBConnection() == 0 {
		log.Fatal("No DB connection")
		return
	}
	hnd.BackendHandlers()
}