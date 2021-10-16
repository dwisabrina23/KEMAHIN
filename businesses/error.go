package businesses

import "errors"

var (
	ErrDuplicateData           = errors.New("data already exist")
	ErrUsernamePasswordInvalid = errors.New("invalid Username or password")
	ErrUserNotFound            = errors.New("cannot find user")
	ErrInternalServer          = errors.New("something went wrong")
	ErrIDNotFound              = errors.New("id not found")
	ErrInvalidEventID          = errors.New("invalid order id")
	ErrOutOfStock              = errors.New("ticket out of stock")
	ErrInvalidPaymentMeth      = errors.New("invalid payment method")
	ErrOrderNotFound           = errors.New("order not found")
	ErrInvalidOrderID          = errors.New("invalid order id")
	ErrInvalidOrgID            = errors.New("invalid organizer id")
	ErrInvalidTicketID         = errors.New("invalid order id")
	ErrFillData                = errors.New("please fill in the required data")
)
