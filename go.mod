module usermanagersystem

go 1.14

replace UserManageSystem => ./

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/sbinet/go-python v0.1.0 // indirect
	gopkg.in/yaml.v2 v2.2.8
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)
