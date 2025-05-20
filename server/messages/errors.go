package messages

import "errors"

var ErrLoginExists = errors.New("login already exists")
var ErrHashPass = errors.New("can't hash password")
var ErrLoginNotExists = errors.New("login doens't exists")
var ErrIncorrectPass = errors.New("incorrect login or password")
var ErrCreateJWT = errors.New("can't create token")
var ErrMissOrInvToken = errors.New("missing or invalid token")
var ErrInvOrExpToken = errors.New("invalid or expired token")
var ErrBadConvertReqToJSON = errors.New("can't convert request to json")
