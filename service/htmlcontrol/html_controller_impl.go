package htmlcontrol

import (
	"log"
	"net/http"
	"usermanagersystem/consts"
	"usermanagersystem/utils/rediscontrol"

	"github.com/gin-gonic/gin"
)

type htmlControllerImpl struct {
	rc rediscontrol.RedisController
}

func (htmlController *htmlControllerImpl) ToLogin(c *gin.Context) {
	if cookie, err := c.Cookie(consts.UserCookieName); err == nil {
		log.Println("get cookie:", cookie)
		if user, _ := htmlController.rc.Get(consts.RedisCookieHash + cookie); user != "" {
			log.Printf("user [%s] logined", user)
			c.HTML(http.StatusOK, "UserManage.html", "")
			return
		}
		log.Println("cookie is timeout")
	}
	c.HTML(http.StatusOK, "index.html", "")
}

func (htmlController *htmlControllerImpl) ToUserManage(c *gin.Context) {
	c.HTML(http.StatusOK, "UserManage.html", "")
}