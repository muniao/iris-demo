package conf

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var DbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "",
	DbName: "test",
}
