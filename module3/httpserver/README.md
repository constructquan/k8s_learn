### README
>update: 2022.02


**镜像制作**



```shell
docker build -t hellochenpro7799/httpserver:latest   .
```



**推送镜像至dockerhub仓库发布**

```shell
docker push hellochenpro7799/httpserver
```





****

**容器启动**

```shell
docker run -d -p 8080:8080 --name httpserver-1 hellochenpro7799/httpserver:latest
```



**访问功能**

- 获取默认页

  > http://IP:8080/

- 获取version

  > http://IP:8080/version

- 获取healthz

  > http://IP:8080/healthz

  
