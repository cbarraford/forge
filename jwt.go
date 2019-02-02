package forge

import (
	"log"
	"time"
)

type JWT struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	ExpiresAt   time.Time
}

func (jwt *JWT) IsExpired() bool {
	log.Printf("%+v", jwt.ExpiresAt)
	return jwt.ExpiresAt.UnixNano() < time.Now().UnixNano()
}

func (jwt *JWT) SetExpiration() {
	jwt.ExpiresAt = time.Now().Add(time.Second * time.Duration(jwt.ExpiresIn))
}
