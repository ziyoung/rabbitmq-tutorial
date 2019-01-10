#rabbitmq-tutorial

[RabbitMQ实战指南](https://book.douban.com/subject/27591386/) 是一本很适合初学者的书籍，该书对 RabbitMQ 中的许多概念讲解十分透彻。由于书中的示例代码较少（一些例子较为简单，放到书中也不合适），且例子都是通过 Java 语言编写。为了方便 Go 语言开发者使用 RabbbitMQ，把书中的例子用 Go 语言实现一遍，同时为第三章以及第四章补充了一些例子。

#### 安装

为了尽快上手开发，我们可以使用 docker-compose。

```bash
# 安装驱动
go get github.com/streadway/amqp 
# 准备 volumn
mkdir -p ~/docker-data/rabbitmq

docker-compose up

# 浏览器打开 localhost:15672
# 用户名 guest 密码 guest
# 通过管理后台添加用户 root，设置密码为 root123，设置权限
```

#### 开发

go build 构建完成后，直接运行即可。
