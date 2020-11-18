package routers

import (
	"encoding/json"
	"net/http"
	"time"

	db "github.com/jacobd39/twittor/bds"
	model "github.com/jacobd39/twittor/models"
)

//GraboTweet permite grabar el tweet en la base de datos
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje model.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := model.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
