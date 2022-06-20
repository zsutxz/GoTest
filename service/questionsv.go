package service

import (
	"errors"
	conf "hrms/common/config"
	"hrms/model"
	"log"

	"github.com/gin-gonic/gin"
)

type questService struct {
}

var QuestService = new(questService)

// Add 添加
func (mcquestServ *questService) Add(c *gin.Context, params model.Question) (err error) {
	var quest = model.Question{Quest: params.Quest, ChoiceA: params.ChoiceA, ChoiceB: params.ChoiceB, ChoiceC: params.ChoiceC,
		ChoiceD: params.ChoiceD, Answer: params.Answer, Score: 2, Dif: params.Dif, Stat: 1, Owner: "coldeye"}
	// err = global.App.DB.Create(&quest).Error

	if err := conf.HrmsDB(c).Create(&quest).Error; err != nil {
		log.Printf("questService Add err = %v", err)
		return err
	}

	return nil
}

// Get 获取题目信息
func (mcquestServ *questService) Get(id uint) (err error, quest model.Question) {

	// err = global.App.DB.First(&quest, id).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return nil, quest
}
