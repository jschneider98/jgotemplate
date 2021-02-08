package web

import (
	"github.com/gocraft/web"
	"github.com/jschneider98/jgoweb"
	"os"
	"path/filepath"
)

//
func GetWebRouter() *web.Router {
	path, _ := os.Getwd()

	webContext := NewUxtWebContext()

	rootRouter := web.New(*webContext).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware(CacheControlMiddleware(filepath.Join(path, "static"))).
		Middleware(jgoweb.StaticMiddleware(filepath.Join(path, "static"))).
		Middleware((*WebContext).LoadDi).
		Middleware((*WebContext).LoadEndPoint).
		Middleware((*WebContext).LoadJob).
		Middleware((*WebContext).LoadSession).
		NotFound((*WebContext).NotFound).
		Error((*WebContext).Error)

	rootRouter.Subrouter(WebContext{}, "/").
		Middleware((*WebContext).LoadTemplate).
		Middleware((*WebContext).LoadFlashMessages).
		Get("/index", (*WebContext).index)

	return rootRouter
}
