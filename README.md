# gin-vue-essential
<<<<<<< HEAD

## 一、项目内容

基于gin框架的 Gin+Vue 前后端分离实战

**master**分支为后端go代码,**vue**分支为前端vue代码


## 二、怎样运行该项目

### 1. 运行后端程序
> 

从master分支拉取后端golang代码
```shell
# 拉取代码
git clone -b master https://github.com/nongeason/gin-vue-essential.git backend
# 进入项目目录
cd  backend
# 安装项目依赖
go get
```
打开 `config/application.yaml` 文件，修改数据库链接配置，修改项目运行端口，确保端口不被占用，参考如下

启动项目
```
go run routes.go main.go
```

### 2. 运行前端程序
> 先确保你电脑上正确安装了 npm 环境，并安装了 vue、yarn
> 

从vue分支拉取前端vue代码
```shell
# 拉取代码
git clone -b vue https://github.com/nongeason/gin-vue-essential.git vue
# 进入项目目录
cd  vue
# 安装项目依赖
yarn install
```

根据1中的 后端代码的运行端口，修改 `.env.development.local` 和 `.env.development` 两个配置文件，修改配置如下为
```
VUE_APP_BASE_URL = http://localhost:8080/api/
```

在运行项目
```shell
yarn serve
```
=======
>>>>>>> e53f37e0f4f19cb3616753d59a08fc61b4ca9049
