CREATE TABLE `user`
(
    `id`         int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `userid`     varchar(45) NOT NULL COMMENT '用户id',
    `usersecretid`     varchar(45) NOT NULL COMMENT '用户密钥',
    `age`        int(10) unsigned NOT NULL DEFAULT 0 COMMENT '用户年纪',
    `sex`        int(10) unsigned NOT NULL DEFAULT 0 COMMENT '用户性别',
    `headurl`    varchar(45) NOT NULL DEFAULT "" COMMENT '用户头像',
    `nickname`   varchar(45) NOT NULL DEFAULT "" COMMENT '用户昵称',

    `mobile`     varchar(45) NOT NULL DEFAULT "" COMMENT '电话号码',
    `email`      varchar(45) NOT NULL DEFAULT "" COMMENT '邮件地址',

    `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    `delete_at` datetime DEFAULT NULL COMMENT 'Deleted Time',
    UNIQUE KEY (`userid`),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




// redis中存储: (核心存储用户的登陆相关信息), kv结构, key : "rme_loginkey_${userid}"
type LoginInfo struct {
    UserID string 
    LoginType int32         // 登陆类型, 0:wx, 1:mobile
    LoginID string          // 登陆id 
    AccessToken string 
    RefreshToken string 
    RefreshTokenInterval int64

    // todo => ip/mac等端上的登陆信息，目前是没啥用，不知道以后是否有用？
}

// ip,mac等防刷机制!!


