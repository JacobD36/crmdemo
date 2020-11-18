package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/jacobd39/twittor/bds"
	model "github.com/jacobd39/twittor/models"
)

//ConsultaRelacion chequea si hay relación entre dos usuarios
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t model.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp model.RespuestaConsultaRelacion

	status, err := db.ConsultaRelacion(t)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
