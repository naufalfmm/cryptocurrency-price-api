package consts

import "errors"

var (
	ErrInvalidToken                = errors.New("invalid token")
	ErrPasswordConfirmationNotSame = errors.New("password confirmation is not same")
	ErrEmailHasBeenUsed            = errors.New("email has been used")
	ErrWrongPassword               = errors.New("wrong password")
	ErrEmailMissing                = errors.New("email missing")
	ErrCoinSymbolMissing           = errors.New("coin symbol is missing")
	ErrCoinHasBeenAdded            = errors.New("coin has been added for the user")
	ErrCoinMissing                 = errors.New("coin is missing")
	ErrCoinTrackMissing            = errors.New("tracking coin is missing")

	ErrUnknownConstant = errors.New("unknown constant")
)
