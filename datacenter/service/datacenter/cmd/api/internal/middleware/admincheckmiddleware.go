package middleware

import "net/http"

type AdmincheckMiddleware struct {
}

func NewAdmincheckMiddleware() *AdmincheckMiddleware {
	return &AdmincheckMiddleware{}
}

func (m *AdmincheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
