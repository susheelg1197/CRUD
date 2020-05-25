package models

type Address struct {
	City          string `json:"city"`
	StreetAddress string `json:"streetAddress"`
	State         string `json:"state"`
	Country       string `json:"country"`
	Pincode       string `json:"pincode"`
}
type UserDetails struct {
	FormNo             string  `json:"formNo"`
	FirstName          string  `json:"firstName"`
	LastName           string  `json:"lastName"`
	Address            Address `json:"address"`
	BloodGroup         string  `json:"bloodGroup"`
	PhoneNumber        string  `json:"phoneNumber"`
	Email              string  `json:"email"`
	Feedback           string  `json:"feedback"`
	IsDonated          bool    `json:"isDonated"`
	DateOfDonation     string  `json:"dateOfDonation"`
	ImageUrl           string  `json:"imageUrl"`
	IsDocumentUploaded bool    `json:"isDocumentUploaded"`
	IsVerified         bool    `json:"isVerified"`
}
