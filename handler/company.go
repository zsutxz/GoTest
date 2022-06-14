package handler

import (
	conf "hrms/common/config"
	"hrms/model"
	"log"

	"github.com/gin-gonic/gin"
)

func BranchCompanyQuery(c *gin.Context) {
	var list []*model.BranchCompany
	if err := conf.DefaultDb.Find(&list).Error; err != nil {
		log.Println("GetBranchCompanyList err = %v", err)
		log.Printf("BranchCompanyQuery err = %v", err)
		c.JSON(200, gin.H{
			"status": 5000,
			"msg":    err,
		})
	}
	c.JSON(200, gin.H{
		"status": 2000,
		"msg":    list,
	})
}
