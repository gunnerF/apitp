/**********************************************
** @Des: key value map定义
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package utils

var (
	FmtDate = map[int]string{
		1: "Y-m-d H:i:s", 2: "Y/m/d H:i:s", 3: "Y/m/d H:i", 4: "Y-m-d", 5: "H:i:s",
	}
	NodeStatus = map[int]string{
		1: "春天", 3: "秋天", 4: "冬天", 2: "夏天",
	}
)
