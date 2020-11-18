package models

//LoginReturn tiene el token que se devuelve con el login
type LoginReturn struct {
	Token string `json:"token,omitempty"`
	ID    string `json:"id,omitempty"`
}
