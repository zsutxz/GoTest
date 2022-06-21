package handler

import (
	"hrms/model"
	"hrms/service"
	"log"

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
