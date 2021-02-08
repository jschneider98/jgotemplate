package web

import (
	"errors"
	"fmt"
	"github.com/gocraft/web"
	"github.com/jschneider98/jgoweb"
	"github.com/jschneider98/jgoweb/util"
	"html/template"
	"net/http"
)

type WebContext struct {
	jgoweb.WebContext
}

//
func NewWebContext() *WebContext {
	ctx := &WebContext{}
	ctx.User, _ = jgoweb.NewUser(ctx)

	return ctx
}

// Not found handler
func (ctx *WebContext) NotFound(rw web.ResponseWriter, r *web.Request) {
	rw.WriteHeader(http.StatusNotFound)

	params := struct {
		Title   string
		Message string
	}{}

	params.Title = "Page Not Found"
	params.Message = "We could not find the page you were looking for."

	tmpl, err := ctx.GetTemplate("error.html")

	if err != nil {
		ctx.JobWarning(util.WhereAmI(), err)
		fmt.Fprintf(rw, "Page not found")

		return
	}

	err = tmpl.Execute(rw, params)

	if err != nil {
		ctx.JobWarning(util.WhereAmI(), err)
		fmt.Fprintf(rw, "Page not found")

		return
	}

	ctx.JobSuccess()
}

// Error handler
func (ctx *WebContext) Error(rw web.ResponseWriter, r *web.Request, err interface{}) {
	rw.WriteHeader(http.StatusInternalServerError)

	logErr := errors.New(fmt.Sprintf("%v", err))
	ctx.JobWarning(util.WhereAmI(), logErr)

	params := struct {
		Title   string
		Message string
	}{}

	params.Title = "Error"
	params.Message = "We're sorry, but we've encountered an unexpected error. Our fault not yours..."

	tmpl, newErr := ctx.GetTemplate("error.html")

	if newErr != nil {
		ctx.JobWarning(util.WhereAmI(), newErr)
		fmt.Fprintf(rw, "Application error.")

		return
	}

	newErr = tmpl.Execute(rw, params)

	if newErr != nil {
		ctx.JobWarning(util.WhereAmI(), newErr)
		fmt.Fprintf(rw, "Application error.")

		return
	}

	ctx.JobSuccess()
}

// Save flash messages to session
func (ctx *WebContext) SetFlashMessages(rw web.ResponseWriter, messages template.HTML) {
	ctx.SessionPutString(rw, "flash_messages", string(messages))
}

// flash message middleware
func (ctx *WebContext) LoadFlashMessages(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	var err error

	if ctx.Template == nil {
		next(rw, req)
		return
	}

	messages, err := ctx.SessionGetString("flash_messages")

	if err != nil {
		ctx.JobError(util.WhereAmI(), err)
	}

	// clear flash messages
	ctx.SessionPutString(rw, "flash_messages", "")

	// Add messages to main layout template
	if messages != "" {
		str := `[[define "flash-messages"]]` + messages + "[[end]]"
		ctx.Template, err = ctx.Template.Parse(str)

		if err != nil {
			ctx.JobError(util.WhereAmI(), err)
		}
	}

	next(rw, req)
}
