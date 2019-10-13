package main

import (
	"github.com/gomodule/redigo/redis"
	"github.com/kullcrom/squanchy/controllers"
	"log"
	"net/http"
	"time"
)

var pool *redis.Pool

func main() {
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	loginHandler := controllers.LoginHandler{Pool: pool}
	registrationHandler := controllers.RegistrationHandler{Pool: pool}
	profileHandler := controllers.ProfileHandler{Pool: pool}

	http.HandleFunc("/login", loginHandler.HandleLogin)
	http.HandleFunc("/register", registrationHandler.HandleRegistration)
	http.HandleFunc("/profile", profileHandler.HandleProfile)
	http.HandleFunc("/profile/update", profileHandler.HandleProfileUpdate)

	// API endpoints
	http.HandleFunc("/users", controllers.GetUsers)
	http.HandleFunc("/users/", controllers.GetUserByID)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
