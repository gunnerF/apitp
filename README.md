# apitp
对beego api接口开发实现了一些封装
项目中编写了简单的实例

## 初始化项目
- 创建mysql数据库，数据库名称api_tp，将项目中 api_tp.sql导入库中
```shell
// 1 项目中应用了mysql,validation,errors等组件,启动前需要go get 相关组件
go get github.com/go-sql-driver/mysql
go get github.com/patrickmn/go-cache
go get github.com/pkg/errors
go get github.com/astaxie/beego/validation
go get baliance.com/gooxml/  -- 支持word、excel处理
go get github.com/smartwalle/alipay  -- 引入开源sdk，对接支付宝支付
// 进入到项目目录，启动项目
bee run
// 访问示例： http://localhost:8080/v1/node/getNodeType
```

## 项目结构
    |-- conf
    |   `-- app.conf
    |-- controllers
    |   `-- base.go
    |-- models
    |   `-- base.go
    |-- services
    |   `-- base.go
    |-- routers
    |   `-- router.go
    |-- swagger
    |-- tests
    |   `-- default_test.go
    |-- utils
    |-- main.go

## 项目配置
- 配置文件放在conf文件夹中，项目启动后默认加载app.conf配置文件
- 项目数据库配置信息放在db.conf中

## 路由的使用
- 项目中使用的是注解路由的方式
```Go
    // @Title 节点列表
    // @Description 获取节点列表
    // @Param nodeName body string false "节点名称"
    // @Param subGroupName body string false "二级分类名称"
    // @Param groupName body string false "一级分类名称"
    // @Success 20000 {[]map[string]interface{}}
    // @Failure 403 body is empty
    // @router / [get]
    func (c *NodeController) GetList() {
        ...
    }
```
首先在controller中创建函数，注解路由通过注释中@router进行匹配，@router必须放在函数名的正上方，不能有空格。

```Go
    func (c *NodeController) URLMapping() {
    	c.Mapping("GetList", c.GetList)
    	...
    }
```
然后在controller中实现URLMapping方法，如果没有在router文件中注册路由，则会通过URLMapping方法通过反射来执行对应的方法，与注册路由方式有性能差异

```Go
    beego.NSNamespace("/node",
            beego.NSInclude(
                &controllers.NodeController{},
            ),
        ),
```
在router中绑定对应的namespace
具体路由配置方法参考beego 官方文档 https://beego.me/docs/mvc/controller/router.md

## controller
- controller主要做参数的验证，不做具体业务的实现
- base.go中实现对请求权限的验证，以及请求基础参数的初始化
- 封装了json类型返回函数，post请求验证函数以及一些常用的方法

```Go
func (c *BaseController) ParamsValidate(obj interface{}){
    ...
}
```
ParamsValidate函数结合validation组件使用，该方法实现了对参数的验证，验证后直接将错误信息返回。
使用方法
- 在业务controller中实例化model对象后直接调用ParamsValidate方法
```Go
    node := new(models.Node)
    node.Status = 2
    ...
    c.ParamsValidate(node)
```
- validation具体使用参数等参见beego官方文档 https://beego.me/docs/mvc/controller/validation.md

## services
- services主要做具体业务的实现，复杂业务的封装
- base.go中封装了对key value map排序的方法，以及时间格式处理方法
key value排序函数使用方法
- 在utils文件夹的keyvalues中定义map，直接调用GetKeyValues传入map，返回正序的map。
```Go
    func (s *BaseService) GetKeyValues(mapArr map[int]string) utils.KeyValues {
        ...
    }
```
由于map是无序的，对map排序可以使用sort结构体中的方法，要使用sort中的方法需要实现相应的接口，项目中在utils文件夹中做了key value相应的实现

## model
- 要使用 RelatedSel()方法进行关联查询需要先设置Tag
- base.go中做了对mysql数据库连接的初始化，以及数据库表的重命名
- 新增model后需要在init()方法中注册对应的model

```Go
    func (m *Node) TableName() string {
        ...
    }
```
在业务的model中一定要实现TableName方法，在model实例化时会根据业务model的该方法来获取表名称。

## utils
- 工具类以及公共参数的定义
```Go
        |-- utils
        |   `-- base.go 定义一些方法
        |   `-- cache.go 缓存设置
        |   `-- error.go 错误信息及状态码定义
        |   `-- keyvalues.go key value关系定义
        |   `-- types.go 公共参数类型定义
```
