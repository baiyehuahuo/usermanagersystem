package model

type ConfigModel struct {
	UserAccount string `yaml: "useraccount"`
	Password    string `yaml: "password"`
	Host        string `yaml: "host"`
	Port        int    `yaml: "port"`
	DbName      string `yaml: "dbname"`
}
