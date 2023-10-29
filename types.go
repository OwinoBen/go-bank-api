package main

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Account struct {
	ID            int       `json:"id"`
	FirstName     string    `json:"firstname"`
	LastName      string    `json:"lastname"`
	AccountNumber int64     `json:"account"`
	Balance       int64     `json:"balance"`
	Created_at    time.Time `json:"created_at"`
}

func NewAccount(firstName, lastNme string) *Account {
	return &Account{
		// ID:            rand.Intn(10000),
		FirstName:     firstName,
		LastName:      lastNme,
		AccountNumber: int64(rand.Intn(1000000)),
		Created_at:    time.Now().UTC(),
	}
}
