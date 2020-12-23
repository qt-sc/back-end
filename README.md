# back-end
the back end for zhihu_coffee

## 简介

Zhihu-coffee后端基于Golang、MySql和Redis，可以部署到Windows、Linux和MacOS上，提供数据存储、知乎日报数据拉取、文章发布和分享功能。

## Using

### 部署MySql

| 数据库参数名 | 参数      |
| ------------ | --------- |
| Name         | mydb      |
| Username     | root      |
| Password     | root      |
| Address      | localhost |
| Port         | 3306      |

注意：需要调整MySql的编码为utf8或utf8mb4以适应中文字符。

### 部署Redis

| 数据库参数名 | 参数           |
| ------------ | -------------- |
| Address      | localhost:6379 |
| Password     |                |
| DB           | 0              |

### 运行

```
go run main.go
```

### docker打包镜像

```
docker build -t server .
```

### docker运行

```
docker -d --name back-end --network host server:latest
```

## 设计与功能



## 数据库设计

### article

```go
type Article struct {

	Id int64 `json:"id,omitempty" gorm:"id"`

	Title string `json:"title,omitempty" gorm:"title"`

	ReadNum int64 `json:"readNum,omitempty" gorm:"read_num"`
 
	LikeNum int64 `json:"likeNum,omitempty" gorm:"like_num"`
 
	Content string `json:"content,omitempty" gorm:"type:text;content"`

	UserID int64 `json:"user_id,omitempty" gorm:"user_id"`

	Replies []Reply `json:"replies,omitempty" gorm:"replies"`
 
	Tags []Tag `json:"tags,omitempty" gorm:"tags;many2many:article_tags`

	Url string `json:"url,omitempty" gorm:"url"`
}
```

## 关于测试

### 鉴权模块测试

对于鉴权模块来说，直接用postman和数据库、redis的效果查看测试结果远比用代码编写单元测试要直观有效。因此，鉴权模块使用postman做测试。

##### 注册接口：/users/signup

![image-20201222142407791](img/image-20201222142407791.png)

可见信息能够成功记录，注册成功。

##### 登录接口：/users/login

![image-20201222143153846](img/image-20201222143153846.png)

可见用户用相应的用户名和密码登录，能够成功响应，获取相应的cookie。查看redis也能确定多了一条token记录：

![image-20201222145310362](img/image-20201222145310362.png)

##### 登出接口：/users/logout

首先需要访问受限资源：

![image-20201222145429262](img/image-20201222145429262.png)

可见能够正常访问。然后携带cookie向localhost:8080/users/logout发get请求登出，再次试图访问受限资源:

![image-20201222153050142](img/image-20201222153050142.png)

显然，已经无法访问受限资源了。可见鉴权模块能够正常工作。
