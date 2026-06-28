package app_error

type Kind string

const (
	KindNotFound        Kind = "not_found"
	KindConflict        Kind = "conflict"
	KindUnauthorized    Kind = "unauthorized"
	KindForbidden       Kind = "forbidden"
	KindInvalidArgument Kind = "invalid_argument"
)
