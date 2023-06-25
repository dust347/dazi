-- 用户信息表

CREATE TABLE `t_user_info` (
    `id` VARCHAR(50) NOT NULL PRIMARY KEY COMMENT '用户id',
    `identity_number` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '身份证号码',
    `phone` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '电话号码',
    `birthday` BIGINT NOT NULL DEFAULT 0 COMMENT '生日',
    `gender` TINYINT NOT NULL DEFAULT 0 COMMENT '性别 0-未知, 1-男, 2-女'
    `city` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '城市',
    `nick_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '昵称',
    `tags` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '标签',
    `location` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '位置',
    UNIQUE KEY `idx_identity_number` (`identity_number`),
    UNIQUE KEY `idx_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='psm 拓扑关系表';