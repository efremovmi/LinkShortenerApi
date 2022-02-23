package errors

import "errors"

var (
	RecordNotFound           = errors.New("Record not found")
	NotShortUrlInRequest     = errors.New("Missing short link in request")
	RangeOutLenShortUrl      = errors.New("The short URL must be 10 characters long")
	RangeOutLenUrl           = errors.New("The length of the url must be greater than 0 and less than 500")
	BDnotWorking             = errors.New("The database is down. Ping error")
	IncorrectParamsConnectBD = errors.New("Error in database connection parameters")
	RightLenLessLeftLen      = errors.New("The length of the left line must be less than or equal to the length of the right")
)
