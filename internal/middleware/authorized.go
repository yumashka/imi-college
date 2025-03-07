package middleware

import (
	"context"
	"imi/college/internal/ctx"
	"imi/college/internal/httpx"
	"imi/college/internal/permissions"
	"imi/college/internal/writer"
	"net/http"

	"gorm.io/gorm"
)

func writeError(w http.ResponseWriter) {
	data := httpx.Unauthorized()
	writer.JSON(w, data.Status, data)
}

func RequireUser(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user, err := httpx.GetCurrentUserFromRequest(db, r)
			if err != nil {
				writeError(w)
				return
			}

			c := context.WithValue(r.Context(), ctx.UserKey, user)

			next.ServeHTTP(w, r.WithContext(c))
		}

		return http.HandlerFunc(fn)
	}
}

func writeBadPermissions(w http.ResponseWriter) {
	data := httpx.Forbidden()
	writer.JSON(w, data.Status, data)
}

func RequirePermissions(required int64) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user, err := ctx.GetCurrentUser(r)
			if err != nil {
				writeBadPermissions(w)
				return
			}

			if permissions.HasPermissions(user.Permissions, permissions.PermissionAdmin) {
				next.ServeHTTP(w, r)
				return
			}

			if permissions.HasPermissions(user.Permissions, required) {
				next.ServeHTTP(w, r)
				return
			}

			writeBadPermissions(w)
		}

		return http.HandlerFunc(fn)
	}
}
