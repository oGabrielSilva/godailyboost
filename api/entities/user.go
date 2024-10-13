package entities

import (
	"time"

	"github.com/oGabrielSilva/godailyboost/api/auth"
)

type Social struct {
	Identifier string `json:"identifier"`
	URL        string `json:"URL"`
}

type User struct {
	Id  uint64 `json:"-"`
	Uid string `json:"uid"`

	Name           string   `json:"name"`
	Username       string   `json:"username"`
	AvatarURL      string   `json:"avatarURL"`
	AvatarFilePath string   `json:"-"`
	Bio            string   `json:"bio"`
	Social         []Social `json:"social"`

	Email    string `json:"email"`
	Password string `json:"-"`

	Authorities []auth.Authority

	JoinedAt  time.Time `json:"joinedAt"`
	UpdatedAt time.Time `json:"-"`
	Enabled   bool      `json:"-"`
	Locked    bool      `json:"locked"`
}
