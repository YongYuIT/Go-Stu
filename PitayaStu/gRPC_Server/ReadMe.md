# 编译proto文件

下载protoc工具：https://github.com/protocolbuffers/protobuf/releases/download/v25.1/protoc-25.1-win64.zip

配置protoc工具到path

生成协议文件的Go语言版本代码

~~~
protoc.exe .\proto_src\GetFullNameParams.proto --go_out .\pb
~~~

报错：'protoc-gen-go' 不是内部或外部命令，也不是可运行的程序或批处理文件。

安装 protoc-gen-go

~~~
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
~~~

再次生成协议文件的Go语言版本代码

~~~
protoc.exe .\proto_src\GetFullNameParams.proto --go_out .\pb
~~~

协议文件的Go语言版本代码生成成功之后，生成服务端接口Go语言版本代码

要注意指定import目录（-I）
~~~
protoc.exe .\proto_src\GetFullNameService.proto -I .\proto_src\ --go-grpc_out .\pb
~~~

报错：'protoc-gen-go-grpc' 不是内部或外部命令，也不是可运行的程序或批处理文件。

安装 protoc-gen-go-grpc

~~~
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
~~~

再次生成服务端接口Go语言版本代码

~~~
protoc.exe .\protos\GetFullNameService.proto -I .\protos\ --go-grpc_out .\pb
~~~


