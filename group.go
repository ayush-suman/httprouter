// Implementation similar to PR https://github.com/julienschmidt/httprouter/pull/89/files

package httprouter

import (
	"context"
	"net/http"
)

type RouterGroup struct {
	router *Router
	path string
	middlewares []Middleware
	HasMiddlewares bool
}

func NewGroup(path string) *RouterGroup {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	//Strip traling / (if present) as all added sub paths must start with a /
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return &RouterGroup {
		router: New(),
		path: path,
		middlewares: []Middleware{},
		HasMiddlewares: false,
	}	
}

func (g *RouterGroup) NewGroup(path string) *RouterGroup {
	return g.router.NewGroup(g.subPath(path))
}

func (g *RouterGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.router.ServeHTTP(w, r)
}

func (g *RouterGroup) AddMiddlewares(middlewares []Middleware) {
	g.middlewares = append(g.middlewares, middlewares...)
	g.HasMiddlewares = true
}

func (g *RouterGroup) Handle(method, path string, handle Handle) {
	g.router.Handle(method, g.subPath(path), applyMiddlewares(handle, g.middlewares, g.HasMiddlewares))
}

func (g *RouterGroup) Handler(method, path string, handler http.Handler) {
	g.router.Handle(method, path, applyMiddlewares(HandleFromHandler(handler), g.middlewares, g.HasMiddlewares))
}

func (g *RouterGroup) HandlerFunc(method, path string, handler http.HandlerFunc) {
	g.router.HandlerFunc(method, g.subPath(path), handler)
}

func (g *RouterGroup) GET(path string, handle Handle) {
	g.router.Handle("GET", g.subPath(path), applyMiddlewares(handle, g.middlewares, g.HasMiddlewares))
}
func (g *RouterGroup) HEAD(path string, handle Handle) {
	g.router.Handle("HEAD", g.subPath(path), applyMiddlewares(handle, g.middlewares, g.HasMiddlewares))
}
func (g *RouterGroup) OPTIONS(path string, handle Handle) {
	g.router.Handle("OPTIONS", g.subPath(path), applyMiddlewares(handle, g.middlewares, g.HasMiddlewares))
}

func (g *RouterGroup) POST(path string, handle Handle) {
	g.router.Handle("POST", g.subPath(path), applyMiddlewares(handle, g.middlewares, g.HasMiddlewares))
}

func (g *RouterGroup) PUT(path string, handle Handle) {
	g.router.Handle("PUT", g.subPath(path), applyMiddlewares(handle, g.middlewares, g.HasMiddlewares))
}

func (g *RouterGroup) PATCH(path string, handle Handle) {
	g.router.Handle("PATCH", g.subPath(path), applyMiddlewares(handle, g.middlewares, g.HasMiddlewares))
}

func (g *RouterGroup) DELETE(path string, handle Handle) {
	g.router.Handle("DELETE", g.subPath(path), applyMiddlewares(handle, g.middlewares, g.HasMiddlewares))
}

func (g *RouterGroup) subPath(path string) string {
	if path[0] != '/' {
		panic("path must start with a '/'")
	}
	return g.path + path
}
