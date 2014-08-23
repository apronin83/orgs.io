package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/yosssi/ace"
	"github.com/yosssi/ace-proxy"
	"github.com/yosssi/orgs.io/controllers"
	"github.com/yosssi/orgs.io/models"
)

// New creates and returns a router.
func New(config *models.Config) http.Handler {
	// Create an Ace proxy.
	aceProxy := proxy.New(&ace.Options{
		BaseDir:       "views",
		DynamicReload: config.App.Development(),
	})

	// Create controllers.
	top := controllers.NewTop(config, aceProxy)

	// Create a router.
	router := httprouter.New()

	router.GET("/", top.Index)

	return router
}
