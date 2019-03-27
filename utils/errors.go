/**********************************************
** @Des: 错误信息及状态码定义
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package utils

var (
	RequestSuccess = map[string]interface{}{"code": 20000, "message": ""}
	AddSuccess     = map[string]interface{}{"code": 20000, "message": "新增成功"}
	DeleteSuccess  = map[string]interface{}{"code": 20000, "message": "删除成功"}
	LoginSuccess   = map[string]interface{}{"code": 20000, "message": "登录成功"}

	AlreadyLogin    = map[string]interface{}{"code": 30000, "message": "已经登陆"}
	UserOrPassword  = map[string]interface{}{"code": 30001, "message": "用户名或密码错误"}
	UserPwdEmpty    = map[string]interface{}{"code": 30001, "message": "用户名或密码不能为空"}
	AccountDisabled = map[string]interface{}{"code": 30002, "message": "该账号已经被禁用"}

	MissToken    = map[string]interface{}{"code": 40000, "message": "token缺失"}
	OverdueToken = map[string]interface{}{"code": 40001, "message": "token过期"}

	NoPermission = map[string]interface{}{"code": 50000, "message": "您没有权限访问"}

	ParamsError = map[string]interface{}{"code": 60000, "message": ""}

	SystemError = map[string]interface{}{"code": 60000, "message": ""}
)
