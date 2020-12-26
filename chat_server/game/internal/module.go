package internal

import (
	"gin-micro/chat_server/base"
	"gin-micro/chat_server/conf"
	"gin-micro/model"
	"github.com/hero1s/golib/cache"
	"github.com/hero1s/golib/connsvr/module"
	"github.com/hero1s/golib/db"
	"github.com/hero1s/golib/db/dbsql"
	"github.com/hero1s/golib/log"
	"time"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	//SubscribeRedisMsg()
	//订阅监听队列消息
	//链接mysql
	for _, v := range conf.Config.Mysql {
		err := db.InitDBConf(v, 10, 100)
		if err != nil {
			log.Error("数据库初始化失败", err.Error())
		}
	}
	t := db.TableOper{Tb: model.TbChannel}
	o := t.NewOrm()
	o.Raw("select * from channel where 1").Exec()
	o.Raw("update channel set a=1 where b=c").Exec()
	o.Raw("select count(1) from user where 1").Exec()

	cache.InitCache(conf.Config.Redis, "test")
	cache.InitRedis(conf.Config.Redis)
	count := t.CountRecord("1")
	log.Info(count)
	cids := make([]int64, 0)
	total, err := dbsql.MultiRecordAndTotal("select cid from channel where cid>0 limit 1,1", &cids, o)
	//cids := make([]struct {
	//	Cid int64 `json:"cid"`
	//}, 0)
	//total, err := t.MultiRecordByAnyOrderLimit("1", "cid desc", 0, 1, &cids)

	//o.Raw("select SQL_CALC_FOUND_ROWS * from channel where 1 limit 0,10").Exec()
	//o.Raw("SELECT FOUND_ROWS() AS total").QueryRow(&count)
	log.Infof("total:%v,%v,%v", cids, total, err)

	cache.Redis.Set("test", count, time.Minute)
	cache.SetCache(cache.MemCache, "test", count, time.Minute)
	cache.SetCache(cache.RedisCache, "test", count, time.Minute)
}

func (m *Module) OnDestroy() {

}
