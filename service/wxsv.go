package service

import (
	conf "ims/common/config"
	"ims/model"
	"log"

	"gorm.io/gorm"
)

type wxService struct{}

var WxService = wxService{}

func (s *wxService) Login(id string) (wxuser model.WxUser, err error) {

	// err = conf.HrmsDB(c).Find(&user).Error
	// return user, err

	var hrmsDB *gorm.DB
	var ok bool
	dbName := conf.GlobalConf.DbName
	if hrmsDB, ok = conf.DbMapper["imdb"]; !ok {
		log.Printf("[Login err, 无法获取到db, name = %v]", dbName)

		return
	}

	err = hrmsDB.Where("id = ", id).First(&wxuser).Error

	return wxuser, err

}

func (s *wxService) Register(wxuser model.WxUser) (err error) {

	var hrmsDB *gorm.DB
	var ok bool
	dbName := conf.GlobalConf.DbName
	if hrmsDB, ok = conf.DbMapper[dbName]; !ok {
		log.Printf("[Login err, 无法获取db名称, name = %v]", dbName)

		return nil
	}

	if err = hrmsDB.Create(&wxuser).Error; err != nil {
		log.Printf("questService Add err = %v", err)
	}

	return err
}
