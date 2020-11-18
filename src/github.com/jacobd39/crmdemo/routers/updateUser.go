package routers

import (
	"encoding/json"
	"net/http"

	"github.com/jacobd39/crmdemo/bd"
	model "github.com/jacobd39/crmdemo/models"
)

//UpdateUser actualiza la información de un usuario en la base de datos
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t model.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Data Error "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = bd.UpdateRecord(t, ID)

	if err != nil {
		http.Error(w, "An error occurred while trying to modify the registry. Retry again. "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "The user registry has not been modified", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
