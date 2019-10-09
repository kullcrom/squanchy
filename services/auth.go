package auth

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"net/http"
)

//Authenticate will authenticate a user based on a session token
func Authenticate(r *http.Request, pool *redis.Pool) (bool, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return false, err
	}

	conn := pool.Get()
	defer conn.Close()

	response, err := conn.Do("GET", cookie.Value)
	if err != nil {
		return false, err
	}

	if response == nil {
		return false, errors.New("Error: Redis response cannot be nil")
	}
	return true, nil
}
