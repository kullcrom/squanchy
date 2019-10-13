package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
	"os"
	"strconv"
)

//Authenticate will authenticate a user based on a session token
func Authenticate(r *http.Request, pool *redis.Pool) (string, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}

	conn := pool.Get()
	defer conn.Close()

	response, err := redis.String(conn.Do("GET", cookie.Value))
	if err != nil {
		log.Println(err)
	} else if err == redis.ErrNil {
		return response, errors.New("ERROR: Redis response cannot be nil")
	} else {
		log.Printf("%v = %v", cookie.Value, response)
	}
	return response, nil
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
