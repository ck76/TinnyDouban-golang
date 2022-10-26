-- CREATE SCHEMA `tittytainment` ;

#
# -- ----------------------------
# -- Table structure for user
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_user`;
# CREATE TABLE `ttmt_user`
    # (
    #     `id`           int(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `phone`        varchar(100)     NOT NULL COMMENT '手机号',
    #     `password`     varchar(100)     NOT NULL COMMENT '加密后的密码',
    #     `name`         varchar(100)              DEFAULT '' COMMENT '用户名',
    #     `cover_url`    varchar(200)              DEFAULT '' COMMENT '头像图片地址',
    #     `status`       tinyint(1)       NOT NULL DEFAULT 1 COMMENT '1有效账号 0无效账号',
    #     `created_time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';
#
#
# -- ----------------------------
# -- Table structure for role
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_role`;
# CREATE TABLE `ttmt_role`
    # (
    #     `id`           int(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `name`         varchar(100)              DEFAULT '' COMMENT '角色名',
    #     `description`  varchar(255)              DEFAULT '' COMMENT '描述',
    #     `created_time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time` timestamp        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='角色表';
#
# -- ----------------------------
# -- Records of my_role
# -- ----------------------------
# INSERT INTO `ttmt_role` (id, name, description)
# VALUES (1, 'ADMIN', '超级管理员，拥有所有权限');
# INSERT INTO `ttmt_role` (id, name, description)
# VALUES (2, 'USER', '普通用户');
#
#
# -- ----------------------------
# -- Table structure for role and user
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_user_role`;
# CREATE TABLE `ttmt_user_role`
    # (
    #     `user_id` int(11) NOT NULL,
    #     `role_id` int(11) NOT NULL,
    #     PRIMARY KEY (`user_id`, `role_id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='用户-角色关系表';

#
# -- ----------------------------
# -- Table structure for movie
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_movie`;
# CREATE TABLE `ttmt_movie`
    # (
    #     `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `douban_id`    bigint(20) unsigned NOT NULL default 0 COMMENT '豆瓣id',
    #     `name`         varchar(100)        NOT NULL DEFAULT '' COMMENT '电影名称',
    #     `alias`        varchar(100)                 DEFAULT '' COMMENT '电影别名',
    #     `cover_url`    varchar(200)                 DEFAULT '' COMMENT '电影封面地址',
    #     `douban_score` double unsigned              DEFAULT 0 COMMENT '豆瓣评分',
    #     `douban_vote`  int(10) unsigned             DEFAULT 0 COMMENT '豆瓣打分人数',
    #     `mins`         int(5) unsigned              DEFAULT 0 COMMENT '电影时长(min)',
    #     `release_date` datetime                     DEFAULT CURRENT_TIMESTAMP COMMENT '电影上映时间',
    #     `year`         int(5) unsigned              DEFAULT 0 COMMENT '电影所属年份(比上映时间更全)',
    #     `storyline`    text(1000) COMMENT '电影情节',
    #     `tags`         text(1000) COMMENT '标签',
    #     `created_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='电影表';
#
# ############################ movie - category ###########################
#
# -- ----------------------------
# -- Table structure for category
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_category`;
# CREATE TABLE `ttmt_category`
    # (
    #     `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `category_name` varchar(100)        NOT NULL DEFAULT '' COMMENT '分类名称',
    #     `created_time`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='分类表';
#
# -- ----------------------------
# -- Table structure for movie-category
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_movie_category`;
# CREATE TABLE `ttmt_movie_category`
    # (
    #     `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `movie_id`        bigint(20) unsigned NOT NULL default 0 COMMENT '电影id',
    #     `movie_douban_id` bigint(20) unsigned NOT NULL default 0 COMMENT '电影豆瓣id',
    #     `category_id`     bigint(20) unsigned NOT NULL default 0 COMMENT '分类id',
    #     `created_time`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='电影-分类表';
#
#
# ############################ movie - language ###########################
#
# -- ----------------------------
# -- Table structure for language
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_language`;
# CREATE TABLE `ttmt_language`
    # (
    #     `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `language_name` varchar(100)        NOT NULL DEFAULT '' COMMENT '语言名称',
    #     `created_time`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='语言表';
