/**********************************************
** @Des: 节点服务
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package services

import (
	"apitp/models"
	"apitp/utils"
	"net/url"
	"strconv"
)

type NodeService struct {
	BaseService
}

func (s *NodeService) NodeGetList(page int, pageSize int, query url.Values) (utils.ResultJsonArr, int64) {
	nodes, count := models.NodeGetList(page, pageSize, query)
	result := make(utils.ResultJsonArr, len(*nodes))
	if len(*nodes) > 0 {
		for k, v := range *nodes {
			row := make(utils.ResultJson)
			row = v
			row["create_time"] = s.FormatDate(v["create_time"].(string), kYMDHIS_A)
			row["update_time"] = s.FormatDate(v["update_time"].(string), kYMDHIS_A)
			result[k] = row
		}
	}
	return result, count
}

func (s *NodeService) NodeGetDetail(nodeId int) utils.ResultJsonArr {
	node := models.NodeGetDetail(nodeId)
	result := make(utils.ResultJsonArr, 1)
	if node != nil {
		row := make(utils.ResultJson)
		row["id"] = node.Id
		row["nodeType"] = node.NodeType
		row["nodeName"] = node.NodeName
		row["detail"] = node.Detail
		row["status"] = node.Status
		row["statusText"] = utils.NodeStatus[node.Status]
		row["subGroupName"] = node.SubGroup.SubGroupName
		row["groupName"] = node.SubGroup.Group.GroupName
		row["createTime"] = s.FormatDate(strconv.FormatInt(node.CreateTime, 10), kYMDHIS_A)
		row["updateTime"] = s.FormatDate(strconv.FormatInt(node.CreateTime, 10), kYMDHIS_A)
		result[0] = row
	}
	return result
}

func (s *NodeService) NodeAdd(node *models.Node) (utils.ResultJsonArr, error) {
	res, err := models.NodeAdd(node)
	if err != nil {
		return nil, err
	}
	result := make(utils.ResultJsonArr, 1)
	row := make(utils.ResultJson)
	row["id"] = res
	result[0] = row
	return result, nil
}
