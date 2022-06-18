package handler

import (
	"hrms/model"
	"hrms/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//提交问题界面
func AddQuestion(c *gin.Context) {
	//解析模板文件
	c.HTML(http.StatusOK, "addquestion.tpml", gin.H{
		"title": "HTML 模板渲染样例",
		"body":  "这里是内容",
	})
}

// 表单提交
func PostQuest(c *gin.Context) {
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

	service.QuestService.Add(form)

}
