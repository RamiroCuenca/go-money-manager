package models

import (
	"errors"
	"time"
)

// Transaction model struc for Transaction
type Transaction struct {
	ID          string    `json:"id"`
	UserID      int64     `json:"user_id"`
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	State       bool      `json:"state"`     // 1 = done - 0 = pending
	Recurrent   bool      `json:"recurrent"` // 1 = yes - 0 = no
}

// Create Transaction Command
type CreateTransactionCMD struct {
	UserID      int64     `json:"user_id"`
	Date        time.Time `json:"date"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Amount      float64   `json:"amount"`
	State       bool      `json:"state"`     // 1 = done - 0 = pending
	Recurrent   bool      `json:"recurrent"` // 1 = yes - 0 = no
}

// Validate CreateTransactionCMD
func (c *CreateTransactionCMD) validate() error {
	// Check that has a positive amount
	if c.Amount < 0.0 {
		return errors.New("Amount must be higher than zero")
	}

	// Check it has a title a description
	if c.Title == "" || c.Description == "" {
		return errors.New("Title and/or description must be completed")
	}

	// Should check if Category exists

	return nil
}
