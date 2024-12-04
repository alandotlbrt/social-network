package config

type Response struct {
	Id      int
	Success bool   `json:"success"`
	Message string `json:"message"`
	Info    InformationUser
	Errors  []string `json:"errors,omitempy"`
}

// Receive
type InformationRegister struct {
	Id              string
	Username        string
	Firstname       string
	Lastname        string
	Password        string
	ConfirmPassword string
	NewPassword     string
	Birthday        string
	Email           string
	About           string
	Photo           string
	Privacy         string
}

// Send
type InformationUser struct {
	Id                string
	Username          string
	Firstname         string
	Lastname          string
	Birthday          string
	Email             string
	About             string
	Profil_Pictures   string
	Privacy           string
	Follower          int
	Follow            int
	ConfirmPassword   string
	NewPassword       string
	Password          string
	IsFollowing       bool
	IsWaitingFollow   bool
	FollowNumberInbox int
}

type PostStruct struct {
	Id_post string
	Id_user string
	Content string
	Image   string
	Privacy string
	Like    string
}

type ChatStruct struct {
	Pictures string
	Id_chat  int
	From_Id  int
	To_Id    int
	Content  string
	Type     string
	Image    string
}

type UserChatStruct struct {
	Pictures  string
	Id_User   int
	Username  string
	Firstname string
	Lastname  string
}

type UserFollowStruct struct {
	Username      string
	Id            string
	Pictures      string
	Privacy       string
	HeIsFollowing bool
	HeIsWaiting   bool
}
