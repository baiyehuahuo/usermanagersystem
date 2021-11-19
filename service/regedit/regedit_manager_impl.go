package regedit

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
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
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "success")
}
