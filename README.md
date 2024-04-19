# Tisea Backend 开发事项列表

Tisea Backend 分为两大平台，目前正在开发的是 Lab（2024.4-）

常用参考

- [云服务器 ECS_API调试-阿里云OpenAPI开发者门户](https://api.aliyun.com/api/Ecs/2014-05-26)

## Lab 开发事项

- Auth 统一认证
  - [x] *POST* `/register` 注册
  - [x] *POST* `/login` 登录+Token 生成与发放
- ECS Aliyun 交互
  - *POST* `/ecs/describe` 实例状态查询 DescribeInstance
      - [x] 实现
      - [ ] 测试
  - *POST* `/ecs/create` 创建实例 CreateInstance
    - [x] 实现
    - [ ] 测试
  - *POST* `/ecs/delete` 删除实例 DeleteInstance
    - [x] 实现
    - [ ] 测试
  - *POST* `/ecs/start`, `/ecs/stop`, `/ecs/reboot` <br/> 开启、关闭和重启实例 StartInstance, StopInstance, RebootInstance
    - [x] 实现
    - [ ] 测试
  - *POST* `/ecs/price-history` 抢占式价格查询 DescribeSpotPriceHistory
    - [ ] 实现
    - [ ] 测试
- BSS Aliyun 交互
  - *GET* `/bss/balance` 账户余额 · [参考 API: QueryAccountBalance](https://api.aliyun.com/api/BssOpenApi/2017-12-14/QueryAccountBalance)
    - [ ] 实现
    - [ ] 测试
  - *GET* `/bss/bill` 消费记录 · [参考 API: QueryBill](https://api.aliyun.com/api/BssOpenApi/2017-12-14/QueryBill)
    - [ ] 实现
    - [ ] 测试
- Monitor 监控与自动化任务
  - 实例初始化完毕检测任务 PendingInstanceMonitor
    - [ ] 实现
    - [ ] 测试
  - 实例停机空闲检测任务 StoppedInstanceMonitor
    - [x] 实现
    - [ ] 测试
  - 抢占式价格记录任务 SpotPriceRecordMonitor
    - [ ] 实现
    - [ ] 测试