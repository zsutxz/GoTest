package service

import (
	"errors"
	"fmt"
	conf "hrms/common/config"
	"hrms/common/util"
	"hrms/model"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateSalary(c *gin.Context, dto *model.SalaryCreateDTO) error {
	var total int64
	conf.HrmsDB(c).Model(&model.Salary{}).Where("staff_id = ? and deleted_at is null", dto.StaffId).Count(&total)
	if total != 0 {
		return errors.New(fmt.Sprintf("该员工薪资数据已经存在"))
	}
	var salary model.Salary
	util.Transfer(&dto, &salary)
	salary.SalaryId = RandomID("salary")
	if err := conf.HrmsDB(c).Create(&salary).Error; err != nil {
		log.Printf("CreateSalary err = %v", err)
		return err
	}
	return nil
}

func DelSalaryBySalaryId(c *gin.Context, salaryId string) error {
	if err := conf.HrmsDB(c).Where("salary_id = ?", salaryId).Delete(&model.Salary{}).
		Error; err != nil {
		log.Printf("DelSalaryBySalaryId err = %v", err)
		return err
	}
	return nil
}

func UpdateSalaryById(c *gin.Context, dto *model.SalaryEditDTO) error {
	var salary model.Salary
	util.Transfer(&dto, &salary)
	if err := conf.HrmsDB(c).Model(&model.Salary{}).Where("id = ?", salary.ID).
		Update("staff_id", salary.StaffId).
		Update("staff_name", salary.StaffName).
		Update("base", salary.Base).
		Update("subsidy", salary.Subsidy).
		Error; err != nil {
		log.Printf("UpdateSalaryById err = %v", err)
		return err
	}
	return nil
}

func GetSalaryByStaffId(c *gin.Context, staffId string, start int, limit int) ([]*model.Salary, int64, error) {
	var salarys []*model.Salary
	var err error
	if start == -1 && limit == -1 {
		// 不加分页
		if staffId != "all" {
			err = conf.HrmsDB(c).Where("staff_id = ?", staffId).Find(&salarys).Error
		} else {
			err = conf.HrmsDB(c).Find(&salarys).Error
		}

	} else {
		// 加分页
		if staffId != "all" {
			err = conf.HrmsDB(c).Where("staff_id = ?", staffId).Offset(start).Limit(limit).Find(&salarys).Error
		} else {
			err = conf.HrmsDB(c).Offset(start).Limit(limit).Find(&salarys).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}
	var total int64
	conf.HrmsDB(c).Model(&model.Salary{}).Count(&total)
	if staffId != "all" {
		total = int64(len(salarys))
	}
	return salarys, total, nil
}
