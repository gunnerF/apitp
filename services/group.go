/**********************************************
** @Des: 组服务
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package services

import "apitp/models"

type GroupService struct {
	BaseService
}

func (s *GroupService) GetGroupByName(groupName string) bool {
	return new(models.Group).GetGroupByName(groupName)
}

func (s *GroupService) GroupAdd(group *models.Group) (int64, error) {
	return new(models.Group).GroupAdd(group)
}

func (s *GroupService) GroupDelete(id int) error {
	result, err := new(models.Group).GroupDelete(id)
	if result {
		return nil
	}
	return err
}
