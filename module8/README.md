### README
> update: 2022.02



**镜像制作说明**

使用module3制作的镜像: [httpserver 镜像](https://gitee.com/ahchpr/xunlian/tree/main/module3/httpserver)



**创建Service**

```shell
kubectl create -f httpsvc.yaml

```



**创建ConfigMap**

```shell
kubectl create -f  httpserver-cm.yaml

```



**创建Deployment**

```shell
kubectl create -f  online-deployment.yaml

```



**优雅启动**

使用readinessProbe探针,延时30秒检测,获取 /healthz 的状态码为200,pod才就绪.


**优雅终止**


创建Deployment时，使用 lifecycle.preStop，添加脚本退出httpserver程序。




**配置和代码分离**

修改ConfigMap，能获得不同的版本号：
第一次获取version:
```shell
# curl  -I  10.96.9.250:8080/version
HTTP/1.1 200 OK
Version: v1.0.6
Date: Sun, 27 Feb 2022 10:52:37 GMT
```

修改configMap:
```shell
# kubectl edit cm httpserver-cm

将version字段修改为 v1.0.7
```

再次获取version:
```shell
# curl  -I  10.96.9.250:8080/version
HTTP/1.1 200 OK
Version: v1.0.7
Date: Sun, 27 Feb 2022 10:55:19 GMT
```
