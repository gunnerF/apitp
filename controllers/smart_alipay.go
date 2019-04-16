/**********************************************
** @Des: aliPay控制器
	支付宝对接流程：
		1 在开放平台注册账号
		2 创建应用，生成公钥、私钥
		3 配置沙箱应用准备开发接入（案例中回调配置内网穿透地址）
		4 开发需要使用沙箱版支付宝app，在沙箱工具中下载
		5 服务端对接支付宝使用的是沙箱开放平台的支付宝公钥（一定注意，不然验签不通过）
		6 调用native生成web付款链接、手机扫码付款后，支付宝同步回调return方法或异步回调notify方法
** @Author: jgn
** @Date:   2019/4/16 11:37
** @Last Modified by:   jgn
** @Last Modified time: 2019/4/16 11:37
***********************************************/
package controllers

import (
	"github.com/smartwalle/alipay"
	"fmt"
	"apitp/utils"
)

type SmartAliPayController struct {
	BaseController
}

func (c *SmartAliPayController) URLMapping() {
	c.Mapping("HandleUserPay", c.HandleUserPay)
	c.Mapping("Return", c.Return)
	c.Mapping("Notify", c.Notify)
}

func newAliPayClient() *alipay.AliPay {
	var appId = "2016080300157179"
	//这个是沙箱公匙，此处应该使用支付宝公匙
	//var aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAybUyPX97uy0OkI8o4lspuFyAnMF5wi4EfDZhjLcoIpsbHy3P+fYIOq+dqQW1edOAiGKssI9Y6izULSIiCjdLeNobnEjnSfDyt94DZ0q6WIzn23Boj7UcbzKe1535hni1hJolty7/CB9SkgER7pkQ967oCyaae3o0X/15u+MVh/xrnCfVt2jQi1CkLuxaQexdEKHBRujU3JVLnweIP2Ao+8UYg7WB6MYSJ56ZnEIkxzu7yEnRyIFg08KyweyF10p/PRgySco1LTvOQQ7r2OVw+MNTIZxXa9MV2oty5g77WE6OXCYtgczx8zhDg/t6u4Kdj8VZ4w+CB4on4+NAKVB5kwIDAQAB"
	//支付宝公匙
	var aliPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqI2gT/OzMNy7pSHGnH8Cchhfy8FtkhYZlLb+8BhfkQUYMt52W3q6jNkWChII182QlR/6iacg32VLk+6qRgj6+JedcjZOUKKIwdrf8s2O4+FmXvQ2l1FJKTPHt+6ArcI/bYFyXuIPYD2eVmP59iFvY/R6n4wjlzvD3pw4KgzYVfKyJjlnMjf5fFLdNTkUUTalJtOrpBbSh1klRVXRcm/zJA5RH2hSx8Q7vryX0wpIFd3Ydv8EuzaLupq6mp80mZufUGkJpIxZ9vmEaOHTOCQTKWkJ/0Sgiqf+jY0CHrO49bslFCoCwjOMp39cRm+c8vQic4jeyjIiFxhP6MOgbqQQCQIDAQAB"
	var privateKey = "MIIEpAIBAAKCAQEAybUyPX97uy0OkI8o4lspuFyAnMF5wi4EfDZhjLcoIpsbHy3P+fYIOq+dqQW1edOAiGKssI9Y6izULSIiCjdLeNobnEjnSfDyt94DZ0q6WIzn23Boj7UcbzKe1535hni1hJolty7/CB9SkgER7pkQ967oCyaae3o0X/15u+MVh/xrnCfVt2jQi1CkLuxaQexdEKHBRujU3JVLnweIP2Ao+8UYg7WB6MYSJ56ZnEIkxzu7yEnRyIFg08KyweyF10p/PRgySco1LTvOQQ7r2OVw+MNTIZxXa9MV2oty5g77WE6OXCYtgczx8zhDg/t6u4Kdj8VZ4w+CB4on4+NAKVB5kwIDAQABAoIBAF+SJvfUi3+oXZpI+oolv6LG0Xl0gohq3V0tNxNBRRcAft1LC5vX4xZLV1xahB3xBJmah1Aw22Q5UV8BKTfF81CTpzlpoYz7SUWtP9eVJK4INWNX43MT0SEuzuxCwZPhZcAqeWxOLDBhBvJ+50wx4kQ472E5LUEMKfzBUQihE/aV30I4udNWyrWhIqPH7lc+j5Y7X7OaqF/8jdgCFaVHUwLE43XRmFaHTxVcGaevLwBsI/glQ9PrILL01QTbz5zcdJtUGbJWkX66/kwXQbW9uWEYUXoxN3MhV0qW5GZiT6KUHua5HReujRwOI/a+aXAAQyUOWBJpipeh4CGlqmIAXgECgYEA9Rys9Svj/vzNPJtJSVCxO4QQiS5FQQ2aQjdWYrzkPiGK2Hco8K6vDFtTk7kA37L1CaUoPaNBCCSPrfrWegVsdLpLkC+dCyHjvAoZxhcESHF/T5mVTkfGXKDa8WGXZolOUgkjC4KSskcGnQ3TuX8Pr6SL2L0UJeL61dafG2wXO0ECgYEA0qrxl/7EpSIejxNV4fT7oIlo+KgzfSofCE/qi7XjGZBC1rs4IOgL1IZjgADLnnMw3KRj/pXX3tFAeTSQDpWVvlYCxdy0iTR+xxFZ4JBA/x5743UCs8K6bRUIukKfbwrsCqM2SBQru3rEh4WaFZqu2wMe4f7TL4xeDg6rXHCg49MCgYB4yAaaUWRp89yWbtawH+kVsMANORW7Np663lXdFcNKsnCetikgYJO1fRM0Ccfac263sACiTt3uy0VZr+8b9aw54Mr3Y97461wy/q8wo4riv7rCWKXwCwI+Bq2coEBgsGw9lG8GbvMmkkU6AuQ/1dK+9FBNHu7ctHpyAO4qLxtLgQKBgQCvXB5Y8HiOIPWhEW5G9Z9Tha48df7OLHbCe4t4moIc8GnCxvs04ROf5/IBRig4kBhXQfv+moqj7jkl62wSx89rLbuHKm2ZXHo/iKDw/gMsmhp0Px1ttyVp0b1/FiAS1dugCzYPD8NLfykCymJ1o1nz8wgNBQqUJNrSBsH748D5gwKBgQDCAfPw3Ck4imszZLAgeYGz7N0YGveWqKVyTz9oy6n0U+WSLC9PZh5BqX7hVh8dpKloK6Mcoro9w229T3fguwEJFkjUEkftd+nNd9eRxA+aAGLi4D+V/i/dN4o+M0KBs4I6GL74c+BF4xBg9OS2kWRowf+BNNZ1KtkNV91vGTtcmg=="
	return alipay.New(appId, aliPublicKey, privateKey, false)
}

