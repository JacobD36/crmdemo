package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/jacobd39/twittor/bds"
	model "github.com/jacobd39/twittor/models"
)

//Registro es la funcion para crear en la BD el registro de un usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t model.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe contener como mínimo 6 caracteres", 400)
		return
	}

	_, encontrado, _ := db.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := db.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro del usuario "+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
