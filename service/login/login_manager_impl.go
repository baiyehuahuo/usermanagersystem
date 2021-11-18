package login

import (
	"log"
	"net/http"
	"usermanagersystem/model"
	"usermanagersystem/utils/database"

	"github.com/gin-gonic/gin"
)

type loginManagerImpl struct {
}

func (login *loginManagerImpl) UserLogin(c *gin.Context) {
	db := database.DB
	user := model.User{Account: "8208180115"}
	if err := db.Take(&user).Error; err != nil {
		log.Print(err, user)
		c.JSON(http.StatusInternalServerError, "LoginFail")
		return
	}
	c.JSON(http.StatusOK, user)
}
