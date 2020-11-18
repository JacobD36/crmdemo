package routers

import (
	"net/http"

	db "github.com/jacobd39/twittor/bds"
	model "github.com/jacobd39/twittor/models"
)

//BajaRelacion realiza el borrado de la relación entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe ingresar el parámetro id", http.StatusBadRequest)
		return
	}

	var t model.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := db.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
