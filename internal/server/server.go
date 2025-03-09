package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/apierror"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/config"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/middleware"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/routes"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/pkg"
	"github.com/gorilla/mux"
)

// Interface Server
type APIServer interface {
	Run()
}

// Struct Server
type apiServer struct {
	ListenAndServeString string
}

// Function Setup Server
func NewAPIServer() APIServer {
	return &apiServer{
		ListenAndServeString: fmt.Sprintf("%v:%v", pkg.GetEnv("APP_HOST", "127.0.0.1"), pkg.GetEnv("APP_PORT", "8000")),
	}
}

// Function Start Server
func (s *apiServer) Run() {
	// Database Connection
	db := config.NewDatabaseConnection()
	defer db.Close()

	// Declare Mux Routing
	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.LoggingMiddleware)

	// Routing
	routes.UserRouter(r, db)

	// Health Check
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Application Healthy!"))
	})

	// Not Found Handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pkg.HttpErrorResponse(w, apierror.ErrPageNotFound)
	})

	// Running Mux Router
	fmt.Println("Server Running On :", s.ListenAndServeString)
	log.Fatalln(http.ListenAndServe(s.ListenAndServeString, r))
}
