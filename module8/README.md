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

![image-20220227155143105](../Library/Application Support/typora-user-images/image-20220227155143105.png)



**优雅终止**





**配置和代码分离**

修改ConfigMap，能获得不同的版本号：

![image-20220227155403439](../Library/Application Support/typora-user-images/image-20220227155403439.png)
