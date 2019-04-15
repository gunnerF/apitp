/**********************************************
** @Des: 基础服务
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package services

import (
	"apitp/utils"
	"github.com/astaxie/beego"
	"sort"
	"strconv"
	"time"
)

type BaseService struct {
}

//格式化时间
func (s *BaseService) FormatDate(createTime string, format string) string {
	timeUnix, _ := strconv.ParseInt(createTime, 10, 64)
	return beego.Date(time.Unix(timeUnix, 0), format)
}

//map返回值排序方法
func (s *BaseService) GetKeyValues(mapArr map[int]string) utils.KeyValues {
	keyValue := make(utils.KeyValues, 0)
	for k, v := range mapArr {
		row := utils.KeyValue{Key: k, Value: v}
		keyValue = append(keyValue, row)
	}
	sort.Sort(utils.KeyValues(keyValue))
	return keyValue
}
