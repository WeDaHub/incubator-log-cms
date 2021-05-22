package service

import (
	"log-api/common/anc"
	"log-api/common/result"
)

type IRule interface {
	AddRule(ctx anc.Context, level int, name string) *result.R
	UpdateRule(ctx anc.Context, id int, level int, name string) *result.R
	DeleteRule(ctx anc.Context, id int) *result.R
	GetRule(ctx anc.Context, id int) *result.R
	GetRuleList(ctx anc.Context, page int, size int) *result.R
}

func GetRuleService() IRule {
	return &rule{}
}
