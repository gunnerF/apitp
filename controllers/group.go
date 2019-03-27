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
}

// @Title 新增组
// @Description 新增组
// @Success 20000 {[]map[string]interface{}}
// @Failure 403 body is empty
//@router /groupAdd [post]
func (c *GroupController) GroupAdd() {

	group := new(models.Group)
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
	if id == 0 || err != nil {
		c.jsonMsgResult("id不能为空，且必须为整型", utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	err = new(services.GroupService).GroupDelete(id)
	if err != nil {
		c.jsonMsgResult(err.Error(), utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	c.jsonMsgResult(utils.DeleteSuccess["message"], utils.DeleteSuccess["code"].(int), 1, c.resultJsonArr)
}
