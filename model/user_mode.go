package model

type UserModel struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	GlobalName string `json:"global_name"`
}
