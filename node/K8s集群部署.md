## 部署master节点

1. 服务器准备

   >master 节点： 2核4G ubuntu 18.04
   >
   >node1 节点： 1核2G ubuntu 18.04
   >
   >

2. 首先进行换源（阿里云的源）

   ```shell
   ]$ apt-get update && apt-get install -y apt-transport-https
   
   ]$ curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -
   
   // 将deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main 加入到/etc/apt/sources.list.d/kubernetes.list文件中，文件不存在就创建
   
   ]$ echo deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main > /etc/apt/sources.list.d/kubernetes.list
   
   ```

3. 安装 kubelet、kubeadm、kubectl 以及 docker

   ```shell
   ]$ apt-get updat
   
   ]$ apt-get install -y kubelet kubeadm kubectl
   
   ]$ apt-get install -y docker.io
   ```

4. 准备master节点要用的 yaml 文件，vim kubeadm.yaml：

   ```shell
   apiVersion: kubeadm.k8s.io/v1beta3 #版本信息参考kubeadm config print init-defaults命令结果
   kind: ClusterConfiguration
   kubernetesVersion: 1.23.1 #根据自己安装的k8s版本来写,版本信息参考kubeadm config print init-defaults命令结果
   imageRepository: registry.aliyuncs.com/google_containers #配置国内镜像
   
   apiServer:
       extraArgs:
           runtime-config: "api/all=true"
   
   #controllerManager:
   ## extraArgs:
   ## horizontal-pod-autoscaler-use-rest-clients: "true" #开启kube-controller-manager能够使用自定义资源（Custom Metrics）进行自动水平扩展,但是高版本不支>持该参数需要去掉。
   ## horizontal-pod-autoscaler-sync-period: "10s"
   ## node-monitor-grace-period: "10s"
   
   #etcd:
   #  local:
   #      dataDir: /data/k8s/etcd # 由于服务器只有一个硬盘，就不区分了，直接用默认/var/lib/etc 
   ```

5. 利用写好的kubeadm.yaml部署master节点

   ```shell
   kubeadm init --config kubeadm.yaml
   
   ```

   这一步如有报错：

   ```shell
   [wait-control-plane] Waiting for the kubelet to boot up the control plane as static Pods from directory "/etc/kubernetes/manifests". This can take up to 4m0s
   [kubelet-check] Initial timeout of 40s passed.
   [kubelet-check] It seems like the kubelet isn't running or healthy.
   [kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.
   [kubelet-check] It seems like the kubelet isn't running or healthy.
   [kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.
   [kubelet-check] It seems like the kubelet isn't running or healthy.
   [kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.
   [kubelet-check] It seems like the kubelet isn't running or healthy.
   [kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.
   [kubelet-check] It seems like the kubelet isn't running or healthy.
   [kubelet-check] The HTTP call equal to 'curl -sSL http://localhost:10248/healthz' failed with error: Get "http://localhost:10248/healthz": dial tcp [::1]:10248: connect: connection refused.
   
           Unfortunately, an error has occurred:
                   timed out waiting for the condition
   
           This error is likely caused by:
                   - The kubelet is not running
                   - The kubelet is unhealthy due to a misconfiguration of the node in some way (required cgroups disabled)
   
           If you are on a systemd-powered system, you can try to troubleshoot the error with the following commands:
                   - 'systemctl status kubelet'
                   - 'journalctl -xeu kubelet'
   
           Additionally, a control plane component may have crashed or exited when started by the container runtime.
           To troubleshoot, list all containers using your preferred container runtimes CLI.
   
           Here is one example how you may list all Kubernetes containers running in docker:
                   - 'docker ps -a | grep kube | grep -v pause'
                   Once you have found the failing container, you can inspect its logs with:
                   - 'docker logs CONTAINERID'
   
   error execution phase wait-control-plane: couldn't initialize a Kubernetes cluster
   To see the stack trace of this error execute with --v=5 or higher
   
   
   ```

    应该是安装kubeadm和docker后，这俩使用的cgroup驱动不一致导致。需要在指定docker的cgroup驱动为system. 具体做法为:

   ```shell
   vim /etc/docker/daemon.json
   ```

   json 文件的内容为：

   ```shell
   {
     
             "live-restore": true,
             "exec-opts": ["native.cgroupdriver=systemd"],
             "log-driver": "json-file",
             "log-opts": {
                    "max-size": "100m"
            }
   }
   ```
   
   保存退出后，继续：
   
   ```shell
   systemctl restart docker
   ```
   
   再次应用 kubectl  init --config  kubeadm.yaml 成功后，显示如下信息：
   
   ```shell
   Your Kubernetes control-plane has initialized successfully!
   
   To start using your cluster, you need to run the following as a regular user:
   
     mkdir -p $HOME/.kube
     sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
     sudo chown $(id -u):$(id -g) $HOME/.kube/config
   
   Alternatively, if you are the root user, you can run:
   
     export KUBECONFIG=/etc/kubernetes/admin.conf
   
   You should now deploy a pod network to the cluster.
   Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
     https://kubernetes.io/docs/concepts/cluster-administration/addons/
   
   Then you can join any number of worker nodes by running the following on each as root:
   
   kubeadm join 172.26.153.236:6443 --token z9pr14.s1is23wqh3514epb \
           --discovery-token-ca-cert-hash sha256:d52392f5f240babddb108a3d28be58d5e95dc8a197f09108c5fe8687ba64d4ec
   
   ```
   
   **注意上面的 export 命令和 kubeadm join 两条命令，可以先将它们复制保存。**
   
