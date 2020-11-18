package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jacobd39/crmdemo/bd"
	"github.com/jacobd39/crmdemo/jwt"
	model "github.com/jacobd39/crmdemo/models"
)

//Login realiza el login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t model.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		//http.Error(w, "INVALID_USER_PASSWORD", http.StatusBadRequest)
		http.Error(w, "INVALID_REQUEST", http.StatusBadRequest)
		return
	}

	if len(t.CodUser) == 0 {
		http.Error(w, "USERCODE_MISSING", http.StatusBadRequest)
		return
	}

	document, exists, res := bd.LoginIntent(t.CodUser, t.Password)

	if exists == false {
		//http.Error(w, "INVALID_LOGIN", http.StatusBadRequest)
		if res == 1 {
			http.Error(w, "USER_NOT_FOUND", http.StatusBadRequest)
			return
		}
		if res == 2 {
			http.Error(w, "INVALID_USER_PASSWORD", http.StatusBadRequest)
			return
		}
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "TOKEN_ERROR", http.StatusBadRequest)
		return
	}

	resp := model.LoginReturn{
		Token: jwtKey,
		ID:    document.ID.Hex(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
