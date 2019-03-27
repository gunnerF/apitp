/**********************************************
** @Des: 组模型
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
)

type Group struct {
	Id         int
	GroupName  string `valid:"Required"`
	Detail     string `valid:"Required"`
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
	SubGroups  []*SubGroup `orm:"reverse(many)"`
}

//所有验证通过后执行该方法
func (m *Group) Valid(v *validation.Validation) {
	fmt.Println()
}

func (m *Group) TableName() string {
	return TableName("group")
}

func (m *Group) GetGroupByName(groupName string) bool {
	query := orm.NewOrm().QueryTable(TableName("group")).Filter("GroupName", groupName)
	return query.Exist()
}

func (m *Group) GroupAdd(group *Group) (int64, error) {
	return orm.NewOrm().Insert(group)
}

func (m *Group) GroupDelete(id int) (bool, error) {
	query := orm.NewOrm().QueryTable(TableName("group"))
	_, err := query.Filter("id", id).Delete()
	if err != nil {
		return false, err
	}
	return true, nil
}
