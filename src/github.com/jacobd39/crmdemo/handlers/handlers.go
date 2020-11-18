package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jacobd39/crmdemo/middlew"
	"github.com/jacobd39/crmdemo/routers"
	"github.com/rs/cors"
)

//BackendHandlers setean mi puerto, el handler y luego pongo a escuchar al servidor
func BackendHandlers() {
	router := mux.NewRouter()

	router.HandleFunc("/newUser", middlew.CheckDB(middlew.CheckJWT(routers.CreateUser))).Methods("POST")
	router.HandleFunc("/usersList", middlew.CheckDB(middlew.CheckJWT(routers.UsersList))).Methods("GET")
	router.HandleFunc("/getUser", middlew.CheckDB(middlew.CheckJWT(routers.GetUser))).Methods("GET")
	router.HandleFunc("/updateUser", middlew.CheckDB(middlew.CheckJWT(routers.UpdateUser))).Methods("PUT")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8090"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
