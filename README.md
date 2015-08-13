#goMeikongSpider
一个利用[golang](https://golang.org)语言实现抓取[美空网](http://www.moko.cc/)模特信息的简单网络爬虫

##About
这个项目主要用于增强自己对于[golang](https://golang.org)和[MongoDB2.6.4](https://www.mongodb.org/)的掌握程度，顺便的学习一下[WebSocket](https://www.websocket.org/)等其他的web开发技术

##Require
* [Golang1.4.2](https://golang.org/)
* [MongoDB2.6.4](https://www.mongodb.org/)
* [BootstrapV3](http://getbootstrap.com/)
* [WebSocket](https://www.websocket.org/)
##Features
* 从[美空网](http://www.moko.cc/)抓取模特图片，将模特信息存入[MongoDB2.6.4](https://www.mongodb.org/)数据库中，将模特图片下载存放到本地目录中
* 利用[WebSocket](https://www.websocket.org/)来实时的显示爬虫抓取模特的信息

##ToDo
* 将爬虫的抓取过程修改为并发实现，缩短爬虫抓取的时间
* 采用[beego](http://beego.me/)作为MVC框架的实现
* 提供[REST](https://zh.wikipedia.org/wiki/REST)方式的对外Web Service接口
* 利用[美空网](http://www.moko.cc/)的账号进行模拟登录，从而对每个模特抓取完整的图片
* 模特图片分页与评论功能的添加