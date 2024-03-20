package a

import "errors"

var ErrArg error = errors.New("argument error")
var ErrDuplicateKeyValue error = errors.New("duplicate key value error")

type number interface {
	int | int32 | int64 | float32 | float64
}
