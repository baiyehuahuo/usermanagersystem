# UserManagerSystem

1. 前端暂用 https://gitee.com/xminghua/Login.git ，待小组成员补充
2. 后端用 `go` 开发，`gin` 为服务器， `gorm` 查询 `mysql` 为数据库
3. 运行用 `go run .`

### 目录分级

| 目录        | 作用                               |
| ----------- | ---------------------------------- |
| configs     | 配置文件、SQL结构文件              |
| consts      | 常量规范 如登录成功/失败等提示信息 |
| model       | 常用结构体                         |
| service     | 服务层，实现功能                   |
| static      | 前端静态文件                       |
| templates   | html界面                           |
| uploadfiles | 用户上传文件及头像(`gitignore`)    |
| utils       | 工具包，读取配置文件、连接数据库等 |

### 后端接口

| 接口名         | 说明         | 方法 |
| -------------- | ------------ | ---- |
| GetUserMessage | 获取用户信息 | GET  |
| ModifyPassword | 修改密码     | POST |
| UploadAvatar   | 上传头像     | POST |
| UploadFile     | 上传文件     | POST |
| UserLogin      | 用户登录     | GET  |
| UserRegedit    | 用户注册     | GET  |



#### GetUserMessage

GET：获取用户信息 （无参数 通过 `cookie` 获取用户信息）



#### ModifyPassword

POST：修改密码

| 字段名      | 类型   | 说明   | 必选 |
| ----------- | ------ | ------ | ---- |
| oldPassword | string | 旧密码 | √    |
| newPassword | string | 新密码 | √    |



#### UploadAvatar

POST：上传头像图片（前后端都应该需要进行简单检测）

| 字段名 | 类型 | 说明     | 必选 |
| ------ | ---- | -------- | ---- |
| avatar | file | 图片文件 | √    |



#### UploadFile

POST：上传文件（测试文件上传功能 后期删掉）

| 字段名 | 类型 | 说明     | 必选 |
| ------ | ---- | -------- | ---- |
| file   | file | 各种文件 | √    |



#### UserLogin 

GET：用户登录 

| 字段名   | 类型   | 说明   | 必选 |
| -------- | ------ | ------ | ---- |
| account  | string | 账户id | √    |
| password | string | 密码   | √    |



#### UserRegedit

GET：用户注册（需要界面，然后改成POST，前后端都需要检测邮箱是否合法 后端还需要检测邮箱是否已被注册）

| 字段名   | 类型   | 说明   | 必选 |
| -------- | ------ | ------ | ---- |
| account  | string | 账户id | √    |
| password | string | 密码   | √    |
| email    | string | 邮箱   | √    |
| nick_name | string | 昵称   | √    |



#### 
