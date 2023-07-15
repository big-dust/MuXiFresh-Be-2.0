# 仅供参考
CREATE TABLE `user`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `nickname` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户名',
    `email` VARCHAR(255) NOT NULL DEFAULT '0' COMMENT '用户邮箱',
    `gender` VARCHAR(255) NOT NULL DEFAULT ''COMMENT '性别',
    `age` INTEGER NOT NULL DEFAULT 18 COMMENT '年龄',
    `hash_password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_email_unique`(`email`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;