6. 集群安全访问配置，**每次重新init master节点，这一步都要重新执行**。

   ```shell
   mkdir -p $HOME/.kube
   sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
   sudo chown $(id -u):$(id -g) $HOME/.kube/config
   
   ```

   或者，如使用的是root用户：

   ```shell
   ]$ export KUBECONFIG=/etc/kubernetes/admin.conf
   ```

7. 部署网络插件

   如果这一步不部署，kubectl  的命令会显示网络连接失败。

   通过如下命令，部署网络插件：

   ```shell
   ]$ kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')"
   #kubeadm config print init-defaults可以告诉我们kubeadm.yaml版本信息。
   
   ```

   如果你确定系统使用的kubernetes的版本，例如现在安装的是 1.23.1，那么可以使用如下的命令：

   ```shell
   ]$ kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=1.23.1"
   
   ```

   部署网络插件成功后，查看集群的 pod 状态（需要等待一小会时间）

   ```shell
   ]$ kubectl get pods -n kube-system
   
   ```

   **如果结果中，所有的 pod 都显示为running状态，则表示 master 节点已经部署成功。**

## 部署slave节点

1. 在slave节点，先安装kubectl、kubeadm、kubelet

   ```shell
   }$ apt-get update && apt-get install -y apt-transport-https
   ]$ curl https://mirrors.aliyun.com/kubernetes/apt/doc/apt-key.gpg | apt-key add -
   ]$ echo deb https://mirrors.aliyun.com/kubernetes/apt/ kubernetes-xenial main > /etc/apt/sources.list.d/kubernetes.list
   ]$ apt-get update
   ]$ apt-get install -y kubelet kubeadm kubectl
   ]$ apt-get install -y docker.io
   
   ```

2. 配置docker的cgroup驱动：

   vim   /etc/docker/damon.json 

   ```sehll
   {
     
             "live-restore": true,
             "exec-opts": ["native.cgroupdriver=systemd"],
             "log-driver": "json-file",
             "log-opts": {
                    "max-size": "100m"
            }
   }
   ```
   
3. 把slave节点加入集群

   在部署 master 成功的时候，有一个kubeadm join 的命令提示，里面有token(令牌)，但是令牌是24小时有效的。过期了需要再生成token。

   查看token的命令：

   ``` shell
   ]$ kubeadm token list
   ```

   创建新的token：

   ```shell
   ]$ kubeadm token create
   ```

   如果“--discovery-token-ca-cert-hash” 也忘记了，可以用如下命令查看：

   ```shelll
   ]$ openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'
   ```

   **slave 加入集群前，需要保持节点之间的网络能互通互联，如果是云服务器，则可以开放防火墙，让节点在内网IP之间保持互相访问。**

   slave 节点加入集群：

   ```shell
   ]$ kubeadm join 10.0.12.2:6443 --token mljn91.g3wg36bbil71tf15 \
           --discovery-token-ca-cert-hash sha256:7c0d1882fb6a882749eeed52aa03b2148bd1802da6b605b8aa1131762c90e1d9 
           
   ```

4. 在master节点检查节点是否成功加入

   ```shell
   ]$ kubectl get nodes
   ```

   如果还有其他服务器需要配置成slave节点，则可以如法炮制。

## 为Kubernetes集群安装dashboard

1. 安装 dashboard ：

   也可以自行到github 的项目地址，找到对应的版本，进行安装。

```shell
]$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-rc6/aio/deploy/recommended.yaml
```

​		或者，先将recommended.yml 文件下载到本地。(有可能网络不顺畅)

```shell
]$ wget https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0/aio/deploy/recommended.yaml
```

​		查看是否安装成功：

