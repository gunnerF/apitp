package controllers

import (
	"apitp/models"
	"apitp/services"
	"apitp/utils"
	"time"
)

type NodeController struct {
	BaseController
}

func (c *NodeController) URLMapping() {
	c.Mapping("GetList", c.GetList)
	c.Mapping("GetNodeType", c.GetNodeType)
	c.Mapping("NodeDetail", c.NodeDetail)
}

// @Title 节点列表
// @Description 获取节点列表
// @Param nodeName body string false "节点名称"
// @Param subGroupName body string false "二级分类名称"
// @Param groupName body string false "一级分类名称"
// @Success 20000 {[]map[string]interface{}}
// @Failure 403 body is empty
// @router / [get]
func (c *NodeController) GetList() {
	result, count := new(services.NodeService).NodeGetList(c.page, c.pageSize, c.query)
	c.jsonMsgResult(utils.RequestSuccess["message"], utils.RequestSuccess["code"].(int), count, result)

}

// @Title 节点下拉列表
// @Description 获取节点下拉列表
// @Success 20000 {[]map[string]interface{}}
// @Failure 403 body is empty
// @router /getNodeType [get]
func (c *NodeController) GetNodeType() {
	result := new(services.NodeService).GetKeyValues(utils.NodeStatus)
	c.jsonMsgResult(utils.RequestSuccess["message"], utils.RequestSuccess["code"].(int), int64(len(result)), result)
}

// @Title 节点详情
// @Description 获取节点详情
// @Param id body int true "节点id"
// @Success 20000
// @Failure 403 body is empty
// @router /nodeDetail [get]
func (c *NodeController) NodeDetail() {
	nodeId, _ := c.GetInt("id")
	node := new(models.Node)
	node.SetScene("detail")
	node.Id = nodeId
	c.ParamsValidate(node)
	result := new(services.NodeService).NodeGetDetail(nodeId)
	c.jsonMsgResult(utils.RequestSuccess["message"], utils.RequestSuccess["code"].(int), int64(len(result)), result)
}

// @router /nodeAdd [post]
func (c *NodeController) NodeAdd() {
	node := new(models.Node)
	node.SetScene("add")
	currentTime := time.Now().Unix()
	nodeType, _ := c.GetInt("nodeType")
	subGroupId, _ := c.GetInt("subGroupId", 0)
	//验证是否存在对应分类
	subGroup, err := new(services.SubGroupService).GetInfoById(subGroupId)
	if err != nil {
		c.jsonMsgResult(err.Error(), utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	node.NodeName = c.GetString("nodeName")
	node.SubGroup = subGroup
	node.NodeType = nodeType
	node.Detail = c.GetString("detail")
	node.Status = 2
	node.CreateTime = currentTime
	node.UpdateTime = currentTime
	node.CreateId = c.userId
	node.UpdateId = c.userId
	c.ParamsValidate(node)
	result, err := new(services.NodeService).NodeAdd(node)
	if err != nil {
		c.jsonMsgResult(err.Error(), utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	c.jsonMsgResult(utils.AddSuccess["message"], utils.AddSuccess["code"].(int), 1, result)
}
