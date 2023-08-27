-- 数据库文件，新创建文件的话需要执行此SQL用以初始化数据库

CREATE TABLE `app_info` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `router_fix` varchar(50) DEFAULT NULL COMMENT '路由前缀',
  `name` varchar(200) DEFAULT NULL COMMENT 'APP名称',
  `info` text DEFAULT NULL COMMENT 'APP简介',
  `router_timeout` int(10) unsigned DEFAULT 60 COMMENT '路由超时时间',
  `keep_timeout` int(10) unsigned DEFAULT 60 COMMENT '长连接超时时间',
  `tls_out` int(11) DEFAULT 10 COMMENT 'TLS握手超时时间',
  `autos` tinyint(3) unsigned DEFAULT 0 COMMENT '登录认证方式 0-无认证 1-JWT认证',
  `keys` tinyint(3) unsigned DEFAULT 0 COMMENT '是否启用KEY认证',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='APP详情信息，所有的路由都是基于APP进行发布的';

CREATE TABLE `authority_config` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `u_key` varchar(50) DEFAULT NULL COMMENT 'key',
  `source` varchar(200) DEFAULT NULL COMMENT 'key来源',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `frequency` int(10) NOT NULL DEFAULT 0 COMMENT '次数',
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除 0-正常 1-删除',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `u_key` (`u_key`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='授权key配置表';

CREATE TABLE `authority_exclusive` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `authority_config_id` int(11) NOT NULL DEFAULT 0,
  `use_type` int(11) DEFAULT 0,
  `frequency` int(10) NOT NULL DEFAULT 0 COMMENT '专用次数',
  `start_time` datetime DEFAULT NULL COMMENT '专用有效开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '专用有效结束时间',
  `is_deleted` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否删除 0-正常 1-删除',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='特殊配置表';

CREATE TABLE `authority_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `authority_config_id` int(10) DEFAULT NULL COMMENT 'authority_config_id',
  `type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '类型 0-无 1-专用 2-通用',
  `use_type` int(11) DEFAULT NULL,
  `desc` varchar(100) DEFAULT NULL COMMENT '描述',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `u_key` (`authority_config_id`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='请求记录表';

ALTER TABLE `app_info`   
  ADD COLUMN `autos_jwtUid` VARCHAR(100) NULL  COMMENT 'JWT中存放ID的键名' AFTER `autos`;

CREATE TABLE `app_router` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL COMMENT 'APP_ID',
  `router` varchar(200) DEFAULT NULL COMMENT '路由地址，需拼接所属APP的前缀',
  `targets` text DEFAULT NULL COMMENT '转发地址，[]string{}格式的JSON字符串',
  `is_login` tinyint(3) unsigned DEFAULT 0 COMMENT '是否需要登录授权【APP中的JWT类型TOKEN】',
  `is_key` tinyint(3) unsigned DEFAULT 0 COMMENT '是否需要KEY授权（权限鉴定）',
  `key_type` tinyint(3) unsigned DEFAULT 0 COMMENT 'KEY授权类型 0-无需授权 1-通用授权 2-专用授权',
  `key_use` int(11) DEFAULT 0 COMMENT '专用授权的授权值',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='APP路由表，用来存储路由的转发信息及是否需要授权信息';

CREATE TABLE `app_router_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `rid` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '路由ID',
  `r_uri` varchar(200) DEFAULT NULL COMMENT '请求URL地址，从targets中读取的某个地址列表',
  `code` int(10) unsigned DEFAULT NULL COMMENT '响应的HTTPCODE值',
  `respone_length` int(10) unsigned DEFAULT NULL COMMENT '响应的字节大小',
  `respone_time` int(10) unsigned DEFAULT NULL COMMENT '响应耗时 单位：毫秒',
  `client_time` datetime DEFAULT NULL COMMENT '请求时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='转发路由命中详情表';

ALTER TABLE `app_router_log`   
  ADD COLUMN `gatwat_time` INT UNSIGNED NULL  COMMENT '网关处理耗时 单位：毫秒' AFTER `respone_time`;
