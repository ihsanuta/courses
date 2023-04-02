package model

type TypeUser string

const (
	TypeUserAdmin    TypeUser = "admin"
	TypeUserNonAdmin TypeUser = "user"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID        int      `json:"id"`
	Username  string   `json:"username"`
	Password  string   `json:"-"`
	Type      TypeUser `json:"type"`
	CreatedAt string   `json:"created_at"`
}

type RespLogin struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}
