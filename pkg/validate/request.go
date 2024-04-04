package validate

import (
	"github.com/cbotte21/microservice-common/pkg/response"
	"net/url"
)

const template = "request must contain the following field(s): "

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
	return response.GetError(res), template == res
}
