package router

import (
	"github.com/JumpSama/aug-blog/handler/sd"
	"github.com/JumpSama/aug-blog/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 中间件
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	scvd := g.Group("/sd")
	{
		scvd.GET("/health", sd.HealthCheck)
		scvd.GET("/disk", sd.DiskCheck)
		scvd.GET("/cpu", sd.CPUCheck)
		scvd.GET("/ram", sd.RAMCheck)
	}

	return g
}
