package schema

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct { //Payload
	Id               string `bson:"_id,omitempty" json:"_id,omitempty" redis:"Id"`
	Email            string `bson:"email,omitempty" json:"email,omitempty" redis:"Email"`
	Password         string `bson:"password,omitempty" json:"password,omitempty" redis:"Password"`
	InitialTimestamp string `bson:"initial_timestamp,omitempty" json:"initial_timestamp,omitempty" redis:"InitialTimestamp"`
	RecentTimestamp  string `bson:"recent_timestamp,omitempty" json:"recent_timestamp,omitempty" redis:"RecentTimestamp"`
	Role             int    `bson:"role,omitempty" json:"role,omitempty" redis:"Role"`
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

func (user User) MarshalBinary() ([]byte, error) {
	return json.Marshal(user)
}

func (user User) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, user); err != nil {
		return err
	}
	return nil
}
