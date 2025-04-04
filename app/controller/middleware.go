package controller

import (
	"github.com/gocanto/blog/app/env"
	"net/http"
)

func (s MiddlewareStack) Logging(next BaseController) BaseController {
	return func(w http.ResponseWriter, r *http.Request) *HttpError {
		println("Incoming request:", r.Method, r.URL.Path)

		err := next(w, r)

		if err != nil {
			println("Middleware returned error:", err.Message)
		} else {
			println("Middleware completed successfully")
		}

		return err
	}
}

func (s MiddlewareStack) AdminUser(next BaseController) BaseController {
	return func(w http.ResponseWriter, r *http.Request) *HttpError {
		salt := r.Header.Get(env.ApiKeyHeader)

		if s.isAdminUser(salt) {
			return next(w, r)
		}

		return Unauthorised("Unauthorized", nil)
	}
}

func (s MiddlewareStack) isAdminUser(seed string) bool {
	return s.userAdminResolver(seed)
}

func (s MiddlewareStack) Push(controller BaseController, middlewares ...Middleware) BaseController {
	// Apply middleware in reverse order, so the first middleware in the list is executed first.
	for i := len(middlewares) - 1; i >= 0; i-- {
		controller = middlewares[i](controller)
	}

	return controller
}
