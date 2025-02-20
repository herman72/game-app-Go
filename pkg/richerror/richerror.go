package richerror

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	
)

type RichError struct {
	wrappedError error
	message string

}