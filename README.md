# 记事本

### 安装流程(install)

(以0.1.0为例)

下载合适源代码版本并解压。

```
cd notebook-0.1.0
```

安装相应依赖并build。

```
go mod tidy
go build main.go
```

后续根据需要自行操作即可。

### 注意(notice)

默认开启端口8089，如需修改，可前往 notebook/config/config.go 修改。

目前尚无 user 数据库，即所有用户共享数据库。

暂时无删除功能。