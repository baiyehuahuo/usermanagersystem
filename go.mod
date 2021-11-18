module usermanagersystem

go 1.14

replace UserManageSystem => ./

require (
	UserManageSystem v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.4
	github.com/kylelemons/go-gypsy v1.0.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)
