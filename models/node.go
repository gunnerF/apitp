/**********************************************
** @Des: 节点模型
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package models

import "github.com/astaxie/beego/orm"

type Node struct {
	Id         int
	SubGroup   *SubGroup `orm:"rel(fk)" valid:"Required"` //orm定义model间关联关系，用于关联查询, valid定义参数验证
	NodeType   int       `valid:"Required;Range(1,4)"`
	NodeName   string    `json:"nodeName" valid:"Required"`
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

func NodeGetList(page int, pageSize int, filters ...interface{}) ([]*Node, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Node, 0)
	l := len(filters)
	query := orm.NewOrm().QueryTable(TableName("node"))
	if l > 0 {
		for i := 0; i < l; i += 2 {
			query = query.Filter(filters[i].(string), filters[i+1])
		}
	}
	query = query.RelatedSel()
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
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
