[docker](http://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html)
[docker部署go web程序](https://www.liwenzhou.com/posts/Go/how_to_deploy_go_app_using_docker/#autoid-0-2-1)
### image文件是什么？如何使用？
+ image文件：
Docker 把应用程序及其依赖，打包在 image 文件里面。
只有通过这个文件，才能生成 Docker 容器。
image 文件可以看作是容器的模板。Docker 根据 image 文件生成容器的实例。
同一个 image 文件，可以生成多个同时运行的容器实例。
image 是二进制文件。实际开发中，一个 image 文件往往通过继承另一个 image 文件，加上一些个性化设置而生成。

+ 使用image：

将image文件从仓库抓到本地：

    $ docker image pull library/image_name

在本机查看image文件：

    $ docker image ls

运行image文件：

    $ docker container run image_name

docker container run命令会从 image 文件，生成一个正在运行的容器实例。

有些容器不会自动终止，因为提供的是服务。对于那些不会自动终止的容器，必须使用docker container kill 命令手动终止。

    $ docker container kill [containID]

+ 容器文件：
image 文件生成的容器实例，本身也是一个文件，称为容器文件。
一旦容器生成，就会同时存在两个文件： image 文件和容器文件。
关闭容器并不会删除容器文件，只是容器停止运行而已。

```
# 列出本机正在运行的容器
$ docker container ls

# 列出本机所有容器，包括终止运行的容器
$ docker container ls --all

# 终止运行的容器文件，依然会占据硬盘空间，可以使用docker container rm命令删除。
$ docker container rm [containerID]
```
### 如何制作image文件
+ Dockerfile 文件
如果你要推广自己的软件，势必要自己制作 image 文件。
这就需要用到 Dockerfile 文件。它是一个文本文件，用来配置 image。
Docker 根据 该文件生成二进制的 image 文件。

+ 制作Dockerfile 文件

在项目的根目录下

+ + 文本文件.dockerignore
    示例：

        .git
        node_modules
        npm-debug.log
    上面代码表示，这三个路径要排除，不要打包进入 image 文件。如果你没有路径要排除，这个文件可以不新建。

+ + 文本文件 Dockerfile

```
# 使用基础镜像golang:alpine来创建我们的镜像。这和我们要创建的镜像一样是一个我们能够访问的存储在Docker仓库的基础镜像。
FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 指定接下来的工作路径为：/build
WORKDIR /build

# 将当前目录下的所有文件（除了.dockerignore排除的路径），都拷贝进入 image 文件的/build目录。
COPY ..

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /build 目录复制到当前的 /dist 目录
RUN cp /build/app .

# 将容器 8888 端口暴露出来， 允许外部连接这个端口。
EXPOSE 8888

# 启动容器时第一个运行的命令
CMD ["/dist/app"]
```
+ 创建image文件

在项目目录下，执行下面的命令创建镜像，并指定镜像名称为goweb_app：

    docker build . -t goweb_app
    或
    docker image build . -t goweb_app

+ 运行镜像，生成容器

执行下面的命令来运行镜像，从 image 文件生成容器。

    docker run -p 8080:8888 -it goweb_app
    或
    docker container run -p 8080:8888 -it goweb_app
    或
    docker container run -p 8080:8888 -it goweb_app /dist/app

-p参数：容器的 8888 端口映射到本机的 8080 端口。
-it参数：容器的 Shell 映射到当前的 Shell，然后你在本机窗口输入的命令，就会传入容器。

+ 举例：
```
# 镜像别名为builder
FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 app
RUN go build -o app .

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/app /

# 需要运行的命令
ENTRYPOINT ["/app"]
```