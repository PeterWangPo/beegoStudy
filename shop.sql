CREATE TABLE `sh_user`(
	`id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '用户名',
	`nickname` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '昵称',
	`password` CHAR(32) NOT NULL DEFAULT '' COMMENT '密码',
	`email` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '邮箱',
	`cell_phone` VARCHAR(11) NOT NULL DEFAULT '' COMMENT '手机号',
	`is_email_verified` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '邮箱是否验证',
	`is_cell_phone_verified` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '手机是否验证',
	`sex` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '0未知 1男 2女 3保密',
	`status` TINYINT(1)	NOT NULL DEFAULT 1 COMMENT '1正常 2无效',
	`create_time` DATETIME NOT NULL COMMENT '注册时间',
	`last_login_time` DATETIME NOT NULL COMMENT '最后登陆时间',
	`last_login_ip` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '最后登陆ip',
	`birth_day` DATE NOT NULL DEFAULT '0000-00-00' COMMENT '生日',
	PRIMARY KEY (`id`),
	KEY `username` (`username`) USING BTREE,
	KEY `cell_phone` (`cell_phone`) USING BTREE
)ENGINE=InnoDB CHARSET=utf8;

CREATE TABLE `sh_admin`(
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '用户名',
	`password` CHAR(32) NOT NULL DEFAULT '' COMMENT '密码',
	`status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态：1正常,2删除',
	`create_time` DATETIME NOT NULL COMMENT '添加时间',
	`last_login_time` DATETIME NOT NULL COMMENT '最后登陆时间',
	`last_login_ip` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '最后登陆ip',
	PRIMARY KEY (`id`),
	KEY `username` (`username`) USING BTREE
)ENGINE=InnoDB CHARSET=utf8; 