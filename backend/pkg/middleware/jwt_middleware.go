package middlewares

import (
	"errors"
	"os"
	user "quick-note/internal/user"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(u *user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate((time.Now().Add(1 * time.Hour))),
		},
	})
	ss, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func VerifyJWT(tokenString string) (*MyJWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyJWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
