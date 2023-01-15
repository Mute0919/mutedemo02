package api

import (
	"WestOnline2/serivice"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegisterService serivice.UserService

	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
	}
}

func UserLogin(c *gin.Context) {
	var userLogin serivice.UserService

	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, gin.H{
			"err": err.Error(),
		})
	}
}
