package sdk

import "errors"

var (
	ErrTokenInvalid = errors.New("You must pass the token as arg 0. The token must be a valid JWT bearer. e.g.: caradhras init eyJra...bEw")
	ErrAuthRequired = errors.New("you must initialize the SDK before performing an action. Run the command caradhras init <token>")
)
