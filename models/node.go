/**********************************************
** @Des: 节点模型
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"net/url"
)

type Node struct {
	scene      string
	Id         int
	SubGroup   *SubGroup `orm:"rel(fk)"` //orm定义model间关联关系，用于关联查询, valid定义参数验证
	NodeType   int
	NodeName   string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (m *Node) TableName() string {
	return TableName("node")
}

func (m *Node) SetScene(scene string) {
	m.scene = scene
}

func (m *Node) Valid(v *validation.Validation) {
	if m.scene == "detail" {
		fmt.Println("m.Id:", m.Id)
		if res := v.Required(m.Id, "Id"); !res.Ok {
			res.Message("Id不能为空")
			v.SetError("Id", res.Error.Message)
		}
	}
	if m.scene == "add" {
		if res := v.Required(m.SubGroup, "subGroup"); !res.Ok {
			res.Error.Message = res.Error.Name + res.Error.Message
			v.SetError("SubGroup", res.Error.Message)
		}
		if res := v.Required(m.NodeType, "nodeType"); !res.Ok {
			res.Error.Message = res.Error.Name + res.Error.Message
			v.SetError("NodeType", res.Error.Message)
		}
		if res := v.Range(m.NodeType, 1, 4, "nodeType"); !res.Ok {
			res.Error.Message = res.Error.Name + res.Error.Message
			v.SetError("NodeType", res.Error.Message)
		}
		if res := v.Required(m.NodeName, "nodeName"); !res.Ok {
			res.Error.Message = res.Error.Name + res.Error.Message
			v.SetError("NodeName", res.Error.Message)
		}
	}
}

//func NodeGetList(page int, pageSize int, filters ...interface{}) ([]*Node, int64) {
//	offset := (page - 1) * pageSize
//	list := make([]*Node, 0)
//	l := len(filters)
//	query := orm.NewOrm().QueryTable(TableName("node"))
//	if l > 0 {
//		for i := 0; i < l; i += 2 {
//			query = query.Filter(filters[i].(string), filters[i+1])
//		}
//	}
//	query = query.RelatedSel()
//	total, _ := query.Count()
//	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
//	return list, total
//}

func NodeGetList(page int, pageSize int, query url.Values) (*[]orm.Params, int64) {
	offset := (page - 1) * pageSize
	var result []orm.Params
	o := orm.NewOrm()
	err := o.Begin()
	var conditions []interface{}
	filters := make([]interface{}, 0)
	if query.Get("nodeName") != "" {
		nodeName := "%" + query.Get("nodeName") + "%"
		filters = append(filters, "node_name like ?", nodeName)
	}
	if query.Get("subGroupName") != "" {
		subGroupName := "%" + query.Get("subGroupName") + "%"
		filters = append(filters, "tp_sub_group.sub_group_name like ?", subGroupName)
	}
	if query.Get("groupName") != "" {
		groupName := "%" + query.Get("groupName") + "%"
		filters = append(filters, "tp_group.group_name like ?", groupName)
	}
	l := len(filters)
	qb, _ := orm.NewQueryBuilder("mysql")
	fields := `tp_node.*, tp_sub_group.sub_group_name, tp_group.group_name`
	qb = qb.Select(fields).From("tp_node")
	qb = qb.LeftJoin("tp_sub_group").On("tp_node.sub_group_id = tp_sub_group.id")
	qb = qb.LeftJoin("tp_group").On("tp_sub_group.group_id = tp_group.id")

	if l > 0 {
		for i := 0; i < l; i += 2 {
			qb = qb.Where(filters[i].(string))
			conditions = append(conditions, filters[i+1])
		}
	}
	sql := qb.String()
	o.Raw(sql, conditions).Values(&result)
	count := len(result)
	qb = qb.OrderBy("tp_node.id").Desc()
	qb = qb.Limit(pageSize)
	if offset > 0 {
		qb = qb.Offset(offset)
	}
	sql = qb.String()
	o.Raw(sql, conditions).Values(&result)
	if err != nil {
		err = o.Rollback()
	} else {
		err = o.Commit()
	}
	return &result, int64(count)
}

func NodeGetDetail(nodeId int) *Node {
	node := new(Node)
	query := orm.NewOrm().QueryTable(TableName("node"))
	query.Filter("id", nodeId)
	query = query.RelatedSel() //关联表查询，需要先定义好关联关系
	query.One(node)
	return node
}

func NodeAdd(node *Node) (int64, error) {
	res, err := orm.NewOrm().Insert(node)
	if err != nil {
		return 0, err
	}
	return res, nil
}
