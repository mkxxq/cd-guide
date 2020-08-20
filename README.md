## 容器化部署应用指北

## 常用命令
### compose版本
* 查看应用

```
$ docker-compose ps
        Name                       Command               State            Ports         
----------------------------------------------------------------------------------------
cd-guide_helloworld_1   /bin/sh -c ./helloworld          Up      0.0.0.0:11323->1323/tcp
cd-guide_redis_1        docker-entrypoint.sh redis ...   Up      6379/tcp
```


* 查看日志

```bash
# 直接查看
$ docker-compose logs helloworld
Attaching to cd-guide_helloworld_1
helloworld_1  | 
helloworld_1  |    ____    __
helloworld_1  |   / __/___/ /  ___
helloworld_1  |  / _// __/ _ \/ _ \
helloworld_1  | /___/\__/_//_/\___/ v4.1.16
helloworld_1  | High performance, minimalist Go web framework
helloworld_1  | https://echo.labstack.com
helloworld_1  | ____________________________________O/_______
helloworld_1  |                                     O\
helloworld_1  | ⇨ http server started on [::]:1323

# 查看尾部5行
$ docker-compose logs --tail 5 helloworld
Attaching to cd-guide_helloworld_1
helloworld_1  | High performance, minimalist Go web framework
helloworld_1  | https://echo.labstack.com
helloworld_1  | ____________________________________O/_______
helloworld_1  |                                     O\
helloworld_1  | ⇨ http server started on [::]:1323

# 监视日志最近10行, 当有新的日志时,会滚蛋
$ docker-compose logs --tail 10 -f helloworld
Attaching to cd-guide_helloworld_1
helloworld_1  | 
helloworld_1  |    ____    __
helloworld_1  |   / __/___/ /  ___
helloworld_1  |  / _// __/ _ \/ _ \
helloworld_1  | /___/\__/_//_/\___/ v4.1.16
helloworld_1  | High performance, minimalist Go web framework
helloworld_1  | https://echo.labstack.com
helloworld_1  | ____________________________________O/_______
helloworld_1  |                                     O\
helloworld_1  | ⇨ http server started on [::]:1323
```

* 执行容器内部命令

```
# 直接执行
$ docker-compose exec helloworld ls     
helloworld

# 进入容器
$ docker-compose exec helloworld /bin/sh
/app # pwd
/app
/app # ls
helloworld
/app # 
```