package entity

type Profile struct {
	Model
	ProfileCode    string `json:"profileCode"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Address        string `json:"address"`
	PostalCode     int    `json:"postalCode"`
	DrivingLicense string `json:"drivingLicense"`
	Nationality    string `json:"nationality"`
	PlaceOfBirth   string `json:"placeOfBirth"`
	DateOfBirth    string `json:"dateOfBirth"`
	Photo          string `json:"photo"`
}

type Response struct {
	ProfileCode string `json:"profileCode"`
	ID          int    `json:"id"`
}

type ResponsePhoto struct {
	ProfileCode string `json:"profileCode"`
	Photo       string `json:"photo"`
}
