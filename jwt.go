package forge

import (
	"fmt"
	"time"
)

type JWT struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	ExpiresAt   time.Time
}

func (jwt *JWT) IsExpired() bool {
	return jwt.ExpiresAt.UnixNano() < time.Now().UnixNano()
}

func (jwt *JWT) SetExpiration() {
	jwt.ExpiresAt = time.Now().Add(time.Second * time.Duration(jwt.ExpiresIn))
}

func (jwt *JWT) GetAuthHeader() string {
	return fmt.Sprintf("%s %s", jwt.TokenType, jwt.AccessToken)
}
