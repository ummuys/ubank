package tools

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var mockSecret = []byte("super-secret")

func mockKeyFunc(token *jwt.Token) (interface{}, error) {
	return mockSecret, nil
}

func mockGenerateJWT(login string) string {
	claims := jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(time.Second * 30).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	res, _ := token.SignedString(mockSecret)
	return res
}

func TestValidateJWT(t *testing.T) {
	tests := []struct {
		name  string
		token string
		err   error
	}{
		{"All good", mockGenerateJWT("ummuys"), nil},
		{"Invalid Signature", mockGenerateJWT("ummuys") + "a", jwt.ErrSignatureInvalid},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := ValidateJWT(tc.token, mockKeyFunc)

			if err != nil && tc.err != nil {
				if errors.Is(err, tc.err) {
					return
				} else {
					t.Fatalf("get err -> %v ||| need err ->  %v:", err, tc.err)
				}

			} else if err == nil && tc.err != nil {
				t.Fatalf("get err -> %v ||| need err ->  %v:", err, tc.err)
				return
			}
		})
	}
}
