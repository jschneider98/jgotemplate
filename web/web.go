package web

import (
	"github.com/jschneider98/jgoweb"
)

//
func Start() {
	router := GetWebRouter()
	jgoweb.SetConfigEnvVar("APP_CONFIG")
	jgoweb.Start(router)
}
