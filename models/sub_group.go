/**********************************************
** @Des: 子分类模型
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package models

import "github.com/astaxie/beego/orm"

type SubGroup struct {
	Id           int
	Group        *Group `orm:"rel(fk)"` //orm定义model间关联关系，用于关联查询，具体使用方法参照beego官方文档
	SubGroupName string
	Status       int
	CreateId     int
	UpdateId     int
	CreateTime   int64
	UpdateTime   int64
	Nodes        []*Node `orm:"reverse(many)"`
}

func (m *SubGroup) TableName() string {
	return TableName("sub_group")
}

func (m *SubGroup) GetInfoById(id int) *SubGroup {
	subGroup := new(SubGroup)
	query := orm.NewOrm().QueryTable(TableName("sub_group"))
	query = query.Filter("id", id)
	query = query.RelatedSel()
	query.One(subGroup)
	return subGroup
}
