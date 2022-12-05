package models

type Accueil struct {
	Id        int    `json:"id"`
	Titre     string `json:"titre"`
	Createurs string `json:"createurs"`
}

type User struct {
	Id                    int    `json:"id"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	FirstName             string `json:"firstName"`
	LastName              string `json:"lastName"`
	Email                 string `json:"email"`
	EmailVerificationDate int    `json:"emailVerificationDate"`
}

type Team struct {
	Id              int    `json:"id"`
	Name            string `json:"teamName"`
	UserCount       int    `json:"userCount"`
	GameName        string `json:"game"`
	TournamentCount int    `json:"tournamentCount"`
	CreatorName     string `json:"creatorName"`
	CreationDate    string `json:"creationDate"`
}

type Tournament struct {
	Id         int    `json:"id"`
	Name       string `json:"tournamentName"`
	Count      int    `json:"count"`
	Price      int    `json:"price"`
	Game       string `json:"game"`
	TeamsCount int    `json:"teamCount"`
	IsFinished int    `json:"isFinished"`
	Mode       string `json:"mode"`
	BeginDate  string `json:"beginDate"`
	BeginTime  string `json:"beginTime"`
}
