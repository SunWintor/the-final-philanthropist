CREATE TABLE `tfp_user` (
                            `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
                            `username` varchar(16) NOT NULL DEFAULT '' COMMENT '用户名',
                            `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
                            `token` varchar(255) NOT NULL DEFAULT '' COMMENT '用户鉴权用的token',
                            `ver` int(11) NOT NULL COMMENT '版本号，用于幂等',
                            `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
                            `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            PRIMARY KEY (`id`),
                            UNIQUE KEY uk_username(`username`),
                            KEY ix_ctime(`ctime`),
                            KEY ix_mtime(`mtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户信息表';


CREATE TABLE `tfp_user_game_history` (
                                         `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
                                         `room_id` varchar(16) NOT NULL DEFAULT '' COMMENT '房间的唯一标识',
                                         `game_id` varchar(16) NOT NULL DEFAULT '' COMMENT '游戏的唯一标识',
                                         `user_id` int(11) NOT NULL COMMENT '用户id',
                                         `hero` varchar(16) NOT NULL DEFAULT '' COMMENT '英雄名',
                                         `end_money` int(11) NOT NULL DEFAULT '0' COMMENT '游戏结束时剩余的金钱',
                                         `rank` int(11) NOT NULL DEFAULT '0' COMMENT '本局排名',
                                         `ranking` int(11) NOT NULL DEFAULT '0' COMMENT '本局获得了多少排位分',
                                         `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
                                         `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                         PRIMARY KEY (`id`),
                                         UNIQUE KEY uk_game_id(`game_id`),
                                         KEY ix_room_id(`room_id`),
                                         KEY ix_user_id(`user_id`),
                                         KEY ix_ctime(`ctime`),
                                         KEY ix_mtime(`mtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户游戏历史';


CREATE TABLE `tfp_user_rank` (
                                         `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
                                         `user_id` int(11) NOT NULL COMMENT '用户id',
                                         -- `season` int(11) NOT NULL DEFAULT '' COMMENT '赛季', 当前还不需要这个字段
                                         `ranking` int(11) NOT NULL DEFAULT '0' COMMENT '排位分',
                                         `ver` int(11) NOT NULL COMMENT '版本号，用于幂等',
                                         `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
                                         `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                         PRIMARY KEY (`id`),
                                         UNIQUE KEY uk_user_id(`user_id`),
                                         KEY ix_ranking(`ranking`),
                                         KEY ix_ctime(`ctime`),
                                         KEY ix_mtime(`mtime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户排名';


