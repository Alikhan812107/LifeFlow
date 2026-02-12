package middleware

import (
	"net/http"
	"sync"
)

type SessionStore struct {
	sessions map[string]string
	mu       sync.RWMutex
}

var Store = &SessionStore{
	sessions: make(map[string]string),
}

func (s *SessionStore) Set(sessionID, userID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessions[sessionID] = userID
}

func (s *SessionStore) Get(sessionID string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	userID, exists := s.sessions[sessionID]
	return userID, exists
}

func (s *SessionStore) Delete(sessionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, sessionID)
}

func GetUserID(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", false
	}
	return Store.Get(cookie.Value)
}

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := GetUserID(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}
