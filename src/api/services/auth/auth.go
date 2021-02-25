package auth

import (
	"time"

	"gofiber-boilerplate/api/models/user"

	cfg "gofiber-boilerplate/api/configs"

	jwt "github.com/form3tech-oss/jwt-go"
)

// Claims represents JWT claims
type Claims struct {
	ExternalID string `json:"userID"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

// IssueToken generate tokens used for auth
func IssueToken(u user.User) (string, error) {
	expireTime := time.Now().Add((24 * time.Hour) * 14) // 14 days

	// Generate encoded token
	claims := Claims{
		u.ExternalID,
		u.Role,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    cfg.GetConfig().JWTIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := tokenClaims.SignedString([]byte(cfg.GetConfig().JWTSecret))
	return t, nil
}
