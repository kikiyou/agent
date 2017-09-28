// routes.go

package main

import (
	"github.com/kikiyou/agent/collector"
)

func initializeRoutes() {

	// Use the setUserStatus middleware for every route to set a flag
	// indicating whether the request was from an authenticated user or not
	router.Use(setUserStatus())

	// Handle the index route
	// router.GET("/", ensureLoggedIn(), showIndexPage)
	router.GET("/", showIndexPage)
	router.GET("/server", collector.ModulesRoutes)
}
