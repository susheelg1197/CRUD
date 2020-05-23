package models

type UserDetails struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Address        string `json:"address"`
	BloodGroup     string `json:"bloodGroup"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	Feedback       string `json:"feedback"`
	IsDonated      bool   `json:"isDonated"`
	dateOfDonation string `json:"dateOfDonation"`
	image          string `json:"image"`
}
