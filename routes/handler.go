package routes

import (
	"fmt"
	"net/http"

	"github.com/conclave-dev/celoist-backend/routes/medium"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// SetUpRoutes initiates the setup process for all routes
func SetUpRoutes(router *mux.Router) {
	router.HandleFunc("/health", healthCheck)

	// Add all routes and other actions required for setup
	medium.AddRoutes(router)
}
