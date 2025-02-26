package routes

import (
	"database/sql"

	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/delivery/http"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/middleware"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/repository"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/internal/usecase"
	"github.com/HaikalRFadhilahh/go-auth-api-clean-architecture/pkg"
	"github.com/gorilla/mux"
)

func UserRouter(r *mux.Router, db *sql.DB) {
	// Declare Handler, Usecase, Repository
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := http.NewUserHandler(userUsecase)

	// Sub Routing
	auth := r.PathPrefix("/auth").Subrouter()
	user := r.PathPrefix("/users").Subrouter()

	// Routing Auth (Login,Register)
	auth.HandleFunc("/login", pkg.ConvertToHttpHandleFunc(userHandler.Login)).Methods("POST")
	auth.HandleFunc("/register", pkg.ConvertToHttpHandleFunc(userHandler.Register)).Methods("POST")
	auth.HandleFunc("/validate", pkg.ConvertToHttpHandleFunc(userHandler.Validate)).Methods("POST")

	// User Data Middleware
	user.Use(middleware.CheckJWTMiddleware)

	// Routing Users
	user.HandleFunc("", pkg.ConvertToHttpHandleFunc(userHandler.GetAllUser)).Methods("GET")
}
