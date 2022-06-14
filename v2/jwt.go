package zoom

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func jwtToken(key string, secret string) (string, error) {
	role := "0"
	iat := (time.Now().UnixMilli() / 1000) - 30
	exp := iat + 60*60*2

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sdkKey":   key,
		"role":     role,
		"iat":      iat,
		"appKey":   key,
		"tokenExp": exp,
	})

	return token.SignedString([]byte(secret))
}

func (c *Client) addRequestAuth(req *http.Request, err error) (*http.Request, error) {
	if err != nil {
		return nil, err
	}

	// establish JWT token
	ss, err := jwtToken(c.Key, c.Secret)
	if err != nil {
		return nil, err
	}

	if Debug {
		log.Println("JWT Token: " + ss)
	}

	// set JWT Authorization header
	req.Header.Add("Authorization", "Bearer "+ss)

	return req, nil
}
