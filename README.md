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
| consts      | 常量规范 如登录成功 / 失败 等提示信息 |
| model       | 常用结构体                         |
| service     | 服务层，实现功能                   |
| static      | 前端静态文件                       |
| templates   | HTML 界面                           |
| user_files | 用户上传png  |
| utils       | 工具包，读取配置文件、连接数据库等 |
| systemlogs | 日志文件及数据库备份 |
| avatars | 头像文件 |

### 后端接口

| 接口名              | 说明                               | 方法 |
| ------------------- | ---------------------------------- | ---- |
| CheckEmailAvailable | 检测邮箱可用（格式错误、已被注册） | GET  |
| GetUserFilesPath    | 获取该用户上传的文件的路径         | GET  |
| GetUserMessage      | 获取用户信息                       | GET  |
| PredictPng          | 对png文件进行预测                  | GET  |
| RestoreMySQL        | 恢复数据库                         | GET  |
| UserLogin           | 用户登录                           | GET  |
| UserRegister        | 用户注册                           | GET  |
| SendAuthCode        | 发送邮箱验证码                     | GET  |
| ModifyPassword      | 用户登录后根据旧密码修改密码       | POST |
| ForgetPassword      | 通过邮箱验证修改密码               | POST |
| UploadAvatar        | 上传头像                           | POST |
| UploadPng           | 上传Png文件                        | POST |
| DeletePng           | 删除Png文件                        | POST |

返回结构体

| json结构体 | 类型      | 说明           |
| ---------- | --------- | -------------- |
| Code       | int       | 错误码         |
| Msg        | string    | 错误码对应信息 |
| Data       | interface | 部分接口使用   |

| Code错误码 | Msg错误码对应信息                                            |
| ---------- | ------------------------------------------------------------ |
| 200        | 本次操作成功                                                 |
| 100000     | 输入参数有误，或许是邮箱格式错误，或许是有空输入             |
| 100001     | 未找到该用户的信息，或许是Cookie过期，或许是在数据库中被人为删除（很少出现该错误码） |
| 100002     | 账号或密码错误（可能是账号不存在）                           |
| 100003     | 邮箱验证码校验失败                                           |
| 100004     | 邮箱已被注册                                                 |
| 100005     | 更新密码失败，或许是数据库错误，也可能是新旧密码相同         |
| 100006     | Cookie过期                                                   |
| 100007     | 发送邮箱验证码失败                                           |
| 999998     | 系统错误（各种原因，文件夹创建失败，文件保存失败，框架错误，基本都是不可控行为） |
| 999999     | 数据库错误（记录创建失败，redis数据获取失败，后续可能追加rabbitmq失败） |

#### CheckEmailAvailable

GET：检测邮箱可用（格式错误？已被注册？）

| 字段名 | 类型   | 说明 | 必选 |
| ------ | ------ | ---- | ---- |
| email  | string | 邮箱 | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型 |
| ---------- | --------- | ------------------ |
| Code       | int       | 200 100000 100004  |
| Msg        | string    | 对应错误码信息     |
| Data       | interface | 此处无用           |



#### GetUserFilesPath

GET：获取该用户上传的所有文件的路径

返回信息

| json结构体 | 类型      | 可能返回的信息类型          |
| ---------- | --------- | --------------------------- |
| Code       | int       | 200 100006 999998           |
| Msg        | string    | 对应错误码信息              |
| Data       | interface | Code为200时对应文件路径数组 |



#### GetUserMessage

GET：获取用户信息 （无参数 通过 `cookie` 获取用户信息）

返回信息

| json结构体 | 类型      | 可能返回的信息类型        |
| ---------- | --------- | ------------------------- |
| Code       | int       | 200 100001 100006 999998  |
| Msg        | string    | 对应错误码信息            |
| Data       | interface | Code为200时用户信息结构体 |



#### PredictPng

GET：根据输入图片名预测结果

