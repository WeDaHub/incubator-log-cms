package service

import (
	"log-api/common/anc"
	"log-api/common/result"
	"log-api/model"
)

type rule struct {

}

func (this *rule) AddRule(ctx anc.Context, level int, name string) *result.R {
	return result.CR()

}


func (this *rule) UpdateRule(ctx anc.Context, id int, level int, name string) *result.R {
	if err := model.DB().Model(&model.Rule{}).Where("id = ?", id).UpdateColumns(model.Rule{Level: level, Name: name}).Error; nil != err {
		return result.CR().ERROR(err.Error())
	} else {
		return result.CR().Succeed(nil)
	}
}

func (this *rule) DeleteRule(ctx anc.Context, id int) *result.R {
	if err := model.DB().Delete(&model.Rule{}, "id = ?", id).Error; nil != err {
		return result.CR().ERROR(err.Error())
	} else {
		return result.CR().Succeed(nil)
	}
}
func (this *rule) GetRule(ctx anc.Context, id int) *result.R {
	var ruleData model.Rule
	if err := model.DB().First(&ruleData, "id = ?", id).Error; nil != err {
		return result.CR().Error2(1, err.Error())
	} else {
		return result.CR().Succeed(ruleData)
	}
}
func (this *rule) GetRuleList(ctx anc.Context, page int, size int) *result.R {
	var data struct{
		Count		int64 			`json:"count"`
		Page		int				`json:"page"`
		Size		int 			`json:"size"`
		Data		[]model.Rule	`json:"datas"`
	}

	if err := model.DB().Model(&model.Rule{}).Count(&data.Count).Error; nil != err {
		return result.CR().ERROR(err.Error())
	}

	if err := model.DB().Order("created_at desc").Limit(size).Offset((page - 1) * size).Find(&data.Data).Error; nil != err {
		return result.CR().ERROR(err.Error())
	}
	return result.CR().Succeed(data)
}