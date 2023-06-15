package common_errors

import "net/http"

func AccountNotAuthorizedError(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusUnauthorized)
	_, _ = (*w).Write([]byte(`{ "status": "account not authorized" }`))
}
