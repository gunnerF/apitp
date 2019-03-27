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


