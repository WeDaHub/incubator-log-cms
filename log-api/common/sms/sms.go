package sms

import "log-api/common/result"

type ISmsSender interface {
	SendCode(mobile string, code string) *result.R
}

type smsSender struct {

}

func (this *smsSender) SendCode(mobile string, code string) *result.R{
	return result.CR();
}

var sender *smsSender

func Sender() ISmsSender {
	return sender;
}

func init() {
	sender = &smsSender{}
}