#
# -- ----------------------------
# -- Table structure for movie-language
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_movie_language`;
# CREATE TABLE `ttmt_movie_language`
    # (
    #     `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `movie_id`        bigint(20) unsigned NOT NULL default 0 COMMENT '电影id',
    #     `movie_douban_id` bigint(20) unsigned NOT NULL default 0 COMMENT '电影豆瓣id',
    #     `language_id`     bigint(20) unsigned NOT NULL default 0 COMMENT '语言id',
    #     `created_time`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='电影-分类表';
#
# ############################ movie - region ###########################
#
# -- ----------------------------
# -- Table structure for region
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_region`;
# CREATE TABLE `ttmt_region`
    # (
    #     `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `region_name`  varchar(100)        NOT NULL DEFAULT '' COMMENT '地区名称',
    #     `created_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='地区表';
#
# -- ----------------------------
# -- Table structure for movie-region
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_movie_region`;
# CREATE TABLE `ttmt_movie_region`
    # (
    #     `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `movie_id`        bigint(20) unsigned NOT NULL default 0 COMMENT '电影id',
    #     `movie_douban_id` bigint(20) unsigned NOT NULL default 0 COMMENT '电影豆瓣id',
    #     `region_id`       bigint(20) unsigned NOT NULL default 0 COMMENT '地区id',
    #     `created_time`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='电影-地区表';
#
#
#
# -- ----------------------------
# -- Table structure for filmmaker
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_filmmaker`;
# CREATE TABLE `ttmt_filmmaker`
    # (
    #     `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `douban_id`     bigint(20) unsigned NOT NULL default 0 COMMENT '豆瓣演员id',
    #     `name`          varchar(100)        NOT NULL DEFAULT '' COMMENT '演员名称',
    #     `name_en`       varchar(1000)                DEFAULT '' COMMENT '英文全名',
    #     `name_zh`       varchar(1000)                DEFAULT '' COMMENT '中文全名',
    #     `sex`           varchar(10)         NOT NULL DEFAULT '' COMMENT '性别',
    #     `birth`         datetime                     DEFAULT CURRENT_TIMESTAMP COMMENT '出生日期',
    #     `birth_place`   varchar(100)                 DEFAULT '' COMMENT '出生地点',
    #     `constellation` varchar(10)                  DEFAULT '' COMMENT '星座',
    #     `biography`     text(1000) COMMENT '人物简介',
    #     `created_time`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='电影演员与导演';
#
# -- ----------------------------
# -- Table structure for profession
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_profession`;
# CREATE TABLE `ttmt_profession`
    # (
    #     `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `name`         varchar(100)        NOT NULL DEFAULT '' COMMENT '职业名称',
    #     `created_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='职业表';
#
# -- ----------------------------
# -- Table structure for profession-movie
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_filmmaker_profession`;
# CREATE TABLE `ttmt_filmmaker_profession`
    # (
    #     `id`                  bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `filmmaker_id`        bigint(20) unsigned NOT NULL default 0 COMMENT '演员id',
    #     `filmmaker_douban_id` bigint(20) unsigned NOT NULL default 0 COMMENT '演员豆瓣id',
    #     `profession_id`       bigint(20) unsigned NOT NULL default 0 COMMENT '职业id',
    #     `created_time`        timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`        timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='职业-演员表';
#
# -- ----------------------------
# -- Table structure for filmmaker-movie
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_movie_filmmaker`;
# CREATE TABLE `ttmt_movie_filmmaker`
    # (
    #     `id`                  bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `movie_id`            bigint(20) unsigned NOT NULL default 0 COMMENT '电影id',
    #     `movie_douban_id`     bigint(20) unsigned NOT NULL default 0 COMMENT '电影豆瓣id',
    #     `filmmaker_id`        bigint(20) unsigned NOT NULL default 0 COMMENT '演员id',
    #     `filmmaker_douban_id` bigint(20) unsigned NOT NULL default 0 COMMENT '演员豆瓣id',
    #     `profession_id`       bigint(20) unsigned NOT NULL default 0 COMMENT '职业id',
    #     `created_time`        timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`        timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='电影-演员关系表';


