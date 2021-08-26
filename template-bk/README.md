# algo-golang-dispatcher

## Get Source Code

### 准备开发环境
- consul 是必须的，用于服务发现，可通过docker镜像启动一个consul服务
```
docker run -d --name=consul -p 8500:8500 consul agent -server -ui -bootstrap-expect=1 -client=0.0.0.0
```

- 如需要其他的库或者mq环境，自行准备

### 编译工程并执行
- 运行程序
```bash
// 进入启动目录
cd ./cmd
// 本地编译可执行文件，并将文件命名为exec
go build -v -o ./exec
// 执行编译后的文件
./exec --config ../config
```
或者
```
go run cmd/search/main.go --config="./config"
```

- 执行测试
```bash

```

