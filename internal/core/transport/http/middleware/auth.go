package core_http_middleware

import (
	"context"
	"net/http"
	"strings"

	core_auth "github.com/BladeRunner322/ChatGO/internal/core/auth"
	core_errors "github.com/BladeRunner322/ChatGO/internal/core/errors"
	core_logger "github.com/BladeRunner322/ChatGO/internal/core/logger"
	core_http_response "github.com/BladeRunner322/ChatGO/internal/core/transport/http/response"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(jwtService *core_auth.JWT) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := core_logger.FromContext(r.Context())
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				responseHandler.ErrorResponse(
					core_errors.ErrInvalidToken,
					"Authorization header required",
				)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				responseHandler.ErrorResponse(
					core_errors.ErrInvalidToken,
					"Invalid Authorization header format. Use 'Bearer <token>'",
				)
				return
			}

			if parts[1] == "" {
				responseHandler.ErrorResponse(
					core_errors.ErrInvalidToken,
					"Token is empty",
				)
				return
			}

			userID, err := jwtService.Parse(parts[1])
			if err != nil {
				responseHandler.ErrorResponse(
					err,
					"Invalid or expired token",
				)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(UserIDKey).(int)
	return userID, ok
}
