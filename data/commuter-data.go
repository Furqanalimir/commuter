package data

import (
	"errors"
	"time"
)

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"name"`
	Age       int    `json:"age"`
	Drinks    []int  `json:"drinks"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

func GetNextId() int {
	u := userData[len(userData)-1]
	return u.ID + 1
}
func GetUser(id int) (*User, error) {
	for _, u := range userData {
		if u.ID == id {
			return &u, nil
		}
	}
	return &User{}, errors.New("User not found")
}

func AddUser(u *User) error {
	ul := userData
	u.ID = GetNextId()
	u.CreatedAt = time.Now().UTC().String()
	u.UpdatedAt = time.Now().UTC().String()
	ul = append(ul, *u)
	return nil
}

var userData = []User{
	User{
		ID:       1,
		UserName: "jhon",
		Age:      25,
		Drinks:   []int{1, 2},
	},
}
