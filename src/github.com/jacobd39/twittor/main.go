package main

import (
	"log"

	db "github.com/jacobd39/twittor/bds"
	hnd "github.com/jacobd39/twittor/handlers"
)

func main() {
	if db.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	hnd.Manejadores()
}
