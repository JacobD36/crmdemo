package routers

import (
	"encoding/json"
	"net/http"
	"time"

	db "github.com/jacobd39/twittor/bds"
	"github.com/jacobd39/twittor/jwt"
	model "github.com/jacobd39/twittor/models"
)

//Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	var t model.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := db.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "Usuario y/o contraseña inválidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}

	resp := model.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
