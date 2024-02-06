package middlewares

import (
	"context"
	"net/http"
	"price-tracking-api-gateway/src/constants"
	"price-tracking-api-gateway/src/utils"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the auth token
		tokenString := r.Header.Get(constants.AUTH_HEADER)
		if tokenString == "" {
			http.Error(w, constants.AUTH_HEADER_REQUIRED, http.StatusUnauthorized)
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != constants.BEARER_PREFIX {
			http.Error(w, constants.INVALID_AUTH_FORMAT, http.StatusUnauthorized)
			return
		}

		accessToken := parts[1]
		userData, err := utils.VerifyUserByJWT(accessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !(userData.Success) {
			http.Error(w, constants.UNAUTHORIZED_MESSAGE, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userData", userData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
