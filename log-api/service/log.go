package service

import (
	"log-api/common/anc"
	"log-api/common/result"
	"log-api/model"
)

type log struct {

}

func (this *log) UpdateLog(ctx anc.Context, date string, log string) *result.R {
	var logData model.Log
	err := model.DB().First(&logData, "user_id = ? and date_time = ?", ctx.User.ID, date).Error
	if logData.ID > 0 {
		err = model.DB().Model(&model.Log{}).Where("user_id = ? and date_time = ?").Update("content", log).Error
	} else {
		logData.Content = log
		logData.DateTime = date
		logData.UserId = ctx.User.ID
		err = model.DB().Create(&logData).Error
	}
	return result.CR().Succeed(err)
}

func (this *log) DeleteLog(ctx anc.Context, date string) *result.R{
	err := model.DB().Delete(&model.Log{}, "user_id = ? and date_time = ?", ctx.User.ID, date).Error
	if nil != err {
		return result.CR().Error2(1, err.Error())
	}
	return result.CR().Succeed(nil)
}

func (this *log) GetLog(ctx anc.Context, date string) *result.R{
	var logData model.Log
	if err := model.DB().First(&logData, "user_id = ? and date_time = ?", ctx.User.ID, date).Error; nil != err {
		return result.CR().Error2(1, err.Error())
	} else {
		return result.CR().Succeed(logData)
	}
}

func (this *log) GetLogList(ctx anc.Context, page int, size int) *result.R{

	var data struct{
		Count		int64 		`json:"count"`
		Page		int			`json:"page"`
		Size		int 		`json:"size"`
		Data		[]model.Log	`json:"datas"`
	}

	if err := model.DB().Model(&model.Log{}).Where("user_id = ?", ctx.User.ID).Count(&data.Count).Error; nil != err {
		return result.CR().ERROR(err.Error())
	}

	if err := model.DB().Where("user_id = ?", ctx.User.ID).Order("date_time desc").Limit(size).Offset((page - 1) * size).Find(&data.Data).Error; nil != err {
		return result.CR().ERROR(err.Error())
	}

	return result.CR().Succeed(data)
}
