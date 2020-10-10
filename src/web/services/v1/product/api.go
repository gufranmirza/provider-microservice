package product

import "net/http"

// Product interface
type Product interface {
	GetByID(w http.ResponseWriter, r *http.Request)
}
