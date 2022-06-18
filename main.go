package main

import (
	"context"
	"fmt"
	conf "hrms/common/config"
	"hrms/router"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitConfig() error {
	config := &conf.Config{}
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
	conf.HrmsConf = config
	return nil
}

func InitGin() error {
	server := gin.Default()
	// 静态资源及模板配置
	htmlInit(server)
	// 初始化路由
	router.Init(server)
	err := server.Run(fmt.Sprintf(":%v", conf.HrmsConf.Gin.Port))
	if err != nil {
		log.Printf("[InitGin] err = %v", err)
	}
	log.Printf("[InitGin] success")
	return err
}

func htmlInit(server *gin.Engine) {
	// 静态资源
	server.StaticFS("/static", http.Dir("./static"))
	server.StaticFS("/views", http.Dir("./views"))
	// HTML模板加载
	server.LoadHTMLGlob("views/*")
	// 404页面
	server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})
}

func InitGorm() error {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// 对每个分公司数据库进行连接
	dbNames := conf.HrmsConf.Db.DbName
	dbNameList := strings.Split(dbNames, ",")
	for index, dbName := range dbNameList {
		dsn := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			conf.HrmsConf.Db.User,
			conf.HrmsConf.Db.Password,
			conf.HrmsConf.Db.Host,
			conf.HrmsConf.Db.Port,
			dbName,
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				// 全局禁止表名复数
				SingularTable: true,
			},
			// 日志等级
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Printf("[InitGorm] err = %v", err)
			return err
		}
		// 添加到映射表中
		conf.DbMapper[dbName] = db
		// 第一个是默认DB，用以启动程序选择分公司
		if index == 0 {
			conf.DefaultDb = db
		}
		log.Printf("[InitGorm] 分公司数据库%v注册成功", dbName)
	}
	//fmt.Println(conf.DbMapper["hrms_C001"])
	log.Printf("[InitGorm] success")
	return nil
}

func InitMongo() error {
	mongo := conf.HrmsConf.Mongo
	var err error
	conf.MongoClient, err = qmgo.NewClient(context.Background(), &qmgo.Config{
		Uri:      fmt.Sprintf("mongodb://%v:%v", mongo.IP, mongo.Port),
		Database: mongo.Dataset,
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := InitConfig(); err != nil {
		panic(err)
	}
	if err := InitGorm(); err != nil {
		panic(err)
	}
	if err := InitGin(); err != nil {
		panic(err)
	}
	if err := InitMongo(); err != nil {
		panic(err)
	}
}
