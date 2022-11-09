package models

type Accueil struct {
	Id        int    `json:"id"`
	Titre     string `json:"titre"`
	Createurs string `json:"createurs"`
}

type User struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	EmailVerifiedAt int    `json:"email_verified_at"`
}

type Team struct {
	Id           int    `json:"id"`
	Name         string `json:"team_name"`
	Count        int    `json:"count"`
	GameName     string `json:"game"`
	NbrTournoi   int    `json:"nbr_tournoi"`
	CreatorName  string `json:"name_creator"`
	CreationDate string `json:"date_creator"`
}

type Tournament struct {
	Id         int    `json:"id"`
	Name       string `json:"tournament_name"`
	Count      int    `json:"count"`
	Price      int    `json:"price"`
	Game       string `json:"game"`
	TeamsCount int    `json:"nbr_teams"`
	IsFinished int    `json:"end"`
	Mode       string `json:"mode"`
}