# -- ----------------------------
# -- Table structure for filmmaker-movie
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_movie_user`;
# CREATE TABLE `ttmt_movie_user`
    # (
    #     `id`              bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `user_id`         bigint(20) unsigned NOT NULL default 0 COMMENT '用户id',
    #     `movie_id`        bigint(20) unsigned NOT NULL default 0 COMMENT '电影id',
    #     `movie_douban_id` bigint(20) unsigned NOT NULL default 0 COMMENT '电影豆瓣id',
    #     `wanna_see`       tinyint(1) unsigned NOT NULL default 0 COMMENT '是否想看该电影',
    #     `have_seen`       tinyint(1) unsigned NOT NULL default 0 COMMENT '是否看过该电影',
    #     `score`           double unsigned              default 0 COMMENT '豆瓣评分',
    #     `wanna_see_time`  datetime                     DEFAULT CURRENT_TIMESTAMP COMMENT '想看电影修改时间(想看/取消)',
    #     `have_seen_time`  datetime                     DEFAULT CURRENT_TIMESTAMP COMMENT '看过电影时间',
    #     `score_time`      datetime                     DEFAULT CURRENT_TIMESTAMP COMMENT '给电影打分时间',
    #     `created_time`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`    timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='电影-用户关系表';
#
#
#
# -- ----------------------------
# -- Table structure for user-follow
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_user_follow`;
# CREATE TABLE `ttmt_user_follow`
    # (
    #     `id`             bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `user_id`        bigint(20) unsigned NOT NULL default 0 COMMENT '用户id',
    #     `follow_user_id` bigint(20) unsigned NOT NULL default 0 COMMENT '被关注用户id',
    #     `follow`         tinyint(1) unsigned NOT NULL default 0 COMMENT '是否处于关注状态',
    #     `created_time`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='用户关注表';


#############################################
##
##
##      论坛系统相关表
##
##
#############################################

# -- ----------------------------
# -- Table structure for topic
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_topic`;
# CREATE TABLE `ttmt_topic`
    # (
    #     `id`             bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `name`           varchar(100)        NOT NULL DEFAULT '' COMMENT '主题名称',
    #     `topic_type`     int unsigned        NOT NULL default 0 COMMENT '0为常规主题, 1为电影的默认主题',
    #     `movie_id`       bigint(20) unsigned          default 0 COMMENT '如果是电影的默认主题, 需要对应电影id',
    #     `create_user_id` bigint(20) unsigned NOT NULL default 0 COMMENT '若为常规主题, 则需要创建人id',
    #     `post_num`       bigint(20) unsigned NOT NULL default 0 COMMENT '帖子数量',
    #     `topic_abstract` longtext COMMENT '主题内容介绍',
    #     `user_num`       bigint(20) unsigned NOT NULL default 0 COMMENT '关注用户数量',
    #     `created_time`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`   timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='主题表';
#
#
# -- ----------------------------
# -- Table structure for post
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_post`;
# CREATE TABLE `ttmt_post`
    # (
    #     `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `title`             varchar(100)        NOT NULL DEFAULT '' COMMENT '帖子标题',
    #     `topic_id`          bigint(20) unsigned NOT NULL default 0 COMMENT '对应的主题id',
    #     `content`           longtext COMMENT '帖子内容(一楼内容), 里面可能包含图片, 后台不进行处理',
    #     `post_type`         int unsigned        NOT NULL default 0 COMMENT '0为常规帖子, 1 为电影默认帖子',
    #     `movie_id`          bigint(20) unsigned NOT NULL default 0 COMMENT '如果是电影的默认主题, 需要对应电影id',
    #     `create_user_id`    bigint(20) unsigned NOT NULL default 0 COMMENT '如果不是电影默认主题, 则需要发帖人id',
    #     `like_num`          bigint(20) unsigned NOT NULL default 0 COMMENT '点赞数量',
    #     `comment_num`       bigint(20) unsigned NOT NULL default 0 COMMENT '评论数量',
    #     `last_comment_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最新评论时间',
    #     `created_time`      timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`      timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='帖子表';
