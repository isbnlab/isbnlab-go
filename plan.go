package isbnlab

import (
	"net/http"
	"strings"
)

// Headers the API adds when a response leaves out fields that need a higher
// plan. A book lookup on a Free or Standard plan, for example, returns
// otherEditions, previousEdition and nextEdition as null and names them in
// PlanRestrictedFieldsHeader.
const (
	PlanRestrictedFieldsHeader = "X-Plan-Restricted-Fields"
	PlanUpgradeHeader          = "X-Plan-Upgrade"
)

// RestrictedFields returns the fields your plan left empty in resp, by their
// API names (otherEditions), or nil when nothing was left out. Pass the
// *http.Response that Execute returns:
//
//	book, resp, err := client.BooksAPI.GetBookByIsbn(ctx, "9780765302304").Execute()
//	isbnlab.RestrictedFields(resp) // [otherEditions previousEdition nextEdition]
func RestrictedFields(resp *http.Response) []string {
	if resp == nil {
		return nil
	}
	var fields []string
	for _, value := range resp.Header.Values(PlanRestrictedFieldsHeader) {
		for _, field := range strings.Split(value, ",") {
			if field = strings.TrimSpace(field); field != "" {
				fields = append(fields, field)
			}
		}
	}
	return fields
}

// IsRestricted reports whether your plan left field empty in resp.
func IsRestricted(resp *http.Response, field string) bool {
	for _, f := range RestrictedFields(resp) {
		if f == field {
			return true
		}
	}
	return false
}

// UpgradeURL returns the page that lists the plans which include the left-out
// fields, or "" when nothing was left out.
func UpgradeURL(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	return strings.TrimSpace(resp.Header.Get(PlanUpgradeHeader))
}
