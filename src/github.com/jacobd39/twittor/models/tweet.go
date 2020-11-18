package models

//Tweet captura del body el mensaje de nos llega
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
