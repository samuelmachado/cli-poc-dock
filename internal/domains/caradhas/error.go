package caradhas

import "errors"

var (
	ErrTokenInvalid = errors.New("You must pass the token as arg 0. The token must be a valid JWT bearer. e.g.: caradhas init eyJra...bEw")
)
