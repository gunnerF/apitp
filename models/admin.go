/**********************************************
** @Des: 用户模型
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package models

import "github.com/astaxie/beego/orm"

type Admin struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (m *Admin) TableName() string {
	return TableName("admin")
}

func AdminGetByName(loginName string) (*Admin, error) {
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableName("admin")).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (m *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
