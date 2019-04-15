/**********************************************
** @Des: 组控制器
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package controllers

import (
	"apitp/models"
	"apitp/services"
	"apitp/utils"
	"strings"
	"time"
)

type GroupController struct {
	BaseController
}

func (c *GroupController) URLMapping() {
	c.Mapping("GroupAdd", c.GroupAdd)
	c.Mapping("GroupDelete", c.GroupDelete)
	c.Mapping("WordDown", c.WordDown)
}

// @Title 新增组
// @Description 新增组
// @Success 20000 {[]map[string]interface{}}
// @Failure 403 body is empty
//@router /groupAdd [post]
func (c *GroupController) GroupAdd() {

	group := new(models.Group)
	group.SetScene("add")
	group.GroupName = strings.TrimSpace(c.GetString("groupName"))
	group.Detail = strings.TrimSpace(c.GetString("detail"))
	group.Status = 1
	group.CreateId = c.userId
	group.UpdateId = c.userId
	group.CreateTime = time.Now().Unix()
	group.UpdateTime = time.Now().Unix()
	c.ParamsValidate(group)
	if _, err := new(services.GroupService).GroupAdd(group); err != nil {
		c.jsonMsgResult(err.Error(), utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	c.jsonMsgResult(utils.AddSuccess["message"], utils.AddSuccess["code"].(int), 1, c.resultJsonArr)
}

// @Title 删除组
// @Description 删除组
// @Param id body int true "节点id"
// @Success 20000 {[]map[string]interface{}}
// @Failure 403 body is empty
//@router /groupDelete [post]
func (c *GroupController) GroupDelete() {
	id, err := c.GetInt("id")
	group := new(models.Group)
	group.SetScene("delete")
	c.ParamsValidate(group)
	err = new(services.GroupService).GroupDelete(id)
	if err != nil {
		c.jsonMsgResult(err.Error(), utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	c.jsonMsgResult(utils.DeleteSuccess["message"], utils.DeleteSuccess["code"].(int), 1, c.resultJsonArr)
}

// @Title word文件生成
// @Description word文件
// @Param id body int true "组id"
// @Success 20000 {[]map[string]interface{}}
// @Failure 403 body is empty
//@router /word-down [get]
func (c *GroupController) WordDown() {
	id, _ := c.GetInt("id")
	group := new(models.Group)
	group.SetScene("delete")
	group.Id = id
	host := c.Ctx.Request.Host + "/down/"
	//fmt.Println("URLPATH:", c.Ctx.Request.Host)
	c.ParamsValidate(group)
	fileName, err := new(services.GroupService).WordDown(id)
	if err != nil {
		c.jsonMsgResult(err.Error(), utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	result := make(utils.ResultJson)
	result["down_url"] = host + fileName
	c.jsonMsgResult(utils.RequestSuccess["message"], utils.RequestSuccess["code"].(int), 1, result)
}
