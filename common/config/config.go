package conf

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 全局配置文件
var GlobalConf *Config

// 分公司数据库映射表
var DbMapper = make(map[string]*gorm.DB)

// 默认DB，不作为业务使用
var DefaultDb *gorm.DB

type Gin struct {
	Port int64 `json:"port"`
}

// 解析cookie中的分公司Id，获取对应数据库实例
func HrmsDB(c *gin.Context) *gorm.DB {
	cookie, err := c.Cookie("user_cookie")
	if err != nil || cookie == "" {
		c.HTML(http.StatusOK, "login.html", nil)
		return nil
	}
	branchId := strings.Split(cookie, "_")[2]
	dbName := fmt.Sprintf("hrms_%v", branchId)
	if db, ok := DbMapper[dbName]; ok {
		return db
	}
	c.HTML(http.StatusOK, "login.html", nil)
	return nil
}

type Db struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	DbName   string `json:"dbNname"`
}

type Mongo struct {
	IP      string `json:"ip"`
	Port    int64  `json:"port"`
	Dataset string `json:"dataset"`
}

type Weixin struct {
	Site     string `json:"site"`
	Appid    string `json:"appid"`
	Secret   string `json:"secret"`
	Httptail string `json:"httptail"`
}

var MongoClient *qmgo.Client

type Config struct {
	Gin    `json:"gin"`
	Db     `json:"db"`
	Mongo  `json:"mongo"`
	Weixin `json:"weixin"`
}

func InitConfig() error {
	config := &Config{}
	vip := viper.New()
	vip.AddConfigPath("./config")
	vip.SetConfigType("yaml")
	// 环境判断
	env := os.Getenv("HRMS_ENV")
	if env == "" || env == "dev" {
		// 开发环境
		vip.SetConfigName("config-dev")
	}
	if env == "prod" {
		// 生产环境
		vip.SetConfigName("config-prod")
	}
	err := vip.ReadInConfig()
	if err != nil {
		log.Printf("[config.Init] err = %v", err)
		return err
	}
	if err := vip.Unmarshal(config); err != nil {
		log.Printf("[config.Init] err = %v", err)
		return err
	}
	log.Printf("[config.Init] 初始化配置成功,config=%v", config)
	GlobalConf = config
	return nil
}
