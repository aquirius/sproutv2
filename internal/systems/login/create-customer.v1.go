package login

// Customer
type Customer struct {
	CustomerID          uint64 `json:"customer_id"`
	Email               string `json:"email"`
	RegisteredTS        uint64 `json:"registered_ts"`
	LastLoginTS         uint64 `json:"last_login_ts"`
	EmailToken          string `json:"email_toke"`
	EmailConfirmationTS uint64 `json:"email_confirmation_ts"`
	Status              bool   `json:"status"`
	DisplayName         string `json:"display_name"`
	ImageID             string `json:"image_id"`
	CoverID             string `json:"cover_id"`
	Title               string `json:"title"`
	Salutation          string `json:"salutation"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Birthday            string `json:"birthday"`
	DisplayLanguage     string `json:"language"`
	Country             string `json:"country"`
}

type CreateCustomerV1Params struct {
}

type CreateCustomerV1Result struct {
	Customer *Customer `json:"customer"`
}

//CreateCustomerV1 creates a customer object with given arguments
func (l *Login) CreateCustomerV1(p *CreateCustomerV1Params, res *CreateCustomerV1Result) error {
	customer := &Customer{
		CustomerID:          1,
		Email:               "asdf@poo.com",
		RegisteredTS:        1,
		LastLoginTS:         1,
		EmailToken:          "asdf",
		EmailConfirmationTS: 1,
		Status:              true,
		DisplayName:         "poo",
		ImageID:             "uploads/poo.jpg",
		CoverID:             "uploads/cover/poo.jpg",
		Title:               "sir",
		Salutation:          "mr",
		FirstName:           "poo",
		LastName:            "pimpel",
		Birthday:            "11.1.1112",
		DisplayLanguage:     "ger",
		Country:             "de",
	}
	*res = CreateCustomerV1Result{Customer: customer}
	return nil
}
