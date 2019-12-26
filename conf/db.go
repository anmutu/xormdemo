package conf

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "graduation",
	Pwd:    "graduation",
	DbName: "golangtestdb",
}

var SlaveDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "graduation",
	Pwd:    "graduation",
	DbName: "golangtestdb",
}
