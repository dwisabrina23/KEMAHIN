package businesses

import "errors"

var (
	ErrDuplicateData           = errors.New("data already exist")
	ErrUsernamePasswordInvalid = errors.New("invalid Username or password")
	ErrUserNotFound            = errors.New("cannot find user")
	ErrInternalServer          = errors.New("something went wrong")
	ErrIDNotFound              = errors.New("id not found")
)
