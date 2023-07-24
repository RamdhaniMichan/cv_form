package entity

type Skill struct {
	Model
	Skill     string `json:"skill"`
	Level     string `json:"level"`
	ProfileID int    `json:"profileID"`
}
