package main

import (
	"log"
	"net/http"
	"usermanagersystem/consts"
	"usermanagersystem/service/logincontrol"
	"usermanagersystem/service/usercontrol"

	"github.com/gin-gonic/gin"
)

type handleManager struct {
	lm logincontrol.LoginController
	um usercontrol.UserController
}

func (handle *handleManager) UserLogin(c *gin.Context) {
	if err := handle.lm.UserLogin(c); err != nil {
		log.Printf("Login Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.LoginFail)
		return
	}
	c.JSON(http.StatusOK, consts.LoginSuccess)
}

func (handle *handleManager) UserRegedit(c *gin.Context) {
	if err := handle.lm.UserRegedit(c); err != nil {
		log.Printf("Regedit Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.RegeditFail)
		return
	}
	c.JSON(http.StatusOK, consts.RegeditSuccess)
}

func (handle *handleManager) FileUpload(c *gin.Context) {
	if err := handle.um.FileUpload(c); err != nil {
		log.Printf("FileUpload Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.UploadFail)
		return
	}
	c.JSON(http.StatusOK, consts.UploadSuccess)
}

func (handle *handleManager) ModifyPassword(c *gin.Context) {
	if err := handle.um.ModifyPassword(c); err != nil {
		log.Printf("ModifyPassword Fail: %v", err)
		c.JSON(http.StatusInternalServerError, consts.ModifyPasswordFail)
		return
	}
	c.JSON(http.StatusOK, consts.ModifyPasswordSuccess)
}
