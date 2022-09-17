package handlers

import (
	"gorm.io/gorm"
	"log"
)

// KeyAccount for serialization/deserialization
type KeyAccount struct{}

// Accounts handler for getting, creating and updating accounts
type Accounts struct {
	logger *log.Logger
	db     *gorm.DB
}

// NewAccounts returns a new accounts handler with the given logger
func NewAccounts(logger *log.Logger, db *gorm.DB) *Accounts {
	return &Accounts{logger, db}
}
