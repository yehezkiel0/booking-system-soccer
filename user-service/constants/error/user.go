package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrPasswordIncorrect    = errors.New("password incorrect")
	ErrUsernameExist        = errors.New("username already exist")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
	ErrEmailExist           = errors.New("email already exist")
)

var UserError = []error{
	ErrUserNotFound,
	ErrPasswordIncorrect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
	ErrEmailExist,
}
