# GitOps Demo App [![Build Status](https://drone.surunhao.com/api/badges/surh-kube/demo-app/status.svg?ref=refs/heads/main)](https://drone.surunhao.com/surh-kube/demo-app)

<br/>

GitOps演示App，使用golang和gin搭建的web项目，简单阐述Drone的使用方法。

<br/>

<br/>

## Dockerfile

<br/>

golang能够直接编译成可执行的二进制文件，所以运行镜像选择alpine即可。

``` dockerfile
FROM alpine
WORKDIR /app
COPY main /app/server
ENTRYPOINT ["/app/server"]
```

<br/>

<br/>

## Drone

<br/>

### I 触发器

我们希望在master、develop分支进行提交动作时，或者进行tag时触发构建流程

``` yaml
# 触发器, 当分支为Master、Develop或进行Tag时触发
trigger:
  ref:
    - refs/heads/master
    - refs/heads/develop
    - refs/tags/**


```

<br/>

### II 步骤

#### Tag

我们希望不同的分支打上不同的镜像Tag，需要如下步骤进行处理

``` yaml
# 镜像Tag
# 默认使用Git Commit SHA作为容器Tag
# 当分支为Develop时, 增加latest作为镜像Tag
# 存在Git Tag时, 增加Git Tag作为镜像Tag
- name: tags
  image: alpine
  commands:
    - echo -n "$DRONE_COMMIT_SHA" >> .tags
    - $([ "$DRONE_BRANCH" = "develop" ] && echo -n ",latest" >> .tags || echo "")
    - $([ -n "$DRONE_TAG" ] && echo -n ",$DRONE_TAG" >> .tags || echo "")
    - cat .tags
```

<br/>

### 测试

我们希望对代码进行测试，需要如下步骤处理

``` yaml
# 测试代码
- name: test
  image: golang
  environment:
    GOPROXY: https://goproxy.cn
  # 挂载目录, 共享Go依赖, 避免重复下载
  volumes:
    - name: deps
      path: /go
  commands:
    - go test
```

<br/>

### 构建

``` yaml
# 构建代码
- name: build
  image: golang
  environment:
    GOPROXY: https://goproxy.cn
  # 挂载目录, 共享Go依赖, 避免重复下载
  volumes:
    - name: deps
      path: /go
  commands:
    - CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o main .
  # 等待tags和test阶段, tags和test阶段并行执行
  depends_on:
    - tags
    - test
```

``` yaml
# 构建镜像
- name: docker
  image: plugins/docker
  settings:
    dockerfile: Dockerfile
    repo: harbor.surunhao.com/demo/app
    registry: harbor.surunhao.com
    username:
      # 敏感信息使用密文, 在Drone中对密文进行配置
      from_secret: REGISTRY_USERNAME
    password:
      # 敏感信息使用密文, 在Drone中对密文进行配置
      from_secret: REGISTRY_PASSWORD
  depends_on:
    - build
```

