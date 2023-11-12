package service

import "errors"

var (
	ErrorCouldNotHashPassword = errors.New("could not hash password")
	ErrorInvalidCredentials   = errors.New("invalid credentials")
	ErrorUnauthorized         = errors.New("unauthorized")
	ErrorArtistNotExist       = errors.New("artist does not exist")
	ErrorVenueNotExist        = errors.New("venue does not exist")
	ErrorUpdateVenueFailed    = errors.New("could not update venue")
	ErrorSlotNotExist         = errors.New("timeslot does not exist")
	ErrorInvalidStatus        = errors.New("invalid status")
	ErrorAlreadyApplied       = errors.New("already applied")
)
