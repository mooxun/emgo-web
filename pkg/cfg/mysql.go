package cfg

// Mysql 配置
type Mysql struct {
	Host      string
	Port      string
	Database  string
	Username  string
	Password  string
	Charset   string
	Collation string
	Prefix    string
	Engine    string
	Debug     bool
}
