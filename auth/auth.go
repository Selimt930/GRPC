package auth

import (
	//"MailService/service"
	"context"
	//"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	ID       int    `json:"id"`
	jwt.StandardClaims
}

type favContextKey string

const KeyContext favContextKey = "user"

//get user ID from token and return it with context-response
func GetUserFromToken(r *http.Request) (*http.Request, error) {
	var user string
	token, err := parseToken(r)

	if err != nil {
		return nil, err
	}
	if token == nil {
		return r, nil
	}
	if claims, valid := token.Claims.(*Claims); valid {
		user = strconv.Itoa(claims.ID)
	} else {
		return nil, errors.New("token invalid")
	}
	cont := r.Context()
	cont = context.WithValue(cont, KeyContext, user)
	return r.WithContext(cont), nil
}

//get token from requests header by "Authorization" key
func parseToken(r *http.Request) (*jwt.Token, error) {
	reqAuth := r.Header.Get("Authorization")

	token, err := jwt.ParseWithClaims(reqAuth, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getKey(), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token, nil
}

/*
func Verification(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	token, err := getToken(r)
	if err != nil {
		fmt.Errorf("invalid token")
	}
	if len(token) == 0 {
		fmt.Errorf("invalid token")
	}
	decode := json.NewDecoder(r.Body).Decode(&creds)
	if decode != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from in memory map
	expectedPassword, ok := service.Users[creds.Username]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	//expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
}

*/