// @Title 发起订单付款
// @Description 发起订单付款
// @Param orderNo body int true "订单编号"
// @Success 20000
// @Failure 403 body is empty
// @router /native [get]
func (c *SmartAliPayController) HandleUserPay() {
	client := newAliPayClient()
	oid := c.GetString("oid")
	totalPrice := c.GetString("totalPrice")

	var p = alipay.AliPayTradePagePay{}
	//本地开发内网穿透api地址
	p.NotifyURL = "http://21t874i952.iask.in:34761/v1/smart-pay/notify"
	p.ReturnURL = "http://21t874i952.iask.in:34761/v1/smart-pay/return"
	p.Subject = "jgn_order"
	p.OutTradeNo = oid
	p.TotalAmount = totalPrice
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err = client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}
	var payURL = url.String()
	//fmt.Println("payURL:", payURL)
	result := utils.ResultJson{"payURL": payURL,}
	c.jsonMsgResult(utils.RequestSuccess["message"], utils.RequestSuccess["code"].(int), 1, result)
}

// @Title 订单同步回调
// @Description 订单回调请求
// @Success 20000
// @Failure 403 body is empty
// @router /return [get]
func (c *SmartAliPayController) Return() {
	fmt.Println("return action")
	c.Ctx.Request.ParseForm()
	client := newAliPayClient()
	ok, err := client.VerifySign(c.Ctx.Request.Form)
	fmt.Println( "alipay 回调return=================================", ok, err)
	//处理业务
}

// @Title 订单异步回调
// @Description 订单回调请求
// @Success 20000
// @Failure 403 body is empty
// @router /notify [post]
func (c *SmartAliPayController) Notify() {
	client := newAliPayClient()
	var noti, _ = client.GetTradeNotification(c.Ctx.Request)
	if noti != nil {
		fmt.Println("支付成功")
	} else {
		fmt.Println("支付失败")
	}
	//处理业务
	alipay.AckNotification(c.Ctx.ResponseWriter) // 确认收到通知消息
}