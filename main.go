package main

import (
	"github.com/kullcrom/squanchy/api"
	"github.com/kullcrom/squanchy/controllers"
	"github.com/kullcrom/squanchy/middleware"
	"log"
	"net/http"
	"os"
)

var jwtSecretKey string

func main() {
	jwtSecretKey, exists := os.LookupEnv("JWT_SECRET_KEY")
	if !exists {
		panic("ERROR: JWT_SECRET_KEY not found")
	}

	authMiddleware := middleware.AuthMiddleware{JWTSecretKey: jwtSecretKey}

	loginHandler := controllers.LoginHandler{JWTSecretKey: jwtSecretKey}
	registrationHandler := controllers.RegistrationHandler{JWTSecretKey: jwtSecretKey}
	profileHandler := controllers.ProfileHandler{}
	updateProfileHandler := controllers.UpdateProfileHandler{}

	http.HandleFunc("/login", loginHandler.HandleLogin)
	http.HandleFunc("/register", registrationHandler.HandleRegistration)

	http.Handle("/profile", authMiddleware.Authenticate(profileHandler))
	http.Handle("/profile/update", authMiddleware.Authenticate(updateProfileHandler))

	// API endpoints
	http.HandleFunc("/users", api.GetUsers)
	http.HandleFunc("/users/", api.GetUserByID)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
