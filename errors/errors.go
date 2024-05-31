package Errors

import "errors"

var ErrForbidden403 = errors.New("forbidden")
var ErrNotFound404 = errors.New("not found - requested entity is not found in database")
var ErrServerError500 = errors.New("internal server error - request is valid but operation failed at server side")
