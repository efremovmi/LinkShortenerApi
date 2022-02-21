package pkg

import "errors"

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound    = errors.New("Record not found")
	NotShortUrlInRequest = errors.New("Missing short link in request")
	RangeOutLenShortUrl  = errors.New("The length of the short url must be greater than 0 and less than 10")
	RangeOutLenUrl       = errors.New("The length of the url must be greater than 0 and less than 500")
)
