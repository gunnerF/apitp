/**********************************************
** @Des: 子分类服务
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package services

import (
	"apitp/models"
	"github.com/pkg/errors"
	"reflect"
)

type SubGroupService struct {
	BaseService
}

func (s *SubGroupService) GetInfoById(subGroupId int) (*models.SubGroup, error) {
	subGroup := new(models.SubGroup).GetInfoById(subGroupId)
	if reflect.DeepEqual(subGroup, &models.SubGroup{}) {
		err := errors.New("二级分类不存在")
		return nil, err
	}
	return subGroup, nil
}
