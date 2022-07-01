package main

import (
	"context"
	"fmt"
	conf "ims/common/config"
	"ims/common/db"
	"ims/router"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/qmgo"
)

func htmlInit(server *gin.Engine) {
	// 静态资源
	server.StaticFS("/static", http.Dir("./static"))
	server.StaticFS("/views", http.Dir("./views"))
	// HTML模板加载
	server.LoadHTMLGlob("./views/*")
	// 404页面
	server.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})
}

func InitGin() error {
	server := gin.Default()
	// 静态资源及模板配置
	htmlInit(server)
	// 初始化路由
	router.Init(server)
	err := server.Run(fmt.Sprintf(":%v", conf.GlobalConf.Gin.Port))
	if err != nil {
		log.Printf("[InitGin] err = %v", err)
	}
	log.Printf("[InitGin] success")
	return err
}

func InitMongo() error {
	mongo := conf.GlobalConf.Mongo
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
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}
	if err := db.InitGorm(); err != nil {
		panic(err)
	}
	if err := InitGin(); err != nil {
		panic(err)
	}
	if err := InitMongo(); err != nil {
		panic(err)
	}
}
