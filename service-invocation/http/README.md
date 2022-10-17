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
