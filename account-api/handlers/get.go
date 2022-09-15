package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

// ListSingle handles get requests and returns single user by id
func (p *Accounts) ListSingle(rw http.ResponseWriter, r *http.Request) {
	// TODO: get single account from database
	params := mux.Vars(r)
	id := params["id"]

	p.logger.Println(id)
}
