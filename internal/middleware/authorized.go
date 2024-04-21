package middleware

import (
	"context"
	"imi/college/internal/models"
	"imi/college/internal/writers"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type userKey string

const UserKey userKey = userKey("user")

func writeError(w http.ResponseWriter) {
	writers.Error(w, "Forbidden", http.StatusForbidden)
}

func EnsureUserSession(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		h := func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if len(authHeader) == 0 {
				writeError(w)
				return
			}

			if !strings.HasPrefix(authHeader, "Bearer ") {
				writeError(w)
				return
			}

			headerParts := strings.SplitN(authHeader, " ", 2)
			if headerParts == nil {
				writeError(w)
				return
			}

			providedToken := headerParts[1]
			if len(providedToken) < 1 {
				writeError(w)
				return
			}

			var session models.UserSession

			if err := db.Where(&models.UserSession{Token: providedToken}).First(&session).Error; err != nil {
				writeError(w)
				return
			}

			ctx := context.WithValue(r.Context(), UserKey, session)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(h)
	}
}
