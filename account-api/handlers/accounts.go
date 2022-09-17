// Package classification of Account API
//
// Documentation for Account API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta
package handlers

import (
	"gorm.io/gorm"
	"instagram-clone/internal"
	"log"
)

// KeyAccount for serialization/deserialization
type KeyAccount struct{}

// Accounts handler for getting, creating and updating accounts
type Accounts struct {
	logger    *log.Logger
	db        *gorm.DB
	validator *internal.Validation
}

// NewAccounts returns a new accounts handler with the given logger
func NewAccounts(logger *log.Logger, db *gorm.DB, validator *internal.Validation) *Accounts {
	return &Accounts{logger, db, validator}
}
