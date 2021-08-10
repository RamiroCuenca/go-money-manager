package models

import "errors"

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

func (c *CreateUserCMD) Validate() error {
	// Validate it has a username
	if c.Username == "" {
		return errors.New("Username can not be empty")
	}

	if c.Age < 0 {
		return errors.New("Age must be positive")
	}

	return nil
}
