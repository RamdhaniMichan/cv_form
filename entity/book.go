package entity

type Book struct {
	Model

	Title  string `gorm:"index" json:"title"`
	Author string `json:"author"`
	Page   string `json:"page"`
	Genre  string `json:"genre"`
}
