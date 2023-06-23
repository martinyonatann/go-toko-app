package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type List []interface{}

// RecoverWith will catch error value passed when panic.
func RecoverWith(catch func(recv interface{})) {
	if r := recover(); r != nil && catch != nil {
		catch(r)
	}
}

// Require only once.
func Require(required ...string) (err error) {
	var requireHasCalled bool
	onceEnv := new(sync.Once)

	if len(required) < 1 {
		return nil
	} else if requireHasCalled {
		return errors.New("env: require has called")
	}

	onceEnv.Do(func() {
		r := make(map[string]struct{})
		for _, v := range required {
			v = strings.TrimSpace(v)
			if v != "" {
				r[v] = struct{}{}
			}
		}
		requireHasCalled = true
	})

	return nil
}

type UserInfo struct {
	ID       int64
	FullName string
	Email    string
}
type MyJWTClaims struct {
	*jwt.RegisteredClaims
	UserInfo
}

// Generate your own secret key!
var secret = []byte(os.Getenv("JWT_SECRET_KEY"))

func CreateToken(sub string, userInfo UserInfo) (string, error) {
	// Get the token instance with the Signing method
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Choose an expiration time. Shorter the better
	exp := time.Now().Add(time.Minute * 30)
	// Add your claims
	token.Claims = &MyJWTClaims{
		&jwt.RegisteredClaims{
			// Set the exp and sub claims. sub is usually the userID
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   sub,
		},
		userInfo,
	}
	// Sign the token with your secret key
	val, err := token.SignedString(secret)

	if err != nil {
		// On error return the error
		return "", err
	}
	// On success return the token string
	return val, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func GeneratePassword(raw string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}
