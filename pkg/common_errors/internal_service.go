package common_errors

import "net/http"

func InternalServiceError(w *http.ResponseWriter) {
	(*w).WriteHeader(http.StatusInternalServerError)
	_, _ = (*w).Write([]byte("Please try again later.\n"))
}
