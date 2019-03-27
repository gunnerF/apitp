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
	page, err := c.GetInt("page")
	if err != nil {
		page = c.page
	}

	pageSize, err := c.GetInt("pageSize")
	if err != nil {
		pageSize = c.pageSize
	}
	nodeName := c.GetString("nodeName")
	subGroupName := c.GetString("subGroupName")
	groupName := c.GetString("groupName")
	filters := make([]interface{}, 0)
	if nodeName != "" {
		filters = append(filters, "NodeName__contains", nodeName)
	}
	if subGroupName != "" {
		filters = append(filters, "SubGroup__SubGroupName__contains", subGroupName)
	}
	if groupName != "" {
		filters = append(filters, "SubGroup__Group__GroupName__contains", groupName)
	}
	result, count := new(services.NodeService).NodeGetList(page, pageSize, filters...)
	new(services.NodeService).NodeGetList(page, pageSize, filters...)
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
	nodeId, err := c.GetInt("id")
	if err != nil {
		c.jsonMsgResult("节点id不能为空", utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	if nodeId <= 0 {
		c.jsonMsgResult("节点id必须大于0", utils.ParamsError["code"].(int), 1, c.resultJsonArr)
	}
	result := new(services.NodeService).NodeGetDetail(nodeId)
	c.jsonMsgResult(utils.RequestSuccess["message"], utils.RequestSuccess["code"].(int), int64(len(result)), result)
}

// @router /nodeAdd [post]
func (c *NodeController) NodeAdd() {
	node := new(models.Node)
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
