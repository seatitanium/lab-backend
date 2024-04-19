# Tisea Backend 开发事项列表

Tisea Backend 分为两大平台，目前正在开发的是 Lab（2024.4-）

常用参考

- [云服务器 ECS_API调试-阿里云OpenAPI开发者门户](https://api.aliyun.com/api/Ecs/2014-05-26)

## Lab 开发事项

- Auth 统一认证
  - [x] 注册
  - [x] 登录+Token 生成与发放
- ECS Aliyun 交互
  - 创建实例 CreateInstance
    - [x] 实现
    - [ ] 测试
  - 删除实例 DeleteInstance
    - [x] 实现
    - [ ] 测试
  - 开启、关闭和重启实例 StartInstance, StopInstance, RebootInstance
    - [x] 实现
    - [ ] 测试
  - 抢占式价格查询 DescribeSpotPriceHistory
    - [x] 实现
    - [ ] 测试
- Monitor 监控与自动化任务
  - 实例停机空闲检测任务 StoppedInstanceMonitor
    - [x] 实现
    - [ ] 测试
  - 抢占式价格记录任务 SpotPriceRecordMonitor
    - [ ] 实现
    - [ ] 测试