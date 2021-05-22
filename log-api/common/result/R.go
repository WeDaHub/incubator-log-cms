package result

import "encoding/json"

type R struct {
	Code 		int			`json:"code"`
	Msg			string		`json:"msg"`
	Data		interface{}	`json:"data"`
}

func CR() *R {
	return &R{
		Code: 0,
		Msg: "Succeed",
		Data: nil,
	}
}

func (r *R) Succeed(data interface{}) *R {
	r.Code = 0
	r.Msg = "Succeed"
	r.Data = data
	return r
}

func (r *R) Error(code int) *R {
	r.Code = 1
	r.Msg = "Error"
	r.Data = nil
	return r
}

func (r *R) Error2(code int, msg string) *R{
	r.Code = code
	r.Msg = msg
	r.Data = nil
	return r
}

func (r*R) ERROR(msg string) *R{
	r.Code = 1
	r.Msg = msg
	r.Data = nil
	return r
}

func (r *R) Json() string {
	by, _ := json.Marshal(r)
	return string(by)
}

func (r *R) IsSucceed() bool {
	return 0 == r.Code
}