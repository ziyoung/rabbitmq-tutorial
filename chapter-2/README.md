### 队列

#### work-queue

![工作队列](https://www.rabbitmq.com/img/tutorials/python-two.png)

`./producer -repeat 100` 发送 100 条信息，用两个 consumer 去接受消息，可以看到消息被发送到两个 consumer 中。

#### work-queue-1

![工作队列-1](https://www.rabbitmq.com/img/tutorials/prefetch-count.png)

producer：声明 queue 时设置 durable 为 true。
consumer：声明 queue 时设置 durable 为 true，消费队列时设置 auto-ack 为 false，手动确认消息是否发送成功。调用 Qos 设置 prefetch，来控制消息如何传送。
