CREATE TABLE `tfp_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(16) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `token` varchar(255) NOT NULL DEFAULT '',
  `ver` int(11) NOT NULL,
  `ctime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `mtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY ix_username(`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;