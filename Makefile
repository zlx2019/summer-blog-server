## 编译Protobuf文件
## -I=. 指定源文件目录,从哪个目录开始寻找Proto依赖。
## --go_out=. 表示go文件生成在当前目录
## --go-grpc_out=. 表示go-grpc文件生成在当前目录
build:
	protoc -I=. --go_out=. --go-grpc_out=. ./proto/*.proto