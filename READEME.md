

[TOC]

## Introduction

- Movie management system (similar to [IMDB](https://www.imdb.com/)) developed  in the Go language for practice.
- Administrators can create movie items through the system
- Users can mark movies they want to see, watch, and rate



## Technology used

- Golang，Gin
- MySQL，Gorm
- Swagger
- validator，JWT



## API Description（Eng）

### admin

| URI                     | Description                  |
| ----------------------- | ---------------------------- |
| /admin/user/addRole     | Associate a role to the user |
| /admin/user/delete/{id} | delete user                  |
| /admin/user/lock/{id}   | Account lock                 |
| /admin/user/unlock/{id} | Unlock account               |



### message

| URI                    | Description                        |
| ---------------------- | ---------------------------------- |
| /message/type          | Get all message types              |
| /message/read/{msg_id} | Read message                       |
| /message/system        | Create system message              |
| /message/list          | Get all messages of a certain type |



### movie

| URI                                | Description                                           |
| ---------------------------------- | ----------------------------------------------------- |
| /category/all                      | Query all categories                                  |
| /category/movie                    | Paging to query all movies under the current category |
| /filmmaker/info/{id}               | Get information about the specified actor             |
| /language/all                      | Get all languages                                     |
| /language/movie                    | Paging to query all movies in the current language    |
| /movie/feed                        | Feed page recommendation interface                    |
| /movie/add                         | Insert movie entity                                   |
| /movieinfo/{id}                    | Query movies by id                                    |
| /movieinfo/douban/{db_id}          | Query movies by Douban id                             |
| /movie/delete/{id}                 | Delete the specified movie information                |
| /movie/deleteAll                   | Delete all movie information                          |
| /movie/update                      | Update movie information                              |
| /movie/feedMock                    | Only query with cover, mock data                      |
| movie/user/wanna/{movie_id}        | Want to watch a movie                                 |
| movie/user/wanna/cancel/{movie_id} | Cancel want to watch a movie                          |
| movie/user/seen/{movie_id}         | watched a movie                                       |
| movie/user/rating                  | Rate a movie                                          |
| movie/user/wanna/all               | Want to watch list                                    |
| movie/user/seen/all                | watched list                                          |
| /region/all                        | Get all regions                                       |
| /region/movie                      | Paging to query all movies by region                  |



### search

| URI           | Description             |
| ------------- | ----------------------- |
| /search/all   | Global Aggregate Search |
| /search/movie | Movie Search            |
| /search/post  | Post Search             |
| /search/user  | User search             |



### social

| URI                      | Description                                               |
| ------------------------ | --------------------------------------------------------- |
| /bbs/comment             | Query all categories                                      |
| /bbs/comment             | Pagination to query all movies under the current category |
| /bbs/comment             | Get information about the specified actor                 |
| /follow/{user_id}        | Follow users                                              |
| /follow/cancel/{user_id} | Unfollow                                                  |
| /follow/following        | Get the list of people the current user follows           |
| /follow/followed         | Get the list of people followed by the current user       |
| /bbs/like                | Liked successfully                                        |
| /bbs/unlike              | Unlike success                                            |
| /bbs/post/create         | Post a post                                               |
| /bbs/post/update         | New post content                                          |
| /bbs/post/delete/{id}    | delete post                                               |
| /bbs/post/topic/list     | Get all posts under a topic                               |
| /bbs/post/movie/list     | Get all posts under a movie                               |
| /bbs/reply/create        | Add reply                                                 |
| /bbs/reply/list          | Get a list of replies under a comment                     |
| /bbs/reply/delete/{id}   | Delete a comment                                          |
| /bbs/topic/create        | Topic created successfully                                |
| /bbs/topic/update        | Update topic                                              |
| /bbs/topic/delete/{id}   | Delete topic                                              |
| /bbs/topic/commons       | Get all general topics, sorted by number of posts         |



## API Description(Chinese)

### admin

| URI                     | Description        |
| ----------------------- | ------------------ |
| /admin/user/addRole     | 给用户关联一个角色 |
| /admin/user/delete/{id} | 删除用户           |
| /admin/user/lock/{id}   | 账号锁定           |
| /admin/user/unlock/{id} | 解除账号锁定       |



### message

| URI                    | Description            |
| ---------------------- | ---------------------- |
| /message/type          | 获取全部消息类型       |
| /message/read/{msg_id} | 读取消息               |
| /message/system        | 创建系统消息           |
| /message/list          | 获取某个类型的全部消息 |



### movie

| URI                                | Description                |
| ---------------------------------- | -------------------------- |
| /category/all                      | 查询全部分类               |
| /category/movie                    | 分页查询当前分类下全部电影 |
| /filmmaker/info/{id}               | 获取指定演员的信息         |
| /language/all                      | 获取全部语言               |
| /language/movie                    | 分页查询当前语言下全部电影 |
| /movie/feed                        | feed页面推荐接口           |
| /movie/add                         | 插入电影实体               |
| /movieinfo/{id}                    | 根据id查询电影             |
| /movieinfo/douban/{db_id}          | 根据豆瓣id查询电影         |
| /movie/delete/{id}                 | 删除指定电影信息           |
| /movie/deleteAll                   | 删除全部电影信息           |
| /movie/update                      | 更新电影信息               |
| /movie/feedMock                    | 只查询有封面的，mock数据   |
| movie/user/wanna/{movie_id}        | 想看某个电影               |
| movie/user/wanna/cancel/{movie_id} | 取消想看某个电影           |
| movie/user/seen/{movie_id}         | 看过某个电影               |
| movie/user/rating                  | 给电影打分                 |
| movie/user/wanna/all               | 想看的列表                 |
| movie/user/seen/all                | 看过的列表                 |
| /region/all                        | 获取全部地区               |
| /region/movie                      | 分页查询地区分类下全部电影 |



### search

| URI           | Description  |
| ------------- | ------------ |
| /search/all   | 全局聚合搜索 |
| /search/movie | 电影搜索     |
| /search/post  | 帖子搜索     |
| /search/user  | 用户搜索     |



### social

| URI                      | Description                        |
| ------------------------ | ---------------------------------- |
| /bbs/comment             | 查询全部分类                       |
| /bbs/comment             | 分页查询当前分类下全部电影         |
| /bbs/comment             | 获取指定演员的信息                 |
| /follow/{user_id}        | 关注用户                           |
| /follow/cancel/{user_id} | 取消关注                           |
| /follow/following        | 获取当前用户关注人的列表           |
| /follow/followed         | 获取当前用户关注人的列表           |
| /bbs/like                | 点赞成功                           |
| /bbs/unlike              | 取消点赞成功                       |
| /bbs/post/create         | 发表帖子                           |
| /bbs/post/update         | 新帖子内容                         |
| /bbs/post/delete/{id}    | 删除帖子                           |
| /bbs/post/topic/list     | 获取一个主题下的全部帖子           |
| /bbs/post/movie/list     | 获取一个电影下的全部帖子           |
| /bbs/reply/create        | 添加回复                           |
| /bbs/reply/list          | 获取评论下的回复列表               |
| /bbs/reply/delete/{id}   | 删除某条评论                       |
| /bbs/topic/create        | 话题创建成功                       |
| /bbs/topic/update        | 更新话题                           |
| /bbs/topic/delete/{id}   | 删除主题                           |
| /bbs/topic/commons       | 获取全部常规话题, 按照帖子数量排序 |





```go
.
├── READEME.md
├── scheme.sql
├── configs
│   └── config.yaml
├── docs
│   ├── docs.go
│   ├── douban250.txt
│   ├── sql
│   │   ├── apply.sql
│   │   ├── scheme.sql
│   │   └── init_schema.sql
│   ├── swagger.json
│   └── swagger.yaml
├── global
│   ├── db.go
│   ├── setting.go
│   └── validator.go
├── go.mod
├── go.sum
├── internal
│   ├── dao
│   │   ├── AdminDao.go
│   │   ├── AppDao.go
│   │   ├── AuthDao.go
│   │   ├── MovieDao.go
│   │   ├── SearchDao.go
│   │   ├── UserDao.go
│   │   └── dao.go
│   ├── middleware
│   │   ├── access_log.go
│   │   ├── app_info.go
│   │   ├── context_timeout.go
│   │   ├── cors.go
│   │   ├── jwt.go
│   │   ├── limiter.go
│   │   ├── recovery.go
│   │   └── translations.go
│   ├── model
│   │   ├── AdminModel.go
│   │   ├── AppModel.go
│   │   ├── AuthModel.go
│   │   ├── MessageModel.go
│   │   ├── MovieModel.go
│   │   ├── SearchModel.go
│   │   ├── SocialModel.go
│   │   ├── UserModel.go
│   │   └── model.go
│   ├── routers
│   │   ├── api
│   │   └── router.go
│   └── service
│       ├── AdminService.go
│       ├── AppService.go
│       ├── AuthService.go
│       ├── MovieService.go
│       ├── SearchService.go
│       ├── UserService.go
│       └── service.go
├── main
├── main.go
├── pkg
│   ├── app
│   │   ├── app.go
│   │   ├── form.go
│   │   ├── jwt.go
│   │   └── pagination.go
│   ├── convert
│   │   └── convert.go
│   ├── email
│   │   └── email.go
│   ├── errcode
│   │   ├── common_code.go
│   │   ├── errcode.go
│   │   └── module_code.go
│   ├── limiter
│   │   ├── limiter.go
│   │   └── method_limiter.go
│   ├── logger
│   │   └── logger.go
│   ├── setting
│   │   ├── section.go
│   │   └── setting.go
│   ├── upload
│   │   └── file.go
│   ├── util
│   │   └── md5.go
│   └── validator
│       └── custom_validator.go
└── storage
    └── uploads

24 directories, 66 files
```
