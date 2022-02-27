package html_control

import (
	"net/http"
	"usermanagersystem/utils"

	"github.com/gin-gonic/gin"
)

type htmlControllerImpl struct {
	rc utils.RedisController
}

func (htmlController *htmlControllerImpl) ToLogin(c *gin.Context) {
	// if cookie, err := c.Cookie(consts.CookieNameOfUser); err == nil {
	// 	log.Println("get cookie:", cookie)
	// 	if user, _ := htmlController.rc.Get(consts.RedisCookieHashPrefix + cookie); user != "" {
	// 		log.Printf("user [%s] logined", user)
	// 		c.HTML(http.StatusOK, "index.html", "")
	// 		return
	// 	}
	// 	log.Println("cookie is timeout")
	// }
	c.HTML(http.StatusOK, "index.html", "")
}
