package model

import "strconv"

type Mc_Question struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Question string `json:"quest" gorm:"size:512;not null;comment:问题"`
	ChoiceA  string `json:"choicea" gorm:"size:256;not null;index;comment:选项"`
	ChoiceB  string `json:"choiceb" gorm:"size:256;not null;index;comment:"`
	ChoiceC  string `json:"choicec" gorm:"size:256;not null;index;comment:"`
	ChoiceD  string `json:"choiced" gorm:"size:256;not null;index;comment:"`
	Answer   int8   `json:"answer,string" gormobilem:"not null;default:'';comment:答案"`
	Score    int8   `json:"score,string" gormobilem:"not null;default:'';comment:"`
	Grade    int8   `json:"grade,string" gormobilem:"not null;default:'';comment:等级"`
	Kind     int8   `json:"kind,string" gormobilem:"not null;default:'';comment:类型"`
	Dif      int8   `json:"dif,string" gormobilem:"not null;default:'';comment:"`
	Stat     int8   `json:"stat,string" gormobilem:"not null;default:'';comment:状态"`
	Owner    int16  `json:"owner" gormobilem:"not null;default:'';comment:"`
	Timestamps
	SoftDeletes
}

func (question Mc_Question) GetUid() string {
	return strconv.Itoa(int(question.ID))
}
