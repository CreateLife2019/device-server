package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/device-server/controller"
	"github.com/device-server/domain/constants"
	"github.com/device-server/global"
	"github.com/device-server/internal/repository/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Register() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(global.Cfg.ServerCfg.Key),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: constants.KeyId,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.Account); ok {
				return jwt.MapClaims{
					constants.KeyId:       v.UserName,
					constants.KeyPassword: v.Password,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &entity.Account{
				UserName: claims[constants.KeyId].(string),
				Password: claims[constants.KeyPassword].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			user := data.(*entity.Account)
			return controller.GetInstance().AccountService().CheckUser(user.UserName, user.Password)
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
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		logrus.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}
