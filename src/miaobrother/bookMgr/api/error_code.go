package main

const (
	ErrSuccess          = 0
	ErrInvalidParameter = 1001
)

func getMessage(code int) {
	switch code {
	case ErrSuccess:
		msg = "success"
	case ErrInvalidParameter:
		msg = "invalid parament"
	default:
		msg = "Ukown"
	}
	return
}
