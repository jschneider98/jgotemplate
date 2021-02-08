package web

import (
	"github.com/gocraft/web"
	"github.com/jschneider98/jgoweb/util"
)

// GET: Index route
func (ctx *WebContext) index(rw web.ResponseWriter, req *web.Request) {
	err := ctx.Template.Execute(rw, nil)

	if err != nil {
		ctx.JobError(util.WhereAmI(), err)
	}

	ctx.JobSuccess()
}
