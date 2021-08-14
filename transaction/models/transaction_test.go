package models

import (
	"testing"
	"time"
)

func NewTransaction(
	userId int64,
	date time.Time,
	title string,
	description string,
	category string,
	amount float64,
) *CreateTransactionCMD {
	return &CreateTransactionCMD{
		UserID:      userId,
		Date:        date,
		Title:       title,
		Description: description,
		Amount:      amount,
		Category:    category,
	}
}

func Test_ValidateCreateTransactionWithCorrectParams(t *testing.T) {
	transac := NewTransaction(23, time.Now(), "The title", "A description", "My category", 35.3)

	err := transac.validate()

	if err != nil {
		t.Error("Fail. The transaction has correct params.")
		t.Fail()
	} else {
		t.Log("Success")
	}
}
