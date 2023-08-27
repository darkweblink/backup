# backup gatway server for personal study use

#### 介绍
网关服务，支持get、post请求，接口统计、授权识别等
Gateway service, support get, post request, interface statistics, authorization identification, etc.

#### 软件架构
软件架构说明


#### 此网关支持的功能
1. API管理【仅支持固定API，不支持通配符匹配】
2. 安全控制

    1. 目前仅支持JWT登录授权
    2. 目前仅支持固定key的接口授权、时间授权、次数授权

3. 统一监控
4. 负载均衡

    1. 多客户端管理【随机命中】

5. 超时默认返回
6. 大流量熔断【每分钟最大多少次请求】

#### 注意事项

1. 本程序使用了endless来进行平滑重启，即如果有路由更新的话，请执行 `kill -1 xxxx` 来进行在不影响用户访问的情况下进行重载配置
2. 如果有新版本程序的话，等同于条目1，即先覆盖掉源程序，然后执行 `kill -1 xxxx` 来进行平滑重启
3. 本程序的配置文件均属于数据库内配置，所以入口配置中仅保留redis配置和数据库配置
4. 本程序使用了redis作为主数据存储，数据库内容仅在启动时扫描一遍，所以使用时请注意不要随意关闭redis库


