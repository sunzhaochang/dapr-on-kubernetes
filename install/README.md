# dapr install

1. 搭建kubernetes集群
这里使用kind来搭建k8s集群, kind安装参考[kind installation](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)

创建集群:
```shell
kind create cluster
```

2. 安装dapr
```shell
# install
dapr init -k

# uninstall
dapr uninstall -k
```