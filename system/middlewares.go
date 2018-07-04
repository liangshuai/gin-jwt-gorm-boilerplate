package system

import (
	"net/http"
	"github.com/Sirupsen/logrus"
	"github.com/liangshuai/retu/model"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ContextData() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if uID := session.Get("UserID"); uID != nil {
			user := model.User{}
			model.GetDB().First(&user, uID)
			if user.ID != 0 {
				c.Set("User", &user)
			}
		}
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("User"); user != nil {
			c.Next()
		} else {
			logrus.Warnf("User is not authorized to visit %s", c.Request.RequestURI)
			c.HTML(http.StatusForbidden, "errors/403", nil)
			c.Abort()
		}
	}
}