-- 用户信息表

CREATE TABLE `t_user_info` (
    `id` VARCHAR(50) NOT NULL PRIMARY KEY COMMENT '用户id',
    `open_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'open_id',
    `phone` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '电话号码',
    `birthday` DATE DEFAULT NULL COMMENT '生日',
    `gender` TINYINT NOT NULL DEFAULT 0 COMMENT '性别 0-未知, 1-男, 2-女',
    `city` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '城市',
    `city_name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '城市名称',
    `nick_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '昵称',
    `avatar_url` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '头像 url',
    `tags` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '标签',
    `location` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '位置',
    `detail` VARCHAR(512) NOT NULL DEFAULT '' COMMENT '详情',
    UNIQUE KEY `idx_openid` (`open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';