#
# -- ----------------------------
# -- Table structure for comment
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_comment`;
# CREATE TABLE `ttmt_comment`
    # (
    #     `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `post_id`      bigint(20) unsigned NOT NULL default 0 COMMENT '对应的主题id',
    #     `content`      longtext COMMENT '评论内容',
    #     `user_id`      bigint(20) unsigned NOT NULL default 0 COMMENT '评论者id',
    #     `like_num`     bigint(20) unsigned NOT NULL default 0 COMMENT '点赞数量',
    #     `reply_num`    bigint(20) unsigned NOT NULL default 0 COMMENT '回复数量',
    #     `status`       tinyint(1) unsigned NOT NULL default 0 COMMENT '0为删除状态, 1为生效状态',
    #     `created_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='评论表';
#
#
# -- ----------------------------
# -- Table structure for reply
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_reply`;
# CREATE TABLE `ttmt_reply`
    # (
    #     `id`            bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `reply_type`    int unsigned        NOT NULL default 0 COMMENT '0回复评论, 1为对回复回复',
    #     `comment_id`    bigint(20) unsigned NOT NULL default 0 COMMENT '回复的评论id',
    #     `reply_id`      bigint(20) unsigned          default 0 COMMENT '回复的回复id',
    #     `content`       longtext COMMENT '回复内容',
    #     `user_id`       bigint(20) unsigned NOT NULL default 0 COMMENT '回复者id',
    #     `reply_user_id` bigint(20) unsigned NOT NULL default 0 COMMENT '回复的人id',
    #     `status`        tinyint(1) unsigned NOT NULL default 0 COMMENT '0为删除状态, 1为生效状态',
    #     `created_time`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='回复表';
#
#
# -- ----------------------------
# -- Table structure for like
# -- ----------------------------
# DROP TABLE IF EXISTS `ttmt_like`;
# CREATE TABLE `ttmt_like`
    # (
    #     `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `type`         int unsigned        NOT NULL default 0 COMMENT '0为帖子点赞, 1为评论点赞',
    #     `post_id`      bigint(20) unsigned NOT NULL default 0 COMMENT '点赞的帖子id',
    #     `comment_id`   bigint(20) unsigned NOT NULL default 0 COMMENT '点赞的评论id',
    #     `user_id`      bigint(20) unsigned NOT NULL default 0 COMMENT '点赞者id',
    #     `status`       tinyint(1) unsigned NOT NULL default 0 COMMENT '0取消点赞, 1为点赞生效',
    #     `created_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='点赞表';


# #############################################
# ##
# ##      站内信相关表
# ##
# #############################################
#
# DROP TABLE IF EXISTS `ttmt_message`;
# CREATE TABLE `ttmt_message`
    # (
    #     `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    #     `type`         int unsigned        NOT NULL default 0 COMMENT '0为系统类型信息, 1为点赞类型信息，2为评论类型信息，3为回复类型信息，4为关注类型信息',
    #     `type_id`      bigint(20) unsigned NOT NULL default 0 COMMENT '不同种信息对应的事件id',
    #     `send_user_id` bigint(20) unsigned NOT NULL default 0 COMMENT '需要推送给的用户',
    #     `from_user_id` bigint(20) unsigned NOT NULL default 0 COMMENT '产生该事件的用户',
    #     `has_read`     tinyint(1) unsigned NOT NULL default 0 COMMENT '0为未读, 1为已读',
    #     `message`      text(1000) COMMENT '消息内容',
    #     `created_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    #     `updated_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    #     PRIMARY KEY (`id`)
    # ) ENGINE = InnoDB
    #   DEFAULT CHARSET = utf8mb4 COMMENT ='点赞表';


#############################################
##
##      App 版本相关信息  主版本号  次版本号  灰度版本号  下载地址 更新信息
##
#############################################

DROP TABLE IF EXISTS `ttmt_app_config`;
CREATE TABLE `ttmt_app_config`
(
    `id`                bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `main_version_code` int unsigned        NOT NULL default 0 COMMENT '主版本号',
    `sub_version_code`  int unsigned        NOT NULL default 0 COMMENT '次版本号',
    `gray_version_code` int unsigned        NOT NULL default 0 COMMENT '灰度版本号',
    `version_code`      varchar(20)         NOT NULL DEFAULT '' COMMENT '版本号字符串',
    `resource_url`      varchar(200)        NOT NULL DEFAULT '' COMMENT '当前版本apk下载地址',
    `update_info`       text(1000) COMMENT '版本更新内容',
    `created_time`      timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_time`      timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4 COMMENT ='点赞表';
