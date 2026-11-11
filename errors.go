package isbnlab

import (
	"encoding/json"
	"errors"
)

// ErrorDetail returns the explanation the API gives in an error response's
// "detail" field, for example "This endpoint needs a Standard or Pro plan.",
// or "" if err is not an API error or the body has no detail.
//
// The HTTP status is on the *http.Response every method returns alongside err.
func ErrorDetail(err error) string {
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		return ""
	}
	var body struct {
		Detail string `json:"detail"`
	}
	if json.Unmarshal(apiErr.Body(), &body) != nil {
		return ""
	}
	return body.Detail
}
