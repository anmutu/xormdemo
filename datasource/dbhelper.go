package datasource

import (
	"github.com/go-xorm/xorm"
	"sync"
)

var (
	masterEnginer *xorm.Engine
	slaveEnginer  *xorm.Engine
	lock          sync.Mutex
)

func InstanceMaster() *xorm.Engine {
	//TODO
	return masterEnginer
}

func InstanceSlave() *xorm.Engine {
	//TODO
	return slaveEnginer
}
