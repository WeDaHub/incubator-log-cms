package service

import (
	"log-api/common/anc"
	"log-api/common/result"
)

type ILog interface {
	UpdateLog(ctx anc.Context, date string, log string) *result.R
	DeleteLog(ctx anc.Context, date string) *result.R
	GetLog(ctx anc.Context, date string) *result.R
	GetLogList(ctx anc.Context, page int, size int) *result.R
}

func GetLogService() ILog {
	return &log{}
}