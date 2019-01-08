### 队列

#### work-queue

![工作队列](https://www.rabbitmq.com/img/tutorials/python-two.png)

`./producer -repeat 100` 发送 100 条信息，用两个 consumer 去接受消息，可以看到消息被发送到两个 consumer 中。

#### work-queue-1

![工作队列-1](https://www.rabbitmq.com/img/tutorials/prefetch-count.png)

producer：声明 queue 时设置 durable 为 true。
consumer：声明 queue 时设置 durable 为 true，消费队列时设置 auto-ack 为 false，手动确认消息是否发送成功。调用 Qos 设置 prefetch，来控制消息如何传送。

#### publish-subscrible

![发布订阅](https://www.rabbitmq.com/img/tutorials/exchanges.png)

使用 exchange 把消息传送到两个 queue 中。producer 创建 exchange，consumer 也需要创建 exchange，创建匿名 queue，并且将 queue 与 exchange 绑定。其他与上述例子大致相同。

#### routing

![路由](https://www.rabbitmq.com/img/tutorials/python-four.png)

指定 routing-key，就可以 exchange 可以把消息发送到指定的 queue 中。在 producer 发布消息时指定 routing-key，把 type 设置为 direct。consumer 消费消息前，绑定 queue 时设置 routing-key 即可。
