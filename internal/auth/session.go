package auth

import (
	"os"

	"github.com/gorilla/sessions"
)

func getSessionKey() string {
	sk := os.Getenv("SESSION_KEY")
	if sk == "" {
		panic("SESSION_KEY not set in .env file")
	}
	return sk
}

func CreateStore() *sessions.CookieStore {
	sk := getSessionKey()
	store := sessions.NewCookieStore([]byte(sk))
	return store
}
