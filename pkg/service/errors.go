package service

import "errors"

var (
	ErrorCouldNotHashPassword = errors.New("could not hash password")
	ErrorInvalidCredentials   = errors.New("invalid credentials")
	ErrorUnauthorized         = errors.New("unauthorized")
	ErrorArtistNotExist       = errors.New("artist does not exist")
	ErrorSlotNotExist         = errors.New("timeslot does not exist")
	ErrorInvalidStatus        = errors.New("invalid status")
)
