# 基于 gRPC 的聊天室实现

## 架构

### API

两个 rpc service： `MessageService` 与 `PushService`：

`MessageService` 处理客户端发送新消息以及历史消息的范围获取，接口定义在[这里](pkg/message/message.proto)。

`PushService` 使用了 gRPC streaming 接口做长连接新消息实时推送，接口定义在[这里](pkg/push/push.proto)。

### 服务端逻辑

实现 `MessageService` 接口的服务端接收到客户端的发送新消息请求后，生成消息对象并持久化到数据库中（当前用的是 mongo），并将消息对象提交至该消息所属聊天室的消息队列。

实现 `PushService` 接口的服务端根据客户端的连接情况，从消息队列订阅相应聊天室的消息队列。消息队列中得到的消息对象通过与客户端使用 streaming 方法的长连接推送给客户端。

消息对象的 `id` 规定为细粒度增序（当前实现使用 ulid 达到粒度为 1ms 的增序），客户端可以在 `PushService` 接口暂时不可用时通过 `MessageService` 接口用起始消息 `id` 范围来获取历史消息，避免消息丢失。

实现 `MessageService` 接口的服务端是无状态的，实现 `PushService` 接口的服务端因为长连接的存在是有状态的，两者之间不耦合。目前这两个接口暂时都由 `oneserver` 同时提供，之后会真正拆分这两部分（很简单）以在实际环境中分离部署。

## Docker Compose 演示

运行服务端

```bash
$ docker-compose -f demo/oneserver/compose-server.yml up -d
```

运行客户端

```bash
$ docker-compose -f demo/oneserver/compose-client.yml run --rm client oneserverX:8000 oneserverX:8001 chatroom username
```

其中 `oneserverX` 可以替换为 `oneserver1` 或者 `oneserver2`，代表所连接的服务端 `oneserver `实例编号。此处模拟服务端的横向扩容，实际场景中客户端经 load balencr 连接到不同的服务器实例均能正常工作。

`chatroom` 与 `username` 可换成目标聊天室与用户名。

![screencast](screencast.gif)

## Roadmap

- 分离消息服务与推送服务（各自 scale）

- Kubernetes 部署演示

- Observability
