package login

// Customer
type Customer struct {
	customerID          uint64
	email               string
	registeredTS        uint64
	lastLoginTS         uint64
	emailToken          string
	emailConfirmationTS uint64
	status              bool
	displayName         string
	imageID             string
	coverID             string
	title               string
	salutation          string
	firstName           string
	lastName            string
	birthday            string
	displayLanguage     string
	country             string
}

// NewLoginProvider returns a new Login provider
func CreateCustomerV1() *Customer {
	return &Customer{
		customerID:          1,
		email:               "asdf@poo.com",
		registeredTS:        1,
		lastLoginTS:         1,
		emailToken:          "asdf",
		emailConfirmationTS: 1,
		status:              true,
		displayName:         "poo",
		imageID:             "uploads/poo.jpg",
		coverID:             "uploads/cover/poo.jpg",
		title:               "sir",
		salutation:          "mr",
		firstName:           "poo",
		lastName:            "pimpel",
		birthday:            "11.1.1112",
		displayLanguage:     "ger",
		country:             "de",
	}
}
