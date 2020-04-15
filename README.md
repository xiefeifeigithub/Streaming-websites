# 流媒体点播网站

# 流媒体网站整体介绍与架构梳理

## 1.1 为什么选择视频网站

> ```
> 1. Go是一门网络编程语言
> 2. 视频网站包含Go在实战项目中绝大部分技能要点
> 3. 优良的native http库以及模板引擎（无需任何第三方框架）
> ```

## 1.2 总体架构：前后端分离

> ```
> 1. API会将一些业务的数据往DB里写。
> 2. Streaming模块：API会将视频的播放请求、上传和下载请求传入该模块。
> 3. Scheduler模块：会处理一些删除、软删除和定期清理等之类的事情。
> 4. Streaming和Scheduler都会同时访问DB，然而最重要的是会直接去文件系统里找相应的视频文件并做相应的处理。
> ```



<img src="README.assets/Snipaste_2020-04-14_20-44-03.png" alt="Snipaste_2020-04-14_20-44-03" style="zoom:80%;" />

## 1.3 前后端解耦

### 1.3.1 什么是前后端解耦

> ```
> 1. 前后端解耦是时下流行的web网站架构。
> 2. 前端页面和服务通过普通web引擎渲染。
> 3. 后端数据通过渲染后的页面脚本调用后处理和呈现。
> ```

### 1.3.2 前后端解耦的优势

> ```
> 1. 解放生产力，提高合作效率。
> 2. 松耦合的架构更灵活，部署更方便，更符合微服务的设计特征。
> 3. 性能的提升，可靠性的提升。
> ```

### 1.3.3 前后端解耦的缺点

> ```
> 1. 工作量大。
> 2. 前后端分离带来的团队成本以及学习成本。
> 3. 系统复杂度加大。
> ```





# API设计与架构

## 2.1 api

> ```
> 1. REST(Representational Status Transfer)API。
> 2. REST是一种设计风格，不是任何架构标准。
> 3. 当今RESTful API通常使用HTTP作为通信协议，JSON作为数据格式。
> 4. 统一接口(Uniform Interface)。
> 5. 无状态(Stateless)。
> 6. 可缓存(Cacheable)。
> 7. 分层(Layered System)。
> 8. CS模式(Client-server Atchitecture)。
> ```

## 2.2 API设计原则

> ```
> 1. 以URL(统一资源定位符)风格设计API。
> 2. 通过不同的METHOD(GET,POST,PUT,DELETE)来区分对资源的CRUD。
> 3. 返回码(Status Code)符合HTTP资源描述的规定。
> ```

## 2.3 API设计图

<img src="README.assets/image-20200414211121188.png" alt="image-20200414211121188" style="zoom:80%;" />

### 2.3.1 API设计：用户

> ```
> 1. 创建（注册）用户：URL:/user Method:POST,SC:201,400,500
> 2. 用户登录：URL:/user/:username Method:POST,SC:200,400,500
> 3. 获取用户基本信息：URL:/user/:username METHOD:GET,SC:200,400,401,403,500
> 4. 用户注销：URL:/user/:username Method:DELETE,SC:204,400,401,403,500
> ```

### 2.3.2 API设计：资源（视频）

> ```
> 1. List all videos: URL:/user/:username/videos Method:GET,SC:200,400,500
> 2. Get one video：URL:/user/:username/videos/:vid-id Method:GET,SC:200,400,500
> 3. DELETE one video：URL/user/:username/videos/:vid-id Method:DELETE,SC:204,400,401,403,500
> ```

### 2.3.3 API设计：评论

> ```
> 1. Show comments: URL:/videos/:vid-id/comments Method:GET,SC:200,400,500
> 2. Post a comment: URL:/videos/:vid-id/comments Method:POST,SC:201,400,500
> 3. Delete a comment: URL:/videos/:vid-id/comment/:comment-id Method:DELETE,SC:2O4,400,401,403,500
> ```

## 2.4 api之httphandler层设计

> ```
> API请求过程：handler->validation{1.request, 2.user}->business logic->response.
> 1. data model
> 2. error handling
> 注意：
> 1. 针对API请求过程对API请求采用分层架构的方式对其编写代码，可扩展、高可用、易于维护。
> 2. 对于request的处理采用这种分层架构对于编写test case是很容易的，
> 而且更能照顾到它的可扩展性，对工程上的效率也是非常高的。
> ```

## 2.5 api之数据库层设计

```
因为数据库模型的设计、数据库表的设计直接关系到最终业务逻辑的处理方式，所以首先对数据库层进行设计。
```

<img src="README.assets/image-20200415074349139.png" alt="image-20200415074349139" style="zoom:80%;" />

### 2.5.1 表结构

> ```mysql
> CREATE TABLE `users` (
>   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
>   `login_name` varchar(64) DEFAULT NULL,
>   `pwd` text NOT NULL,
>   PRIMARY KEY (`id`),
>   UNIQUE KEY `login_name` (`login_name`)
> ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
> 
> CREATE TABLE `video_info` (
>   `id` varchar(64) NOT NULL,
>   `author_id` int(10) unsigned DEFAULT NULL,
>   `name` text,
>   `display_ctime` text,
>   `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
>   PRIMARY KEY (`id`)
> ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
> 
> CREATE TABLE `comments` (
>   `id` varchar(64) NOT NULL,
>   `video_id` varchar(64) DEFAULT NULL,
>   `author_id` int(10) unsigned DEFAULT NULL,
>   `content` text,
>   `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
>   PRIMARY KEY (`id`)
> ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
> 
> CREATE TABLE `sessions` (
>   `session_id` varchar(64) NOT NULL,
>   `TTL` tinytext,
>   `login_name` varchar(64) DEFAULT NULL,
>   PRIMARY KEY (`session_id`)
> ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
> ```

