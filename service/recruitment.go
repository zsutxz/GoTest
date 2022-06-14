package service

import (
	conf "hrms/common/config"
	"hrms/common/util"
	"hrms/model"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateRecruitment(c *gin.Context, dto *model.RecruitmentCreateDTO) error {
	var recruitmentRecord model.Recruitment
	util.Transfer(&dto, &recruitmentRecord)
	recruitmentRecord.RecruitmentId = RandomID("recruitment")
	if err := conf.HrmsDB(c).Create(&recruitmentRecord).Error; err != nil {
		log.Printf("CreateRecruitment err = %v", err)
		return err
	}
	return nil
}

func DelRecruitmentByRecruitmentId(c *gin.Context, recruitmentId string) error {
	if err := conf.HrmsDB(c).Where("recruitment_id = ?", recruitmentId).Delete(&model.Recruitment{}).
		Error; err != nil {
		log.Printf("DelRecruitmentByRecruitmentId err = %v", err)
		return err
	}
	return nil
}

func UpdateRecruitmentById(c *gin.Context, dto *model.RecruitmentEditDTO) error {
	var recruitment model.Recruitment
	util.Transfer(&dto, &recruitment)
	if err := conf.HrmsDB(c).Model(&model.Recruitment{}).Where("id = ?", recruitment.ID).
		Updates(&recruitment).Error; err != nil {
		log.Printf("UpdateRecruitmentById err = %v", err)
		return err
	}
	return nil
}

func GetRecruitmentByJobName(c *gin.Context, jobName string, start int, limit int) ([]*model.Recruitment, int64, error) {
	var records []*model.Recruitment
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if jobName != "all" {
			err = conf.HrmsDB(c).Where("job_name like ?", "%"+jobName+"%").Find(&records).Error
		} else {
			err = conf.HrmsDB(c).Find(&records).Error
		}

	} else {
		// 加分页
		if jobName != "all" {
			err = conf.HrmsDB(c).Where("job_name like ?", "%"+jobName+"%").Offset(start).Limit(limit).Find(&records).Error
		} else {
			err = conf.HrmsDB(c).Offset(start).Limit(limit).Find(&records).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	conf.HrmsDB(c).Model(&model.Recruitment{}).Count(&total)
	if jobName != "all" {
		total = int64(len(records))
	}
	return records, total, nil
}
