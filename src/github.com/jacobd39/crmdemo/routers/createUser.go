package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jacobd39/crmdemo/bd"
	model "github.com/jacobd39/crmdemo/models"
)

//CreateUser función que registra un usuario en la BD
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var t model.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in the received data "+err.Error(), 400)
		return
	}

	if len(t.Name1) == 0 {
		http.Error(w, "The First Name field is required", 400)
		return
	}

	if len(t.LastName1) == 0 {
		http.Error(w, "The First Last Name field is required", 400)
		return
	}

	if len(t.LastName2) == 0 {
		http.Error(w, "The Second Last Name field is required", 400)
		return
	}

	if t.Campaign == 0 {
		http.Error(w, "You must select one campaign at least", 400)
		return
	}

	_, userFound, _ := bd.UserVerification(t.CodUser)

	if userFound == true {
		http.Error(w, "User already exists", 400)
		return
	}

	t.CreateAt = time.Now()
	t.ModifiedAt = time.Now()

	idUsr, status, err := bd.InsertNewUser(t)

	if err != nil {
		http.Error(w, "An error occurred while trying to register the user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "The user registration could not be inserted", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(idUsr)
}
