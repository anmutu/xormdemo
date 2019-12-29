package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"sync"
	"xormdemo191226/conf"
)

var (
	masterEngineer *xorm.Engine
	slaveEngineer  *xorm.Engine
	lock           sync.Mutex
)

//创建master的实例
func InstanceMaster() *xorm.Engine {
	if masterEngineer != nil {
		return masterEngineer
	}
	//考虑到并发，需要用到互斥锁
	lock.Lock()
	defer lock.Unlock()

	//如果现在来了10个，有一个进入并创建了，其他9个被锁住了，则他们还会再创建，这样就不单例了。所以需要再这里再判断一下。
	if masterEngineer != nil {
		return masterEngineer
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.MySqlDriver, driveSource)
	if err != nil {
		//log.Fatal("创建Master的engineer的实例出错")
		log.Fatalf("err：%s。小杜同学定位：创建Master的engineer的实例出错。", err)
		return nil
	} else {
		masterEngineer = engine
		return masterEngineer
	}
	return masterEngineer
}

func InstanceSlave() *xorm.Engine {
	if slaveEngineer != nil {
		return slaveEngineer
	}
	//考虑到并发，需要用到互斥锁
	lock.Lock()
	defer lock.Unlock()

	//如果现在来了10个，有一个进入并创建了，其他9个被锁住了，则他们还会再创建，这样就不单例了。所以需要再这里再判断一下。
	if slaveEngineer != nil {
		return slaveEngineer
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.MySqlDriver, driveSource)
	if err != nil {
		log.Fatal("创建Slave的engineer的实例出错")
		return nil
	} else {
		slaveEngineer = engine
		return slaveEngineer
	}
	return slaveEngineer
}
