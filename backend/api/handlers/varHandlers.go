package handlers

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Register struct {
	Id              string `json:"Id"`
	Username        string `json:"Username"`
	Firstname       string `json:"Firstname"`
	Lastname        string `json:"Lastname"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	NewPassword     string `json:"newPassword"`
	Email           string `json:"email"`
	Photo           string `json:"Photo"`
	Birthday        string `json:"Birthday"`
	About           string `json:"About"`
	Privacy         string `json: "Privacy"`
}

type followJSONStruct struct {
	IdUser       string `json:"iduser"`
	IdToFollow   string `json:"idtofollow"`
	IdToUNFollow string `json:"idunfollow"`
	IdToDecline  string `json:"idtodecline"`
	IdToAccept   string `json:"idtoaccept"`
	IdUserState  string `json:"iduserstate"`
	Type         string `json:"type"`
}
