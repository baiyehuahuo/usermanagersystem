package regedit

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"usermanagersystem/consts"
	"usermanagersystem/model"
	"usermanagersystem/utils/database"

	"github.com/gin-gonic/gin"
)

type regeditManagerImpl struct {
}

func (regeditManager regeditManagerImpl) UserRegedit(c *gin.Context) {
	user := model.User{
		Account:  c.Query("account"),
		Password: fmt.Sprintf("%x", md5.Sum([]byte(c.Query("password")))),
	}
	if err := database.DB.Create(&user).Error; err != nil {
		log.Print(err, user)
		c.JSON(http.StatusInternalServerError, consts.RegeditFail)
		return
	}
	c.JSON(http.StatusOK, consts.RegeditSuccess)
}
