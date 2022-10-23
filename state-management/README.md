# State Management

1. 安装redis
```shell
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
helm install redis bitnami/redis --set image.tag=6.2
```

安装完成后可以通过下面的域名访问redis集群:
* 读写: redis-master.default.svc.cluster.local:6379
* 只读: redis-replicas.default.svc.cluster.local:6379

使用下面的命令获取redis密码:
```
export REDIS_PASSWORD=$(kubectl get secret --namespace default redis -o jsonpath="{.data.redis-password}" | base64 -d)
```

创建一个redis client实例，方便后面使用:
```shell
# 创建实例
kubectl run --namespace default redis-client --restart='Never'  --env REDIS_PASSWORD=$REDIS_PASSWORD  --image docker.io/bitnami/redis:6.2 --command -- sleep infinity
# 登录容器
kubectl exec --tty -i redis-client --namespace default -- bash
# 查询数据
REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h redis-master
REDISCLI_AUTH="$REDIS_PASSWORD" redis-cli -h redis-replicas
```

参考: 
- [Tutorial: Configure state store and pub/sub message broker](https://docs.dapr.io/getting-started/tutorials/configure-state-pubsub/#step-1-create-a-redis-store)
- [hello-kubernetes quickstart](https://github.com/dapr/quickstarts/tree/master/tutorials/hello-kubernetes)

2. 创建component

```shell
kubectl apply -f statestore.yaml
```

3. 创建应用
```
kubectl apply -f deployment.yaml
```

4. 直接从redis查询

```shell
redis-master:6379> keys *
1) "state-demo||orderType"
2) "state-demo||orderId"
3) "nodeapp||order"
redis-master:6379> type state-demo||orderId
hash
redis-master:6379> hgetall state-demo||orderId
1) "data"
2) "345"
3) "version"
4) "1781"
redis-master:6379> hget state-demo||orderId data
"646"
```