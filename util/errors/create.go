package errors

import "fmt"

// NewConvErr return customError{cade: bad_params, err: "failed to convert <item> because <reason>"}
func NewConvErr(item, reason string) error {
	return New(BadParams, fmt.Sprintf("failed to convert %s because %s", item, reason))
}
