package models

// Transaction model struc for Transaction
type Transaction struct {
	ID          string  `json:"id"`
	UserID      int64   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
	State       bool    `json:"state"`     // 1 = done - 0 = pending
	Recurrent   bool    `json:"recurrent"` // 1 = yes - 0 = no
}

// Create Transaction Command
type CreateTransactionCMD struct {
	UserID      int64   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Amount      float64 `json:"amount"`
	State       bool    `json:"state"`     // 1 = done - 0 = pending
	Recurrent   bool    `json:"recurrent"` // 1 = yes - 0 = no
}
