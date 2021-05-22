# 系统相关
root uri: /api/

默认域名: http://localhost:8848/api/
# 账号相关
控制器：account/a(需要登录态)、account/f(无需登录态)
####账号密码登录
POST account/f/login

请求body
```json
{
    "account": "piadi",       //账号
    "password": "password",   //密码
    "code": "1234",           //用户输入的验证码
    "key": "key"              //图片验证码返回的id
}
```
响应header
```
Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MjE2ODc3NTl9LCJVc2VyIjp7IklEIjoxLCJDcmVhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJVcGRhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJEZWxldGVkQXQiOm51bGwsIk1vYmlsZSI6IjEzNjQ5MjkwMTQ3IiwiQWNjb3VudCI6ImxlbyIsIk5hbWUiOiLosKLlrZDpvpkiLCJQYXNzd29yZCI6IjEyMzQ1NiIsIlJ1bGVJZCI6MCwiQ3JlYXRvciI6MH19.Fl3vnhCRkn6n2v-RxAgvCblff0q_3RhaXUoTz3vGlvE
```
响应body
```json
{
    "code": 0,
    "msg": "Succeed",
    "data": null
}
```

#### 注册账号
POST account/f/register

请求body
```json
{
    "account": "piadi",       //账号
    "password": "password",   //密码
    "mobile": "mobile",       //手机号码
    "code": "1234",           //用户输入的验证码
    "key": "key"              //图片验证码返回的id
}
```
响应body
```json
{
    "code": 0,
    "msg": "Succeed",
    "data": null
}
```
# 角色相关
控制器：rule/a(需要登录态)
#### 获取指定权限信息
GET rule/a/${id}

请求包header
```
Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MjE2ODc3NTl9LCJVc2VyIjp7IklEIjoxLCJDcmVhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJVcGRhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJEZWxldGVkQXQiOm51bGwsIk1vYmlsZSI6IjEzNjQ5MjkwMTQ3IiwiQWNjb3VudCI6ImxlbyIsIk5hbWUiOiLosKLlrZDpvpkiLCJQYXNzd29yZCI6IjEyMzQ1NiIsIlJ1bGVJZCI6MCwiQ3JlYXRvciI6MH19.Fl3vnhCRkn6n2v-RxAgvCblff0q_3RhaXUoTz3vGlvE
```
响应body
```json
{
    "code": 0,
    "msg": "Succeed",
    "data": {
        "id": 1,
        "level":1,
        "name": "管理员"
    }
}
```

# 日志相关
控制器：log/a(需要登录态)
#### 提交日志
PUT log/a/${date}

请求包header
```
Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MjE2ODc3NTl9LCJVc2VyIjp7IklEIjoxLCJDcmVhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJVcGRhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJEZWxldGVkQXQiOm51bGwsIk1vYmlsZSI6IjEzNjQ5MjkwMTQ3IiwiQWNjb3VudCI6ImxlbyIsIk5hbWUiOiLosKLlrZDpvpkiLCJQYXNzd29yZCI6IjEyMzQ1NiIsIlJ1bGVJZCI6MCwiQ3JlYXRvciI6MH19.Fl3vnhCRkn6n2v-RxAgvCblff0q_3RhaXUoTz3vGlvE
```
请求body
```json
{
    "content": "这是一条测试日志"
}
```
响应body
```json
{
    "code": 0,
    "msg": "Succeed",
    "data": null
}
```
#### 获取指定日期日志
GET log/a/${date}

请求包header
```
Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MjE2ODc3NTl9LCJVc2VyIjp7IklEIjoxLCJDcmVhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJVcGRhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJEZWxldGVkQXQiOm51bGwsIk1vYmlsZSI6IjEzNjQ5MjkwMTQ3IiwiQWNjb3VudCI6ImxlbyIsIk5hbWUiOiLosKLlrZDpvpkiLCJQYXNzd29yZCI6IjEyMzQ1NiIsIlJ1bGVJZCI6MCwiQ3JlYXRvciI6MH19.Fl3vnhCRkn6n2v-RxAgvCblff0q_3RhaXUoTz3vGlvE
```
响应body
```json
{
    "code": 0,
    "msg": "Succeed",
    "data": {
        "id": 1,
        "userId": 1,
        "dateTime": "20210522"
    } 
}
```
#### 分页获取日志列表
GET log/a/${page}/${size}

请求包header
```
Authorization=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdGFuZGFyZENsYWltcyI6eyJleHAiOjE2MjE2ODc3NTl9LCJVc2VyIjp7IklEIjoxLCJDcmVhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJVcGRhdGVkQXQiOiIyMDIxLTA1LTIyVDExOjE5OjQ1LjM5N1oiLCJEZWxldGVkQXQiOm51bGwsIk1vYmlsZSI6IjEzNjQ5MjkwMTQ3IiwiQWNjb3VudCI6ImxlbyIsIk5hbWUiOiLosKLlrZDpvpkiLCJQYXNzd29yZCI6IjEyMzQ1NiIsIlJ1bGVJZCI6MCwiQ3JlYXRvciI6MH19.Fl3vnhCRkn6n2v-RxAgvCblff0q_3RhaXUoTz3vGlvE
```
响应body
```json
{
    "code": 0,
    "msg": "Succeed",
    "data": {
        "count": 100,
        "page": 1,
        "size": "10",
        "datas": [
            {
                "id": 1,
                "userId": 1,
                "dateTime": "20210522"
            }
        ]
    } 
}
```
# 图片验证码相关
控制器：captcha/f

GET captcha/f

响应body
```json
{
    "code":0,
    "msg":"Succeed",
    "data":{
        "id":"e0e1523c-f614-97e4-d418-d1a604e0ca29",
        "img":"data:image/png;base64,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"
    }
}
```