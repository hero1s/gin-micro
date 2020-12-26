package main

import (
	"flag"
	"gin-micro/api/clients"
	"gin-micro/api/routers"
	"github.com/hero1s/golib/cache"
	"github.com/hero1s/golib/conf"
	"github.com/hero1s/golib/db"
	"github.com/hero1s/golib/log"
	smicro "github.com/hero1s/golib/micro"
	"github.com/hero1s/golib/third_sdk/nacos"
	"github.com/hero1s/golib/utils"
	sweb "github.com/hero1s/golib/web"
	"github.com/gin-gonic/gin"
)

type Conf struct {
	Etcd   smicro.EtcdRegistry `json:"etcd"`
	Mysql  []db.DbConf         `json:"mysql"`
	Redis  cache.RedisConf     `json:"redis"`
	Web    sweb.Config         `json:"web"`
	Jaeger smicro.JaegerConf   `json:"jaeger"`
}

var (
	confPath = flag.String("conf", "", "config path, example: -conf /config.toml")
)

// @title gin-micro 服务框架测试
// @version 0.0.1
// @description  接口文档
// @BasePath /v1/
func main() {
	if flag.Parsed() {
		log.Error("已经解析")
	} else {
		flag.Parse()
	}

	var config Conf
	err := conf.AutoParseFile(*confPath, &config)
	if err != nil {
		log.Errorf(err.Error())
	} else {
		log.Infof("%+v", config)
	}
	//测试配置中心
	var endpoint = "acm.aliyun.com"
	var namespaceId = "4d2e32d3-ffa4-4668-b213-495eecfaf27c"
	// 推荐使用 RAM 用户的 accessKey、secretKey
	var accessKey = "nafgL5OLBZiDgbfU"
	var secretKey = "zpfC1OBProcFxEhXbe7pEuMD5L0MO0"
	nacos.InitConfigClient(endpoint, namespaceId, accessKey, secretKey)
	nacos.GetConfigToStruct("sysconfig.mysql", "dev", &config.Mysql)
	nacos.GetConfigToStruct("sysconfig.redis", "dev", &config.Redis)
	nacos.GetConfigToStruct("sysconfig.etcd", "dev", &config.Etcd)

	log.Infof("%+v", config)

	smicro.InitGlobalTracer(config.Jaeger)

	clients.InitClient(config.Etcd)

	g := sweb.InitGinServer(&config.Web)
	g.Gin.Use(sweb.RequestId())
	g.Gin.Use(sweb.LogRequest(true))
	//g.Gin.Use(g.JWTAuth())
	g.Register(routers.RegisterUser)
	if gin.EnvGinMode != gin.ReleaseMode {
		// 需要提前准备好swagger相关资源到api/swagger, 直接make doc会自动下载
		g.Gin.Static("swagger", "swagger")
	}

	httpAddr := config.Web.Host + config.Web.Port
	log.Infof("%v,%v,listen httpAddr:%v", config.Web.Host, config.Web.Port, httpAddr)

	utils.RunMain(func() error {
		g.Start()
		return nil
	}, func() {

	}, "")

}
