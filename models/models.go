package models

type Accueil struct {
	Id        int    `json:"id"`
	Titre     string `json:"titre"`
	Createurs string `json:"createurs"`
}
