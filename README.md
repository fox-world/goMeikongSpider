# goMeikongSpider
一个利用[golang](https://golang.org)语言实现抓取[美空网](http://www.moko.cc/)模特信息的简单网络爬虫

## About
这个项目主要用于增强自己对于[golang](https://golang.org)和[MongoDB2.6.4](https://www.mongodb.org/)的掌握程度，顺便学习一下[WebSocket](https://www.websocket.org/)等其他的web开发技术

## Require
* [Golang1.4.2](https://golang.org/)
* [MongoDB2.6.4](https://www.mongodb.org/)
* [BootstrapV3](http://getbootstrap.com/)
* [WebSocket](https://www.websocket.org/)
* [jQuery1.11.0](https://jquery.com/)

## Features
* 从[美空网](http://www.moko.cc/)抓取模特图片，将模特信息存入[MongoDB2.6.4](https://www.mongodb.org/)数据库中，将模特图片下载存放到本地目录中
* 利用[WebSocket](https://www.websocket.org/)来实时的显示爬虫抓取模特的信息
* 以分页的方式显示模特照片

## ToDo
* 将爬虫的抓取过程修改为并发实现，缩短爬虫抓取的时间
* 将Http请求的路由过程更改为统一的路由处理，避免在程序中定义过多的处理函数
* 提供[REST](https://zh.wikipedia.org/wiki/REST)方式的对外Web Service接口
* 利用[美空网](http://www.moko.cc/)的账号进行模拟登录，从而对每个模特抓取完整的图片
* 对模特照片进行添加评论与显示

## Display
* 爬虫操作页面<br/>
![爬虫操作页面](http://i.imgur.com/HXvzD2c.png)
* 模特列表页面<br/>
![模特列表页面](http://i.imgur.com/mFdg9Ph.png)