# apitp
对beego api接口开发实现了一些封装
项目中编写了简单的实例

## 初始化项目
- 创建mysql数据库，数据库名称api_tp，将项目中 api_tp.sql导入库中
```shell
// 进入到项目目录，启动项目
bee run
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

## controller
- base.go中实现对请求权限的验证，以及请求基础参数的初始化
- 封装了json类型返回函数，post请求验证函数以及一些常用的方法

```Go
func (c *BaseController) ParamsValidate(obj interface{}){
    ...
}
```
ParamsValidate函数结合validation组件使用，该方法实现了对参数的验证，验证后直接将错误信息返回。
### 使用方法
- 在业务controller中实例化model对象后直接调用ParamsValidate方法
```Go
    node := new(models.Node)
    node.Status = 2
    ...
    c.ParamsValidate(node)
```
- validation具体使用参数等参见beego官方文档 https://beego.me/docs/mvc/controller/validation.md

