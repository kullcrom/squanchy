package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"os"
	"strconv"
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

//GenerateSessionToken generates a random session token based on the SESSION_TOKEN_MODIFIER
func GenerateSessionToken() (string, error) {
	sessionTokenModifier, err := strconv.Atoi(os.Getenv("SESSION_TOKEN_MODIFIER"))
	if err != nil {
		return "", err
	}

	byteArray := make([]byte, sessionTokenModifier)

	_, err = rand.Read(byteArray)
	if err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(byteArray)
	return token, nil
}
