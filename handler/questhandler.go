package handler

import (
	conf "hrms/common/config"
	"hrms/model"
	"hrms/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddQuestion(c *gin.Context) {
	// 参数绑定
	var dto model.Mc_Question
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[AddQuestion11] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}
	log.Printf("AddQuestion,dto=%v", dto)
	// 业务处理
	err := service.QuestService.Add(c, dto)
	if err != nil {
		c.JSON(200, gin.H{
			"status": 5002,
			"result": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func SearchQuestion(c *gin.Context) {
	var total int64 = 1

	code := 2000
	quest_id := c.Param("id")
	log.Printf("SearchQuestion,id=%v", quest_id)

	var qustions []model.Mc_Question
	if quest_id == "all" {
		// 查询全部
		conf.HrmsDB(c).Find(&qustions)

		if len(qustions) == 0 {
			// 不存在
			code = 2001
		}
		// 总记录数
		conf.HrmsDB(c).Model(&model.Staff{}).Count(&total)
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"total":  total,
			"msg":    qustions,
		})
		return
	}
}
