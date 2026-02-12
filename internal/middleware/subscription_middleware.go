package middleware

import (
	"Assignment3/internal/service"
	"net/http"
)

func RequirePremium(userService *service.UserService, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := GetUserID(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		user, err := userService.GetByID(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		if user.Role != "premium" {
			http.Error(w, "Premium subscription required", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

func GetUserRole(userService *service.UserService, userID string) string {
	user, err := userService.GetByID(userID)
	if err != nil {
		return "free"
	}
	if user.Role == "" {
		return "free"
	}
	return user.Role
}
