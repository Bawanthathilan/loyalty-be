package models

type LoyaltyAccount struct {
	ID             string `json:"id"`
	ProgramID      string `json:"program_id"`
	Balance        int    `json:"balance"`
	LifetimePoints int    `json:"lifetime_points"`
	CustomerID     string `json:"customer_id"`
	EnrolledAt     string `json:"enrolled_at"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Mapping        struct {
		ID          string `json:"id"`
		CreatedAt   string `json:"created_at"`
		PhoneNumber string `json:"phone_number"`
		Password    string `json:"password"` // ← new!
	} `json:"mapping"`
}


