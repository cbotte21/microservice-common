package validate

import "net/url"

const template = "Request must contain the following field(s): "

func ValidateRequestWithErrorMessage(payload url.Values, params ...string) ([]byte, bool) {
	res := template
	for i, str := range params {
		if !payload.Has(str) {
			res += str
			if i != len(params)-1 {
				res += ", "
			}
		}
	}
	return []byte(res), template == res
}
