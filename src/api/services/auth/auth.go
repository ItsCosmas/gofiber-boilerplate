package auth

import (
	"time"

	"gofiber-boilerplate/api/models/user"

	cfg "gofiber-boilerplate/api/configs"

	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

// Claims represents JWT claims
type Claims struct {
	ExternalID string `json:"id"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

// IssueToken generate tokens used for auth
func IssueToken(u user.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add((24 * time.Hour) * 14) // 14 days

	// Create token
	// token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// claims := token.Claims.(jwt.MapClaims)
	// claims["userId"] = u.ExternalID
	// claims["role"] = u.Role
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.

	claims := Claims{
		u.ExternalID,
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(cfg.GetConfig().JWTSecret))
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(cfg.GetConfig().JWTSecret), nil
		},
	)

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
