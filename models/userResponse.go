package models

type UserResponse struct {
	Id       int    `json:"id"`
	PhotoUrl string `json:"photoUrl"`
	UserName string `json:"userName"`
	FullName string `json:"fullName"`
	Token    string `json:"token"`
}
