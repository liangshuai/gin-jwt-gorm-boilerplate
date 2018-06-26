package main

import (
	"os"
	"strconv"
	"syscall"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/Sirupsen/logrus"
	"github.com/liangshuai/retu/system"
)

func SavePid(address string) {
	config := system.GetConfig()
	pid := syscall.Getpid()

	pidFile, err := os.OpenFile(config.PidFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		logrus.Errorf("Error creating pid file: %v", err)
	}

	_, err = pidFile.Write([]byte(strconv.Itoa(pid)))
	if err != nil {
		logrus.Errorf("Error write to pid file: %v", err)
	}

	pidFile.Close()

	logrus.Infof("Listening and serving HTTP on", address)
	logrus.Infof("Actual pid is %d", pid)
}

func StartServer(engine *gin.Engine) {
	config := system.GetConfig()
	address := config.Domain
	server := endless.NewServer(address, engine)

	server.BeforeBegin = SavePid

	err := server.ListenAndServe()
	if err != nil {
		logrus.Errorf("Server interrupted: %v", err)
	}
}