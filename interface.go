package httprouter

import "net/http"

type IRouter interface {
	http.Handler
	GET(string, Handle)
	HEAD(string, Handle)
	OPTIONS(string, Handle)
	POST(string, Handle) 
	PUT(string, Handle) 
	PATCH(string, Handle) 
	DELETE(string, Handle)
	Handle(string, string, Handle)
	Handler(string, string, http.Handler)
	HandlerFunc(string, string, http.HandlerFunc)
	NewGroup(string) *RouterGroup
	AddMiddlewares([]Middleware)
}