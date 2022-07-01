package db

import (
	"fmt"
	conf "hrms/common/config"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func Close() {
	// DB.Close()
}

func InitGorm() error {
	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// 对每个分公司数据库进行连接
	dbNames := conf.GlobalConf.Db.DbName
	dbNameList := strings.Split(dbNames, ",")
	for index, dbName := range dbNameList {
		dsn := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			conf.GlobalConf.Db.User,
			conf.GlobalConf.Db.Password,
			conf.GlobalConf.Db.Host,
			conf.GlobalConf.Db.Port,
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
		log.Printf("[InitGorm] 数据库%v注册成功", dbName)
	}
	//fmt.Println(conf.DbMapper["imdb"])
	log.Printf("[InitGorm] success")
	return nil
}
