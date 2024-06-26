# Tisea Backend 开发事项列表

Tisea Backend 分为两大平台，目前正在开发的是 Lab（2024.4-）

## 参考

- [云服务器 ECS_API调试-阿里云OpenAPI开发者门户](https://api.aliyun.com/api/Ecs/2014-05-26)
- [阿里云 Billing_API调试-阿里云OpenAPI开发者门户](https://api.aliyun.com/api/BssOpenApi/2017-12-14)

## Lab 开发事项

- Auth 统一认证
  - [x] *POST* `/register` 注册
  - [x] *POST* `/login` 登录+Token 生成与发放
- ECS Aliyun 交互
  - *GET* `/ecs/describe` 实例状态查询 DescribeInstance
      - [x] 实现
      - [x] 测试
  - *GET* `/ecs/create` 创建实例 CreateInstance
    - [x] 实现
    - [x] 测试
  - *DELETE* `/ecs/delete` 删除实例 DeleteInstance
    - [x] 实现
    - [x] 测试
  - *GET* `/ecs/start`, `/ecs/stop`, `/ecs/reboot` <br/> 开启、关闭和重启实例 StartInstance, StopInstance, RebootInstance
    - [x] 实现
    - [x] 测试
  - *GET* `/ecs/price-history` 抢占式价格查询 DescribeSpotPriceHistory
    - [ ] 实现
    - [ ] 测试
  - *GET* `/ecs/deploy-status` 部署状态查询
    - [x] 实现
    - [x] 测试
- BSS Aliyun 交互
  - *GET* `/bss/balance` 账户余额
    - [x] 实现
    - [x] 测试
  - *GET* `/bss/transactions` 消费记录
    - [x] 实现
    - [x] 测试
- Monitor 监控与自动化任务
  - 自动部署任务 DeployMonitor
    - [x] 实现
    - [x] 测试
  - 实例停机空闲检测任务 StoppedInstanceMonitor
    - [x] 实现
    - [x] 测试
  - 实例部署状态自动更新任务 DeployStatusMonitor
    - [x] 实现
    - [x] 测试