# Service Invocation

1. 构建镜像
```shell
# client镜像
sh client/build.sh

# server镜像
sh server/build.sh
```

2. 创建应用

```shell
kubectl apply -f deployment.yaml
```

3. 开启trace功能

需要加一个注解`dapr.io/config: "tracing"`

参考: [config zipkin for kubernetes](https://docs.dapr.io/operations/monitoring/tracing/zipkin/#configure-kubernetes)

4. 设置resiliency

```yaml
kubectl apply -f resiliency.yaml
```

参考: [resiliency overview](https://docs.dapr.io/operations/resiliency/resiliency-overview/)