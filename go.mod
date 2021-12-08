module usermanagersystem

go 1.14

replace UserManageSystem => ./

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.17.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.2.0
	gorm.io/gorm v1.22.3
)
