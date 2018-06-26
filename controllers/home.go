package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/liangshuai/retu/model"
)

func HomeGet(c *gin.Context) {
	user := model.User{Name: "liangshuai", Password: "ls123456", Email: "liangshuais@qq.com", AvatarURL: "https://avatars3.githubusercontent.com/u/7534777?s=460&v=4"}
	model.GetDB().Create(&user)
	c.String(http.StatusOK, "Home");
}