package models

// User model structure for User
type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Age         int64  `json:"age"`
	Gender      bool   `json:"gender"` // 1 = Male - 0 = Female
	Nationality string `json:"nationality"`
}

// Create User Command. Does not include the ID
type CreateUserCMD struct {
	Username    string `json:"username"`
	Age         int64  `json:"age"`
	Gender      bool   `json:"gender"` // 1 = Male - 0 = Female
	Nationality string `json:"nationality"`
}
