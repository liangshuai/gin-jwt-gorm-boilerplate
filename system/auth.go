package system

import (
	"time"
	"github.com/gin-gonic/gin"
	"gopkg.in/appleboy/gin-jwt.v2"
	"github.com/liangshuai/retu/model"
)

var JWT *jwt.GinJWTMiddleware
var jwtRealm, jwtSecret string
var jwtTTL int

func Init(realm string, secret string, ttl int) {
	jwtRealm = realm
	jwtSecret = secret
	jwtTTL = ttl
}

func NewAuthMiddleware(min_level int) *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      jwtRealm,
		Key:        []byte(jwtSecret),
		Timeout:    time.Hour * time.Duration(jwtTTL),
		MaxRefresh: time.Hour * time.Duration(jwtTTL),
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			user, err := model.GetUser(username, password)
			if err != nil {
				return username, false
			}

			verified := user.Verify(password)

			return username, verified
		},
		Authorizator: func(username string, c *gin.Context) bool {
			correlation, err := model.GetUserRoleByUsername(username)
			if (err != nil) {
				return false
			}
			if (correlation.ID2 >= uint16(min_level)) {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		TokenLookup: "header:Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
	}
}