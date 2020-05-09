package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/joho/godotenv"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/conclave-dev/celoist-backend/routes"
	"github.com/conclave-dev/celoist-backend/util"
)

var rateLimiter *limiter.Limiter
var sentryHandler *sentryhttp.Handler

// Wrapper for rate-limiter handler
func rateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpError := tollbooth.LimitByRequest(rateLimiter, w, r)

		if httpError != nil {
			rateLimiter.ExecOnLimitReached(w, r)
			w.Header().Add("Content-Type", rateLimiter.GetMessageContentType())
			w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))

			// Direct the error to Sentry
			if hub := sentry.GetHubFromContext(r.Context()); hub != nil {
				hub.CaptureException(httpError)
			}
		}

		// Proceed to the next middleware/handler when the request rate is safe
		next.ServeHTTP(w, r)
	})
}

// Wrapper for sentry handler
func sentryMiddleware(next http.Handler) http.Handler {
	return sentryHandler.Handle(next)
}

func main() {
	startServer()
}

func startServer() {
	// Load .env and set the networkID for the backend server
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Setup clients for all supported networks
	util.SetupClients()

	router := mux.NewRouter().StrictSlash(true)
	routes.SetUpRoutes(router)
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	port, isPortSet := os.LookupEnv("PORT")
	if !isPortSet {
		// Default to Port 3001
		port = "3001"
	}

	server := &http.Server{
		Addr:         "localhost:" + port,
		Handler:      cors(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting at %s", port)
	log.Fatal(server.ListenAndServe())
}

func setUpRouter() *mux.Router {
	// Setup rate-limiter to 1 request per second
	rateLimiter = tollbooth.NewLimiter(1, nil)

	// Setup Sentry for error/events logging
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://dce55a06dc394be8a4275562e770bd9e@sentry.io/5171009",
		Debug:            true,
		AttachStacktrace: true,
	}); err != nil {
		log.Fatal("Sentry initialization failed: ", err.Error())
	}

	sentryHandler = sentryhttp.New(sentryhttp.Options{
		Repanic:         true,
		WaitForDelivery: false,
	})

	router := mux.NewRouter().StrictSlash(true)
	router.Use(sentryMiddleware)
	router.Use(rateLimitMiddleware)
	routes.SetUpRoutes(router)
	return router
}
