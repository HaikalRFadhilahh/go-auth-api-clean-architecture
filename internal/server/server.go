package server

import (
	"fmt"
	"log"
	"net/http"

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
	// Declare Mux Routing
	r := mux.NewRouter()

	// Middleware

	// Routing

	// Not Found Handler

	// Running Mux Router
	fmt.Println("Server Running On :", s.ListenAndServeString)
	log.Fatalln(http.ListenAndServe(s.ListenAndServeString, r))
}
