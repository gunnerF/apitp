/**********************************************
** @Des: 组服务
** @Author: jgn
** @Date:   2019-03-11 14:17:37
** @Last Modified by:   jgn
** @Last Modified time: 2019-03-11 14:17:37
***********************************************/
package services

import (
	"apitp/models"
	"baliance.com/gooxml/measurement"
	"baliance.com/gooxml/document"
	"baliance.com/gooxml/schema/soo/wml"
	"strconv"
	"time"
	"baliance.com/gooxml/color"
	"os"
	"path/filepath"
	"strings"
)

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

//生成word文档
func (s *GroupService) WordDown(id int) (string, error) {


	doc := document.New()	//新建文档实例
	para := doc.AddParagraph()	//文字建立在Paragraph之上
	para.Properties().SetAlignment(wml.ST_JcCenter)	//设置段落居中
	para.Properties().Spacing().SetAfter(10)	//设置段落下边距
	run := para.AddRun()	//在段落上创建run实例

	//设置文字大小
	run.Properties().SetStyle("Heading1")
	run.Properties().SetBold(true)	//设置文字加粗
	run.Properties().SetCharacterSpacing(2)	//设置字间距
	run.AddText("拱墅区物业管理项目现场检查考核表")	//添加文字

	table := doc.AddTable()	//添加表格
	table.Properties().SetAlignment(wml.ST_JcTableCenter)	//设置表格文字居中
	table.Properties().SetWidth(6 * measurement.Inch)	//设置表格宽度
	borders := table.Properties().Borders()	//设置边框样式
	borders.SetAll(wml.ST_BorderSingle, color.Auto, measurement.Zero)

	row := table.AddRow()	//添加row
	cell := row.AddCell()
	cell.Properties().SetWidth(1 * measurement.Inch)
	cell.Properties().SetVerticalMerge(wml.ST_MergeRestart)	//设置合并row
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)	//设置cell居中
	cell.AddParagraph().AddRun().AddText("管理项目主任(经理)基本情况")
	cell = row.AddCell()
	cell.Properties().SetWidth(2 * measurement.Inch)
	cell.AddParagraph().AddRun().AddText("姓名及联系方式")
	cell = row.AddCell()
	cell.Properties().SetWidth(2 * measurement.Inch)
	cell.AddParagraph().AddRun().AddText("")

	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetWidth(1 * measurement.Inch)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)	//ST_MergeContinue表示继续合并
	cell.AddParagraph().AddRun()	//注意：此例子第一列不合并，第二、第三列合并，此行作用是保持cell数量一致
	cell = row.AddCell()
	cell.Properties().SetWidth(2 * measurement.Inch)
	cell.AddParagraph().AddRun().AddText("管理项目名称")
	cell = row.AddCell()
	cell.Properties().SetWidth(2 * measurement.Inch)
	cell.AddParagraph().AddRun().AddText("")

	row = table.AddRow()
	cell = row.AddCell()
	cell.Properties().SetWidth(1 * measurement.Inch)
	cell.Properties().SetVerticalMerge(wml.ST_MergeContinue)
	cell.AddParagraph().AddRun()
	cell = row.AddCell()
	cell.Properties().SetWidth(2 * measurement.Inch)
	cell.AddParagraph().AddRun().AddText("物业企业名称")
	cell = row.AddCell()
	cell.Properties().SetWidth(2 * measurement.Inch)
	cell.AddParagraph().AddRun().AddText("")

	row = table.AddRow()
	row.Properties().SetHeight(1 * measurement.Inch, wml.ST_HeightRuleAtLeast)
	cell = row.AddCell()
	cell.Properties().SetWidth(1 * measurement.Inch)
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.AddParagraph().AddRun().AddText("检查考核内容")
	cell = row.AddCell()
	cell.Properties().SetWidth(4 * measurement.Inch)
	cell.Properties().SetColumnSpan(2)	//合并cell，合并前需考虑好要合并cell数量，要严格对齐
	cell.AddParagraph().AddRun().AddText("具体见《杭州市物业管理项目检查考核记分内容与标准》")

	row = table.AddRow()
	row.Properties().SetHeight(4 * measurement.Inch, wml.ST_HeightRuleAtLeast)
	cell = row.AddCell()
	cell.Properties().SetWidth(1 * measurement.Inch)
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.AddParagraph().AddRun().AddText("检查项目存在问题及整改要求告知情况")
	cell = row.AddCell()
	cell.Properties().SetColumnSpan(2)
	cell.Properties().SetWidth(4 * measurement.Inch)
	cell.AddParagraph().AddRun().AddText("暂无问题")

	row = table.AddRow()
	row.Properties().SetHeight(1 * measurement.Inch, wml.ST_HeightRuleAtLeast)
	cell = row.AddCell()
	cell.Properties().SetWidth(1 * measurement.Inch)
	cell.Properties().SetVerticalAlignment(wml.ST_VerticalJcCenter)
	cell.AddParagraph().AddRun().AddText("相关人员签章")
	cell = row.AddCell()
	cell.Properties().SetColumnSpan(2)
	cell.Properties().SetWidth(4 * measurement.Inch)
	cell.AddParagraph().AddRun().AddText("以上检查情况属实")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	path := strings.Replace(dir, "\\", "/", -1) + "/download/"
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "simple.docx"
	doc.SaveToFile(path + fileName)
	return fileName, nil
}
