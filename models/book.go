package models

type Book struct {
	Id      uint   `gorm:"primaryKey autoIncrement" json:"id"`
	Judul   string `json:"judul"`
	Content string `json:"content"`
}
