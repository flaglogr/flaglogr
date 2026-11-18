package controller

import (
	"github.com/TwiN/logr"
	"github.com/flaglogr/flaglogr/docs"
	"github.com/gin-gonic/gin"
	"github.com/janishutz/oidclogin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (c *Controller) AddRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	projects := api.Group("/projects", oidclogin.EnsureLogin(false))
	projects.GET("/get", c.GetProjects)
	projects.PUT("/create", c.CreateProject)
	projects.PATCH("/update/:id", c.UpdateProject)
	projects.DELETE("/delete/:id", c.DeleteProject)

	// ── Healthcheck ──────────────────────────────────────────────────
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": true}) })

	// ── Swagger ──────────────────────────────────────────────────────
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	docs.SwaggerInfo.BasePath = "/"

	logr.Info("\n\n\n==> API documentation available at http://localhost:8080/swagger/index.html <==\n\n\n")
}
