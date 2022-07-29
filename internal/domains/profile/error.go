package profile

import "errors"

var (
	ErrProfileInvalid = errors.New("You must pass the name of profile as arg 0. e.g.: dock profile create foobar")
)
