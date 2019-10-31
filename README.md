# cloudwatch_exporter
aws cloudwatch exporter


fork from yet another cloudwatch exporter

主要用于一些aws服务的监控

- 定时获取置顶metric信息，而不是访问是被动获取
- 多账号，多region
- 补充Tag
- 汇总与分开显示，例如`/metric`, `region/type/metric`
- 资源Tag过滤
