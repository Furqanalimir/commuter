package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type Authentication struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7"`
}
type User struct {
	ID       int    `json:"id"`
	UserName string `json:"name" validate:"required,min=3"`
	// Authentication
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=7"`
	Phone     int       `json:"phone" validate:"required,numeric,min=10"`
	Age       int       `json:"age" validate:"required,gt=12"`
	Role      string    `json:"role" validate:"required,oneof=admin manager user"`
	createdAt time.Time `json:"-"`
	updatedAt time.Time `json:"-"`
	deletedAt time.Time `json:"-"`
}

func (u *User) ToJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func (u *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err.Field(), err.Tag())
		}
	}
	return err
}

func (u *User) VerifyRole(role string) bool {
	for _, user := range userData {
		if user.ID == u.ID && user.Role == role {
			return true
		}
	}
	return false
}

// UserStructLevelValidation contains custom struct level validations that don't always
// make sense at the field validation level. For Example this function validates that either
// FirstName or LastName exist; could have done that with a custom field validation but then
// would have had to add it to both fields duplicating the logic + overhead, this way it's
// only validated once.
//
// NOTE: you may ask why wouldn't I just do this outside of validator, because doing this way
// hooks right into validator and you can combine with validation tags and still have a
// common error output format.
func UserStructLevelValidation(sl validator.StructLevel) {

	user := sl.Current().Interface().(User)

	if len(user.UserName) == 3 || len(user.Password) == 0 {
		sl.ReportError(user.UserName, "fname", "FirstName", "fnameorlname", "")
		sl.ReportError(user.Password, "lname", "LastName", "fnameorlname", "")
	}

	// plus can do more, even with different tag than "fnameorlname"
}

func GetNextId() int {
	if len(userData) == 0 {
		return 1
	}
	u := userData[len(userData)-1]
	return u.ID + 1
}

func GetUserById(id int) (*User, error) {
	for _, u := range userData {
		if u.ID == id {
			return &u, nil
		}
	}
	return &User{}, fmt.Errorf("User with id %d not found", id)
}

func (u *User) VerifyUser() (uint64, error) {

	for _, user := range userData {
		if user.Email == u.Email && user.Password == u.Password {
			return uint64(user.ID), nil
		}
	}
	return 0, errors.New("User not found")
}

func GetUserByEmail(email string) (*User, error) {
	for _, u := range userData {
		if u.Email == email {
			return &u, nil
		}
	}
	return &User{}, errors.New("User not found")
}

// func GetUser(param map[string]any) (*User, error) {
// 	id, ok := param["id"]
// 	for _, u := range userData {
// 	if u.ID == id && ok {
// 			return &u, nil
// 		}
// 	}
// 	return &User{}, errors.New("User not found")
// }

func AddUser(u *User) error {
	user, _ := GetUserByEmail(u.Email)
	if user.ID > 0 {
		return errors.New("User already exists")
	}
	u.ID = GetNextId()
	u.createdAt = time.Now().UTC()
	u.updatedAt = time.Now().UTC()
	userData = append(userData, *u)
	return nil
}

var userData = []User{
	User{
		ID:        100,
		UserName:  "jhon doe",
		Email:     "jhondoe@gmail.com",
		Password:  "password123",
		Phone:     1234567890,
		Age:       19,
		Role:      "user",
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
		deletedAt: time.Now().UTC(),
	},
	User{
		ID:        1,
		UserName:  "joe",
		Email:     "joe@gmail.com",
		Password:  "1234567890",
		Phone:     12345678,
		Age:       11,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
		deletedAt: time.Now().UTC(),
	},
}
