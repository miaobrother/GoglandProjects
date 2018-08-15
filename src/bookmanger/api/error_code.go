package main

//import "fmt"

const (
	ErrSuccess    = 0
	ErrInvalidNum = 1001
)

func GetMessage(code int) (msg string) {
	switch code {
	case ErrSuccess:
		return "The success.."
	case ErrInvalidNum:
		return "Get error "
	default:
		return "参数无效"
	}

}
