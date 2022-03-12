**安装loki stack**

loki-prometheus，loki-grafana 这两个 svc 需要使用nodePort方式，让它在服务器能够被外网访问。



安装：

```shell
# helm repo add grafana https://grafana.github.io/helm-charts


# helm upgrade --install loki grafana/loki-stack --set grafana.enabled=true,prometheus.enabled=true,prometheus.alertmanager.persistentVolume.enabled=false,prometheus.server.persistentVolume.enabled=false

```

结果：

```shell
root@k8smaster:~# k get ep
NAME                            ENDPOINTS                                 AGE
kubernetes                      172.26.102.173:6443                       5d10h
loki                            10.42.0.1:3100                            3m14s
loki-grafana                    10.42.0.4:3000                            3m14s
loki-headless                   10.42.0.1:3100                            3m14s
loki-kube-state-metrics                                                   3m14s
loki-prometheus-alertmanager    10.42.0.3:9093                            3m14s
loki-prometheus-node-exporter   172.26.102.174:9100,172.26.102.175:9100   3m14s
loki-prometheus-pushgateway     10.44.0.3:9091                            3m14s
loki-prometheus-server          10.44.0.2:9090                            3m14s
root@k8smaster:~# k get pods 
NAME                                            READY   STATUS             RESTARTS   AGE
loki-0                                          1/1     Running            0          3m26s
loki-grafana-68b589d7c4-qmrft                   2/2     Running            0          3m26s
loki-kube-state-metrics-6d8fdf5fd8-ncfvg        0/1     ImagePullBackOff   0          3m26s
loki-prometheus-alertmanager-5d9f5b4dd7-hwz9s   2/2     Running            0          3m26s
loki-prometheus-node-exporter-5xxld             1/1     Running            0          3m26s
loki-prometheus-node-exporter-95mb8             1/1     Running            0          3m26s
loki-prometheus-pushgateway-7cf69b85f7-5lrn7    1/1     Running            0          3m26s
loki-prometheus-server-7566499d5c-4d84b         2/2     Running            0          3m26s
loki-promtail-29fqc                             1/1     Running            0          3m26s
loki-promtail-cwh2r                             1/1     Running            0          3m26s
loki-promtail-l4dsf                             1/1     Running            0          3m26s
root@k8smaster:~# k get svc 
NAME                            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
kubernetes                      ClusterIP   10.96.0.1       <none>        443/TCP    5d10h
loki                            ClusterIP   10.109.93.127   <none>        3100/TCP   3m46s
loki-grafana                    ClusterIP   10.106.254.63   <none>        80/TCP     3m46s
loki-headless                   ClusterIP   None            <none>        3100/TCP   3m46s
loki-kube-state-metrics         ClusterIP   10.108.72.76    <none>        8080/TCP   3m46s
loki-prometheus-alertmanager    ClusterIP   10.107.27.129   <none>        80/TCP     3m46s
loki-prometheus-node-exporter   ClusterIP   None            <none>        9100/TCP   3m46s
loki-prometheus-pushgateway     ClusterIP   10.99.64.124    <none>        9091/TCP   3m46s
loki-prometheus-server          ClusterIP   10.99.229.119   <none>        80/TCP     3m46s
root@k8smaster:~# 
```

修改 NodePort的方式(自定义它们的端口)：

```shell

i. kubectl   edit svc loki-prometheus-server -oyaml -n default

ii. kubectl  edit svc loki-grafana -oyaml -n default 
```

查看 grafana 的密码：

```shell
查看密码：

kubectl   get secret loki-grafana -oyaml

查看： admin-password :  echo "xxxxx" | base64 -d 

```



**创建 httpserver 的 pod**

```shell
kubectl create namespace httpserver

kubectl -n httpserver  create configmap httpserver-cm --from-file=../module8/httpserverconfig.yaml 

//创建deployment 
kubectl -n httpserver create -f online-deployment.yaml 

```



**实验效果**

- 获取 httpserver pod 的地址

  ```shell
  root@k8smaster:~# k get pods -n httpserver  -owide 
  NAME                          READY   STATUS    RESTARTS   AGE   IP          NODE       NOMINATED NODE   READINESS GATES
  httpserver-7c568fcb98-t7gv7   1/1     Running   0          16h   10.42.0.5   k8snode2   <none>           <none>
  ```

- 访问 /delay uri

  ```shell
  curl 10.42.0.5:8080/delay 
  ```

- 在prometheus 查看数据