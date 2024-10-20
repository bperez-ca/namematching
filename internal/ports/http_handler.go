package ports

import "net/http"

// HTTPHandler defines the methods for handling HTTP requests
type HTTPHandler interface {
	NameMatchHandler(w http.ResponseWriter, r *http.Request)
	EmailMatchHandler(w http.ResponseWriter, r *http.Request)
}