| 字段名           | 类型   | 说明                              | 必选 |
| ---------------- | ------ | --------------------------------- | ---- |
| predict_png_name | string | 预测的png文件名（只要文件名.png） | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型                                           |
| ---------- | --------- | ------------------------------------------------------------ |
| Code       | int       | 200 100000 100006 999998                                     |
| Msg        | string    | 对应错误码信息                                               |
| Data       | interface | Code为200时为预测结果png路径，预测结果png文件会保存在用户文件下 |



#### RestoreMySQL

GET：恢复最近保存的数据库信息（可用，但暂时不考虑）



#### UserLogin 

GET：用户登录 

| 字段名   | 类型   | 说明    | 必选 |
| -------- | ------ | ------- | ---- |
| account  | string | 账户 ID | √    |
| password | string | 密码    | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型       |
| ---------- | --------- | ------------------------ |
| Code       | int       | 200 100000 100006 999999 |
| Msg        | string    | 对应错误码信息           |
| Data       | interface | 此处无用                 |



#### UserRegister

GET：用户注册

| 字段名    | 类型   | 说明       | 必选 |
| --------- | ------ | ---------- | ---- |
| account   | string | 账户 ID    | √    |
| password  | string | 密码       | √    |
| email     | string | 邮箱       | √    |
| auth_code | int    | 邮箱验证码 | √    |
| nick_name | string | 昵称       | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型       |
| ---------- | --------- | ------------------------ |
| Code       | int       | 200 100000 100003 999999 |
| Msg        | string    | 对应错误码信息           |
| Data       | interface | 此处无用                 |



#### SendAuthCode

GET：发送邮箱验证码 

| 字段名 | 类型   | 说明 | 必选 |
| ------ | ------ | ---- | ---- |
| email  | string | 邮箱 | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型 |
| ---------- | --------- | ------------------ |
| Code       | int       | 200 100000 100007  |
| Msg        | string    | 对应错误码信息     |
| Data       | interface | 此处无用           |



#### ModifyPassword

POST：修改密码

| 字段名      | 类型   | 说明   | 必选 |
| ----------- | ------ | ------ | ---- |
| oldPassword | string | 旧密码 | √    |
| newPassword | string | 新密码 | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型       |
| ---------- | --------- | ------------------------ |
| Code       | int       | 200 100000 100005 100006 |
| Msg        | string    | 对应错误码信息           |
| Data       | interface | 此处无用                 |





#### ForgetPassword

POST：通过邮箱验证修改密码

| 字段名       | 类型   | 说明       | 必选 |
| ------------ | ------ | ---------- | ---- |
| email        | string | 邮箱       | √    |
| auth_code    | int    | 对应验证码 | √    |
| new_password | string | 新密码     | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型       |
| ---------- | --------- | ------------------------ |
| Code       | int       | 200 100000 100003 100005 |
| Msg        | string    | 对应错误码信息           |
| Data       | interface | 此处无用                 |



#### UploadAvatar

POST：上传头像图片（前后端都应该需要进行简单检测）

| 字段名 | 类型 | 说明     | 必选 |
| ------ | ---- | -------- | ---- |
| avatar | file | 图片文件 | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型              |
| ---------- | --------- | ------------------------------- |
| Code       | int       | 200 100000 100006 999998 999999 |
| Msg        | string    | 对应错误码信息                  |
| Data       | interface | 此处无用                        |



#### UploadPng

POST：上传png图像文件

| 字段名   | 类型 | 说明    | 必选 |
| -------- | ---- | ------- | ---- |
| png_name | file | png文件 | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型 |
| ---------- | --------- | ------------------ |
| Code       | int       | 200 100003 999998  |
| Msg        | string    | 对应错误码信息     |
| Data       | interface | 此处无用           |



#### DeletePng

POST：上传png图像文件

| 字段名          | 类型   | 说明            | 必选 |
| --------------- | ------ | --------------- | ---- |
| delete_png_name | string | 删除的png文件名 | √    |

返回信息

| json结构体 | 类型      | 可能返回的信息类型 |
| ---------- | --------- | ------------------ |
| Code       | int       | 200 100000 999998  |
| Msg        | string    | 对应错误码信息     |
| Data       | interface | 此处无用           |





todo:

1. 修改密码直接获取用户cookie然后发送到对应邮箱
2. 添加predict接口
3. 部署模型