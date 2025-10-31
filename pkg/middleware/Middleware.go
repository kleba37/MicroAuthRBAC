package Middleware

import "net/http"

type Middleware struct {
	Handlers []func(http.Handler) http.Handler
}

func (m *Middleware) Apply(next http.Handler) http.Handler {
	for _, han := range m.Handlers {
		next = han(next)
	}
	return next
}

func New(handlers ...func(handlers http.Handler) http.Handler) *Middleware {
	return &Middleware{
		Handlers: handlers,
	}
}
