package services

import (
	"github.com/c-petrini/bookstore_users-api/domain/users"
	"github.com/c-petrini/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{
		Id: userId,
	}
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	if err := result.Get(userId); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
