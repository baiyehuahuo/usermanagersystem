package model

type ConfigModel struct {
	MysqlConfig    mysqlConfig `yaml: "mysqlconfig"`
	RedisConfig    redisConfig `yaml: "redisconfig"`
	RabbitMQConfig rabbitMQConfig
	EmailConfig    emailConfig `yaml: "emailconfig"`
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

type rabbitMQConfig struct {
	// Conn *amqp.Connection
}

type emailConfig struct {
	Email    string `yaml: "email"`
	Addr     string `yaml: "addr"`
	UserName string `yaml: "username"`
	Password string `yaml: "password"`
	Host     string `yaml: "host"`
}
