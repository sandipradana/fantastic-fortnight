package model

type Admin struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
