package controller

import (
	"github.com/TwiN/logr"
	"github.com/gin-gonic/gin"
	"github.com/janishutz/oidclogin"
)

// IMPORTANT: Also load gin.BasicAuth, see https://gin-gonic.com/en/docs/middleware/using-basicauth-middleware/
// FIXME: Probably should be used on the login route to then store in session
func CheckBasicAuth(c *gin.Context) {
	userRaw, ok := c.Get(gin.AuthUserKey)
	if !ok {
		c.AbortWithStatus(401)
	} else {
		user, ok := userRaw.(string)
		if !ok {
			c.AbortWithStatus(401)
		} else {
			c.Set("uid", user)
		}
		c.Next()
	}
}

func SetUpOidcAuth(r *gin.Engine, url string, prod bool) {
	oidclogin.Configure(r, url, url, false, func(c *gin.Context) {}, prod)
}

func CheckOidcAuth() func(c *gin.Context) {
	return oidclogin.EnsureLogin(true)
}

type GroupMapping map[string]string

func CheckTrustedHeaderAuth(secret string, usernameClaimHeader string, userGroupsClaimHeader string, groupMappings GroupMapping) func(c *gin.Context) {
	// TODO: Probably want to make group mappings editable via UI (and claims as well)
	return func(c *gin.Context) {
		// TODO: Decide on what happens with user table? Probably nothing
		if c.GetHeader("X-Proxy-Secret") == secret {
			username := c.GetHeader(usernameClaimHeader)
			// groups := c.GetHeader(userGroupsClaimHeader)
			c.Set("uid", username)
			c.Next()
		} else {
			logr.Warn("[Trusted Header Auth] Proxy secret did not match. Make sure to set the X-Proxy-Secret header!")
			c.AbortWithStatus(401)
		}
	}
}
