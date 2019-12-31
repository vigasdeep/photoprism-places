package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism-places/internal/api"
	"github.com/photoprism/photoprism-places/internal/config"
)

func registerRoutes(router *gin.Engine, conf *config.Config) {
	log.Debugf("path: %s", conf.HttpStaticPath()+"/favicon.ico")
	// Favicon
	router.StaticFile("/favicon.ico", conf.HttpStaticPath()+"/favicon.ico")

	// Static assets like js and css files
	router.Static("/static", conf.HttpStaticPath())

	// JSON-REST API Version 1
	v1 := router.Group("/v1")
	{
		api.GetLocation(v1, conf)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "operational"})
	})
}
