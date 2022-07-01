package handler

import (
	conf "ims/common/config"
	"ims/model"
	"ims/service"
	"log"
	"net/http"
	"strconv"

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

	// 分页
	start, limit := service.AcceptPage(c)
	code := 2000
	quest_id := c.Param("id")

	var qustions []model.Mc_Question
	if quest_id == "all" {
		// 查询全部
		conf.HrmsDB(c).Find(&qustions)

		// 查询全部
		if start == -1 && start == -1 {
			conf.HrmsDB(c).Find(&qustions)
		} else {
			conf.HrmsDB(c).Offset(start).Limit(limit).Find(&qustions)
		}

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
		// log.Printf("question: %v", qustions)
		return
	}

	conf.HrmsDB(c).Where("id = ? ", quest_id).Find(&qustions)
	if len(qustions) == 0 {
		// 不存在
		code = 2001
	}
	// log.Printf("question: %v", qustions)
	total = int64(len(qustions))
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"total":  total,
		"msg":    qustions,
	})
}

func EditQuestion(c *gin.Context) {
	// 参数绑定
	var dto model.Mc_Question
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[EditQuestion] err = %v", err)
		c.JSON(200, gin.H{
			"status": 5001,
			"result": err.Error(),
		})
		return
	}

	conf.HrmsDB(c).Model(&model.Mc_Question{}).Where("id = ?", strconv.Itoa(int(dto.ID))).
		Updates(&dto)
	c.JSON(200, gin.H{
		"status": 2000,
	})
}

func DelQuestion(c *gin.Context) {
	quest_id := c.Param("id")
	log.Printf("DelQuestion,quest_id=%v", quest_id)
	// 业务处理
	num, err := strconv.Atoi(quest_id)
	service.QuestService.DelById(c, num)
	if err != nil {
		log.Printf("[DelExample] err = %v", err)
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
