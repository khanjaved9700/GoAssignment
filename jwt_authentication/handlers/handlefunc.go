package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key") // creating jwt key used in jwt token
var users = map[string]string{    //  using to comparing username and password because i have not used databse right now
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct { // used to pass the data form the api
	Username string `json:"username"`
	Password string `json: "password"`
}

type Claims struct { // creating for payload for our JWT so inside the payload we can user name and claims struct
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login so user can post login request to the server...
func Login(w http.ResponseWriter, r *http.Request) {

	var credential Credentials // declear credentials for using ref of username and pass

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // if getting error so there is a bad request passed
		return
	}

	expectedPassword, ok := users[credential.Username]

	if !ok || expectedPassword != credential.Password { // check password are okay or not if not the return unauthorized access
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if okay then create claims so creating claims here

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: credential.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// creating token using claims function with method signing method hs256 and passing claims algo as a arguments
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // this throws entranal sever error
		return
	}

	// set the cookies

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}

// home page

func Home(w http.ResponseWriter, r *http.Request) {

	// getting cookie from the request
	cookie, err := r.Cookie("token")
	// check error and No Cookie err also
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value // creating token string
	claims := &Claims{}      //assign cliams

	// parsingwith claims and passed the toknstring, cliams and func
	// so function returns jwt token thats assign in tkn variable
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// check error and errorsignatureinvalid error also
	if err != nil {

		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check weather toker are valid or invalid
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if token is valid then print hello with user name or any message

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	// getting cookie from the request
	cookie, err := r.Cookie("token")
	// check error and No Cookie err also
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value // creating token string
	claims := &Claims{}      //assign cliams

	// parsingwith claims and passed the toknstring, cliams and func
	// so function returns jwt token thats assign in tkn variable
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// check error and errorsignatureinvalid error also
	if err != nil {

		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check weather toker are valid or invalid
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// check expiration time if expiration time is arround 30 sec remaining at that time token will be refreshed
	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)
	claims.ExpiresAt = expirationTime.Unix()

	// creating token using claims function with method signing method hs256 and passing claims algo as a arguments
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // this throws entranal sever error
		return
	}

	// set the cookies

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})

}
