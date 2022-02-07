# single-sign-on-system

## 一：SSO单点登录系统开发总结

### (一)：整体架构分析

- 基于go-oauth2/oauth2库实现的前端分离SSO单点登录系统

### (二)：系统技术点分析

- 当前系统的业务技术栈如下

  - Vue3 、ElementUI 作为前端页面
  - Nginx 用于解决系统之间的跨域请求，和前端登录的负载均衡
  - Gin 提供Web层HTTP请求处理
  - go-oauth2/oauth2 提供 Oauth2.0授权码认证模型
  - casbin 为user-server提供RBAC权限模型
  - Redis 提供session会话和Token的存储
  - MySQL 提供数据存储
  - ...... 持续开发中
