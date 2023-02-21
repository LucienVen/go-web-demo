# go-web-demo
试验go-web项目搭建（包括，docker, k8s, grpc）



### 使用环境

* mysql
* GRPC （TODO 转换 rpcx ）

### 服务提供：

* user 用户服务
* order 订单服务



## USER 用户服务

> 当前版本 v1

#### web接口功能：

| 请求方式 | 接口（路由） | 功能         | 描述 |
| -------- | ------------ | ------------ | ---- |
| POST     | user         | 新增用户     |      |
| PUT      | user/:id     | 用户信息编辑 |      |
| DELETE   | user/:id     | 删除用户     |      |
| GET      | user/:id     | 查询用户信息 |      |

rpc服务：

| 接口（路由） | 功能         | 描述          |
| ------------ | ------------ | ------------- |
| GetUserData  | 获取用户信息 | Param: userId |
|              |              |               |
|              |              |               |



## ORDER 订单服务

>  当前版本 v1

#### web接口功能：

| 请求方式 | 接口（路由） | 功能         | 描述 |
| -------- | ------------ | ------------ | ---- |
| POST     | order        | 新增订单     |      |
| PUT      | order/:id    | 订单信息编辑 |      |
| DELETE   | order/:id    | 删除订单     |      |
| GET      | order/:id    | 查询订单信息 |      |
|          |              |              |      |

rpc服务：

| 接口（路由） | 功能         | 描述           |
| ------------ | ------------ | -------------- |
| GetOrder     | 获取订单信息 | Param: orderId |
|              |              |                |
|              |              |                |



### GRPC 备注
* 因当前为了方便，预计将全部服务的proto文件放在同一目录下，因此可能会产生生成的 pb 文件中有 数据结构被重复定义（因为命名相同）。故方法命名方式暂定义为 
  > (server_name)(func_name|message_name)
