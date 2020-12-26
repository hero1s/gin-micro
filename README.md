gin + micro 框架测试

# install protoc-gen-go
go get github.com/golang/protobuf/{proto,protoc-gen-go}
# install protoc-gen-micro
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master

#查看etcd key
etcdctl get / --prefix --keys-only --endpoints=172.16.3.21:2379

#安装糗百qbtool工具
go install git.moumentei.com/plat_go/qbtool



go mod 私有仓库配置(目前golib没有设置成public,所以要开启ssh 及配置 GOPRIVATE)
https://blog.csdn.net/xuduorui/article/details/103753155