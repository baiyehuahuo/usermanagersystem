# UserManagerSystem

1. 前端暂用 https://gitee.com/xminghua/Login.git ，待小组成员补充
2. 后端用 `go` 开发，`gin` 为服务器， `gorm` 查询， MySQL 为数据库
3. 运行用 `go run .`
4. 系统日志文件在主目录下的 `systemlogs` 下，同时数据库备份文件也在该目录下
5. 配置文件在 `config` 目录下，配置信息在 `config/config.yaml` 中
6. [redis密码设置教程](https://www.runoob.com/redis/redis-security.html)

### 目录分级

| 目录        | 作用                               |
| ----------- | ---------------------------------- |
| configs     | 配置文件、SQL 结构文件              |
| consts      | 常量规范 如登录成功 / 失败等提示信息 |
| model       | 常用结构体                         |
| service     | 服务层，实现功能                   |
| static      | 前端静态文件                       |
| templates   | HTML 界面                           |
| uploadfiles | 用户上传文件及头像（`gitignore`）    |
| utils       | 工具包，读取配置文件、连接数据库等 |
| systemlogs | 日志文件及数据库备份 |

### 后端接口

| 接口名              | 说明                               | 方法 |
| ------------------- | ---------------------------------- | ---- |
| CheckAuthCode       | 检测验证码                         | GET  |
| CheckEmailAvailable | 检测邮箱可用（格式错误、已被注册） | GET  |
| ForgetPassword      | 通过邮箱验证修改密码               | POST |
| GetUserFilesPath    | 获取该用户上传的文件的路径         | GET  |
| GetUserMessage      | 获取用户信息                       | GET  |
| ModifyPassword      | 修改密码                           | POST |
| RestoreMySQL        | 恢复数据库                         | GET  |
| UploadAvatar        | 上传头像                           | POST |
| UploadFile          | 上传文件                           | POST |
| UserLogin           | 用户登录                           | GET  |
| UserRegister        | 用户注册                           | GET  |
| SendAuthCode        | 发送邮箱验证码                     | GET  |

#### CheckAuthCode

GET：检测邮箱验证码是否正确

| 字段名    | 类型   | 说明       | 必选 |
| --------- | ------ | ---------- | ---- |
| email     | string | 邮箱       | √    |
| auth_code | int    | 对应验证码 | √    |

#### CheckEmailAvailable

GET：检测邮箱可用（格式错误？已被注册？）

| 字段名 | 类型   | 说明 | 必选 |
| ------ | ------ | ---- | ---- |
| email  | string | 邮箱 | √    |

#### ForgetPassword

POST：通过邮箱验证修改密码

| 字段名       | 类型   | 说明       | 必选 |
| ------------ | ------ | ---------- | ---- |
| email        | string | 邮箱       | √    |
| auth_code    | int    | 对应验证码 | √    |
| new_password | string | 新密码     | √    |

#### GetUserFilesPath

GET：获取该用户上传的所有文件的路径

#### GetUserMessage

GET：获取用户信息 （无参数 通过 `cookie` 获取用户信息）

#### ModifyPassword

POST：修改密码

| 字段名      | 类型   | 说明   | 必选 |
| ----------- | ------ | ------ | ---- |
| oldPassword | string | 旧密码 | √    |
| newPassword | string | 新密码 | √    |

#### RestoreMySQL

GET：恢复最近保存的数据库信息

#### UploadAvatar

POST：上传头像图片（前后端都应该需要进行简单检测）

| 字段名 | 类型 | 说明     | 必选 |
| ------ | ---- | -------- | ---- |
| avatar | file | 图片文件 | √    |

#### UploadFile

POST：上传文件

| 字段名 | 类型 | 说明     | 必选 |
| ------ | ---- | -------- | ---- |
| file   | file | 各种文件 | √    |

#### UserLogin 

GET：用户登录 

| 字段名   | 类型   | 说明   | 必选 |
| -------- | ------ | ------ | ---- |
| account  | string | 账户 ID | √    |
| password | string | 密码   | √    |

#### UserRegister

GET：用户注册

| 字段名   | 类型   | 说明   | 必选 |
| -------- | ------ | ------ | ---- |
| account  | string | 账户 ID | √    |
| password | string | 密码   | √    |
| email    | string | 邮箱   | √    |
| auth_code | int | 邮箱验证码 | √ |
| nick_name | string | 昵称   | √    |

#### SendAuthCode

GET：发送邮箱验证码 

| 字段名 | 类型   | 说明 | 必选 |
| ------ | ------ | ---- | ---- |
| email  | string | 邮箱 | √    |

todo:
1. 修改密码直接获取用户cookie然后发送到对应邮箱
2. 添加predict接口
3. 部署模型