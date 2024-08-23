package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"trip-planner/internal/database"

	"github.com/gorilla/sessions"
)

type UserDto struct {
	UserID   int32
	Email    string
	Username string
}

type UserSessionStore interface {
	GetUserFromSession(r *http.Request, w http.ResponseWriter) (*UserDto, error)
	CreateUserSession(r *http.Request, w http.ResponseWriter, user *database.User) error
}

const (
	userSessionKey = "user-session"
	userKey        = "user"
)

type userSessionStore struct {
	store *sessions.CookieStore
}

func getSessionKey() string {
	sk := os.Getenv("SESSION_KEY")
	if sk == "" {
		panic("SESSION_KEY not set in .env file")
	}
	return sk
}

func NewUserSessionStore() *userSessionStore {
	sk := getSessionKey()
	store := sessions.NewCookieStore([]byte(sk))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // valid for a week
		HttpOnly: true,
	}
	return &userSessionStore{
		store,
	}
}

func (u *userSessionStore) getSession(r *http.Request) (*sessions.Session, error) {
	return u.store.Get(r, userSessionKey)
}

func (u *userSessionStore) saveSession(r *http.Request, w http.ResponseWriter, session *sessions.Session) error {
	return session.Save(r, w)
}

func (u *userSessionStore) CreateUserSession(r *http.Request, w http.ResponseWriter, user *database.User) error {
	userDto := &UserDto{
		UserID:   user.UserID,
		Email:    user.Email,
		Username: user.Username,
	}
	userDtoJson, err := json.Marshal(userDto)
	if err != nil {
		return err
	}

	userSession, err := u.store.Get(r, userSessionKey)
	if err != nil {
		return err
	}

	userSession.Values[userKey] = userDtoJson
	err = u.saveSession(r, w, userSession)
	if err != nil {
		return err
	}
	return nil
}

func (u *userSessionStore) GetUserFromSession(r *http.Request, w http.ResponseWriter) (*UserDto, error) {
	userSession, err := u.getSession(r)
	if err != nil {
		return nil, err
	}
	userDtoJson := userSession.Values[userSessionKey]
	if userDtoBytes, ok := userDtoJson.([]byte); !ok {
		return nil, fmt.Errorf("userDtoJson from user session not correct type: %s", userDtoJson)
	} else {
		var userDto *UserDto
		err = json.Unmarshal(userDtoBytes, &userDto)
		if err != nil {
			return nil, err
		}
		return userDto, nil
	}
}
