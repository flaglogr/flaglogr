package main

import (
	"os"
	"time"

	"github.com/TwiN/logr"
	"github.com/flaglogr/flaglogr/config"
	"github.com/flaglogr/flaglogr/controller"
	"github.com/flaglogr/flaglogr/model"
	"github.com/flaglogr/flaglogr/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/janishutz/oidclogin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
)

// @title flaglogr backend
// @version         0.1.0
// @description.markdown
// @termsOfService  https://janishutz.com/legal/tos/
//
// @contact.name   Support
// @contact.url    https://github.com/flaglogr/flaglogr#readme
// @contact.email  support@janishutz.com
//
// @license.name  GNU Affero General Public License 3.0
// @license.url   https://www.gnu.org/licenses/agpl-3.0.en.html
// @schemes 	  http https
//
// @tag.name projects
// @tag.description Asset CRUD (assets are any type of file the user can upload)
//
// @BasePath  /
func main() {
	logr.SetThreshold(logr.LevelDebug)
	logr.Debug("Starting...")
	r := gin.Default()

	conf := config.Load()
	prod := os.Getenv("PROD") == "true"

	// Load templates
	r.LoadHTMLGlob("templates/*.tmpl")

	// ┌                                               ┐
	// │            Security configuration             │
	// └                                               ┘
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{conf.AppUrl},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(util.DefaultHeaders)

	r.Use(util.RateLimiter())
	// r.SetTrustedProxies(conf.Urls.TrustedProxies)

	// ┌                                               ┐
	// │                    Session                    │
	// └                                               ┘
	secret, ok := os.LookupEnv("SESSION_SECRET")
	if !ok {
		logr.Warn("Missing secret (SESSION_SECRET is unset), using unsafe default")
		secret = "secret"
	}
	store := memstore.NewStore([]byte(secret))
	r.Use(sessions.Sessions("session", store))

	// ┌                                               ┐
	// │               Routes, DB, Login               │
	// └                                               ┘
	db := model.Init(prod, &conf)
	oidclogin.ConfigureFull(r, conf.AppUrl, conf.AppUrl, func(userid, name, email string, c *gin.Context) {
		uid := db.UpdateOrCreateUser(userid, name, email)
		ses := sessions.Default(c)
		ses.Set("userid", uid)
		ses.Set("username", name)
		ses.Set("email", email)
		ses.Save()
	}, true, nil, prod)

	// Routes
	c := controller.NewController(db)

	r.GET("/", func(c *gin.Context) {
		c.Redirect(307, "/swagger/index.html")
	})

	c.AddRoutes(r)

	r.Run()
}
