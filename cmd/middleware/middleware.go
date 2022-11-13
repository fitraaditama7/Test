package middleware

import (
	"net/http"
	"test/cmd/config"
	"test/cmd/constant"
	"test/pkg/responses"
)

func RequiredAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get(constant.HeaderKey)

		if key == "" {
			responses.Error(w, constant.ApiKeyMissingError)
			return
		}

		if key != config.GetKeyConfig() {
			responses.Error(w, constant.InvalidApiKey)
			return
		}
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}
