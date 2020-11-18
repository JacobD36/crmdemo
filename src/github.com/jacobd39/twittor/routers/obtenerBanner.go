package routers

import (
	"io"
	"net/http"
	"os"

	db "github.com/jacobd39/twittor/bds"
)

//ObtenerBanner envía el Banner al HTTP
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	if perfil.Banner == "" {
		http.Error(w, "Imagen de Banner no encontrada", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)

	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}
