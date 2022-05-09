# goWeb-gin



# 项目介绍

基于gin框架搭建和mysql数据库开箱即用的web服务demo。

支持接口日志记录

支持数据库增删改日志记录

支持https映射

支持跨域访问

# 项目结构



```
├── .gitignore
├── config
│   ├── application.yml		//配置文件
│   ├── config.go
│   └── staticParam.go	//http静态错误吗
├── db
│   ├── dao.go	//数据库操作方法
│   └── entity.go	//数据库实体类
├── log
│   └── log.go	//日志相关配置
├── main.go	//项目入口
├── router
│   ├── commonHandler.go	//通用handler
│   ├── demoServiceHandler.go	//具体业务的handler
│   └── router.go	//路由管理
├── service
│   └── demoService.go	//具体业务service
├── utils
│   └── utils.go	//工具包
└── vo
    ├── commonVo.go	//通用组件的Vo
    └── poolVo.go	//具体业务vo
```







