package auth

type User struct {
	Id             string
	Name           string `json:"name"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	CreatedAt      string
	Session_cookie string `json:"session_cookie"`
}
