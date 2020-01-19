package code

type Code string

const (
	OK          Code = "ok"
	BadParams   Code = "bad_params"
	NotFound    Code = "not_found"
	Internal    Code = "internal_error"
	ExternalAPI Code = "external_api_error"
	Database    Code = "database_error"
	Forbidden   Code = "forbidden"
	Unknown     Code = "unknown"
)

func (c Code) String() string { return string(c) }
