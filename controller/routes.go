package controller

import (
	"github.com/TwiN/logr"
	"github.com/flaglogr/flaglogr/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (c *Controller) AddRoutes(r *gin.Engine) {
	// ── Healthcheck ──────────────────────────────────────────────────
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": true}) })

	// ── Swagger ──────────────────────────────────────────────────────
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/"

	logr.Info("\n\n\n==> API documentation available at http://localhost:8080/swagger/index.html <==\n\n\n")
}
