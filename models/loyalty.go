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

var LoyaltyAccounts = []LoyaltyAccount{
	{
		ID:             "278fd068-2b3c-4d09-9c6c-d5cd8c3bd902",
		ProgramID:      "6b12bf92-0469-4071-9432-43d74ca7d74b",
		Balance:        12,
		LifetimePoints: 22,
		CustomerID:     "4YYBWHK2Q9EXDHRWF2RNDTE55W",
		EnrolledAt:     "2025-07-12T14:22:12Z",
		CreatedAt:      "2025-07-11T18:06:52Z",
		UpdatedAt:      "2025-07-12T16:54:28Z",
		Mapping: struct {
			ID          string "json:\"id\""
			CreatedAt   string "json:\"created_at\""
			PhoneNumber string "json:\"phone_number\""
			Password    string "json:\"password\""
		}{
			ID:          "77fd657f-79f3-4b7b-9e10-669e9eb0b5a0",
			CreatedAt:   "2025-07-11T18:06:52Z",
			PhoneNumber: "+94768327247",
			Password:    "helloworld",
		},
	},
	{
		ID:             "e4ff4dcd-3d65-4ac9-8a53-3e0b71c3b16d",
		ProgramID:      "6b12bf92-0469-4071-9432-43d74ca7d74b",
		Balance:        0,
		LifetimePoints: 0,
		CustomerID:     "6RRJWK5DY90C0H6P2Y39HG2BN8",
		CreatedAt:      "2025-07-12T15:06:02Z",
		UpdatedAt:      "2025-07-12T15:06:03Z",
		Mapping: struct {
			ID          string "json:\"id\""
			CreatedAt   string "json:\"created_at\""
			PhoneNumber string "json:\"phone_number\""
			Password    string "json:\"password\""
		}{
			ID:          "4072f645-45fb-4462-99ef-0593932a955a",
			CreatedAt:   "2025-07-12T15:06:02Z",
			PhoneNumber: "+94719095814",
			Password:    "123456",
		},
	},
	{
		ID:             "e042319f-f0b5-41bd-8dd9-7e7781b2bd56",
		ProgramID:      "6b12bf92-0469-4071-9432-43d74ca7d74b",
		Balance:        0,
		LifetimePoints: 0,
		CustomerID:     "QEJP1YH14ZYAHAN33Q63VJV8Z8",
		CreatedAt:      "2025-07-12T16:59:05Z",
		UpdatedAt:      "2025-07-12T16:59:05Z",
		Mapping: struct {
			ID          string "json:\"id\""
			CreatedAt   string "json:\"created_at\""
			PhoneNumber string "json:\"phone_number\""
			Password    string "json:\"password\""
		}{
			ID:          "c59e5362-e365-4d8a-8c02-23e24030edb9",
			CreatedAt:   "2025-07-12T16:59:05Z",
			PhoneNumber: "+94719095815",
			Password:    "12345",
		},
	},
}
