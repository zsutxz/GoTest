package handler

import (
	"hrms/model"
	"hrms/service"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddQuestion(c *gin.Context) {
	// 参数绑定
	var dto model.Question
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Printf("[AddQuestion] err = %v", err)
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
		log.Printf("[CreateExample] err = %v", err)
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

// 提交题目
func AddQuestion11(c *gin.Context) {
	var form model.Question

	form.Quest = c.PostForm("question")
	form.ChoiceA = c.DefaultPostForm("choicea", "aaa")
	form.ChoiceB = c.DefaultPostForm("choiceb", "bbb")
	form.ChoiceC = c.DefaultPostForm("choicec", "ccc")
	form.ChoiceD = c.DefaultPostForm("choiced", "ddd")

	//答案
	num, err := strconv.Atoi(c.PostForm("answer"))
	if err != nil {
		num = 100
	}
	form.Answer = int8(num)

	//难度
	num, err = strconv.Atoi(c.PostForm("dif"))
	if err != nil {
		num = 100
	}
	form.Dif = int8(num)
	log.Printf("form,v=%v", form)

	// service.QuestService.Add(form)

}
