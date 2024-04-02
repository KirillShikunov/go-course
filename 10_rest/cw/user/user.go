package user

import "errors"

type User struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
}

var ErrUserNotFound = errors.New("user not found")

func Get(Id int) (*User, error) {
	for _, user := range users {
		if user.Id == Id {
			return user, nil
		}
	}

	return nil, ErrUserNotFound
}

func List() []*User {
	return users
}

var users = []*User{
	{1, "active"},
	{2, "inactive"},
}
