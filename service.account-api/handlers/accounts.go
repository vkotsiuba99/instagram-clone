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

// A single account response
// swagger:response accountResponse
type accountResponse struct {
	// The id of the account that will be returned
	// in: path
	// required: true
	ID uint `json:"id"`
	// Returned account in the system
	// in: body
	Body []internal.Account
}

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
