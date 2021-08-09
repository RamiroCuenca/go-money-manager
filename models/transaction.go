package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model

	UserID      int64   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
	State       bool    `json:"state"`     // 1 = done - 0 = pending
	Recurrent   bool    `json:"recurrent"` // 1 = yes - 0 = no
}
