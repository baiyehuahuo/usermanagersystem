package model

type ConfigModel struct {
	MysqlConfig mysqlConfig `yaml: "mysqlconfig"`
	RedisConfig redisConfig `yaml: "redisconfig"`
}

type mysqlConfig struct {
	UserAccount string `yaml: "useraccount"`
	Password    string `yaml: "password"`
	Host        string `yaml: "host"`
	Port        int    `yaml: "port"`
	DbName      string `yaml: "dbname"`
}

type redisConfig struct {
	Host     string `yaml: "host"`
	Port     int    `yaml: "port"`
	Password string `yaml: "password"`
	DbNum    int    `yaml: "dbnum"`
}
