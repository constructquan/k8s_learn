### README
> update: 2022.03.06



### 第一部分

**镜像制作说明**

使用module3制作的镜像: [httpserver 镜像](https://gitee.com/ahchpr/xunlian/tree/main/module3/httpserver)



**更新说明：**

由于要和 ingress  controller 的namespace 相同，所以，这次统一使用 ingress 的 namespace 名。



**创建Service**

```shell
kubectl -n ingress create -f online-svc.yaml

```



**创建ConfigMap**

```shell
kubectl -n ingress create configmap httpserver-cm --from-file=httpserverconfig.yaml 

```



**创建Deployment**

```shell
kubectl -n ingress create -f  online-deployment.yaml

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

**探活功能**

Deployment 增加 livenessProbe 探针。



**增加Pod的QoS限制**

Deployment使用resource.request和resource.limit对cpu和memory资源需求配额。



### 第二部分

**安装 helm **

```shell
i. 下载最新版本：helm-v3.8.0-linux-amd64.tar.gz  
ii. 解压缩到：/usr/local/src/
iii. 将可执行文件迁移到$PATH 目录： mv linux-amd64/helm  /usr/local/bin/helm
```

**通过helm安装 nginx-ingress**

- 添加nginx-ingress仓库：

  ```shell
   # helm repo add nginx-stable https://helm.nginx.com/stable
   # helm repo update 
  ```

- 下载 nginx-ingress 到本地

  ```shell
  helm fetch nginx-stable/nginx-ingress
  ```

- 修改 nginx-ingress 参数，再安装。（由于暂时没有LoadBalancer 服务器）

  ```shell
  # tar -xzf  nginx-ingress-0.12.1.tgz
  # cd nginx-ingress
  修改 values.yaml 文件：
    hostNetwork: true  -> 修改为 hostNetWork: true
    type: LoadBalancer -> 修改为 type: ClusterIP
  然后安装 nginx-ingress 
  # helm install --namespace ingress ingress-nginx ./nginx-ingress -f ./nginx-ingress/values.yaml 
  ```

- 查看nginx-ingress 控制器（以pod的形式出现）

  ```shell
  root@k8smaster:~/personer/xunlian/module8# ki get pods -owide 
  NAME                                           READY   STATUS    RESTARTS   AGE   IP               NODE       NOMINATED NODE   READINESS GATES
  ingress-nginx-nginx-ingress-56bc9d5bcf-pxg5d   1/1     Running   0          13h   172.26.102.174   k8snode1   <none>           <none>
  ```

- 创建新的 Service

  ```shell
  kubectl -n ingress create -f online-svc.yaml
  
  ```

- 创建 Ingress (无 tls 签名证书)

  ```shell
  # cat  online-ingress.yaml
  
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: online-ingress
    annotations:
      kubernetes.io/ingress.class: "nginx"
  spec:
    rules:
    - host: fengyiqi.com  # 将域名映射
      http:
        paths:
        - path: /version
          pathType: Prefix
          backend:
             service:
               name: online-svc 
               port:
                 number: 80
  
  ```

  创建：

  ```shell
  kubectl -n ingress create online-ingress.yaml
  
  ```

  绑定域名到 /etc/hosts，因为 nginx ingress controller 在 172.26.102.174 的节点，所以，绑定hosts:

  ```shell
  echo "172.26.102.174  test.fengyiqi.com"
  ```

  通过 ingress 访问pod的web服务：

  ```shell
  root@k8smaster:~/personer/xunlian/module8# curl fengyiqi.com/version -I
  HTTP/1.1 200 OK
  Server: nginx/1.21.6
  Date: Sun, 06 Mar 2022 04:13:01 GMT
  Connection: keep-alive
  Version: v1.0.6
  ```

- 创建 ingress (增加 tls 签名)

  - 首先，创建tls证书：

    ```sell
    # openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=fengyiqi.com/O=fengyiqi" -addext "subjectAltName = DNS:fengyiqi.com"
    ```

  - 然后，在k8s创建secret

    ```shell
    # kubectl -n ingress  create secret tls fengyiqi-tls --cert=./tls.crt --key=./tls.key
    ```

  - 然后，创建新的 online-ingress.yaml

  ```shell
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: online-ingress
    annotations:
      kubernetes.io/ingress.class: "nginx"
  spec:
    tls:
    - hosts:
      - fengyiqi.com
      secretName: fengyiqi-tls
    rules:
    - host: fengyiqi.com  # 将域名映射
      http:
        paths:
        - path: /version
          pathType: Prefix
          backend:
             service:
               name: online-svc 
               port:
                 number: 80
  ```

  创建：

  ```shell
  kubectl -n ingress create -f online-ingress.yaml
  
  ```

  测试结果：

  ```shell
  root@k8smaster:~/personer/xunlian/module8# curl -H "Host: fengyiqi.com" https://test.fengyiqi.com/version -v -k
  *   Trying 172.26.102.174:443...
  * TCP_NODELAY set
  * Connected to test.fengyiqi.com (172.26.102.174) port 443 (#0)
  * ALPN, offering h2
  * ALPN, offering http/1.1
  * successfully set certificate verify locations:
  *   CAfile: /etc/ssl/certs/ca-certificates.crt
    CApath: /etc/ssl/certs
  * TLSv1.3 (OUT), TLS handshake, Client hello (1):
  * TLSv1.3 (IN), TLS handshake, Server hello (2):
  * TLSv1.2 (IN), TLS handshake, Certificate (11):
  * TLSv1.2 (IN), TLS handshake, Server key exchange (12):
  * TLSv1.2 (IN), TLS handshake, Server finished (14):
  * TLSv1.2 (OUT), TLS handshake, Client key exchange (16):
  * TLSv1.2 (OUT), TLS change cipher, Change cipher spec (1):
  * TLSv1.2 (OUT), TLS handshake, Finished (20):
  * TLSv1.2 (IN), TLS handshake, Finished (20):
  * SSL connection using TLSv1.2 / ECDHE-RSA-AES256-GCM-SHA384
  * ALPN, server accepted to use http/1.1
  * Server certificate:
  *  subject: CN=NGINXIngressController
  *  start date: Sep 12 18:03:35 2018 GMT
  *  expire date: Sep 11 18:03:35 2023 GMT
  *  issuer: CN=NGINXIngressController
  *  SSL certificate verify result: self signed certificate (18), continuing anyway.
  > GET /version HTTP/1.1
  > Host: fengyiqi.com
  > User-Agent: curl/7.68.0
  > Accept: */*
  > 
  * Mark bundle as not supporting multiuse
  < HTTP/1.1 200 OK
  < Server: nginx/1.21.6
  < Date: Sun, 06 Mar 2022 04:36:49 GMT
  < Content-Length: 0
  < Connection: keep-alive
  < Version: v1.0.6
  < 
  * Connection #0 to host test.fengyiqi.com left intact
  root@k8smaster:~/personer/xunlian/module8# 
  ```

  

  
