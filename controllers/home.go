package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/Sirupsen/logrus"
	"github.com/liangshuai/retu/model"
)

func Signup (c *gin.Context) {
	user := model.User{Name: "liangshuai", Password: "ls123456", Email: "liangshuais@qq.com", AvatarURL: "https://avatars3.githubusercontent.com/u/7534777?s=460&v=4"}
	tx := model.GetDB().Begin()
	if tx.Error != nil {
		logrus.Info("555999")
		return
	}
	
	if err := tx.Create(&user).Error; err != nil {
		logrus.Info("55566")
		tx.Rollback()
		return
	}
	if err := tx.Create(&model.Correlation{Type: model.CorrelationUserRole, ID1: user.ID, ID2: model.UserRoleAdmin}).Error; err != nil {
		logrus.Info("555")
		tx.Rollback()
		return
	}
	err := tx.Commit().Error
	logrus.Infof("%v", err)
	logrus.Infof("---------")
	c.String(http.StatusOK, "SignIn");
}

func HomeGet(c *gin.Context) {
	c.String(http.StatusOK, "Home Get");
}

func HomeGetAdmin(c *gin.Context) {
	c.String(http.StatusOK, "Home Admin");
}

func HomeGetSudo(c *gin.Context) {
	c.String(http.StatusOK, "Home Sudo");
}