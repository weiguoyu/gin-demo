DROP TABLE IF EXISTS `project_test`;
CREATE TABLE `project_test` (
    `project_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `name` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '项目名称',
    `user_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '用户id',
    `create_time` int(11) unsigned DEFAULT '1363033208' COMMENT '创建时间',
    `update_time` int(11) unsigned DEFAULT '1363033208' COMMENT '修改时间',
    PRIMARY KEY (`project_id`),
    KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COMMENT='项目';

DROP TABLE IF EXISTS `project_item_test`;
CREATE TABLE `project_item_test` (
    `project_item_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
    `resource_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '资源id',
    `resource_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' '资源名称',
    `resource_type` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '资源类型',
    `project_id` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '项目id',
    `create_time` int(11) unsigned DEFAULT '1363033208' COMMENT '最后修改时间',
    PRIMARY KEY (`project_item_id`),
    KEY `idx_name` (`resource_name`),
    KEY `idx_id` (`resource_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COMMENT='项目资源';