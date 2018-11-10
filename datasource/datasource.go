/**
 * 获取数据库连接
 */
package datasource

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"org.mm/iris-demo/conf"
	"org.mm/iris-demo/models"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// 单例
func Instance() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}
	c := conf.DbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("dbhelper.DbInstance,", err)
		return nil
	}
	// Debug模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)
	engine.SetTZLocation(conf.SysTimeLocation)

	// 性能优化的时候才考虑，加上本机的SQL缓存
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)
	engine.Sync2(new(models.Movie))
	masterEngine = engine
	return engine
}