package entity

type WorkingExperience struct {
	Model
	WorkingExperience string `json:"workingExperience"`
	ProfileID         int    `json:"profileID"`
}