```shell
]$ kubectl get svc -n kubernetes-dashboard
NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)         AGE
dashboard-metrics-scraper   ClusterIP   10.100.28.165   <none>        8000/TCP        16h
kubernetes-dashboard        NodePort    10.103.186.17   <none>        443:30349/TCP   16h
# 拥有kubernetes-dashboard service则表示安装成功


```

 2. 更新无效的secrete，通过自签证书重新创建

    ```shell
    ]$ kubectl edit svc kubernetes-dashboard -n kubernetes-dashboard
    1 # Please edit the object below. Lines beginning with a '#' will be ignored,
    2 # and an empty file will abort the edit. If an error occurs while saving this file will be
    3 # reopened with the relevant failures.
    4 #
    5 apiVersion: v1
    6 kind: Service
    7 metadata:
    8   creationTimestamp: "2020-04-21T17:34:01Z"
    9   labels:
    10     k8s-app: kubernetes-dashboard
    11   name: kubernetes-dashboard
    12   namespace: kubernetes-dashboard
    16 spec:
    17   clusterIP: 10.105.195.26
    18   ports:
    19   - port: 443
    20     protocol: TCP
    21     targetPort: 8443
    22   selector:
    23     k8s-app: kubernetes-dashboard
    24   sessionAffinity: None
    25   type: ClusterIP    # < ----- 改为 NodePort
    ```

    验证svc是否修改成功：

    ```shell
    ]$ kubectl get svc -n kubernetes-dashboard
    NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)         AGE
    dashboard-metrics-scraper   ClusterIP   10.100.28.165   <none>        8000/TCP        16h
    kubernetes-dashboard        NodePort    10.103.186.17   <none>        443:30349/TCP   16h
    
    # service 的类型变为NodePort, 表明修改成功
    ```

	3. 删除默认的secret:  kubernetes-dashboard-certs

    ```shell
    ]$ kubectl delete secret -n kubernetes-dashboard kubernetes-dashboard-certs
    ```

	4. 自签发证书（为访问的ip，如果使用的是云主机，想用公网ip访问，则使用公网ip）

    ```shell
    ]$ openssl genrsa -out tls.key 2048
    ]$ openssl req -new -out tls.csr -key tls.key -subj '/CN=114.14.208.126'
    ]$ openssl x509 -req -in tls.csr -signkey tls.key  -out tls.crt
    
    ```

	5. 根据新的证书，创建secret:

    ```shell
    ]$ cd keys # 如果本身就在keys文件夹下，则可以省略该步骤
    
    ]$ kubectl create secret generic kubernetes-dashboard-certs --from-file=./ -n kubernetes-dashboard
    ```

	6. 修改kubernetes-dashboard deployment，启用新的secret：

    ```shell
    ]$ kubectl edit deploy kubernetes-dashboard -n kubernetes-dashboard
    ...
     31   template:
     32     metadata:
     33       creationTimestamp: null
     34       labels:
     35         k8s-app: kubernetes-dashboard
     36     spec:
     37       containers:
     38       - args:
     41         - --auto-generate-certificates
     42         - --namespace=kubernetes-dashboard
     43         image: kubernetesui/dashboard:v2.0.0
     44         imagePullPolicy: Always
    ...
    
    # 在args中添加两行
           containers:
           - args:
             - --tls-cert-file=/tls.crt
             - --tls-key-file=/tls.key
          
          
    # 添加之后
    ...
     31   template:
     32     metadata:
     33       creationTimestamp: null
     34       labels:
     35         k8s-app: kubernetes-dashboard
     36     spec:
     37       containers:
     38       - args:
     39         - --tls-cert-file=/tls.crt   < ----- 这里
     40         - --tls-key-file=/tls.key    < ----- 这里
     41         - --auto-generate-certificates
     42         - --namespace=kubernetes-dashboard
     43         image: kubernetesui/dashboard:v2.0.0
     44         imagePullPolicy: Always
     ...
    ```

	7. 查看dashboard 修改后的端口：(每次修改，dashboard的端口都会改变。)

    ```shell
    ]$ kubctl get svc -n kubernetes-dashboard
    ```

	8. 在浏览器通过ip+端口的形式，访问。（使用安全等级要求低的浏览器）

	9. 在访问dashboard的时候，可能会要求属于令牌。

    所以，需要先创建一个。

- dashboard 登录需要令牌

  **接下来创建访问账号：**

  ```shell
  # 准备一个yaml文件: vim dash.yaml  填入以下内容
  apiVersion: v1
  kind: ServiceAccount
  metadata:
    name: admin-user
    namespace: kubernetes-dashboard
  ---
  apiVersion: rbac.authorization.k8s.io/v1
  kind: ClusterRoleBinding
  metadata:
    name: admin-user
  roleRef:
    apiGroup: rbac.authorization.k8s.io
    kind: ClusterRole
    name: cluster-admin
  subjects:
  - kind: ServiceAccount
    name: admin-user
    namespace: kubernetes-dashboard
  
  ```

  

  ```shell
  ]# kubectl apply -f dash.yaml
  
  执行以下命令，生成登录用的token:
  ]# kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')
  
  ```

  

  

将token复制到登录页面，就可以打开了。





## slave 节点如何运行 kubectl  命令

首先，kubectl config view --raw  # 查看kubectl 配置的文件，可以将它们放到node，然后node节点就可以执行kubectl 命令操作集群了。

例如在节点执行：

kubectl get pods -A --config=kube-config.yml









