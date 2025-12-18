use papers;

create table if not exists `t_tags`(
    `id` int auto_increment DEFAULT NULL,
    `name`  varchar(255) DEFAULT NULL,
    `parent_id` int DEFAULT NULL,
    `created_at` datetime NOT NULL COMMENT '创建时间',
    `updated_at` datetime DEFAULT NULL  COMMENT '更新时间'
)


