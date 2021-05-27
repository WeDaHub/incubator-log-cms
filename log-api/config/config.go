package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Cfg struct {
	Server Server	`yaml:"server"`
	MySQL  MySQL	`yaml:"mysql"`
}

type Server struct {
	Address string `yaml:"address"`
	Prefix  string `yaml:"prefix"`
}


type MySQL struct {
	User 			string		`yaml:"user"`
	Password		string		`yaml:"password"`
	Host			string 		`yaml:"host"`
	Database		string 		`yaml:"db"`
	Charset			string 		`yaml:"charset"`
	MaxIdleCount	int			`yaml:"maxidle"`
	MaxConnCount	int			`yaml:"maxconn"`
}

var (
	cfg Cfg
)

func SrvCfg() (srv *Server) {
	srv = &cfg.Server
	return
}

func MySQLCfg() (mysql *MySQL) {
	mysql = &cfg.MySQL
	return
}

func (this *MySQL) GetDSN() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@%s/%s?charset=%s&parseTime=true", this.User, this.Password, this.Host, this.Database, this.Charset)
	return
}

func init() {
	if file, err := os.Open("./config/config.yaml"); nil == err{
		yaml.NewDecoder(file).Decode(&cfg)
	} else {
		cfg = Cfg{
			MySQL:MySQL{
				User:"logcms",
				Password: "aA123!@#",
				Host:"tcp(118.24.10.209:3306)",
				Database: "logcms",
				Charset: "utf8mb4",
				MaxIdleCount: 100,
				MaxConnCount: 300,
			},
			Server: Server{
				Address: "0.0.0.0:8848",
				Prefix:"/api/",
			},
		}
	}
}
