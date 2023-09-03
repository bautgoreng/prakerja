/*
package models

import (

	"gorm.io/gorm"

)

	type User struct {
		gorm.Model
		Name     string `json:"name" form:"name"`
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}
*/
package models

type User struct {
	Id       int    `json:"id"`
	PhotoUrl string `json:"photoUrl"`
	UserName string `json:"userName"`
	FullName string `json:"fullName"`
	Password string `json:"password"`
}
