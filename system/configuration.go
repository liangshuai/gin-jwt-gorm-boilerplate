package system

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

type Configs struct{
	Debug		Config
	Release 	Config
	Test		Config
}

type Config struct {
	Public			string `json:"public"`
	Domain			string `json:"domain"`
	SessionSecret	string `json:"session_secret"`
	PidFile			string `json:"pid_file"`
	SignupEnabled	bool   `json:"signup_enabled"`
	Secret			string `json:"secret"`
	Realm			string `json:"realm"`
	TTL				int	`json:"ttl"`
	Database		DatabaseConfig
}

type DatabaseConfig struct {
	Host		string
	Port		int
	Name		string
	User		string
	Password	string
}
var config *Config

func LoadConfig() {
	data, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		panic(err)
	}
	configs := &Configs{}
	err = json.Unmarshal(data, configs)
	if  err != nil {
		panic(err)
	}
	switch gin.Mode() {
	case gin.DebugMode:
		config = &configs.Debug
	case gin.ReleaseMode:
		config = &configs.Release
	case gin.TestMode:
		config = &configs.Test
	default:
		panic(fmt.Sprintf("Unknown gin mode %s", gin.Mode()))
	}
	if !path.IsAbs(config.Public) {
		workingDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		config.Public = path.Join(workingDir, config.Public)
	}
}

func GetConfig() *Config{
	return config
}

func PublicPath() string {
	return config.Public
}

func UploadsPath() string {
	return path.Join(config.Public, "uploads")
}

func GetConnectionString() string {
	return fmt.Sprintf("host=%s  port=%d user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name)
}
