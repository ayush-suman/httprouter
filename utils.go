package httprouter

import (
	"context"
	"net/http"
)

func applyMiddlewares(handle Handle, middlewares []Middleware, hasMiddlewares bool) Handle {
	if !hasMiddlewares {
		return handle
	}

	for i := len(middlewares) - 1; i >= 0; i-- {
		handle = middlewares[i](handle)
	}
	return handle
}

func HandleFromHandler(handler http.Handler) Handle {
	return func(w http.ResponseWriter, req *http.Request, ps Params) {
		if len(ps) > 0 {
			ctx := req.Context()
			ctx = context.WithValue(ctx, ParamsKey, ps)
			req = req.WithContext(ctx)
		}
		handler.ServeHTTP(w, req)
	}
}