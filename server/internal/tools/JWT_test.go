package tools

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

var superSecret = []byte("superSecret")

func FakeKeyFunc(token *jwt.Token) (interface{}, error) {
	return superSecret, nil
}

func CreateTestToken(t *testing.T, claims jwt.MapClaims) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	s, err := token.SignedString(superSecret)
	require.NoError(t, err)
	return s
}

func TestValidateJWT(t *testing.T) {
	now := time.Now()

	tests := []struct {
		testName string
		tokenStr string
		errHave  bool
		errWant  error
		ansWant  string
	}{
		{
			testName: "valid token",
			tokenStr: CreateTestToken(t, jwt.MapClaims{
				"login": "bumus",
				"exp":   now.Add(1 * time.Hour).Unix(),
			}),
			errHave: false,
			errWant: nil,
			ansWant: "bumus",
		},
		{
			testName: "token expired",
			tokenStr: CreateTestToken(t, jwt.MapClaims{
				"login": "bumus",
				"exp":   now.Add(-1 * time.Hour).Unix(),
			}),
			errHave: true,
			errWant: jwt.ErrTokenExpired,
			ansWant: "",
		},
		{
			testName: "invalid signature",
			tokenStr: "hallo",
			errHave:  true,
			errWant:  jwt.ErrInvalidKey,
			ansWant:  "",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			claims, err := ValidateJWT(tc.tokenStr, FakeKeyFunc)
			if tc.errHave {
				require.Error(t, err)
				require.ErrorIs(t, tc.errWant, err)
				require.Nil(t, claims)
				return
			}
			require.Equal(t, tc.ansWant, claims["login"])
		})
	}

}
