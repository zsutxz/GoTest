package service

import (
	"errors"
	conf "ims/common/config"
	"ims/model"
	"log"

	"github.com/gin-gonic/gin"
)

type questService struct {
}

var QuestService = new(questService)

// Add 添加
func (mcquestServ *questService) Add(c *gin.Context, params model.Mc_Question) (err error) {
	var question = model.Mc_Question{Question: params.Question, ChoiceA: params.ChoiceA, ChoiceB: params.ChoiceB, ChoiceC: params.ChoiceC,
		ChoiceD: params.ChoiceD, Answer: params.Answer, Kind: params.Kind, Score: params.Score, Dif: params.Dif, Stat: params.Stat, Describe: params.Describe, Owner: 10001}

	if err := conf.HrmsDB(c).Create(&question).Error; err != nil {
		log.Printf("questService Add err = %v", err)
		return err
	}

	return nil
}

// Get 获取题目信息
func (mcquestServ *questService) Get(id uint) (err error, quest model.Mc_Question) {

	// err = global.App.DB.First(&quest, id).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return nil, quest
}

func (mcquestServ *questService) DelById(c *gin.Context, id int) error {
	if err := conf.HrmsDB(c).Where("id = ?", id).Delete(&model.Mc_Question{}).Error; err != nil {
		log.Printf("DelById err = %v", err)
		return err
	}
	return nil
}
