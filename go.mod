module usermanagersystem

go 1.14

replace UserManageSystem => ./

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis v6.15.9+incompatible
	gopkg.in/yaml.v2 v2.2.8
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)
