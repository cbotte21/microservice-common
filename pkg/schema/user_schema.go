package schema

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct { //Payload
	Id               string `bson:"_id,omitempty" json:"_id,omitempty" redis:"_id"`
	Email            string `bson:"email,omitempty" json:"email,omitempty" redis:"email"`
	Password         string `bson:"password,omitempty" json:"password,omitempty" redis:"password"`
	InitialTimestamp string `bson:"initial_timestamp,omitempty" json:"initial_timestamp,omitempty" redis:"initial_timestamp"`
	RecentTimestamp  string `bson:"recent_timestamp,omitempty" json:"recent_timestamp,omitempty" redis:"recent_timestamp"`
	Role             int    `bson:"role,omitempty" json:"role,omitempty" redis:"role"`
}

func (user User) Database() string {
	return "auth"
}

func (user User) Collection() string {
	return "users"
}

func (user User) Key() string {
	return user.Id
}

func (user *User) SetPassword(candidePassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(candidePassword), 10)
	if err != nil {
		return errors.New("could not hash password")
	}
	user.Password = string(hash)
	return nil
}

func (user User) VerifyPassword(candidePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(candidePassword))
}
